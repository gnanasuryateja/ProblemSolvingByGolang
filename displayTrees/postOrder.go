package displayTrees

import (
	"fmt"

	model "probsolbygolang/model"
)

func PostOrder(node *model.BSTNode) {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Printf("%d ", node.Data)
}
