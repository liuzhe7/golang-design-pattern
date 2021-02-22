/*
   @Author:huolun
   @Date:2021/2/19
   @Description
*/
package main

import (
	"fmt"
	"time"
)

// functional options, 通过函数控制对象的可选属性
//直觉式的编程
//高度的可配置化
//很容易维护和扩展
//自文档
//对于新来的人很容易上手
//没有什么令人困惑的事（是nil 还是空）

type server struct {
	Name    string
	Timeout time.Duration
	Tls     string
}

type option func(*server)

func setTimeout(t time.Duration) option {
	return func(s *server) {
		s.Timeout = t
	}
}
func setTls(t string) option {
	return func(s *server) {
		s.Tls = t
	}
}

func NewServer(name string, options ...option) *server {
	server := server{
		Name:    name,
		Timeout: time.Minute * 3,
		Tls:     "",
	}
	for _, option := range options {
		option(&server)
	}
	return &server
}

func main() {
	server := NewServer("aaa", setTimeout(time.Minute), setTls("asdasd"))
	fmt.Println(server)
}
