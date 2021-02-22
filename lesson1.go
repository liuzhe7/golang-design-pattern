/*
   @Author:huolun
   @Date:2021/2/18
   @Description
*/
package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func sliceTest() {
	// demo1
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100
	// 切片操作的cap 会保留到最后一个
	bar := foo[1:4]
	bar[1] = 99
	// foo[2] 变为了99
	fmt.Println(foo)

	// demo2
	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)
	a[2] = 42
	// a重新分配了内存  所以b不变， 注意：如果append操作时 cap 够用 则不会重新分配内存
	fmt.Println(b)

	// demo3
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB

}

func deepEqual() {
	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true
}

// Program to an interface not an implementation
//我们使用了一个叫Stringable 的接口，我们用这个接口把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦了。于是，只要实现了Stringable 接口，都可以传给 PrintStr() 来使用。
type country struct {
	Name string
}
type city struct {
	Name string
}
type StringAble interface {
	Tostring() string
}

func (c country) Tostring() string {
	return c.Name
}
func (c city) Tostring() string {
	return c.Name
}
func Printstr(p StringAble) {
	fmt.Println(p.Tostring())
}

// 检查接口完整性
type Shape interface {
	Sides() int
	Area() int
}
type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}

//时间
//在 Go 语言中，你一定要使用 time.Time 和 time.Duration 两个类型：
//
//在命令行上，flag 通过 time.ParseDuration 支持了 time.Duration
//JSon 中的 encoding/json 中也可以把time.Time 编码成 RFC 3339 的格式
//数据库使用的 database/sql 也支持把 DATATIME 或 TIMESTAMP 类型转成 time.Time
//YAML你可以使用 gopkg.in/yaml.v2 也支持 time.Time 、time.Duration 和 RFC 3339 格式
//如果你要和第三方交互，实在没有办法，也请使用 RFC 3339 的格式。
//
//最后，如果你要做全球化跨时区的应用，你一定要把所有服务器和时间全部使用UTC时间。

// 性能
//Go 语言是一个高性能的语言，但并不是说这样我们就不用关心性能了，我们还是需要关心的。下面是一个在编程方面和性能相关的提示。
//
//如果需要把数字转字符串，使用 strconv.Itoa() 会比 fmt.Sprintf() 要快一倍左右
//尽可能地避免把String转成[]Byte 。这个转换会导致性能下降。
//如果在for-loop里对某个slice 使用 append()请先把 slice的容量很扩充到位，这样可以避免内存重新分享以及系统自动按2的N次方幂进行扩展但又用不到，从而浪费内存。
//使用StringBuffer 或是StringBuild 来拼接字符串，会比使用 + 或 += 性能高三到四个数量级。
//尽可能的使用并发的 go routine，然后使用 sync.WaitGroup 来同步分片操作
//避免在热代码中进行内存分配，这样会导致gc很忙。尽可能的使用 sync.Pool 来重用对象。
//使用 lock-free的操作，避免使用 mutex，尽可能使用 sync/Atomic包。 （关于无锁编程的相关话题，可参看《无锁队列实现》或《无锁Hashmap实现》）
//使用 I/O缓冲，I/O是个非常非常慢的操作，使用 bufio.NewWrite() 和 bufio.NewReader() 可以带来更高的性能。
//对于在for-loop里的固定的正则表达式，一定要使用 regexp.Compile() 编译正则表达式。性能会得升两个数量级。
//如果你需要更高性能的协议，你要考虑使用 protobuf 或 msgp 而不是JSON，因为JSON的序列化和反序列化里使用了反射。
//你在使用map的时候，使用整型的key会比字符串的要快，因为整型比较比字符串比较要快。

func main() {
	sliceTest()
	deepEqual()
	c1 := country{
		Name: "china",
	}
	c2 := city{
		Name: "hangzhou",
	}
	Printstr(c1)
	Printstr(c2)

	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
	//var _ Shape = (*Square)(nil) 类型转换 如果没有实现全部方法会报错
}
