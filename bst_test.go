package tc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeTreeNodeStr(t *testing.T) {
	newTest := func(expectation string) TestDataFunc {
		return func() TestData {
			return TestData{
				Expectation: expectation,
				Input:       MakeTreeNode(expectation),
			}
		}
	}

	NewTestcases(t).
		AddFunc(newTest("[]")).
		AddFunc(newTest("[1]")).
		AddFunc(newTest("[1,2]")).
		AddFunc(newTest("[1,2,3]")).
		AddFunc(newTest("[10,5,15,3,7,13,18,1,null,6]")).
		AddFunc(newTest("[1,2,3,4,5,6,7]")).
		AddFunc(newTest("[1,2,3,4,5,6,7,8,9,10,11]")).
		AddFunc(newTest("[10,5,15,3,7,13,18,1,2,6,9,12,14,17,20]")).
		AddFunc(newTest("[1,null,15,null,null,13,18,null,null,null,null,null,null,17]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := MakeTreeNodeStr(td.Input.(*TreeNode))

			a.Equal(td.Expectation, actual)
		})
}

func TestMakeTreeNode(t *testing.T) {
	var nodeNil *TreeNode

	tc := NewTestcases(t).
		Add(nodeNil, "").
		Add(nodeNil, "[]")
	tc.AddExpectation(&TreeNode{
		Val: 1,
	}).AddInput("[1]")
	tc.AddExpectation(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}).AddInput("[1,2,3]")
	tc.AddExpectation(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}).AddInput("[1, 2, 3]")
	tc.AddExpectation(&TreeNode{
		Val:   1,
		Left:  &TreeNode{IsNil: true},
		Right: &TreeNode{Val: 2},
	}).AddInput("[1,null,2]")
	tc.AddExpectation(&TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}).AddInput("[1,2]")
	tc.AddExpectation(&TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
				Left: &TreeNode{
					Val: 12,
				},
				Right: &TreeNode{
					Val: 14,
				},
			},
			Right: &TreeNode{
				Val: 18,
				Left: &TreeNode{
					Val: 17,
				},
				Right: &TreeNode{
					Val: 20,
				},
			},
		},
	}).AddInput("[10,5,15,3,7,13,18,1,2,6,9,12,14,17,20]")
	tc.AddExpectation(&TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				IsNil: true,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}).AddInput("[10,5,15,3,7,null,18]")
	tc.AddExpectation(&TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					IsNil: true,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}).AddInput("[10,5,15,3,7,13,18,1,null,6]")
	tc.AddExpectation(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			IsNil: true,
			Left: &TreeNode{
				IsNil: true,
				Left: &TreeNode{
					IsNil: true,
				},
				Right: &TreeNode{
					IsNil: true,
				},
			},
			Right: &TreeNode{
				IsNil: true,
				Left: &TreeNode{
					IsNil: true,
				},
				Right: &TreeNode{
					IsNil: true,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
				Left: &TreeNode{
					IsNil: true,
				},
				Right: &TreeNode{
					IsNil: true,
				},
			},
			Right: &TreeNode{
				Val: 18,
				Left: &TreeNode{
					Val: 17,
				},
			},
		},
	}).AddInput("[1,null,15,null,null,13,18,null,null,null,null,null,null,17]")

	tc.Each(func(a *assert.Assertions, td TestData) {
		actual := MakeTreeNode(td.Input.(string))

		a.Equal(td.Expectation, actual)
	})
}
