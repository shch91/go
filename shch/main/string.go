package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
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

//双向
func wordPattern(pattern string, s string) bool {
	if len(pattern) != len(strings.Split(s, " ")) {
		return false
	}
	p2s := map[byte]string{}
	s2p := map[string]byte{}
	str := strings.Split(s, " ")
	for i := range pattern {
		x, y := pattern[i], str[i]
		if p2s[x] != "" && p2s[x] != y || s2p[y] > 0 && s2p[y] != x {
			return false
		}
		p2s[x] = y
		s2p[y] = x
	}
	return true
}

//翻转元音字母
func reverseVowels(s string) string {
	var ret = []byte(s)
	var index []int
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			index = append(index, i)
		}
	}
	for i, j := 0, len(index)-1; i <= j; i, j = i+1, j-1 {
		ret[index[i]], ret[index[j]] = ret[index[j]], ret[index[i]]
	}
	return string(ret)
}

var Vowels = []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func isVowel(b byte) bool {
	for _, ch := range Vowels {
		if ch == b {
			return true
		}
	}
	return false
}

//字符有字符中的magazine组成
func canConstruct(ransomNote string, magazine string) bool {
	var m = make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		val, ok := m[ransomNote[i]]
		if ok && val > 0 {
			m[ransomNote[i]]--
		} else {
			return false
		}
	}
	return true
}

func findTheDifference(s string, t string) byte {
	var m = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	var ch byte
	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; val == 0 || !ok {
			ch = s[i]
			break
		} else {
			m[s[i]]--
		}
	}
	return ch
}

//最长回文字符串
func longestPalindrome(s string) int {
	var m = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	var ans = 0
	for _, v := range m {
		if v%2 == 0 {
			ans += v
		} else if v > 1 {
			ans += v - 1
		}
	}
	if ans < len(s) {
		return ans + 1
	}
	return ans
}

