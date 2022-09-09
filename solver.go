package main

import (
	"encoding/json"
	"fmt"
)

type Solver struct{}

type mazeData struct {
	Room    *map[string]any
	History *[]string
}

// solveMazeVerbose returns the exit path as a slice or a "Sorry" string if there's no exit.
func (s Solver) SolveMazeVerbose(rawJsonMaze string) any {
	path := s.SolveMaze(rawJsonMaze)

	if len(path) == 0 {
		return "Sorry"
	} else {
		return path
	}
}

// SolveMaze returns a path from the starting room to the exit in the format of a slice
func (s Solver) SolveMaze(rawJsonMaze string) []string {
	// Convert raw JSON to Go map
	// Reference: https://stackoverflow.com/questions/33436730/unmarshal-json-with-some-known-and-some-unknown-field-names
	var maze map[string]any
	if err := json.Unmarshal([]byte(rawJsonMaze), &maze); err != nil {
		panic(fmt.Sprintf("Failed to decode JSON string.\nOriginal error: %v\n", err))
	}

	return s.SolveMazeBFS(&maze)
}

// SolveMazeBFS looks for the nearest exit from the starting room and returns the path.
func (s Solver) SolveMazeBFS(maze *map[string]any) []string {

	// Create a queue and put the starting point of the maze in it (as part of mazeData)
	queue := [](*mazeData){&mazeData{maze, &[]string{}}}

	for len(queue) > 0 {
		// Pop a room from the queue to explore
		data := queue[0]

		// Update the queue. This is not the most memory-efficient way to do it, but it saves the computing power
		// of having to copy over the remainder the of queue every iteration
		queue = queue[1:]

		// Look at every available direction in the room and examine its content
		for dir, content := range *data.Room {
			history := append(*(data.History), dir)
			if attrStr, ok := content.(string); ok {
				// The only attribute we're interested in is the exit
				// Ignore other attributes like animals and dead-end
				if attrStr == "exit" {
					return history
				}
			} else {
				// Otherwise, this is a map which represents a room. This is deduced from the problem definition.
				subMaze := content.(map[string]any)
				// Add to queue
				queue = append(queue, &mazeData{&subMaze, &history})
			}
		}
	}

	return []string{}
}

// SolveMazeDFS looks for an exit in the maze recursively, but the exit it finds is not guaranteed to be the nearest.
// This is a backtracking algorithm.
// For demonstration only, this is NOT USED in the current code.
func (s Solver) SolveMazeDFS(maze *map[string]any, path *[]string) bool {
	for dir, content := range *maze {
		// Remember the current room
		(*path) = append((*path), dir)

		if attrStr, ok := content.(string); ok {
			// The only attribute we're interested in is the exit
			// Ignore other attributes like animals and dead-end
			if attrStr == "exit" {
				return true
			}
		} else {
			// Otherwise, this is a map which represents a room. This is deduced from the problem definition.
			subMaze := content.(map[string]any)

			// Explore this direction
			if s.SolveMazeDFS(&subMaze, path) {
				return true
			}
		}

		// Forget the current room because it's not on the exit path. I.e. backtracking
		(*path) = (*path)[0 : len(*path)-1]
	}

	return false
}
