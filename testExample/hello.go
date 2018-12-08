package main

import (
	"fmt"
	"testing"
	"time"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func TestReverse(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func forTest() {
	// 1.
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
	}
	fmt.Println(a)

	// 2.条件语句
	b := 1
	for b <= 3 {
		b++
	}
	fmt.Println(b)

	//3.
	c := 1
	for i := 0; i < 3; i++ {
		c++
	}
	fmt.Println(c)

}

func switchTest() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("Nothing")
	}
	fmt.Println(a)


	b := 1
	switch {
	case b >= 0:
		fmt.Println("b=0")
		fallthrough
	case b >= 1:
		fmt.Println("b=1")
	}
	fmt.Println(b)


	switch a := 1; {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	}

}

func breakTest() {
	// 标签
	LABLE:
		for {
			for i := 0; i < 10; i++ {
				if i > 3 {
					break LABLE
			}
			}
			}
	fmt.Println("ok")

LABLE1:
	for i := 0; i < 10; i++ {
		for {
			continue LABLE1
			fmt.Printf("i:%v\n",i)
		}
	}

	fmt.Println("ok")
}
func varTest() {
	var a int = 65
	b := string(a)
	fmt.Println(b)
	fmt.Println(a)


	// 初始化变量，在块里边初始化，作用域只是在该区块中
	if a, b, c := 1, 2, 3; a + b + c > 6 {
		fmt.Println("大于6")
		fmt.Println(a)

	} else {
		fmt.Println("小于等于6")
		fmt.Println(a)

	}
}

func rangeTest() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}


	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func typeTest() {
	type MyTime = time.Time

	var t MyTime = time.Now()
	fmt.Println(t)


	type D = int
	type I int

	v := 100
	var d D = v
	// var i I = v // error

	fmt.Printf("d: %v, v: %v\n", d, v)


	//
	var v1 interface{}
	var d1 D = 100
	v = d

	switch i := v1.(type) {
	case int:
		fmt.Println("it is an int:", i)
		fmt.Println("d1:", d1)
		// case D:
		// 	fmt.Println("it is D type:", i)
	}

	// 类型循环

}


func main() {
	fmt.Println(Reverse("!selpmaxe oG ,olleH"))

	fmt.Println("7.0/3.0 =", 7.0/3.0)

	typeTest()
}