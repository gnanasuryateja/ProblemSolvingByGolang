package displayTrees

import (
	"fmt"

	model "probsolbygolang/model"
)

func InOrder(node *model.BSTNode) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Printf("%d ", node.Data)
	InOrder(node.Right)
}
