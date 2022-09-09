package main

import (
	_ "fmt"
	"strings"
	"testing"
)

type solverTest struct {
	maze     string
	expected string
	desc     string
}

var solverTests = []solverTest{
	solverTest{
		desc:     "Test case from problem statement PDF",
		maze:     `{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "dragon"}, "right": {"forward": "dead end"}}`,
		expected: "left forward upstairs",
	},
	solverTest{
		desc:     "Test case from problem statement PDF",
		maze:     `{"forward": "exit"}`,
		expected: "forward",
	},
	solverTest{
		desc:     "Test case from problem statement PDF",
		maze:     `{"forward": "tiger", "left": "ogre", "right": "demon"}`,
		expected: "Sorry",
	},
	solverTest{
		desc:     "Check if the exit is the nearest exit.",
		maze:     `{"A": {"B": {"C": {"Z": "exit"}}}, "D": {"E": "exit"}, "F": {"G": {"H": "exit" }} , "I": "random stuff"}`,
		expected: "D E",
	},
	solverTest{
		desc:     "Check if the exit is the nearest exit.",
		maze:     `{"A": {"B": {"C": {"Z": "exit"}}}, "D": {"E": "exit"}, "F": {"G": {"H": "exit" }} , "I": "exit"}`,
		expected: "I",
	},
}

func TestSolver(t *testing.T) {
	var s Solver
	for _, test := range solverTests {
		output := s.SolveMazeVerbose(test.maze)

		// Assertion is needed here because the function above could return different types
		arr, isArr := output.([]string)
		if isArr {
			output = strings.Join(arr, " ")
		}
		if output != test.expected {
			t.Errorf("Expected %q but got %q for maze %q. Test case description: %q", test.expected, output, test.maze, test.desc)
		}
	}
}
