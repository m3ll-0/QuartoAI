package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

var nodeCounter = 0

func runAI(board [4][4]interface{}, pieceToPlace Piece) []interface{} {

	setOptimizedMaxDepth(initialOptimizationScale, board) // get maxDepth from config

	println(ColorWhite, "\nCurrent board: ")
	printBoard(board, true)

	// Get stock
	stock := removePiecesOnBoardFromStock(board, getFullStock())

	// Create root node and generate children
	rootNode := &TreeNode{nil, board, stock, []Move{}, []*TreeNode{}, nil,false, true, false }
	generateInitialChildren(rootNode, pieceToPlace)

	//rootNode.CountTree()
	//println(nodeCounter)
	//os.Exit(1)

	// Evaluate each first child node of root node independently and add first move to endList
	firstLevelNodeList := []*TreeNode{}

	println(fmt.Sprintf("[*] Evaluating children of root node using minimax algorithm."))
	for _, c := range rootNode.children {
		c.Evaluate()
		firstLevelNodeList = append(firstLevelNodeList, c)
	}

	println(fmt.Sprintf("[*] Determining best first level node and piece to give to opponent."))

	bestTuple := []interface{}{}

	// First, get all the first level nodes that contain a 1.
	// If there are multiple, determine best 'best' node by counting the losing/winning positions and use that score
	// If there is just 1, then choose that node
	// If there are none, choose best 'worst' node by counting the losing/winning positions and use that score
	// Or, go one level deeper and see if there is a child indicating a 1...

	// First, get all first level nodes that contain a 1.
	positiveFirstLevelNodeList := []*TreeNode{}

	for _, firstLevelNode := range firstLevelNodeList {
		if *firstLevelNode.score == 1{
			positiveFirstLevelNodeList = append(positiveFirstLevelNodeList, firstLevelNode)
		}
	}

	if len(positiveFirstLevelNodeList) == 1 { // Only one positive node, this node is the best. todo Next, determine best piece to give.
		// Get the best child. All children are enemy nodes in this case.
		// Repeat process. Determine the best enemy node in this case. Score 1 is the best. Score -1 is the worst.
		// ... Naive method, get random node with 1, if none, get random node with -1
		// Second option, sum the score of the children per enemy node and make list:
		// Then, for option in list, make sure enemy has no quarto.
		bestTuple = determineBestFirstLevelNode(positiveFirstLevelNodeList)
	} else if len(positiveFirstLevelNodeList) > 1 { // More than one positive node, determine best 'best' node from all positive nodes.
		bestTuple = determineBestFirstLevelNode(positiveFirstLevelNodeList)
	} else if len(positiveFirstLevelNodeList) == 0 { // Not a single positive node, in this case determine the best 'worst' node from all negative nodes.
		bestTuple = determineBestFirstLevelNode(firstLevelNodeList)
	}

	bestFirstLevelNode := bestTuple[0].(TreeNode)
	bestPieceToGive := bestTuple[1].(Piece)

	print(ColorGreen, fmt.Sprintf("[*] AI is done.\n"))

	if bestPieceToGive.form == "" { // Position always loses
		println(ColorRed, fmt.Sprintf("\nNo best piece to give, loss is imminent!"))
		printStatistics(rootNode)
		os.Exit(1)
	} else {
		println(ColorWhite, fmt.Sprintf("\nPlace piece [%v] on position [%v, %v]", bestFirstLevelNode.currentMoves[0].piece, bestFirstLevelNode.currentMoves[0].rowNumber, bestFirstLevelNode.currentMoves[0].columnNumber))
		println(fmt.Sprintf("Best piece to give: %v", bestPieceToGive))

		println("\nNext board: ")
		printBoard(rootNode.currentBoard, true, bestFirstLevelNode.currentMoves[len(bestFirstLevelNode.currentMoves) - 1])
	}

	printStatistics(rootNode)

	return []interface{}{bestFirstLevelNode.currentBoard, bestPieceToGive}
}

