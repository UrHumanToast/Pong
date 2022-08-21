package objects

import (
	"github.com/UrHumanToast/ScreenHockey/modules/utilities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxBallSpeed float32 = 12
)

/******************** BALL ********************/
type Ball struct {
	X           float32
	Y           float32
	Radius      float32
	SpeedX      float32
	SpeedY      float32
	Color       rl.Color
	BounceSound []rl.Sound
	CurSound    int
}

func (b *Ball) Update() {
	b.X += b.SpeedX
	b.Y += b.SpeedY

	if b.Y > float32(rl.GetScreenHeight()) {
		b.Y = float32(rl.GetScreenHeight())
		b.PlaySound()
		b.SpeedY *= -1
	}
	if b.Y < 0 {
		b.Y = 0
		b.PlaySound()
		b.SpeedY *= -1
	}
	if b.X > float32(rl.GetScreenWidth()) {
		b.X = float32(rl.GetScreenWidth())
		b.PlaySound()
		b.SpeedX *= -1
	}
	if b.X < 0 {
		b.X = 0
		b.PlaySound()
		b.SpeedX *= -1
	}
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.X), int32(b.Y), b.Radius, b.Color)
}

func (b *Ball) UpdateHalt() {
	b.SpeedX = 0
	b.SpeedY = 0
}

func (b *Ball) PlaySound() {
	if b.SpeedX != 0 && b.SpeedY != 0 {
		// Reset the sound index to 0 once the last one is reached
		if b.CurSound >= len(b.BounceSound) {
			b.CurSound = 0
		}

		// Play the current sound effect
		rl.PlaySound(b.BounceSound[b.CurSound])
		b.CurSound++
	}
}

/******************** PADDLE ********************/
type Paddle struct {
	X        int32
	Y        int32
	Width    int32
	Height   int32
	Speed    int32
	Color    rl.Color
	MoveUp   int32
	MoveDown int32
	HitSound []rl.Sound
	CurSound int
}

func (p *Paddle) GetRectangle() rl.Rectangle {
	return rl.NewRectangle(float32(p.X)-float32(p.Width)/2, float32(p.Y)-float32(p.Height)/2, float32(p.Width), float32(p.Height))
}

func (p *Paddle) Draw() {
	rl.DrawRectangleRec(p.GetRectangle(), p.Color)
}

func (p *Paddle) UpdateCollisionBall(b *Ball) {
	if rl.CheckCollisionCircleRec(rl.NewVector2(float32(b.X), float32(b.Y)), b.Radius, p.GetRectangle()) {

		if b.SpeedX < 0 {
			// Ball is moving left
			b.X = float32(p.X+p.Width/2) + float32(p.Width)/2
			b.SpeedY = ((b.Y - float32(p.Y)) / (float32(p.Height) / 2)) * -b.SpeedX
		} else {
			// Ball is moving right
			b.SpeedY = ((b.Y - float32(p.Y)) / (float32(p.Height) / 2)) * b.SpeedX
			b.X = float32(p.X-p.Width/2) - float32(p.Width)/2
		}

		// Reset the sound index to 0 once the last one is reached
		if p.CurSound >= len(p.HitSound) {
			p.CurSound = 0
		}

		// Play the current sound effect
		rl.PlaySound(p.HitSound[p.CurSound])
		p.CurSound++

		// Change ball direction, and increase seed slightly
		b.SpeedX *= -1.05

		// Set speed limit
		if !(utilities.InRange(float64((-1)*maxBallSpeed), float64(b.SpeedX), float64(maxBallSpeed))) {
			if b.SpeedX > maxBallSpeed {
				// Call is moving right after calculation
				b.SpeedX = maxBallSpeed
			} else {
				// Ball is moving left after calculation
				b.SpeedX = -maxBallSpeed
			}
		}
	}
}

func (p *Paddle) UpdateMovement() {
	if rl.IsKeyDown(p.MoveUp) {
		p.Y -= p.Speed
	}
	if rl.IsKeyDown(p.MoveDown) {
		p.Y += p.Speed
	}
}

func (p *Paddle) UpdateWallCollision() {
	if p.Y >= int32(rl.GetScreenHeight())-p.Height/2 {
		p.Y = int32(rl.GetScreenHeight()) - p.Height/2
	}
	if p.Y <= +p.Height/2 {
		p.Y = 0 + p.Height/2
	}
}

/******************** SCREEN TEXT ********************/

type ScreenText struct {
	Text     string
	X        int32
	Y        int32
	FontSize int32
	Flag     bool
	Color    rl.Color
}

func (t *ScreenText) Draw() {
	if t.Flag {
		rl.DrawText(t.Text, t.X, t.Y, t.FontSize, t.Color)
	}
}

func (t *ScreenText) UpdateText(text string) {
	t.Text = text
}

func (t *ScreenText) UpdateBool(flag bool) {
	t.Flag = flag
}

func (t *ScreenText) UpdateColor(color rl.Color) {
	t.Color = color
}

func (t *ScreenText) UpdatePosition(x, y int32) {
	t.X = x
	t.Y = y
}
