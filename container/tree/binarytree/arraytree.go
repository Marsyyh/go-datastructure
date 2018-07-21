package binarytree

// ArrayTree is Array representation of tree
// 1, 2, 3, 4, 5 =>
//        1
//    2      3
//  4  5  nil nil
type ArrayTree struct {
	Array []*TreeNode
}

// NewArrayTree is constructor of creating a ArrayTree with giving slice
func NewArrayTree(input []interface{}) *ArrayTree {
	re := &ArrayTree{}
	for i, v := range input {
		if i == 0 {
			re.Array = append(re.Array, &TreeNode{nil, nil, v})
		} else if v != nil {
			t := &TreeNode{nil, nil, v}
			re.Array = append(re.Array, t)
			if i%2 == 1 {
				re.Array[(i-1)/2].Left = t
			} else {
				re.Array[(i-1)/2].Right = t
			}
		}

	}
	return re
}

// GetRoot return the root TreeNode of tree
func (a *ArrayTree) GetRoot() *TreeNode {
	return a.Array[0]
}
