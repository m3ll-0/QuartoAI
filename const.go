package main

type Piece struct {
	color string
	height string
	solidity string
	form string
}

type Move struct {
	rowNumber int
	columnNumber int
	piece Piece
}

type QuartoCombination struct{
	board [4][4]interface{}
	moveList []Move
}

type TreeNode struct {
	score *int
	currentBoard [4][4]interface{}
	currentStock []Piece
	currentMoves []Move
	children []*TreeNode
	parent *TreeNode
	isQuartoNode bool
	isPlayerTurn bool
	isTerminalNode bool
}

// Piece properties
const (
	black string = "BLACK"
	white string = "WHITE"

	short string = "SHORT"
	tall string = "TALL"

	open string = "OPEN"
	closed string = "CLOSED"

	square string = "SQUARE"
	circle string = "CIRCLE"
)

// All Quarto pieces

var BTCS = Piece{
color: black,
height: tall,
solidity: closed,
form: square,
}

var BTOS = Piece{
color: black,
height: tall,
solidity: open,
form: square,
}

var BTCC = Piece{
color: black,
height: tall,
solidity: closed,
form: circle,
}

var BTOC = Piece{
color: black,
height: tall,
solidity: open,
form: circle,
}

var BSCS = Piece{
color: black,
height: short,
solidity: closed,
form: square,
}

var BSOS = Piece{
color: black,
height: short,
solidity: open,
form: square,
}

var BSCC = Piece{
color: black,
height: short,
solidity: closed,
form: circle,
}

var BSOC = Piece{
color: black,
height: short,
solidity: open,
form: circle,
}

var WTCS = Piece{
color: white,
height: tall,
solidity: closed,
form: square,
}

var WTOS = Piece{
color: white,
height: tall,
solidity: open,
form: square,
}

var WTCC = Piece{
color: white,
height: tall,
solidity: closed,
form: circle,
}

var WTOC = Piece{
color: white,
height: tall,
solidity: open,
form: circle,
}

var WSCS = Piece{
color: white,
height: short,
solidity: closed,
form: square,
}

var WSOS = Piece{
color: white,
height: short,
solidity: open,
form: square,
}

var WSCC = Piece{
color: white,
height: short,
solidity: closed,
form: circle,
}

var WSOC = Piece{
color: white,
height: short,
solidity: open,
form: circle,
}

// Colors to print

const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const ColorWhite = "\033[97m"