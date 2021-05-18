package main

import (
	"math"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

//判断是否存在路径
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	left := hasPathSum(root.Left, targetSum-root.Val)
	right := hasPathSum(root.Right, targetSum-root.Val)
	return left || right
}

//声明
var path [][]int

//求出所有路径
func pathSum(root *TreeNode, targetSum int) [][]int {
	var curPath []int
	//初始化
	path = make([][]int, 0)
	dfs(root, targetSum, curPath)
	return path
}

func dfs(root *TreeNode, targetSum int, curPath []int) {
	if root == nil {
		return
	}
	curPath = append(curPath, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		var val []int
		for _, v := range curPath {
			val = append(val, v)
		}
		path = append(path, val)
		return
	}
	dfs(root.Left, targetSum-root.Val, curPath)
	dfs(root.Right, targetSum-root.Val, curPath)
}

//二叉树的右视图 层次遍历的最右边
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	var result []int
	//入队
	var quene []*TreeNode
	quene = append(quene, root)

	for len(quene) > 0 {
		//当前层次的节点数量
		l := len(quene)
		for l > 0 {
			//队头
			cur := quene[0]
			if cur.Left != nil {
				quene = append(quene, cur.Left)
			}
			if cur.Right != nil {
				quene = append(quene, cur.Right)
			}
			//出队列
			quene = quene[1:]
			if l == 1 {
				result = append(result, cur.Val)
			}
			l--
		}
	}
	return result
}

/**
 * Definition for a Node.
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//补充Next,层次遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var quene []*Node
	quene = append(quene, root)
	for len(quene) > 0 {
		//层节点数量
		l := len(quene)

		for l > 0 {
			cur := quene[0]
			//出队列
			quene = quene[1:]
			l--
			if l > 0 {
				cur.Next = quene[0]
			} else {
				cur.Next = nil
			}

			if cur.Left != nil {
				quene = append(quene, cur.Left)
			}
			if cur.Right != nil {
				quene = append(quene, cur.Right)
			}

		}
	}
	return root
}

func connectNext(root *Node) *Node {
	if root == nil {
		return root
	}
	start := root
	for start != nil {
		var last, nextStart *Node
		for p := start; p != nil; p = p.Next {
			if p.Left != nil {
				handle(&last, &p.Left, &nextStart)
			}
			if p.Right != nil {
				handle(&last, &p.Right, &nextStart)
			}
		}
		//下一层开始起点
		start = nextStart
	}
	return root
}

//链接next  当前层 last上一个节点，p当前节点，nextStart每一层的起始节点
func handle(last, p, nextStart **Node) {
	if *last != nil {
		(*last).Next = *p
	}
	if *nextStart == nil {
		*nextStart = *p
	}
	*last = *p
}

//寻找BST树中的最小差值
func minDiffInBST(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	//递归中序遍历
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

//四叉树
type FourNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *FourNode
	TopRight    *FourNode
	BottomLeft  *FourNode
	BottomRight *FourNode
}

//构建四叉树
func construct(grid [][]int) *FourNode {
	//单个元素
	if len(grid) == 1 && len(grid[0]) == 1 {
		return &FourNode{Val: grid[0][0] == 1, IsLeaf: true}
	}
	//递归构建
	return buildFourTree(grid, 0, 0, len(grid)-1, len(grid)-1)
}

func buildFourTree(grid [][]int, i, j, x, y int) *FourNode {
	if isEqual(grid, i, j, x, y) {
		return &FourNode{Val: grid[i][j] == 1, IsLeaf: true}
	} else {
		root := FourNode{Val: true, IsLeaf: false}
		//分四部分构建
		row, col := x-i+1, y-j+1
		root.TopLeft = buildFourTree(grid, i, j, i+row/2-1, j+col/2-1)
		root.TopRight = buildFourTree(grid, i, j+col/2, i+row/2-1, y)

		root.BottomLeft = buildFourTree(grid, i+row/2, j, x, j+col/2-1)
		root.BottomRight = buildFourTree(grid, i+row/2, j+col/2, x, y)
		return &root
	}
}

//网格中下标的(i,j)->(x,y)是否全部相等
func isEqual(grid [][]int, i, j, x, y int) bool {
	if x-i == 0 || y-j == 0 {
		return true
	}
	for row := i; row <= x; row++ {
		for col := j; col <= y; col++ {
			if grid[row][col] != grid[i][j] {
				return false
			}
		}
	}
	return true
}

//二叉树左叶子和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Leave(root.Left, true) +Leave(root.Right, false)
}

func sumOfRightLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Leave(root.Left, false) +Leave(root.Right, true)
}

func Leave(root *TreeNode, dir bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil && dir {
		return root.Val
	}
	return Leave(root.Left, true) + Leave(root.Right, false)
}


func main() {

}
