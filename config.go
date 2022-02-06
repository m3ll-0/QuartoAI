package main

// Set initial board
var initialBoard = [4][4]interface{}{
	{BTCS, WSCC, BTOC, WTOC} ,
	{nil, WSOC, WSCS, nil} ,
	{nil, BSCS, WTCS, BTOS} ,
	{BTCC, WTOS, BSCC, nil} ,
}

// Set optimization scale
var initialOptimizationScale = 1

// Set generator method
var generatorMethodConfig = "DFS"

// Empty board
//	{nil, nil, nil, nil} ,
//	{nil, nil, nil, nil} ,
//	{nil, nil, nil, nil} ,
//	{nil, nil, nil, nil} ,
//pieceToPlace := BTCC

// Forced move
//	{WSCC, nil, WSCS, BTOC} ,
//	{nil, WSOC, BSCS, WSOS} ,
//	{BTCS, BSCC, BTOS, WTOC} ,
//	{WTOS, WTCS, BSOC, WTCC} ,
//pieceToPlace := BTCC


// Alg build with || Interesting because move that I place, and piece that I give lead to imminently losing position!
//board := [4][4]interface{}{
//	{WSCC, nil, WSCS, nil} ,
//	{nil, WSOC, nil, WSOS} ,
//	{BTCS, BSOC, BTOS, nil} ,
//	{WTOS, nil, nil, WTCC} ,
//pieceToPlace := WTCS


// Nice game: http://quarto.is-great.org/en/choix-case.php?N=1&P=10&B0=-1&B1=4&B2=11&B3=9&B4=0&B5=-1&B6=-1&B7=14&B8=-1&B9=-1&B10=-1&B11=-1&B12=-1&B13=5&B14=8&B15=-1
//board := [4][4]interface{}{
//{nil, WTCC, BSOS, BTOS} ,
//{WTCS, nil, nil, BSCC} ,
//{BSCS, nil, nil, nil} ,
//{nil, WTOC, BTCS, nil} ,
//}
//
////Place piece [{BLACK TALL OPEN CIRCLE}] on position [3, 3]
////Best piece to give: {WHITE SHORT OPEN CIRCLE}
//
//// Get stock
//stock := removePiecesOnBoardFromStock(board, getFullStock())
//
//// Set piece to place
//pieceToPlace := WSOC

// Interesting board, variation on imminent board (build AI):
//{WSCC, nil, WSCS, nil} ,
//{nil, WSOC, nil, WSOS} ,
//{BTCS, nil, BTOS, nil} ,
//{WTOS, nil, nil, WTCC} ,
//
//pieceToPlace := WTCS

//For a given piece I place on the board on some position
//For some piece I give to the opponent
//If the opponent places the piece on some position
//The opponent is able to give me a piece such that every position I place that piece I have no quarto, but every next move results in quarto for the opponent

// First check if child node contains 1,
// From those nodes, select...


// FOR DEBUGGING PURPOSES

//if newNode.isQuartoNode && !newNode.isPlayerTurn && newNode.isTerminalNode && len(newNode.currentMoves) == 2 {
//cm := []Move{{rowNumber: 2, columnNumber: 1, piece: WSOC}, {rowNumber: 1, columnNumber: 1, piece: BTCC}}
//
//if(newNode.currentMoves[0] == cm[0] && newNode.currentMoves[1] == cm[1]){
//println("FIndasojdasi")
//}
//
//}