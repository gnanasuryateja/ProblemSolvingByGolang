package displayTrees

import (
	"fmt"

	model "probsolbygolang/model"
)

func PreOrder(node *model.BSTNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	PreOrder(node.Left)
	PreOrder(node.Right)
}
