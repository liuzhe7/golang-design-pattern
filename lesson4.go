/*
   @Author:huolun
   @Date:2021/2/19
   @Description
*/
package main

import (
	"errors"
	"fmt"
)

type Label struct {
	Text string
}

type Box struct {
	Label
	Name string
}

type Pointer interface {
	Point()
}

func (label Label) Point() {
	fmt.Println(label.Text)
}

// 反转控制
//type IntSet struct {
//	data map[int]bool
//}
//func NewIntSet() IntSet {
//	return IntSet{make(map[int]bool)}
//}
//func (set *IntSet) Add(x int) {
//	set.data[x] = true
//}
//func (set *IntSet) Delete(x int) {
//	delete(set.data, x)
//}
//func (set *IntSet) Contains(x int) bool {
//	return set.data[x]
//}

// 通过包装intset 给他添加一个撤销操作的记录器 就是这个函数切片，但是，这种方式最大的问题是，Undo操作其实是一种控制逻辑，并不是业务逻辑，所以，在复用 Undo这个功能上是有问题。因为其中加入了大量跟 IntSet 相关的业务逻辑。
//type UndoableIntSet struct { // Poor style
//	IntSet    // Embedding (delegation)
//	functions []func()
//}
//
//func NewUndoableIntSet() UndoableIntSet {
//	return UndoableIntSet{NewIntSet(), nil}
//}
//
//
//func (set *UndoableIntSet) Add(x int) { // Override
//	if !set.Contains(x) {
//		set.data[x] = true
//		set.functions = append(set.functions, func() { set.Delete(x) })
//	} else {
//		set.functions = append(set.functions, nil)
//	}
//}
//
//func (set *UndoableIntSet) Delete(x int) { // Override
//	if set.Contains(x) {
//		delete(set.data, x)
//		set.functions = append(set.functions, func() { set.Add(x) })
//	} else {
//		set.functions = append(set.functions, nil)
//	}
//}
//
//func (set *UndoableIntSet) Undo() error {
//	if len(set.functions) == 0 {
//		return errors.New("没有可撤销的操作")
//	}
//	index := len(set.functions) - 1
//	if function := set.functions[index]; function != nil {
//		function()
//		set.functions[index] = nil // For garbage collection
//	}
//	set.functions = set.functions[:index]
//	return nil
//}

// 反转依赖
type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

// 直接给inset加入undo类型 ，通过增加删除undo 记录撤销操作， 吊用undo 进行撤销， 这样的undo 与业务逻辑无关，可以复用
type IntSet struct {
	data map[int]bool
	undo Undo
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

func main() {

}
