package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type generator struct {
	cfg mazeGenConfig
}

// Only depth is supported currently
type mazeGenConfig struct {
	depth int
}

var directions = []string{"forward", "left", "right", "upstairs", "downstairs"}
var numDir = len(directions)

var attributes = []string{"dragon", "tiger", "ogre", "dragon", "dead end", "demon"}
var numAttr = len(attributes)

func NewGenerator() generator {
	return generator{
		cfg: NewMazeGenConfig(),
	}
}

func NewMazeGenConfig() mazeGenConfig {
	return mazeGenConfig{
		depth: 2,
	}
}

type Room map[string]any

// GenerateMazeJson returns a maze in JSON string format.
func (g *generator) GenerateMazeJson() string {
	maze := g.GenerateMaze()
	jsonMaze, err := json.Marshal(maze)

	if err != nil {
		panic(fmt.Sprintf("Failed to convert maze to JSON.\nOriginal error message: %v\n", err))
	}

	return string(jsonMaze)
}

// GenerateMaze generates an acyclic maze -- none of the paths lead back to an earlier path.
// For example, if you're at Room A and you go "upstairs" and then "downstairs", you will end up
// in Room B instead of the original starting point Room A.
func (g *generator) GenerateMaze() *Room {
	depth := g.cfg.depth

	// This is NOT very random but OK for use in a proof-of-concept
	rand.Seed(time.Now().UnixNano())

	exitPlaced := false

	startRoom := g.genRoom()
	queue := []*Room{startRoom}

	// BFS, process all rooms at the same depth before moving to "deeper" rooms in the maze
	for len(queue) > 0 {
		nextQueue := []*Room{}

		// Process all rooms with the same depth
		for i, room := range queue {

			// For each direction in the current room
			for dir := range *room {
				if depth > 0 {
					// 0...99
					chance := rand.Intn(100)

					if chance == 99 {
						// 1% chance of an early exit
						if !exitPlaced {
							(*room)[dir] = "exit"
							exitPlaced = true
						}
					} else if chance > 70 {
						// Place an attribute here. E.g. dragon or dead end
						(*room)[dir] = g.randAttr()
					} else {
						// Place a room in this direction
						var newRoom *Room = g.genRoom()
						(*room)[dir] = newRoom

						// Add room to queue
						nextQueue = append(nextQueue, newRoom)
					}

				} else {
					// Depth = 0, we're now at the outer layer of the maze. Stop adding more rooms to the maze
					if !exitPlaced {
						(*room)[dir] = "exit"
						exitPlaced = true
					} else {
						// Place an attribute here. E.g. dragon or dead end
						(*room)[dir] = g.randAttr()
					}
				}
			}

			// If this is the last room in current depth and there is still no link to the next depth,
			// renovate the first direction in the room to point to a room
			if depth > 0 && len(nextQueue) == 0 && i == len(queue)-1 {
				for dir := range *room {
					if (*room)[dir] == "exit" {
						continue
					}
					var newRoom *Room = g.genRoom()
					(*room)[dir] = newRoom

					// Add room to queue
					nextQueue = append(nextQueue, newRoom)
					break
				}
			}
		}

		queue = nextQueue
		depth--
	}

	return startRoom
}

// genRoom creates a room with one or more directions.
func (g *generator) genRoom() *Room {
	room := Room{}

	for _, dir := range directions {
		// 50/50 chance a direction is added
		if rand.Intn(10) >= 5 {
			// Initialize a dummy room in this direction
			room[dir] = "dead end"
		}
	}

	// The room must have at least one direction
	if len(room) == 0 {
		room[directions[rand.Intn(numDir)]] = "dead end"
	}

	return &room
}

// randAttr returns a randomized attribute
func (g generator) randAttr() string {
	return attributes[rand.Intn(numAttr)]
}

func (g *generator) setDepth(d int) {
	g.cfg.depth = d
}
