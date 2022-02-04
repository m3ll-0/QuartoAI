package main

import (
	"fmt"
	"time"
)

func printBoard(board [4][4]interface{}, isPlayerMove bool, moves ...Move){

	var move *Move

	if len(moves) > 0 {
		move = &moves[0]
	} else {
		move = nil
	}

	for rowCounter, row := range board {
		print(ColorWhite, "{")
		for colCounter, col := range row {
			if col == nil {
				if move != nil {
					if move.rowNumber == rowCounter && move.columnNumber == colCounter {
						if(isPlayerMove){
							print(ColorGreen, move.piece.toString())
						} else {
							print(ColorRed, move.piece.toString())
						}
					} else {
						print(colorYellow, "nil")
					}

				} else { // always print nil
					print(colorYellow, "nil")
				}

			} else {
				print(ColorWhite, col.(Piece).toString())
			}

			if colCounter != 3 {
				print(ColorWhite, ", ")
			}
		}
		print(ColorWhite, "} ,")
		println("")
	}
}

func (piece Piece)toString() string{
	return string(piece.color[0]) + string(piece.height[0]) + string(piece.solidity[0]) + string(piece.form[0])
}

func printStatistics() {
	println(colorCyan, "\n ======== Statistics ========")
	println(fmt.Sprintf("Total amount of nodes generated: %v", counter))
	println(fmt.Sprintf("Total amount of time: %v", time.Since(timeStart)))
}

func printMoveList(moveList []Move){
	println("Movelist: ")
	for _, move := range moveList {
		println(fmt.Sprintf("Row: %v Column: %v ===> %v", move.rowNumber, move.columnNumber, move.piece))
	}
}

func printBanner(){
	print(colorPurple)
	println("  ______                                   __                       ______   ______ ")
	println(" /      \\                                 /  |                     /      \\ /      |")
	println("/$$$$$$  | __    __   ______    ______   _$$ |_     ______        /$$$$$$  |$$$$$$/ ")
	println("$$ |  $$ |/  |  /  | /      \\  /      \\ / $$   |   /      \\       $$ |__$$ |  $$ |  ")
	println("$$ |  $$ |$$ |  $$ | $$$$$$  |/$$$$$$  |$$$$$$/   /$$$$$$  |      $$    $$ |  $$ |")
	println("$$ |_ $$ |$$ |  $$ | /    $$ |$$ |  $$/   $$ | __ $$ |  $$ |      $$$$$$$$ |  $$ |  ")
	println("$$ / \\$$ |$$ \\__$$ |/$$$$$$$ |$$ |        $$ |/  |$$ \\__$$ |      $$ |  $$ | _$$ |_ ")
	println("$$ $$ $$< $$    $$/ $$    $$ |$$ |        $$  $$/ $$    $$/       $$ |  $$ |/ $$   |")
	println(" $$$$$$  | $$$$$$/   $$$$$$$/ $$/          $$$$/   $$$$$$/        $$/   $$/ $$$$$$/ ")
	println("     $$$/                                                                           ")
	print(ColorReset)
}

func setOptimizedMaxDepth(optimizationScale int, board [4][4]interface{}){
	// 0 = No optimization
	// 1 = Optimize lightly, prefer more time less optimization
	// 2 = Optimize heavily, prefer less time more optimization

	nilSpots := 0

	for _, row := range board{
		for _, col := range row {
			if col == nil{
				nilSpots++
			}
		}
	}

	psMaxDepth := 1000
	optimizationScaleString := "NIL"

	switch optimizationScale{
	case 0:
		optimizationScaleString = "UNLIMITED"
		psMaxDepth = 1000
	case 1:
		optimizationScaleString = "OPTIMIZED_LOW"
		if nilSpots <= 9{
			psMaxDepth = 1000
		} else if nilSpots == 10 {
			psMaxDepth = 6
		} else if nilSpots == 11 {
			psMaxDepth = 4
		} else if nilSpots == 12 {
			psMaxDepth = 4
		} else if nilSpots == 13 {
			psMaxDepth = 4
		} else if nilSpots > 13 {
			psMaxDepth = 3
		}
	case 2:
		optimizationScaleString = "OPTIMIZED_HIGH"

		if nilSpots <= 9{
			psMaxDepth = 1000
		} else if nilSpots == 10 {
			psMaxDepth = 5
		} else if nilSpots == 11 {
			psMaxDepth = 4
		} else if nilSpots == 12 {
			psMaxDepth = 4
		} else if nilSpots == 13 {
			psMaxDepth = 3
		} else if nilSpots > 13 {
			psMaxDepth = 3
		}
	}

	println(colorYellow, fmt.Sprintf("\nUsing optimization scheme [%v]. Setting max depth to %v.", optimizationScaleString, psMaxDepth))
	maxDepth = psMaxDepth

}

