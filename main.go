package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game struct holds all your game's data.
type Game struct{}

// Update is called every frame to update game logic.
func (g *Game) Update() error {
	// For now, we'll leave this empty.
	return nil
}

// Draw is called every frame to draw the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw a "Hello, World!" message.
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout defines the game's window size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	// Create a new Game object.
	game := &Game{}

	// Set the window size and title.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("My Go Game")

	// Run the game!
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
