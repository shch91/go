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

//
func halvesAreAlike(s string) bool {
	var l, t = len(s), 0

	s = strings.ToLower(s)
	for i := 0; i < l; i++ {
		var is = aeiou(s[i])
		if i < 2/t && is {
			t++
		} else if is {
			t--
		}
	}
	return t == 0
}
func aeiou(a uint8) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u'
}

//旋转字符串
func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}

//单调递增的数字
func monotoneIncreasingDigits(N int) int {

	var strN = []byte(strconv.Itoa(N))
	var i = 1
	//找到递增的位置
	for i < len(strN) && strN[i] >= strN[i-1] {
		i++
	}
	if i < len(strN) {
		//i-1严格大于
		for i > 0 && strN[i] < strN[i-1] {
			strN[i-1] -= 1
			i--
		}
		for i++; i < len(strN); i++ {
			strN[i] = '0' + 9
		}
	}

	val, _ := strconv.Atoi(string(strN))
	return val
}

//字符最短距离
func shortestToChar(s string, c byte) []int {
	var ans = make([]int, len(s))
	var prev = math.MinInt32
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			prev = i
		}
		ans[i] = i - prev
	}
	prev = math.MaxInt32
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		if prev-i < ans[i] {
			ans[i] = prev - i
		}

	}
	return ans
}

//转换小写字母
func toLowerCase(str string) string {
	var ch = []byte(str)
	for i := 0; i < len(ch); i++ {
		if ch[i] >= 'A' && ch[i] <= 'Z' {
			ch[i] += 32
		}
	}
	return string(ch)
}

//同构字符串
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

func main() {

	fmt.Println(monotoneIncreasingDigits(332))
	fmt.Println(secondHighest("abc1111"))
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	widths := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	S := "bbbcccdddaaa"
	fmt.Println(numberOfLines(widths, S))
	fmt.Println(compareVersion("1.00", "1.1"))

}
