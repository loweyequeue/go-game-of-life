package game

import (
	"bufio"
	"bytes"
	"fmt"
	"gol/util"
	"math/rand/v2"
	"os"
	"strings"
)

const BOARD_SIZE = 32

type Board struct {
	inner      [BOARD_SIZE][BOARD_SIZE]Entity
	Population uint
}

func NewBlank() Board {
	return Board{}
}

func (b *Board) Get(cords Vector2D) Entity {
	cords.Normalize()
	return b.inner[cords.X][cords.Y]
}

// returns the 3x3 grid surrounding the given position including itself
func (b *Board) GetSurrounding(pos Vector2D) [3][3]Entity {
	return [3][3]Entity{
		{b.Get(Vector2D{pos.X - 1, pos.Y - 1}), b.Get(Vector2D{pos.X, pos.Y - 1}), b.Get(Vector2D{pos.X + 1, pos.Y - 1})},
		{b.Get(Vector2D{pos.X - 1, pos.Y}), b.Get(Vector2D{pos.X, pos.Y}), b.Get(Vector2D{pos.X + 1, pos.Y})},
		{b.Get(Vector2D{pos.X - 1, pos.Y + 1}), b.Get(Vector2D{pos.X, pos.Y + 1}), b.Get(Vector2D{pos.X + 1, pos.Y + 1})},
	}
}

func randomBool(prob float32) bool {
	return rand.Float32() < prob
}

func BoardFromFile(path string) (Board, error) {
	b := NewBlank()
	var line_no uint8 = 0
	file, e := os.Open(path)
	if e != nil {
		return b, e
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		fmt.Println("Line: ", line)
		util.Assert(len(line) >= BOARD_SIZE, "Line ", line_no+1, " is too short")
		for c := 0; c < BOARD_SIZE; c++ {
			switch line[c] {
			case '0':
				b.inner[line_no][c] = Entity{Alive: false}
			case '1':
				b.Population++
				b.inner[line_no][c] = Entity{Alive: true}
			case 'L':
				if randomBool(0.25) {
					b.Population++
					b.inner[line_no][c] = Entity{Alive: true}
				}
			case 'M':
				if randomBool(0.50) {
					b.Population++
					b.inner[line_no][c] = Entity{Alive: true}
				}
			case 'H':
				if randomBool(0.75) {
					b.Population++
					b.inner[line_no][c] = Entity{Alive: true}
				}
			}
		}
		line_no++
		if line_no == BOARD_SIZE {
			break
		}
	}
	return b, nil
}

func (b Board) Update() Board {
	board_clone := b
	for x, r := range b.inner {
		for y, e := range r {
			pos := Vector2D{int8(x), int8(y)}
			effect := e.Update(&b, pos)
			util.Assert(effect.Coords == pos, "Entity attempted to modify an entity other than itself")
			switch effect.Type {
			case Born:
				board_clone.Population++
				board_clone.inner[effect.Coords.X][effect.Coords.Y].Alive = true
			case Died:
				board_clone.Population--
				board_clone.inner[effect.Coords.X][effect.Coords.Y].Alive = false
			default:
				continue
			}
		}
	}
	return board_clone
}

func (b Board) Render(buf *bytes.Buffer) {
	for _, r := range b.inner {
		for _, e := range r {
			e.Render(buf)
		}
		buf.Write([]byte("\n"))
	}
}
