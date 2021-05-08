package engine

import (
	"fmt"
	"sync"

	"github.com/faiface/pixel"
)

var PlayerList *playerList

type playerList struct {
	players              map[string]*Player
	recentConnections    []string
	recentDisconnections []string
	mutex                sync.RWMutex
}

func newPlayerList() *playerList {
	return &playerList{
		players: map[string]*Player{},
		mutex:   sync.RWMutex{},
	}
}

func (p *playerList) AddPlayer(username string) error {
	p.mutex.RLock()
	if _, ok := p.players[username]; ok {
		p.mutex.RUnlock()
		return fmt.Errorf("username taken.")
	}
	p.mutex.RUnlock()

	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.players[username] = NewPlayer(pixel.Vec{}, PlayerSpeed, PlayerAcceleration, DefaultCharacterSprite)
	p.recentConnections = append(p.recentConnections, username)

	return nil
}

/*
	Returns true if the player was removed,
	Returns false if the player was not in the list.
*/
func (p *playerList) RemovePlayer(username string) bool {

	p.mutex.RLock()
	if _, ok := p.players[username]; !ok {
		p.mutex.RUnlock()
		return false
	}
	p.mutex.RUnlock()
	p.mutex.Lock()
	defer p.mutex.Unlock()

	delete(p.players, username)
	p.recentDisconnections = append(p.recentDisconnections, username)

	return true
}

func (p *playerList) GetRecents() (*[]string, *[]string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	connects := p.recentConnections
	disconnects := p.recentDisconnections

	p.clearRecents()

	return &connects, &disconnects
}

/*
	ONLY CALL THIS INSIDE A MUTEX LOCK.
*/
func (p *playerList) clearRecents() {
	p.recentConnections = []string{}
	p.recentDisconnections = []string{}
}

func (p *playerList) GetPlayers() map[string]*Player {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.players
}
