package main

func removePieceFromStock(stock []Piece, pieceToBeRemoved Piece) []Piece {

	newStock := []Piece{}

	for _, currentPiece := range stock {
		if currentPiece != pieceToBeRemoved{
			newStock = append(newStock, currentPiece)
		}
	}

	return newStock

}

func removePiecesOnBoardFromStock(board [4][4]interface{}, stock []Piece) []Piece {

	for _, row := range board { // For each row
		for _, colPiece := range row{ // For each element in row

			if colPiece == nil {
				continue
			}

			pieceToBeRemoved := colPiece.(Piece)
			stock = removePieceFromStock(stock, pieceToBeRemoved)
		}
	}

	return stock
}


func getFullStock() []Piece {

	stock := []Piece{}

	stock = append(stock, BTCS, BTOS, BTCC, BTOC,  BSCS, BSOS, BSCC, BSOC, WTCS, WTOS, WTCC, WTOC, WSCS, WSOS, WSCC, WSOC)

	return stock
}

