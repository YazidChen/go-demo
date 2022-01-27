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
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
