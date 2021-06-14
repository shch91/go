package main

import (
	"errors"
	"fmt"
	"math"
	"shch/main/util"
	"strconv"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return math.Sqrt(f), nil
}

func Fact(k int) int {
	if k == 1 {
		return 1
	}
	return k * Fact(k-1)
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//找出数组中只出现一次的数
func singleNumber(nums []int) int {
	var seenOnce, seenTwice = 0, 0
	for _, v := range nums {
		seenOnce = ^seenTwice & (seenOnce ^ v)
		seenTwice = ^seenOnce & (seenTwice ^ v)
	}
	return seenOnce
}

//自除数
func selfDividingNumbers(left int, right int) []int {
	var ret []int
	for k := left; k <= right; k++ {
		c := k
		for c > 0 {
			t := c % 10
			if t == 0 {
				break
			}
			if k%t != 0 {
				break
			}
			c /= 10
		}
		if c == 0 {
			ret = append(ret, k)
		}
	}
	return ret
}

//判定二的幂函数
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}

//计算二进制中1的个数
func countOne(val int) int {
	var res = 0
	for val != 0 {
		val = val & (val - 1)
		res++
	}
	return res
}

//蚂蚁掉落时刻
func getLastMoment(n int, left []int, right []int) int {
	var max = math.MinInt32
	for _, val := range left {
		if val > max {
			max = val
		}
	}
	for _, val := range right {
		if n-val > max {
			max = n - val
		}
	}
	return max
}

//基数数量
func countOdds(low int, high int) int {

	if low%2 == 0 {
		low += 1
	}
	if high%2 == 0 {
		high -= 1
	}

	return (high-low)/2 + 1
}

//12345678910 找到第n位数字
func findNthDigit(n int) int {
	//当前位数的最高值，最低值，以及值的宽度
	var low, high uint64 = 1, 9
	var width, mul uint64 = 1, 10
	//位置
	var pos uint64 = 0
	for {
		pos += (high - low + 1) * width
		if pos > uint64(n) {
			break
		}
		low = high + 1
		high = low*mul - 1
		width++
	}
	//退回位数
	//往回退的步数
	back := pos - uint64(n)
	remain := back % width
	high -= back / width
	for remain != 0 {
		high /= 10
		remain--
	}
	//取当前末尾数字
	return int(high % 10)
}

//排列硬币
func arrangeCoins(n int) int {
	var row = 0
	for n > row {
		row++
		n -= row
	}
	return row
}

//移除k个元素剩余最小元素
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var stack = []byte{num[0]}

	for i := 1; i < len(num); i++ {
		l := len(stack)
		if l < 1 { //空
			stack = append(stack, num[i])
		} else {
			//大于当前元素的弹出
			for l > 0 && k > 0 && num[i] < stack[l-1] {
				stack = stack[:l-1]
				k--
				l--
			}
			stack = append(stack, num[i])
		}
	}
	for k > 0 {
		k--
		l := len(stack)
		stack = stack[:l-1]
	}
	//去掉0
	for i := 0; i < len(stack); i++ {
		if stack[i] != '0' {
			return string(stack[i:])
		}
	}
	return "0"
}

//各位数字相加只剩一位
func addDigits(num int) int {
	for num > 9 {
		val := 0
		for num != 0 {
			val += num % 10
			num /= 10
		}
		num = val
	}
	return num
}

//卡特兰数
func numTrees(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}
	return C
}

//最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//分式化简
func fraction(cont []int) []int {
	var n, m = 1, cont[len(cont)-1]
	for i := len(cont) - 2; i >= 0; i-- {
		n += m * cont[i]
		//交换n,m
		n, m = m, n
	}
	g := gcd(m, n)
	return []int{m / g, n / g}
}

//最小那币次数
func minCount(coins []int) int {
	var ret = 0
	for i := 0; i < len(coins); i++ {
		if coins[i]%2 == 0 {
			ret += coins[i] / 2
		} else {
			ret += coins[i]/2 + 1
		}
	}
	return ret
}

//判断是否丑数
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}
	return true
}

