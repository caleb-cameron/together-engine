package engine

import (
	"sync"

	"github.com/faiface/pixel"
)

func init() {
	PlayerList = newPlayerList()
}

type Player struct {
	Username     string
	Position     pixel.Vec
	Velocity     pixel.Vec
	Speed        float64
	Acceleration float64
	Sprite       string
	mutex        sync.RWMutex
}

func NewPlayer(username string, position pixel.Vec, speed float64, acceleration float64, sprite string) *Player {
	return &Player{
		Username:     username,
		Position:     position,
		Velocity:     pixel.Vec{},
		Speed:        speed,
		Acceleration: acceleration,
		Sprite:       sprite,
		mutex:        sync.RWMutex{},
	}
}

func (p *Player) GetPosition() pixel.Vec {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.Position
}

func (p *Player) GetVelocity() pixel.Vec {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.Velocity
}

func (p *Player) Update(dt float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
}

func (p *Player) ApplyFriction(dt float64, groundFriction float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.Velocity.X < 0 {
		// Player is moving left
		if p.Velocity.X > groundFriction*dt*-1 {
			p.Velocity.X = 0
		} else {
			p.Velocity.X += groundFriction * dt
		}
	} else if p.Velocity.X > 0 {
		// Player is moving right
		if p.Velocity.X < groundFriction*dt {
			p.Velocity.X = 0
		} else {
			p.Velocity.X -= groundFriction * dt
		}
	}

	if p.Velocity.Y < 0 {
		// Player is moving down
		if p.Velocity.Y > groundFriction*dt*-1 {
			p.Velocity.Y = 0
		} else {
			p.Velocity.Y += groundFriction * dt
		}
	} else if p.Velocity.Y > 0 {
		// Player is moving up
		if p.Velocity.Y < groundFriction*dt {
			p.Velocity.Y = 0
		} else {
			p.Velocity.Y -= groundFriction * dt
		}
	}
}