func addStrings(num1 string, num2 string) string {
	ans := ""
	add := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

//bf模式匹配
func bf(s, p string) int {
	var i, j = 0, 0
	for i < len(s) && j < len(p) {
		if s[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == len(p) {
		return i - j
	}
	return -1
}

//字符串匹配
func kmp(query, pattern string) bool {
	n, m := len(query), len(pattern)
	fail := make([]int, m)
	for i := 0; i < m; i++ {
		fail[i] = -1
	}
	for i := 1; i < m; i++ {
		j := fail[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j+1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	match := -1
	for i := 1; i < n-1; i++ {
		for match != -1 && pattern[match+1] != query[i] {
			match = fail[match]
		}
		if pattern[match+1] == query[i] {
			match++
			if match == m-1 {
				return true
			}
		}
	}
	return false
}

func repeatedSubstringPattern(s string) bool {
	return kmp(s+s, s)
}

//岛屿周长
func islandPerimeter(grid [][]int) int {
	var row, col = len(grid), len(grid[0])
	var edge = 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != 1 {
				continue
			}
			//上
			if i-1 < 0 || grid[i-1][j] == 0 {
				edge++
			}
			//右
			if j+1 == col || grid[i][j+1] == 0 {
				edge++
			}
			//下
			if i+1 == row || grid[i+1][j] == 0 {
				edge++
			}
			//左
			if j-1 < 0 || grid[i][j-1] == 0 {
				edge++
			}
		}
	}
	return edge
}

//格式话秘钥
func licenseKeyFormatting(s string, k int) string {

	var con []byte
	for i := 0; i < len(s); i++ {
		if s[i] != '-' {
			con = append(con, s[i])
		}
	}
	var ans []byte

	con = []byte(strings.ToUpper(string(con)))
	if len(con)%k > 0 {
		ans = append(ans, con[0:len(con)%k]...)
		if len(con)%k < len(con) {
			ans = append(ans, '-')
		}
	}
	for i := len(con) % k; i < len(con); i += k {
		ans = append(ans, con[i:i+k]...)
		if i+k < len(con) {
			ans = append(ans, '-')
		}
	}
	return string(ans)
}

//同一行键盘打印字符
func findWords(words []string) []string {
	var one, two, three = "qwertyuiop", "asdfghjkl", "zxcvbnm"
	var oneUp, twoUp, threeUp = strings.ToUpper(one), strings.ToUpper(two), strings.ToUpper(three)

	var ans []string
	for i := 0; i < len(words); i++ {
		//每一个字符串在每行键盘中出现的次数
		var count1, count2, count3 = 0, 0, 0
		for j := 0; j < len(words[i]); j++ {

			if strings.IndexByte(one, words[i][j]) >= 0 || strings.IndexByte(oneUp, words[i][j]) >= 0 {
				count1++
			}
			if strings.IndexByte(two, words[i][j]) >= 0 || strings.IndexByte(twoUp, words[i][j]) >= 0 {
				count2++
			}
			if strings.IndexByte(three, words[i][j]) >= 0 || strings.IndexByte(threeUp, words[i][j]) >= 0 {
				count3++
			}
		}
		if count1 == len(words[i]) || count2 == len(words[i]) || len(words[i]) == count3 {
			ans = append(ans, words[i])
		}
	}
	return ans
}

//结构体排序
type Scu struct {
	Score int
	Index int
}

//相对名次
func findRelativeRanks(score []int) []string {
	var t []Scu
	for i, val := range score {
		t = append(t, Scu{Score: val, Index: i})
	}
	//降序排列
	sort.SliceStable(t, func(i, j int) bool {
		return t[i].Score > t[j].Score
	})

	var ans = make([]string, len(score))
	for i := 0; i < len(t); i++ {
		if i == 0 {
			ans[t[i].Index] = "Gold Medal"
		} else if i == 1 {
			ans[t[i].Index] = "Silver Medal"
		} else if i == 2 {
			ans[t[i].Index] = "Bronze Medal"
		} else {
			ans[t[i].Index] = strconv.Itoa(i + 1)
		}
	}
	return ans
}

//大写检测
func detectCapitalUse(word string) bool {
	var cnt, index = 0, -1
	for i, val := range word {
		if val >= 'A' && val <= 'Z' {
			index = i
			cnt++
		}
	}
	return cnt == 0 || (index == 0 && cnt == 1) || cnt == len(word)
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) >= len(b) {
		return len(a)
	}
	return len(b)
}

//反转前k后续
func reverseStr(s string, k int) string {
	var ans []byte
	var i int
	for i = 0; i+k-1 < len(s); i += 2 * k {
		v := i + k - 1
		for v >= i {
			ans = append(ans, s[v])
			v--
		}
		v = i + k
		for v < i+2*k && v < len(s) {
			ans = append(ans, s[v])
			v++
		}
	}
	//剩下反转
	if i < len(s) {
		t := len(s) - 1
		for t >= i {
			ans = append(ans, s[t])
			t--
		}
	}
	return string(ans)
}

func checkRecord(s string) bool {
	var aCnt, conL = 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			aCnt++
		}
		if s[i] == 'L' {
			conL++
			if conL > 2 || aCnt > 1 {
				return false
			}
		} else {
			conL = 0
		}
	}
	return aCnt <= 1 && conL <= 2
}

func judgeCircle(moves string) bool {
	var left, right, up, down = 0, 0, 0, 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'U' {
			up++
		} else if moves[i] == 'D' {
			down--
		} else if moves[i] == 'L' {
			left++
		} else {
			right--
		}
	}
	return left+right == 0 && up+down == 0
}

//最多删除一个字符,判断是否回文字符串
func validPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		} else {
			return isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1)
		}
	}
	return true
}

func isPalindrome(s string, l, h int) bool {
	for i, j := l, h; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func countBinarySubstrings(s string) int {
	var cnt []int
	var t, p = 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[p] {
			t++
		} else {
			cnt = append(cnt, t)
			t = 1
			p = i
		}
	}
	//最后一个数
	cnt = append(cnt, t)
	var ans = 0
	for i := 1; i < len(cnt); i++ {
		if cnt[i] > cnt[i-1] {
			ans += cnt[i-1]
		} else {
			ans += cnt[i]
		}
	}
	return ans

}

func main() {
	//"PPALLP"
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(judgeCircle("LDRRLRUULR"))
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(strings.IndexByte("ASDFGHJKL", 'A'))

	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(licenseKeyFormatting("2", 2))
	fmt.Println(islandPerimeter([][]int{{1, 0}}))
	repeatedSubstringPattern("abcabcab")
	reverseVowels("hello")

	fmt.Println(wordPattern("abba", "dog cat cat dog"))
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
