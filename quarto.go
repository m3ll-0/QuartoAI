package main

import (
	"fmt"
	"os"
)

func boardHasQuarto(board [4][4]interface{}) bool {

	// Check rows
	flagRows:
	for _, row := range board { // For each row

		pieceListToCheck := []Piece{}

		for _, colPiece := range row{ // For each element in row


			if colPiece == nil {
				continue flagRows
			}

			castedColPiece := colPiece.(Piece)
			pieceListToCheck = append(pieceListToCheck, castedColPiece)
		}

		// Check pieceList
		//println("Checking rows!")
		if checkQuartoPieces(pieceListToCheck) {
			return true
		}
	}

	// Check columns
	flagColumns:
	for i := 0; i < 4; i++{

		pieceListToCheck := []Piece{}

		for j := 0; j < 4; j++{
			piece := board[j][i] // For each row, get element 0

			if piece == nil {
				continue flagColumns
			}

			castedColPiece := piece.(Piece)
			pieceListToCheck = append(pieceListToCheck, castedColPiece)
		}

		// Check pieceList
		//println("Checking columns!")
		if checkQuartoPieces(pieceListToCheck) {
			return true
		}
	}

	// Check diagonals
	pieceListDiagonal1ToCheck := []Piece{}

	for i := 0; i < 4; i++{

		piece := board[i][i] // For each row, get element 0

		if piece == nil {
			continue
		}

		castedColPiece := piece.(Piece)
		pieceListDiagonal1ToCheck = append(pieceListDiagonal1ToCheck, castedColPiece)

	}

	// Check pieceList
	if len(pieceListDiagonal1ToCheck) == 4{
		//println("Checking diagonal 1!")
		if checkQuartoPieces(pieceListDiagonal1ToCheck) {
			return true
		}
	}


	// Check diagonals
	pieceListDiagonal2ToCheck := []Piece{}

	for i := 0; i < 4; i++{

		columnNumber := i
		rowNumber := 3 - columnNumber

		piece := board[rowNumber][columnNumber] // For each row, get element 0

		if piece == nil {
			continue
		}

		castedColPiece := piece.(Piece)
		pieceListDiagonal2ToCheck = append(pieceListDiagonal2ToCheck, castedColPiece)

	}

	// Check pieceList
	if len(pieceListDiagonal2ToCheck) == 4 {
		//println("Checking diagonal 2!")
		if checkQuartoPieces(pieceListDiagonal2ToCheck) {
			return true
		}
	}

	// Check squares
	SquareList := []interface{}{
		[]interface{}{board[0][0], board[0][1], board[1][0], board[1][1]},
		[]interface{}{board[1][0], board[1][1], board[2][0], board[2][1]},
		[]interface{}{board[2][0], board[2][1], board[3][0], board[3][1]},

		[]interface{}{board[0][1], board[0][2], board[1][1], board[1][2]},
		[]interface{}{board[1][1], board[1][2], board[2][1], board[2][2]},
		[]interface{}{board[2][1], board[2][2], board[3][1], board[3][2]},

		[]interface{}{board[0][2], board[0][3], board[1][2], board[1][3]},
		[]interface{}{board[1][2], board[1][3], board[2][2], board[2][3]},
		[]interface{}{board[2][2], board[2][3], board[3][2], board[3][3]},
	}

	squareFlag:
	for _, square := range SquareList {

		pieceListToCheck := []Piece{}
		castedSquare := square.([]interface{})

		for _, piece := range castedSquare {

			if piece == nil {
				continue squareFlag
			}

			castedColPiece := piece.(Piece)
			pieceListToCheck = append(pieceListToCheck, castedColPiece)
		}

		// Check pieceList
		//println("Checking square list!")
		if checkQuartoPieces(pieceListToCheck) {
			return true
		}
	}

	return false
}

// Checks if piece list forms quarto
func checkQuartoPieces(pieceList []Piece) bool {

	if len(pieceList) != 4 {
		println("Error, piecelist is not 4!")
		println(fmt.Sprint(pieceList))
		os.Exit(1)
	}

	whiteCounter := 0
	shortCounter := 0
	openCounter := 0
	squareCounter := 0


	for _, piece := range pieceList {
		if piece.color == "WHITE" {
			whiteCounter++
		}
		if piece.height == "SHORT" {
			shortCounter++
		}
		if piece.solidity == "OPEN" {
			openCounter++
		}
		if piece.form == "SQUARE" {
			squareCounter++
		}
	}

	if whiteCounter == 4 || shortCounter == 4 || openCounter == 4 || squareCounter == 4 || whiteCounter == 0 || shortCounter == 0 || openCounter == 0 || squareCounter == 0 {
		return true
	}

	return false

}