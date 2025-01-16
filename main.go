package main

import (
	"bytes"
	"fmt"
	"gol/game"
	"os"
	"strconv"
	"time"
)

const CLEAR_CONSOLE = "\033[H\033[2J"

func main() {
	// Quick and dirty argument parsing
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run . <path_to_map> [frame_delay_ms (default: 200)]")
		os.Exit(1)
	}
	map_file := args[0]
	var delay uint64 = 200
	if len(args) == 2 {
		input, _ := strconv.ParseUint(args[1], 10, 32)
		delay = input
	}

	// load the board from a (.gol) file
	board, err := game.BoardFromFile(map_file)
	if err != nil {
		fmt.Println("Could not load file", map_file, "due to an error")
		panic(err)
	}

	// main loop
	var frame uint64 = 0
	keeping_up := true
	for {
		next_frame := time.Now().Add(time.Duration(delay) * time.Millisecond)
		var buf bytes.Buffer
		buf.WriteString(CLEAR_CONSOLE)
		board.Render(&buf)
		buf.WriteString(fmt.Sprint("Generation: ", frame, " | Population: ", board.Population))
		if !keeping_up {
			buf.WriteString(" | WARNING: Slow rendering")
		}
		fmt.Print(buf.String())
		frame++
		board = board.Update()
		remaining_time := time.Until(next_frame)
		if remaining_time > 0 {
			keeping_up = true
			time.Sleep(remaining_time)
		} else {
			keeping_up = false
		}
	}
}
