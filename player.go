package engine

import (
	"sync"

	"github.com/abeardevil/together-engine/pb"
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
	Mutex        sync.RWMutex
}

func NewPlayer(username string, position pixel.Vec, speed float64, acceleration float64, sprite string) *Player {
	return &Player{
		Username:     username,
		Position:     position,
		Velocity:     pixel.Vec{},
		Speed:        speed,
		Acceleration: acceleration,
		Sprite:       sprite,
		Mutex:        sync.RWMutex{},
	}
}

func PlayerFromProto(proto *pb.PlayerEvent) *Player {
	return &Player{
		Username:     proto.Username,
		Position:     pixel.Vec{X: float64(proto.Position.Position.X), Y: float64(proto.Position.Position.Y)},
		Velocity:     pixel.Vec{X: float64(proto.Position.Velocity.X), Y: float64(proto.Position.Velocity.Y)},
		Speed:        PlayerSpeed,
		Acceleration: PlayerAcceleration,
		Sprite:       DefaultCharacterSprite,
		Mutex:        sync.RWMutex{},
	}
}

func (p *Player) GetPosition() pixel.Vec {
	p.Mutex.RLock()
	defer p.Mutex.RUnlock()
	return p.Position
}

func (p *Player) GetVelocity() pixel.Vec {
	p.Mutex.RLock()
	defer p.Mutex.RUnlock()
	return p.Velocity
}

func (p *Player) Update(dt float64) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
}

func (p *Player) ApplyFriction(dt float64, groundFriction float64) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

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
