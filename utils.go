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

func computeDist2(distTbl *DistTable, minMatching *MinCostMatch, root1 *Node, root2 *Node) {
	for child1, child2 := root1.LastChild, root2.LastChild; child1.LastChild != nil && child2.LastChild != nil; child1, child2 = child1.LastChild, child2.LastChild {
		for x := child1; x != nil; x = child1.PrevSibling {
			for y := child2; y != nil; y = child2.PrevSibling {
				if x.Signature == y.Signature {
					fmt.Println("signature match")
					fmt.Println("x sig", x.Signature)
					fmt.Println("y sig", y.Signature)
					if bytesEqual(x.Hash, y.Hash) {
						fmt.Println("we found a perfect match")
						continue
					}
				}
			}
		}
	}

}
