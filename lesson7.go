/*
   @Author:huolun
   @Date:2021/2/22
   @Description
*/
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// go语言的装饰器
func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("start")
		f(s)
		fmt.Println("end")
	}
}

func testFunc(s string) {
	fmt.Println(s)
}

// 计算时间的装饰器
type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		// 传入现在时间， f()结束后 执行defer 计算时间差
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n",
				getFunctionName(f), time.Since(t))
		}(time.Now())

		return f(start, end)
	}
}

func Sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

// 范型修饰器
func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

type MyFoo func(int, int, int) int

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func main() {
	// 先装饰 再掉用
	decorator(testFunc)("nihao")
	timedSumFunc(Sum1)(1, 3000000)

	// 泛型装饰器
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1, 1, 1)
}
