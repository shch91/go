package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//比较版本号
func compareVersion(version1 string, version2 string) int {
	var i, j = 0, 0
	for i < len(version1) || j < len(version2) {
		v1, v2 := 0, 0
		for ; i < len(version1) && version1[i] != '.'; i++ {
			v1 = v1*10 + int(version1[i]-'0')
		}
		for ; j < len(version2) && version2[j] != '.'; j++ {
			v2 = v2*10 + int(version2[j]-'0')
		}

		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
		i++
		j++
	}
	return 0
}

//摩尔翻译
var wordIndex = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

//摩尔表示个数
func uniqueMorseRepresentations(words []string) int {
	var key = make(map[string]struct{})

	for _, str := range words {
		result := ""
		for _, ch := range str {
			result += wordIndex[ch-'a']
		}
		if _, ok := key[result]; !ok {
			key[result] = struct{}{}
		}
	}
	return len(key)
}

//字符书写
func numberOfLines(widths []int, s string) []int {
	var line, curLine = 1, 0

	for _, ch := range s {
		if curLine+widths[ch-'a'] > 100 {
			line++
			curLine = widths[ch-'a']
		} else {
			curLine += widths[ch-'a']
		}
	}
	return []int{line, curLine}
}

//不常见单词
func uncommonFromSentences(A string, B string) []string {
	var keyA, keyB = make(map[string]int), make(map[string]int)
	var result []string

	for _, str := range strings.Split(A, " ") {
		keyA[str]++
	}
	for _, str := range strings.Split(B, " ") {
		keyB[str]++
	}

	for key, val := range keyA {
		_, ok := keyB[key]
		if val == 1 && !ok {
			result = append(result, key)
		}
	}

	for key, val := range keyB {
		_, ok := keyA[key]
		if val == 1 && !ok {
			result = append(result, key)
		}
	}
	return result
}

//子域名访问计数
func subdomainVisits(cpdomains []string) []string {
	var key = make(map[string]int)
	var result []string
	for _, str := range cpdomains {
		//单个访问
		visit := strings.Split(str, " ")
		time, _ := strconv.Atoi(visit[0])
		//顶级域名
		com := strings.Split(visit[1], ".")
		var tmp string
		//a.b.c
		for i := len(com) - 1; i >= 0; i-- {
			if i < len(com)-1 {
				tmp = "." + tmp
			}
			tmp = com[i] + tmp
			key[tmp] += time
		}
	}
	for k, val := range key {
		result = append(result, fmt.Sprintf("%d %s", val, k))
	}
	return result
}

//特殊等价字符串
func numSpecialEquivGroups(A []string) int {
	var m = make(map[string]struct{})
	for _, str := range A {
		var cnt [52]int
		for i := 0; i < len(str); i++ {
			cnt[int(str[i])-'a'+26*(i%2)]++
		}
		var str string
		for _, v := range cnt {
			str += strconv.Itoa(v)
		}
		m[str] = struct{}{}
	}
	return len(m)
}

//最常见单词
func mostCommonWord(paragraph string, banned []string) string {
	reg := regexp.MustCompile("[\\s!?',;.]+")

	ban := make(map[string]struct{})
	count := make(map[string]int)
	for _, val := range banned {
		ban[val] = struct{}{}
	}

	for _, str := range reg.Split(paragraph, -1) {
		low := strings.ToLower(str)
		if _, ok := ban[low]; !ok {
			count[low]++
		}
	}
	var ret string
	var max = math.MinInt32
	for key, cnt := range count {
		if cnt > max {
			ret = key
			max = cnt
		}
	}
	return ret
}

//字符串中第二大数字
func secondHighest(s string) int {
	var first, second = -1, -1
	for _, i := range s {
		if i >= '0' && i <= '9' {
			val := int(i - '0')
			if first == -1 {
				first = val
			} else if val > first {
				second = first
				first = val
			} else if val > second && val < first {
				second = val
			}
		}
	}
	return second

}

func main() {

	fmt.Println(secondHighest("abc1111"))
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	widths := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	S := "bbbcccdddaaa"
	fmt.Println(numberOfLines(widths, S))
	fmt.Println(compareVersion("1.00", "1.1"))

}