func determineBestFirstLevelNode(firstLevelNodeList []*TreeNode) []interface{}{
	bestFirstLevelNode := TreeNode{}
	bestPieceToGive := Piece{}
	bestPieceToGiveScore := math.Inf(-1)

	for _, firstLevelNode := range firstLevelNodeList{
		// For each child of first node, create a map with a piece and a score.
		// The score is determined by each move the opponent makes with its piece and the evaluation of the current move.
		scoreCountByPiece := make(map[Piece]int)
		for _, secondLevelNode := range firstLevelNode.children {
			scoreCountByPiece[secondLevelNode.currentMoves[1].piece] += *secondLevelNode.score
		}

		// Determine move with the highest score, and see if it beats a prior move.
		// If yes, set new best firstLevelNode to use as first move.
		// Set pieceToGive from scoreCountyPiece
		f:
		for piece,score := range scoreCountByPiece{
			if float64(score) > bestPieceToGiveScore {

				for _, secondLevelNode := range firstLevelNode.children {

					// Here we check if for each of the second level node and the score we just calculated for the firstLevelNode, if the piece
					// that has the highest score,... that if we give this piece to the opponent, the opponent has a quarto.
					if secondLevelNode.currentMoves[1].piece == piece && secondLevelNode.isQuartoNode { // If enemy move is not currentPieceToGive, continue. If enemy has quarto, disregard node
						continue f
					}
				}

				bestFirstLevelNode = *firstLevelNode
				bestPieceToGive = piece
				bestPieceToGiveScore = float64(score)
				println(colorYellow, fmt.Sprintf("New best piece found [%v] with score [%v].", piece, score))
			}
		}
	}

	tuple := []interface{}{
		bestFirstLevelNode,
		bestPieceToGive,
	}

	return tuple

}

// GeneratePiecePlacementBoards Generates all positions from setting a piece
func generateInitialChildren(rootNode *TreeNode, pieceToPlace Piece){

	println(ColorGreen, fmt.Sprintf("\n[*] Generating full game tree."))
	counter = 0 // Reset amount of nodes generated for statistics
	timeStart = time.Now() // Reset time for statistics

	// Place piece on every open spot
	for rowNumber, row := range rootNode.currentBoard{ // Every row
		for columnNumber, spot := range row {
			if spot == nil { // Spot is free

				// Create new board filling free spot
				newBoard := rootNode.currentBoard
				newBoard[rowNumber][columnNumber] = pieceToPlace

				newStock := removePieceFromStock(rootNode.currentStock, pieceToPlace)
				newMoveList := []Move{{rowNumber: rowNumber, columnNumber: columnNumber, piece: pieceToPlace}}

				newBoardHasQuarto := false

				if boardHasQuarto(newBoard) {
					newBoardHasQuarto = true
					println(ColorWhite, "\nFound Quarto in one move! Next board: ")
					printBoard(rootNode.currentBoard, true, newMoveList[len(newMoveList)-1])
					printStatistics(rootNode)
					os.Exit(1)
				}

				newNode := &TreeNode{
					currentBoard: newBoard,
					currentMoves: newMoveList,
					currentStock: newStock,
					children:     []*TreeNode{},
					parent: rootNode,
					isQuartoNode: newBoardHasQuarto,
					isPlayerTurn: len(newMoveList) %2 == 1,
					isTerminalNode: len(newStock) == 0 || newBoardHasQuarto,
					score: nil,
				}

				// AssignScore
				newNode.AssignScore()
				counter++

				rootNode.children = append(rootNode.children, newNode)

				addChildrenRecursively(newNode)
			}
		}
	}
}

