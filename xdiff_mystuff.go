package xdiff

// Remove a node from a tree by dereferencing it.
// Takes a node and another node that is a sibling and pointing to the node we
// want to remove.
func removeNode(node, sibling *Node) {
	// This function can't do anything if given a root
	if node.Parent == nil {
		return
	}

	if node.Parent.LastChild == node {
		node.Parent.LastChild = node.PrevSibling
	} else {
		sibling.PrevSibling = node.PrevSibling
	}
}

// Exclude subtrees with equal hashes from cost calculation.
// xdiff algorithm recommends just checking top level trees under root for
// equality but theres no reason we can't go deeper into sub-trees/levels,
// this however could make the diff less sensical.
// l is used as slider between performance and quality. Higher number reduces
// quality and increases performance.
func myExcludeEqual(rootX, rootY *Node, l int) {
	if l <= 0 {
		return
	}
	var x, lastX, y, lastY *Node

	for x, lastX = rootX.LastChild, nil; x != nil; lastX, x = x, x.PrevSibling {
		for y, lastY = rootY.LastChild, nil; y != nil; lastY, y = y, y.PrevSibling {
			if bytesEqual(x.Hash, y.Hash) {
				removeNode(x, lastX)
				removeNode(y, lastY)

				break
			}
			if rootX.Signature == rootY.Signature {
				l--
				excludeEqual(x, y, l)
			}
		}
	}
}
