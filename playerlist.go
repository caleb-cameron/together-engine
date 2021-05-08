package engine

import (
	"fmt"
	"log"
	"sync"

	"github.com/faiface/pixel"
)

var PlayerList *playerList

type playerList struct {
	players              map[string]*Player
	recentConnections    []string
	recentDisconnections []string
	recentUpdates        []string
	mutex                sync.RWMutex
}

func newPlayerList() *playerList {
	return &playerList{
		players: map[string]*Player{},
		mutex:   sync.RWMutex{},
	}
}

func (p *playerList) AddPlayer(username string, player *Player) error {
	p.mutex.RLock()
	if _, ok := p.players[username]; ok {
		p.mutex.RUnlock()
		return fmt.Errorf("username taken.")
	}
	p.mutex.RUnlock()

	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.players[username] = player
	p.recentConnections = append(p.recentConnections, username)

	return nil
}

func (p *playerList) UpdatePlayer(username string, player *Player) error {
	p.mutex.RLock()
	if _, ok := p.players[username]; !ok {
		p.mutex.RUnlock()
		return fmt.Errorf("tried to update user that does not exist.")
	}
	p.mutex.RUnlock()

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// We never actually release this lock, because we immediately replace the player.
	// We only lock it to block reads first.
	p.players[username].Mutex.Lock()

	p.players[username] = player
	p.recentConnections = append(p.recentConnections, username)

	return nil
}

func (p *playerList) GetPlayer(username string) *Player {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if _, ok := p.players[username]; !ok {
		return nil
	}

	return p.players[username]
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

	// We never actually release this lock, because we immediately replace the player.
	// We only lock it to block reads first.
	p.players[username].Mutex.Lock()

	delete(p.players, username)
	p.recentDisconnections = append(p.recentDisconnections, username)

	return true
}

func (p *playerList) GetRecents() (*[]string, *[]string, *[]string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	connects := p.recentConnections
	disconnects := p.recentDisconnections
	updates := p.recentUpdates

	p.clearRecents()

	return &connects, &disconnects, &updates
}

/*
	ONLY CALL THIS INSIDE A MUTEX LOCK.
*/
func (p *playerList) clearRecents() {
	p.recentConnections = []string{}
	p.recentDisconnections = []string{}
	p.recentUpdates = []string{}
}

func (p *playerList) GetPlayers() map[string]*Player {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.players
}

func (p *playerList) UpdatePlayerVelocity(username string, vel pixel.Vec) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, ok := p.players[username]; !ok {
		log.Printf("Could not update player %s velocity: player not found.\n", username)
		return
	}

	p.players[username].Velocity = vel

}

func (p *playerList) UpdatePlayers(dt float64, groundFriction float64) {
	players := p.GetPlayers()
	for _, p := range players {
		p.ApplyFriction(dt, groundFriction)
		p.Update(dt)
	}
}
