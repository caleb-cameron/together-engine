package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	Position     pixel.Vec
	Velocity     pixel.Vec
	Speed        float64
	Acceleration float64
	Sprite       *pixel.Sprite
}

func NewPlayer(position pixel.Vec, speed float64, acceleration float64, sprite *pixel.Sprite) *Player {
	return &Player{
		Position:     position,
		Velocity:     pixel.Vec{},
		Speed:        speed,
		Acceleration: acceleration,
		Sprite:       sprite,
	}
}

func (p *Player) GetPosition() pixel.Vec {
	return p.Position
}

func (p *Player) Draw(target *pixelgl.Window) {
	p.Sprite.Draw(target, pixel.IM.Moved(p.Position))
}

func (p *Player) Update(dt float64, window *pixelgl.Window) {
	if window.Pressed(pixelgl.KeyLeft) || window.Pressed(pixelgl.KeyA) {
		if p.Velocity.X-p.Acceleration*dt < (p.Speed * -1) {
			p.Velocity.X = p.Speed * -1
		} else {
			p.Velocity.X -= p.Acceleration * dt
		}
	}

	if window.Pressed(pixelgl.KeyRight) || window.Pressed(pixelgl.KeyD) {
		if p.Velocity.X+p.Acceleration*dt > p.Speed {
			p.Velocity.X = p.Speed
		} else {
			p.Velocity.X += p.Acceleration * dt
		}
	}

	if window.Pressed(pixelgl.KeyUp) || window.Pressed(pixelgl.KeyW) {
		if p.Velocity.Y+p.Acceleration*dt > p.Speed {
			p.Velocity.Y = p.Speed
		} else {
			p.Velocity.Y += p.Acceleration * dt
		}
	}

	if window.Pressed(pixelgl.KeyDown) || window.Pressed(pixelgl.KeyS) {
		if p.Velocity.Y-p.Acceleration*dt < (p.Speed * -1) {
			p.Velocity.Y = p.Speed * -1
		} else {
			p.Velocity.Y -= p.Acceleration * dt
		}
	}

	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y

}

func (p *Player) ApplyFriction(dt float64, groundFriction float64) {
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
