[README]

pxmaze is a library and command-line utility to generate and solve JSON-formatted maze.

## Maze Solver ##

1. Test the solver using pre-defined tests
go build
go test

2. Test the solver using command-line argument
go build
./pxmaze solve <JSON Maze data>

[Example]

Input:

./pxmaze solve '{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "dragon"}, "right": {"forward": "dead end"}}'

Output:

[left forward upstairs]


## Maze Generator ##

Generate a maze in the style of the problem 

1. Test the generator using the command line
go build
./pxmaze generate <Optional:depth>

[Example 1]

Input:

./pxmaze generate

Output:

{"forward":{"downstairs":{"left":"demon"},"forward":{"downstairs":"dead end","upstairs":"ogre"},"left":{"forward":"exit","right":"tiger"},"right":"dead end"}}

[Example 2]

Input:

./pxmaze generate 4

Output:

{"downstairs":"dragon","forward":{"forward":{"downstairs":{"downstairs":"dragon","forward":{"forward":"ogre","left":"dead end","right":"tiger"},"left":{"forward":"tiger","upstairs":"dragon"},"right":{"downstairs":"tiger","upstairs":"dragon"},"upstairs":"ogre"},"left":{"left":{"downstairs":"dead end","forward":"ogre","left":"tiger"},"right":"tiger","upstairs":{"forward":"exit","right":"dragon"}},"right":{"downstairs":{"forward":"dragon","left":"ogre","upstairs":"dragon"}},"upstairs":{"forward":{"downstairs":"dragon","left":"demon","right":"dead end","upstairs":"tiger"},"left":{"forward":"ogre","right":"ogre"},"right":{"downstairs":"dragon","forward":"dragon","right":"dragon","upstairs":"dragon"}}},"left":{"downstairs":{"left":{"upstairs":"dragon"}},"upstairs":{"left":{"downstairs":"tiger","forward":"demon","right":"tiger","upstairs":"ogre"},"right":{"downstairs":"dragon","left":"dragon","right":"demon"}}},"upstairs":{"forward":"ogre","left":"dragon"}},"left":"dead end"}

## Generator-Solver ##

Generate a maze then solve it for an exit.

./pxmaze gensolve <Optional: depth>

[Example 1]

Input:
./pxmaze gensolve

Output:

[Maze]
 {"downstairs":{"downstairs":{"left":"tiger","right":"demon"}},"left":{"downstairs":"exit","forward":"tiger"},"upstairs":{"forward":{"forward":"dragon"}}}
[Solution]
[left downstairs]

[Example 2]

Input:

./pxmaze gensolve 4

Output:

[Maze]
 {"forward":"dragon","left":{"left":"tiger","upstairs":{"forward":{"downstairs":"tiger","forward":{"downstairs":"exit","left":"dead end","right":"demon"},"left":{"upstairs":"ogre"}}}},"right":"ogre","upstairs":{"forward":{"right":"ogre"},"right":{"right":"dragon"}}}
[Solution]
[left upstairs forward left downstairs]
