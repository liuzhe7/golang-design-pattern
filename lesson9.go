/*
   @Author:huolun
   @Date:2021/2/22
   @Description
*/
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

//visitor 模式

//这个模式是一种将算法与操作对象的结构分离的一种方法。这种分离的实际结果是能够在不修改结构的情况下向现有对象结构添加新操作，是遵循开放/封闭原则的一种方法

// 两个实现了shapev接口的结构体， 定义数据结构
type Visitor func(shape Shapev)
type Shapev interface {
	accept(Visitor)
}
type Circle struct {
	Radius int
}

func (c Circle) accept(v Visitor) {
	v(c)
}

type Rectangle struct {
	Width, Heigh int
}

func (r Rectangle) accept(v Visitor) {
	v(r)
}

// 定义两个visitor， 一个是用来做JSON序列化的，另一个是用来做XML序列化的， 定义算法， 这样就结偶了数据结构和算法
func JsonVisitor(shape Shapev) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shapev) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	c := Circle{10}
	r := Rectangle{100, 200}
	shapes := []Shapev{c, r}
	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}
}