func getInputWhatPieceOpponentHasGiven(currentBoard [4][4]interface{}) Piece{

	pieceToReturn := Piece{}

	for true {
		time.Sleep(1 * time.Second)
		fmt.Println("\nEnter the name of the piece opponent has given: ")
		var pieceString string
		fmt.Scan(&pieceString)

		if pieceString == "BTCS" {
			pieceToReturn = BTCS
		} else if pieceString == "BTOS" {
			pieceToReturn = BTOS
		}  else if pieceString == "BTCC" {
			pieceToReturn = BTCC
		} else if pieceString == "BTOC" {
			pieceToReturn = BTOC
		} else if pieceString == "BSCS" {
			pieceToReturn = BSCS
		} else if pieceString == "BSOS" {
			pieceToReturn = BSOS
		}  else if pieceString == "BSCC" {
			pieceToReturn = BSCC
		} else if pieceString == "BSOC" {
			pieceToReturn = BSOC
		} else if pieceString == "WTCS" {
			pieceToReturn = WTCS
		} else if pieceString == "WTOS" {
			pieceToReturn = WTOS
		} else if pieceString == "WTCC" {
			pieceToReturn = WTCC
		} else if pieceString == "WTOC" {
			pieceToReturn = WTOC
		} else if pieceString == "WSCS" {
			pieceToReturn = WSCS
		} else if pieceString == "WSOS" {
			pieceToReturn = WSOS
		} else if pieceString == "WSCC" {
			pieceToReturn = WSCC
		} else if pieceString == "WSOC" {
			pieceToReturn = WSOC
		} else {
			println(ColorRed, fmt.Sprintf("Error: piece does not exist. Try again."))
			continue
		}

		// Check if piece is not on the board already
		if boardContainsPiece(currentBoard, pieceToReturn) {
			println(ColorRed, "Error! Piece entered is already on the board. Try again.")
			continue
		}

		break
	}

	return pieceToReturn
}

func boardContainsPiece(board [4][4]interface{}, piece Piece) bool{
	for _, row := range board {
		for _, col := range row {
			if col != nil {
				if col.(Piece) == piece {
					return true
				}
			}
		}
	}

	return false
}

func getInputPositionOpponentPlacePiece(board [4][4]interface{}, piece Piece) []interface{} {

	positionTuple := []interface{}{nil, nil}

	for true {
		time.Sleep(1 * time.Second)

		fmt.Print(fmt.Sprintf("\nEnter row position opponent placed [%v]: ", piece))
		var row int

		_, errRow := fmt.Scan(&row)
		if errRow != nil {
			println(ColorRed, "Error selecting row position of opponent piece. Try again.")
			continue
		}

		fmt.Print(fmt.Sprintf("Enter column position opponent placed [%v]: ", piece))
		var col int
		_, errCol := fmt.Scan(&col)
		if errCol != nil {
			println(ColorRed, "Error selecting column position of opponent piece. Try again.")
			continue
		}

		// Check if position of piece is within board bounds
		if row < 0 || row >= 4 || col < 0 || col >= 4{
			println(ColorRed, "Board position to place opponent's piece is out of bounds. Try again.")
			continue
		}

		// Check if there is not already a piece on board position
		if board[row][col] != nil {
			println(ColorRed, "Error selecting position opponent placed piece. Board position is already taken. Try again.")
			continue
		}

		positionTuple[0] = row
		positionTuple[1] = col
		break
	}

	return positionTuple
}