//十六进制
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var ch = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	var ans []byte
	t := uint32(num)
	for t != 0 {
		tmp := t & 15
		ans = append(ans, ch[tmp])
		t >>= 4
	}
	return revStr(string(ans))
}

//7进制
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var t = int(math.Abs(float64(num)))
	var ans []byte
	for t != 0 {
		ans = append(ans, byte(t%7)+'0')
		t /= 7
	}
	str := revStr(string(ans))
	if num < 0 {
		return "-" + str
	}
	return str
}

func revStr(str string) string {
	var res []byte
	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, str[i])
	}
	return string(res)
}

func minMoves(nums []int) int {
	m := util.Min(nums...)
	var ans int
	for _, val := range nums {
		ans += val - m
	}
	return ans
}

//构造矩形面积
func constructRectangle(area int) []int {
	var w = 1
	var l = area / w
	w = int(math.Sqrt(float64(area)))
	for w >= 1 {
		if area%w == 0 {
			return []int{area / w, w}
		}
		w -= 1
	}
	return []int{l, w}
}

//完美数
func checkPerfectNumber(num int) bool {
	var s = 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			s += i
			if i*i != num {
				s += num / i
			}
		}
	}
	return s-num == num
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func countPrimeSetBits(left int, right int) int {
	var ans = 0
	for i := left; i <= right; i++ {
		cnt := 0
		for t := i; t > 0; t >>= 1 {
			if t&1 > 0 {
				cnt++
			}
		}
		if isPrime(cnt) {
			ans++
		}
	}
	return ans
}

//质数
//质数还有一个特点，就是它总是等于 6x-1 或者 6x+1，其中 x 是大于等于1的自然数。
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	//与6的倍数不相邻
	if n%6 != 1 && n%6 != 5 {
		return false
	}
	for i := 5; float64(i) <= math.Floor(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func numPrimeArrangements(n int) int {
	var primeCnt = 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeCnt++
		}
	}
	return factorial(n-primeCnt) * factorial(primeCnt)
}

//求阶乘
func factorial(n int) int {
	var result = 1
	for i := 2; i <= n; i++ {
		result *= i
		result %= 1e9 + 7
	}
	return result
}

//旋转数字
func rotatedDigits(n int) int {
	var cnt = 0
	for i := 1; i <= n; i++ {
		if isNice(i) {
			cnt++
		}
	}
	return cnt
}

func isNice(v int) bool {
	var m = map[uint8]int{
		0: 0,
		1: 1,
		2: 5,
		5: 2,
		6: 9,
		8: 8,
		9: 6,
	}
	var str = strconv.Itoa(v)
	var res = 0
	for i := 0; i < len(str); i++ {
		ch := str[i] - '0'
		r, ok := m[ch]
		if !ok {
			return false
		}
		res = res*10 + r
	}
	return res != v
}

func largestTriangleArea(points [][]int) float64 {
	var ans = 0.0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				t := area(points[i], points[j], points[k])
				if t > ans {
					ans=t
				}
			}
		}
	}
	return ans
}

//已知三点坐标（a，b）、（c，d）、（e，f），求三角形面积
func area(P, Q, R []int) float64 {
	a := float64(P[0]*Q[1] + Q[0]*R[1] + R[0]*P[1] - P[1]*Q[0] - Q[1]*R[0] - R[1]*P[0])
	return 0.5 * math.Abs(a)
}

func main() {

	isNice(255)
	fmt.Println(countPrimeSetBits(244, 269))
	fmt.Println(convertToBase7(-7))
	fmt.Println(constructRectangle(4))
	fmt.Println(minMoves([]int{1, 1, 1000}))
	fmt.Println(toHex(16))
	fmt.Println(isUgly(-2147483648))
	fmt.Printf("val=%v \n", fraction([]int{3, 2, 0, 2}))
	fmt.Println(removeKdigits("112", 1))
	fmt.Println(arrangeCoins(5))
	fmt.Println(findNthDigit(1))
	ret := selfDividingNumbers(1, 22)
	fmt.Println(ret)
	fmt.Println(singleNumber([]int{-1, -1, -1, -2}))
	var pos = 1
	fmt.Println(^pos)
	fmt.Println(Fact(10))
}
