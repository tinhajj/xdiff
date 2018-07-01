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

func computeDistance(distTbl *DistTable, minMatching *MinCostMatch, root1 *Node, root2 *Node) {
	child1 := root1.LastChild
	child2 := root2.LastChild

	for x := child1; x != nil; child1.PrevSibling {
		for y := child2; y != nil; child2.PrevSibling {
			if x.Signature == y.Signature {
				fmt.Println("we found a match")
				if bytesEqual(x.Hash, y.Hash) {
					fmt.Println("we found a perfect match")
					continue
				}
			}
		}
	}
}
