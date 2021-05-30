package main

import (
	"math"
	"strconv"
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
	return Leave(root.Left, true) + Leave(root.Right, false)
}

func sumOfRightLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Leave(root.Left, false) + Leave(root.Right, true)
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

func getMinimumDifference(root *TreeNode) int {
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

/**
 * Definition for a MultiNode.
 */
type MultiNode struct {
	Val      int
	Children []*MultiNode
}

//求多叉树的最大深度
func maxDepth(root *MultiNode) int {
	if root == nil {
		return 0
	}
	var deep = 1
	for _, ch := range root.Children {
		chDep := maxDepth(ch)
		if chDep+1 > deep {
			deep = chDep + 1
		}
	}
	return deep
}

//判断二叉树是否子树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

//判断两颗树是否一样
func isSame(a, b *TreeNode) bool {
	//都为空
	if a == nil && b == nil {
		return true
	}
	//只有一个为空
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}

//先序遍历转换为字符串
func tree2str(root *TreeNode) string {
	if root == nil {
		return "()"
	}
	var ans = strconv.Itoa(root.Val)

	//左节点不为空
	if root.Left != nil {
		ans += "("
		ans += tree2str(root.Left)
		ans += ")"
	} else { //左节点为空，存在右节点 补空括号
		if root.Right != nil {
			ans += tree2str(root.Left)
		}
	}

	if root.Right != nil {
		ans += "("
		ans += tree2str(root.Right)
		ans += ")"
	}
	return ans
}

//合并二叉树
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	var root = &TreeNode{}
	if root1 != nil {
		root.Val += root1.Val
	}
	if root2 != nil {
		root.Val += root2.Val
	}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)

	return root
}

//每层节点平均值
func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64
	var q []*TreeNode
	q = append(q, root)
	//层次遍历
	for len(q) > 0 {
		//当前层次节点数量
		var cnt = len(q)
		sum := 0
		for k := cnt; k > 0; k-- {
			t := q[0]
			q = q[1:]
			sum += t.Val
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}

		}
		ans = append(ans, float64(sum)/float64(cnt))
	}
	return ans
}

func main() {

}
