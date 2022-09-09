package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Support command-line argument, as required by the problem statement
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "solve":
			if len(os.Args) < 3 {
				panic(fmt.Sprintf("Error. Please provide maze data."))
			}

			var s Solver
			fmt.Println(s.SolveMazeVerbose(os.Args[2]))
		case "generate":
			g := NewGenerator()
			if len(os.Args) >= 3 {
				depth, err := strconv.Atoi(os.Args[2])
				if err != nil {
					panic(fmt.Sprintf("Invalid depth: %v.\nOriginal error message: %v\n", os.Args[2], err))
				}
				g.setDepth(depth)
			}
			fmt.Println(g.GenerateMazeJson())
		case "gensolve":
			var s Solver
			g := NewGenerator()
			if len(os.Args) >= 3 {
				depth, err := strconv.Atoi(os.Args[2])
				if err != nil {
					panic(fmt.Sprintf("Invalid depth: %v.\nOriginal error message: %v\n", os.Args[2], err))
				}
				g.setDepth(depth)
			}
			maze := g.GenerateMazeJson()
			fmt.Printf("[Maze]\n %v\n", maze)
			fmt.Printf("[Solution]\n%v\n", s.SolveMazeVerbose(maze))
		default:
			fmt.Println("Accepted commands are: solve, generate, gensolve")
		}
	} else {
		fmt.Println("Accepted commands are: solve, generate, gensolve")
	}
}
