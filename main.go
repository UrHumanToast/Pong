package main

import (
	objects "github.com/UrHumanToast/ScreenHockey/modules/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	winWidth  = 1280
	winHeight = 720
	winTitle  = "Screen Hockey"
)

var (
	myBall       objects.Ball
	myPaddle_1   objects.Paddle
	myPaddle_2   objects.Paddle
	myWinnerText objects.ScreenText
	ballSpeed    float32 = -5
	bkgImage     rl.Texture2D
	bkgSound     rl.Music
	hitSounds    []rl.Sound
	bounceSounds []rl.Sound
	scoreSound   rl.Sound
	resetClock   uint8
)

/******************** MAIN LOGIC ********************/
func main() {
	mainInit()

	for !rl.WindowShouldClose() {
		mainUpdate()
		mainRender()
	}

	mainQuit()
}

/******************** FUNCTIONS ********************/

// mainInit will initialize the application contents
func mainInit() {
	// Create a window
	rl.InitWindow(winWidth, winHeight, winTitle)
	// Aim FPS for the monitor refresh rate
	rl.SetWindowState(rl.FlagVsyncHint)

	bkgImage = rl.LoadTexture("resources/textures/Table1280.png")
	rl.InitAudioDevice()

	bkgSound = rl.LoadMusicStream("resources/sounds/bkgTable.mp3")
	rl.PlayMusicStream(bkgSound)

	hitSounds = []rl.Sound{rl.LoadSound("resources/sounds/hit1.mp3"), rl.LoadSound("resources/sounds/hit2.mp3"), rl.LoadSound("resources/sounds/hit3.mp3")}
	bounceSounds = []rl.Sound{rl.LoadSound("resources/sounds/bounce1.mp3"), rl.LoadSound("resources/sounds/bounce2.mp3"), rl.LoadSound("resources/sounds/bounce3.mp3")}
	scoreSound = rl.LoadSound("resources/sounds/score.mp3")

	objInit()
}

// mainQuit will free memory and close the application
func mainQuit() {
	rl.UnloadMusicStream(bkgSound)
	for _, sound := range hitSounds {
		rl.UnloadSound(sound)
	}
	for _, sound := range bounceSounds {
		rl.UnloadSound(sound)
	}
	rl.UnloadSound(scoreSound)
	rl.CloseAudioDevice()
	rl.UnloadTexture(bkgImage)
	rl.CloseWindow()
}

// mainUpdate will update object positions and logic
func mainUpdate() {

	rl.UpdateMusicStream(bkgSound)

	myBall.Update()

	myPaddle_1.UpdateMovement()
	myPaddle_1.UpdateWallCollision()
	myPaddle_1.UpdateCollisionBall(&myBall)

	myPaddle_2.UpdateMovement()
	myPaddle_2.UpdateWallCollision()
	myPaddle_2.UpdateCollisionBall(&myBall)

	if myBall.X == 0 {
		if resetClock == 0 {
			blueWinCondition()
			myWinnerText.UpdateBool(true)
		}
		myBall.UpdateHalt()
		ballSpeed = -5
		resetClock++
	}
	if myBall.X == float32(rl.GetScreenWidth()) {
		if resetClock == 0 {
			redWinCondition()
			myWinnerText.UpdateBool(true)
		}
		myBall.UpdateHalt()
		ballSpeed = 5
		resetClock++
	}
	if resetClock >= 60*3 {
		objInit()
	}
}

// mainRender begins drwaing and updates pull requests
func mainRender() {
	rl.BeginDrawing()

	// Game Config/Static
	rl.ClearBackground(rl.NewColor(0, 0, 0, 255))
	rl.DrawTexture(bkgImage, 0, 0, rl.White)
	rl.DrawFPS(10, 10)

	// Draw game objects
	myWinnerText.Draw()
	myBall.Draw()
	myPaddle_1.Draw()
	myPaddle_2.Draw()

	rl.EndDrawing()
}

func objInit() {
	// Create the game objects
	myBall = objects.Ball{
		X:           float32(rl.GetScreenWidth() / 2),
		Y:           float32(rl.GetScreenHeight() / 2),
		Radius:      10,
		SpeedX:      ballSpeed,
		SpeedY:      0.25,
		Color:       rl.Black,
		BounceSound: bounceSounds,
		CurSound:    0,
	}
	myPaddle_1 = objects.Paddle{
		X:        80,
		Y:        int32(rl.GetScreenHeight() / 2),
		Width:    10,
		Height:   100,
		Speed:    10,
		Color:    rl.Red,
		MoveUp:   rl.KeyW,
		MoveDown: rl.KeyS,
		HitSound: hitSounds,
		CurSound: 1,
	}
	myPaddle_2 = objects.Paddle{
		X:        int32(rl.GetScreenWidth()) - 80,
		Y:        int32(rl.GetScreenHeight() / 2),
		Width:    10,
		Height:   100,
		Speed:    10,
		Color:    rl.Blue,
		MoveUp:   rl.KeyUp,
		MoveDown: rl.KeyDown,
		HitSound: hitSounds,
		CurSound: 0,
	}
	myWinnerText = objects.ScreenText{
		FontSize: 60,
		Flag:     false,
	}

	resetClock = 0

}

func redWinCondition() {
	rl.PlaySound(scoreSound)
	myWinnerText.UpdateText("Red Wins!")
	myWinnerText.UpdateColor(rl.Red)
	myWinnerText.UpdatePosition(int32(rl.GetScreenWidth()/2)-470, int32(rl.GetScreenHeight()/2)-30)
}

func blueWinCondition() {
	rl.PlaySound(scoreSound)
	myWinnerText.UpdateText("Blue Wins!")
	myWinnerText.UpdateColor(rl.Blue)
	myWinnerText.UpdatePosition(int32(rl.GetScreenWidth()/2)+180, int32(rl.GetScreenHeight()/2)-30)
}
