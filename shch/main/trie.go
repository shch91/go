package main

import (
	"fmt"
	"shch/trie"
)

//trie树节点
type TrieNode struct {
	Char     uint8
	Children map[uint8]*TrieNode
	End      int //构建字典树数组的位置+1
}

//新建节点
func NewTrie(ch uint8) *TrieNode {
	root := &TrieNode{Char: ch, Children: make(map[uint8]*TrieNode)}
	return root
}

func insert(root *TrieNode, word string, index int) {
	cur := root
	for i := 0; i < len(word); i++ {
		_, ok := cur.Children[word[i]]
		if !ok {
			cur.Children[word[i]] = NewTrie(word[i])
		}
		cur = cur.Children[word[i]]
	}
	cur.End = index
}

func dfsTrie(root *TrieNode, words []string) string {

	ans := ""
	var stack []*TrieNode
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		if cur.End > 0 || cur == root {
			if cur != root {
				word := words[cur.End-1]
				if len(word) > len(ans) ||
					len(word) == len(ans) && word < ans {
					ans = word
				}
			}
			//子节点
			for _, child := range cur.Children {
				stack = append(stack, child)
			}

		}
	}
	return ans
}

func longestWord(words []string) string {
	var trie = NewTrie('/')
	index := 0
	for _, str := range words {
		index++
		insert(trie, str, index)
	}
	return dfsTrie(trie, words)
}

func main() {
	fmt.Println("apple" < "apply")
	//longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})

	trie := trie.NewTrie()
	words := []string{"Golang", "学院君", "Language", "Trie", "Go"}
	// 构建 Trie 树
	for _, word := range words {
		trie.Insert(word)
	}
	// 从 Trie 树中查找字符串
	term := "学院君"
	if trie.Find(term) {
		fmt.Printf("包含单词\"%s\"\n", term)
	} else {
		fmt.Printf("不包含单词\"%s\"\n", term)
	}

}
