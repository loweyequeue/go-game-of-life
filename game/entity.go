package game

import (
	"bytes"
)

type EntityEffectType int

const (
	Born EntityEffectType = iota
	Died
	Skip
)

type EntityEffect struct {
	Type   EntityEffectType
	Coords Vector2D
}

type Entity struct {
	Alive bool
}

func (e *Entity) Update(b *Board, coords Vector2D) EntityEffect {
	surroundings := b.GetSurrounding(coords) // TODO: i am left over after refactor

	var alive_count uint8 = 0
	var dead_count uint8 = 0

	for x, row := range surroundings {
		for y, entity := range row {
			if x == 1 && y == 1 {
				continue
			}

			if entity.Alive {
				alive_count++
			} else {
				dead_count++
			}
		}
	}

	if e.Alive {
		if alive_count < 2 {
			return EntityEffect{Died, coords}
		} else if alive_count < 4 {
			return EntityEffect{Skip, coords}
		} else if alive_count > 3 {
			return EntityEffect{Died, coords}
		}
	} else {
		if alive_count == 3 {
			return EntityEffect{Born, coords}
		}
	}

	return EntityEffect{Skip, coords}
}

func (e *Entity) Render(buf *bytes.Buffer) {
	if e.Alive {
		buf.Write([]byte(" ■"))
	} else {
		buf.Write([]byte("  "))
	}
}