func addChildrenRecursively(node *TreeNode){
	for _, pieceToPlace := range node.currentStock{
		// Place piece on every open spot
		for rowNumber, row := range node.currentBoard{ // Every row
			for columnNumber, spot := range row {
				if spot == nil { // Spot is free

					// Create new board filling free spot
					newBoard := node.currentBoard
					newBoard[rowNumber][columnNumber] = pieceToPlace
					newStock := removePieceFromStock(node.currentStock, pieceToPlace)

					newMove := Move{rowNumber: rowNumber, columnNumber: columnNumber, piece: pieceToPlace}
					newMoveList := append(node.currentMoves, newMove)

					newBoardHasQuarto := boardHasQuarto(newBoard)

					newNode := &TreeNode{
						currentBoard: newBoard,
						currentMoves: newMoveList,
						currentStock: newStock,
						parent: node,
						children: []*TreeNode{},
						isQuartoNode: newBoardHasQuarto,
						isPlayerTurn: len(newMoveList) %2 == 1,
						isTerminalNode: len(newStock) == 0 || newBoardHasQuarto,
						score: nil,
					}

					if len(newNode.currentMoves) >= maxDepth {
						newNode.isTerminalNode = true
					}

					// Assign score
					newNode.AssignScore()

					// Alpha-beta pruning
					// If a newNode has sibling with quarto, don't add node as child and don't generate children
					if newNode.SiblingHasQuartoWithPiece(){
						continue
					}

					// If current node is enemy node, and there exists a sibling node with the exact same piece that results in no quarto, remove sibling from tree
					if !newNode.isPlayerTurn && newNode.isQuartoNode {
						// Check sibling nodes to see if for current node there exist a sibling node with quarto

						// First get sibling nodes
						siblingNodes := node.children

						for _, siblingNode := range siblingNodes {
							// Check if the sibling node move is the same, in this case:
							// - Both nodes are enemy moves, the current node has quarto, the sibling node has no quarto, and they use the same move
							if siblingNode.currentMoves[len(siblingNode.currentMoves) - 1].piece == newNode.currentMoves[len(siblingNode.currentMoves) - 1].piece{
								// Delete sibling node from tree
								node.RemoveChildFromNode(siblingNode)
							}
						}
					}

					counter++ // Counter for diagnostic purposes
 					node.children = append(node.children, newNode)

					if !(len(newStock) == 0 || newNode.isQuartoNode) && !(len(newNode.currentMoves) >= maxDepth)  { // If no more pieces, stop iterating
						addChildrenRecursively(newNode)
					}
				}
			}
		}
	}
}

func (node *TreeNode) RemoveChildFromNode(childToRemove *TreeNode) {
	newChildrenList := removeChildFromParentsChildrenList(node.children, childToRemove)
	node.children = newChildrenList
}

func (node *TreeNode) AssignScore() {

	// AssignScore
	if node.isPlayerTurn && node.isQuartoNode {
		throw := 1
		node.score = &throw
	} else if !node.isPlayerTurn && node.isQuartoNode {
		throw := -1
		node.score = &throw
	} else if node.isTerminalNode {
		throw := 0
		node.score = &throw
	}
}

// CountTree Count the amount of nodes
func (node *TreeNode) CountTree() {

	nodeCounter++

	for _, cn := range node.children { // Go deep
			cn.CountTree()
	}
}

// Evaluate runs through the tree and caculates the score from the terminal nodes
// all the the way up to the root node
func (node *TreeNode) Evaluate() {
	for _, cn := range node.children { // Go deep
		if !cn.isTerminalNode {
			cn.Evaluate()
		}

		if cn.parent.score == nil {
			cn.parent.score = cn.score
		} else if !cn.isPlayerTurn && *cn.score > *cn.parent.score {
			cn.parent.score = cn.score
		} else if cn.isPlayerTurn && *cn.score < *cn.parent.score {
			cn.parent.score = cn.score
		}
	}
}

func (node *TreeNode) SiblingHasQuartoWithPiece() bool {

	if node.parent == nil {
		return false
	}

	parentNode := node.parent

	// For each sibling of current node
	for _, sibling := range parentNode.children{
		if sibling == node { // Sibling is it's own node
			continue
		}

		if sibling.isQuartoNode {
			// Get piece which forms quarto, latest piece
			lastMovePiece := sibling.currentMoves[len(sibling.currentMoves) - 1].piece

			if lastMovePiece == node.currentMoves[len(sibling.currentMoves) - 1].piece{ // Piece is the same
				return true
			}
		}
	}

	return false
}

func (node *TreeNode)hasForcedMove() bool{
	// Forced move means every child results in quarto
	quartoCounter := 0

	for _, cn := range node.children{
		if cn.isQuartoNode {
			quartoCounter++
		}
	}

	if quartoCounter == len(node.children) && quartoCounter > 1{
		return true
	}

	return false
}

func checkForcedMovesInTree(node *TreeNode){

	if !node.isPlayerTurn {
		if node.hasForcedMove() {
			println("Found Player forced move")
		}
	} else {
		if node.hasForcedMove() {
			println("Found Enemy forced move")
		}
	}

	for _, cn := range node.children{
		checkForcedMovesInTree(cn)
	}
}
