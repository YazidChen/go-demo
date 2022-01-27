package main

import (
	"log"
	"yazidchen.com/go-programming-tour-book/tour/cmd"
)

// > go run main.go word -s=hello -m=1
// 输出结果：HELLO
// > go run main.go word -s=HELLO -m=2
// 输出结果：hello
// > go run main.go word -s=hello_world -m=3
// 输出结果：HelloWorld
// > go run main.go word -s=hello_world -m=4
// 输出结果：helloWorld
// > go run main.go word -s=HelloWorld -m=5
// 输出结果：hello_world
//
// > go run main.go time now
// 输出结果： 2022-01-27 17:10:57, 1643274657, 2022-01-27T17:10:57+08:00
// > go run main.go time calc -c="2022-01-27 17:11:59" -d=22m
// 输出结果：2022-01-27 17:33:59, 1643304839
// > go run main.go time calc -c="2022-01-27 17:11:59" -d=-2h30m
// 输出结果：2022-01-27 14:41:59, 1643294519
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
