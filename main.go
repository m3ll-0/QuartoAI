package main

import (
	"time"
)

var counter int = 0
var timeStart = time.Now()
var maxDepth = 4

func main(){
	printBanner()

	board := initialBoard

	println(ColorWhite, "\nStarting board: ")
	printBoard(board, true)

	// Start AI loop
	for true {

		// Initial line
		println(colorPurple, "\n----------------------------------------------------------------------")

		time.Sleep(1 * time.Second)

		// Ask what piece opponent wants me to place
		nextPieceToPlace := getInputWhatPieceOpponentHasGiven(board)

		// Get next board and piece to give, give current board and piece to place
		gameTuple := runAI(board, nextPieceToPlace)

		board = gameTuple[0].([4][4]interface{})
		nextPieceToGive := gameTuple[1].(Piece)

		// Ask where opponent put piece
		positionToPlaceNextPiece := getInputPositionOpponentPlacePiece(board, nextPieceToGive)

		// Print board and highlight opponent piece
		println(ColorWhite, "\nBoard after opponent move: ")
		opponentMove := Move{rowNumber: positionToPlaceNextPiece[0].(int), columnNumber: positionToPlaceNextPiece[1].(int), piece: nextPieceToGive}
		printBoard(board, false, opponentMove)

		// Create new board with piece
		board[positionToPlaceNextPiece[0].(int)][positionToPlaceNextPiece[1].(int)] = nextPieceToGive
	}
}