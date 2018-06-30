package xdiff

import "fmt"

//Debugging function to make sure that child nodes point to the correct parent
func (t *Tree) check() {
	checkRec(t.Root, nil)
}

func checkRec(node, parent *Node) {
	if parent == nil {
		// at root
		if node.LastChild != nil {
			checkRec(node.LastChild, node)
		}

		return
	}

	if node.LastChild != nil {
		checkRec(node.LastChild, node)
	}

	for n := node.PrevSibling; n != nil; n = n.PrevSibling {
		if n.Parent != parent {
			fmt.Printf("%-15s %s\n%-15s %s\n%-15s %s\n", "looking at", n, "was expecting", parent, "but got", n.Parent)
		}
	}
}
