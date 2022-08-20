package main

import (
	objects "github.com/UrHumanToast/ScreenHockey/modules/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	winWidth  = 1280
	winHeight = 720
	winTitle  = "PONG!"
)

var (
	myBall     objects.Ball
	myPaddle_1 objects.Paddle
	myPaddle_2 objects.Paddle
	resetCount uint    = 60 * 3
	winnerText string  = " "
	ballSpeed  float32 = -5
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

	objInit()
}

// mainQuit will free memory and close the application
func mainQuit() {
	// Free window
	rl.CloseWindow()
}

// mainUpdate will update object positions and logic
func mainUpdate() {

	myBall.Update()

	myPaddle_1.UpdateMovement()
	myPaddle_1.UpdateWallCollision()
	myPaddle_1.UpdateCollisionBall(&myBall)

	myPaddle_2.UpdateMovement()
	myPaddle_2.UpdateWallCollision()
	myPaddle_2.UpdateCollisionBall(&myBall)

	if myBall.X == 0 {
		winnerText = "Right Player Wins!"
		ballSpeed = -5
	}
	if myBall.X == float32(rl.GetScreenWidth()) {
		winnerText = "Left Player Wins!"
		ballSpeed = 5
	}

	winnerUpdate()

}

// mainRender begins drwaing and updates pull requests
func mainRender() {
	rl.BeginDrawing()

	// Some config
	rl.ClearBackground(rl.NewColor(0, 0, 0, 255))
	rl.DrawFPS(10, 10)

	// Draw game objects
	myBall.Draw()
	myPaddle_1.Draw()
	myPaddle_2.Draw()

	rl.EndDrawing()
}

// winCondition checks for a winner
func winnerUpdate() {
	// If a player wins, display the winner text
	if winnerText != "" {
		// Continue to display until the time is up
		if resetCount < 60*5 {
			rl.DrawText(winnerText, int32(rl.GetScreenWidth()/2)-250, int32(rl.GetScreenHeight()/2)-30, 60, rl.Yellow)
			myBall.SpeedX = 0
			myBall.SpeedY = 0
			resetCount++
		} else {
			// Reset the game
			resetCount = 0
			objInit()
		}
	}
}

func objInit() {
	// Create the game objects
	myBall = objects.Ball{
		X:      float32(rl.GetScreenWidth() / 2),
		Y:      float32(rl.GetScreenHeight() / 2),
		Radius: 5,
		SpeedX: ballSpeed,
		SpeedY: 0.25,
		Color:  rl.White,
	}
	myPaddle_1 = objects.Paddle{
		X:        50,
		Y:        int32(rl.GetScreenHeight() / 2),
		Width:    10,
		Height:   100,
		Speed:    10,
		Color:    rl.White,
		MoveUp:   rl.KeyW,
		MoveDown: rl.KeyS,
	}
	myPaddle_2 = objects.Paddle{
		X:        int32(rl.GetScreenWidth()) - 50,
		Y:        int32(rl.GetScreenHeight() / 2),
		Width:    10,
		Height:   100,
		Speed:    10,
		Color:    rl.White,
		MoveUp:   rl.KeyUp,
		MoveDown: rl.KeyDown,
	}

	if resetCount == 0 {
		winnerText = ""
	}
}
