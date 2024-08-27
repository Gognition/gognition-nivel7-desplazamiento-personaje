package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
	"math"
)

const (
	baseScreenWidth  = 640
	baseScreenHeight = 480
	frameWidth       = 64
	frameHeight      = 64
)

type Game struct {
	playerImage *ebiten.Image
	playerX     float64
	playerY     float64
	frameIndex  int
	direction   int // 1 mira hacia la derecha, -1 mira hacia la izquierda
	frameCount  int
	isMoving    bool
	scale       float64
	speed       float64
}

func New() *Game {
	g := &Game{
		playerX:    baseScreenWidth / 2,
		playerY:    baseScreenHeight / 2,
		direction:  1,
		frameCount: 0,
		isMoving:   false,
		scale:      2.0,
		speed:      1.0,
	}

	//Carga la imagen del jugador
	img, _, err := ebitenutil.NewImageFromFile("gognition_gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	g.playerImage = img

	return g
}

func (g *Game) Update() error {
	g.isMoving = false

	moveSpeed := g.speed * g.scale

	// Movimiento de nuestro personaje
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= moveSpeed
		g.direction = -1
		g.isMoving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += moveSpeed
		g.direction = 1
		g.isMoving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.playerY -= moveSpeed
		g.isMoving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.playerY += moveSpeed
		g.isMoving = true
	}

	//Escalamiento de nuestro personaje
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.scale = math.Max(0.1, g.scale*0.99)
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.scale = math.Min(5.0, g.scale*1.01)
	}

	//Ajuste de la velocidad
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.speed = math.Max(0.1, g.scale*0.99)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.speed = math.Min(5.0, g.scale*1.01)
	}

	// Animación del personaje
	if g.isMoving {
		g.frameCount++
		if g.frameCount >= 5 {
			g.frameIndex = (g.frameIndex + 1) % 6 //Avanzamos al siguiente frame, volviendo al 0 después del 5
			g.frameCount = 0
		}
	} else {
		g.frameIndex = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	baseScreen := ebiten.NewImage(baseScreenWidth, baseScreenHeight)

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(float64(g.direction)*g.scale, g.scale)

	if g.direction == -1 {
		op.GeoM.Translate(frameWidth*g.scale, 0)
	}

	//Posicionar el personaje
	op.GeoM.Translate(g.playerX, g.playerY)

	sx := g.frameIndex * frameWidth
	frame := g.playerImage.SubImage(image.Rect(sx, 0, sx+frameWidth, frameHeight)).(*ebiten.Image)

	//Dibujar el personaje en la imagen base
	baseScreen.DrawImage(frame, op)

	//Añadir información de depuración
	ebitenutil.DebugPrint(baseScreen, fmt.Sprintf("Escala: %.2f\nVelocidad: %.2f", g.scale, g.speed))

	// Calcular la escala necesaria para ajustar el juego a la ventana actual
	screenScale := math.Min(
		float64(screen.Bounds().Dx())/float64(baseScreenWidth),
		float64(screen.Bounds().Dy())/float64(baseScreenHeight),
	)

	screenOp := &ebiten.DrawImageOptions{}
	screenOp.GeoM.Scale(screenScale, screenScale)

	screenOp.GeoM.Translate(
		float64(screen.Bounds().Dx())/2-float64(baseScreenWidth)*screenScale/2,
		float64(screen.Bounds().Dy())/2-float64(baseScreenHeight)*screenScale/2,
	)

	screen.DrawImage(baseScreen, screenOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(baseScreenWidth, baseScreenHeight)
	ebiten.SetWindowTitle("¡Aprende haciendo con Gognition")
	if err := ebiten.RunGame(New()); err != nil {
		log.Fatal(err)
	}
}
