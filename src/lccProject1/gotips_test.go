package main

import (
	"fmt"
	"math"
	"strings"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
	"regexp"
	"io/ioutil"
	"io"
	"errors"
	"path/filepath"
)

// 1, hello world
func HelloWorld() {
	fmt.Println("Hello,世界")
}

// 2, 变量类型
func typeDemo() {
	// 变量的声明
	var v1 int
	var (
		v2 int
		v3 string
	)
	// var p *int // 指针类型

	// 变量初始化
	var v4 int = 10
	// 等价于：
	var v5 = 10
	// 但是一般这样写 (直接连var 都忽视掉吧）
	v6 := 10
	fmt.Println(v1, v2, v3, v4, v5, v6)

	// 常量
	const PI float64 = 3.141592653
	const MaxPlayer = 10

	// 枚举
	const (
		// iota从0递增
		Sunday  = iota
		Monday
		Tuesday
	)

	// 类型
	var b1 bool
	b1 = true
	b1 = (1 == 2)
	fmt.Println(b1)

	var i32 int32
	i32 = int32(64)
	fmt.Println(i32)

	var f1 float64 = 1.0001
	var f2 float64 = 1.0002
	// 浮点类型的比较
	isEqual := math.Dim(f1, f2) < 0.001
	fmt.Println(isEqual)

	// 字符串
	var s1 string
	s1 = "abc"
	s1 = s1 + "ddd"
	// 获取长度
	n := len(s1)
	// 获取字符
	c1 := s1[0]
	// 反引号，不转义，常用正则表示是
	fmt.Println(n, c1)

	fmt.Println(strings.HasPrefix("prefix", "pre"))
	fmt.Println(strings.HasSuffix("suffix", "fix"))

	// 字节的遍历
	for i := 0; i < n; i++ {
		ch := s1[i]
		fmt.Println(ch)
	}

	// Unicode 字符串的遍历
	for i, ch := range s1 {
		fmt.Println(i, ch)
	}

	// 数组 (名称，类型)
	var arr1 [32]int
	// 二位数组
	// var arr2 [3][8]int
	arr1 = [32]int{0}
	array := [5]int{1, 2, 3, 4, 5}

	// 临时的结构体数组
	structArray := []struct {
		name string
		age  int
	}{{"time", 18}, {"jime", 20}}

	// 数组的遍历的第一种方式
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// 数组遍历的第二种方式
	for i, v := range structArray {
		fmt.Println(i, v)
	}

	// 数组切片Slice
	var mySlice []int = arr1[:2]
	mySlice1 := make([]int, 5)
	mySlice2 := make([]int, 5, 10)
	fmt.Println("len mySlice2", len(mySlice2))
	fmt.Println("cap mySlice2", cap(mySlice2))

	mySlice3 := append(mySlice, 2, 3, 4)
	mySlice4 := append(mySlice, mySlice1...)

	copy(mySlice3, mySlice4)

	// map
	var m map[int]string
	m[1] = "dddd"
	m1 := make(map[int]string)
	m2 := map[int]string{
		1: "a",
		2: "b",
	}

	delete(m2, 1)
	value, ok := m1[1]
	if ok {
		fmt.Println(value)
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}
}

func flowDemo() {
	// if else 
	a := 10
	if a < 10 {

	} else {

	}

	// switch 
	switch a {
	case 0:
		fmt.Println("0")
	case 10:
		fmt.Println("10")
	default:
		fmt.Println("Default")
	}

	switch {
	case a < 10:
		fmt.Println("<10")
	case a < 20:
		fmt.Println("<20")
	}

	// 循环
	for i := 0; i < 10; i++ {

	}

	// 无限循环
	sum := 0;
	for {
		sum ++
		if sum > 10 {

			break
		}
	}

	// 这是什么鬼？
	goto JLoop
JLoop:
}

// 简单的函数
func sum1(value int, value2 int) (result int, err error) {
	return value + value2, nil
}

func sum2(value1, value2 int) int {
	return value1 + value2
}

// 不定参数
func myFunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

// 任意类型的不定参数
func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is int")
		case string:
			fmt.Println(arg, "is string")
		default:
			fmt.Println(arg, "is unknown")
		}
	}
}

// 匿名函数 (这尼玛是什么鬼？)
func anonymousFunc() {
	f := func(a, b int) int {
		return a + b
	}
	f(1, 2)
}

// defer 延迟
func deferDemo(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	defer f.Close()

	// OR
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Runtime error caught: %v", r)
		}
	}()
}

// 结构体
type Rect struct {
	// 小写的private
	x, y float64
	// 大写的public
	Width, Height float64
}

//大写方法为public 小写的为private
func (r * Rect) Area() float64  {
	return r.Width * r.Height
}

func netRect(x, y, width, height float64) *Rect {
	// 实例化结构体
	// rect1 := new(Rect)
	// rect2 := &Rect{}
	// rect3 := &Rect{Width:100, Height:200}
	return &Rect{x, y, width, height}
}

// 2. 读取文件
// import "io/ioutil"
func readFileDemo(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		//Do something
	}
	lines := strings.Split(string(content), "\n")
	fmt.Println("line count:", len(lines))
}

// 判断目录或文件是否存在
func existsPathCheck(path string) (bool, error) {
	// 判断不存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 不存在
	}

	// 判断是否存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// 文件目录操作
func fileDirDemo() {
	// 级联创建目录
	os.MkdirAll("/path/to/create", 0777)
}

// 拷贝文件
func copyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}

	}
	return
}

// 拷贝目录
func copyDir(source string, dest string) (err error) {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return errors.New(source + " is not a directory")
	}

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(source)
	for _, entry := range entries {
		sfp := filepath.Join(source, entry.Name())
		dfp := filepath.Join(dest, entry.Name())
		if entry.IsDir() {
			err = copyDir(sfp, dfp)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copyFile(sfp, dfp)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return nil
}

// 3. 时间处理
// import "time"
func TestTimeDemo(t *testing.T) {
	// Parse
	postDate, err := time.Parse("2006-01-02 15:04:05", "2015-09-30 19:19:00")
	fmt.Println(postDate, err)

	// Format
	assert.Equal(t, "2015/Sep/30 07:19:00", postDate.Format("2006/Jan/02 03:04:05"))
	assert.Equal(t, "2015-09-30T19:19:00Z", postDate.Format(time.RFC3339))
}

// 4. 正则表达式
// import "regexp"
func TestRegexp(t *testing.T) {
	// 查找匹配
	re := regexp.MustCompile(`(\d+)-(\d+)`)
	r := re.FindAllStringSubmatch("123-666", -1)

	assert.Equal(t, 1, len(r))
	assert.Equal(t, "123", r[0][1])
	assert.Equal(t, "666", r[0][2])

}

func main() {
	//helloWorld()
}

// 来源 https://www.cnblogs.com/sevenyuan/p/6564575.html

