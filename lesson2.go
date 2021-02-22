/*
   @Author:huolun
   @Date:2021/2/18
   @Description
*/
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

//Error Check  Hell
type A struct {
	Name string
	err  error
}

type ReadAble interface {
	read(name string)
}

// 发生错误就把错误标记到对象内部,连续调用后只需要判断一次错误
func (a *A) read(name string) {
	if a.Name == name {
		return
	} else {
		a.err = errors.New("sdfsdfsfsdfsfsdfsd")
	}
}

func Printerr(r ReadAble, name string) {
	r.read(name)
}

// 错误包装
//我们需要包装一下错误，而不是干巴巴地把err给返回到上层，我们需要把一些执行的上下文加入
//if err != nil {
//return fmt.Errorf("something failed: %v", err)
//}

func main() {
	a := A{
		"aba",
		nil,
	}
	Printerr(&a, "bbb")
	Printerr(&a, "aba")
	if a.err != nil {
		// 错误包装
		a.err = errors.Wrap(a.err, "read,faild")
	}
	// cause 解开包装
	switch errors.Cause(a.err).(type) {
	default:
		fmt.Println(a.err)               // 输出原始错误
		fmt.Println(errors.Cause(a.err)) // 输出解开包装对错误
	}
}
