package main

import (
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten"
)

const scale int = 3
const width = 640
const frame_rate = 10
const height = 480

type Game struct{}

var white color.Color = color.RGBA{255, 255, 255, 255}
var blue color.Color = color.RGBA{100, 0, 255, 255}
var grid [width][height]int = [width][height]int{}
var next_gen [width][height]int = [width][height]int{}
var count int = 0

func update() error {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			next_gen[x][y] = 0
			n := grid[x-1][y-1] + grid[x-1][y+0] + grid[x-1][y+1] + grid[x+0][y-1] + grid[x+0][y+1] + grid[x+1][y-1] + grid[x+1][y+0] + grid[x+1][y+1]
			if grid[x][y] == 0 && n == 3 {
				next_gen[x][y] = 1
			} else if n < 2 || n > 3 {
				next_gen[x][y] = 0
			} else {
				next_gen[x][y] = grid[x][y]
			}
		}
	}
	temp := next_gen
	next_gen = grid
	grid = temp
	return nil
}

func display(window *ebiten.Image) {
	window.Fill(blue)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					if grid[x][y] == 1 {
						window.Set(x*scale+i, y*scale+j, white)
					}
				}
			}

		}
	}
}

func frame(window *ebiten.Image) error {
	count++
	var err error = nil
	if count == frame_rate {
		err = update()
		count = 0
	}
	if !ebiten.IsDrawingSkipped() {
		display(window)
	}
	return err
}

func main() {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float64() > 0.5 {
				grid[x][y] = 1
			}
		}
	}
	if err := ebiten.Run(frame, width, height, 2, "Game of Life"); err != nil {
		log.Fatal(err)
	}
}
