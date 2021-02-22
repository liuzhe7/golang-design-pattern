/*
   @Author:huolun
   @Date:2021/2/22
   @Description
*/
package main

// pipeline

// 返回一个只读的channel
func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 读取channel的数据，进行平方 放入一个新的channel
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// range会读取完所有数据， 并阻塞到channel关闭
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// 过滤奇数
func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// 求和
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
