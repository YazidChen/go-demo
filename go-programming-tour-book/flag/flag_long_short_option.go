package main

import (
	"flag"
	"log"
)

// > go run flag_long_short_option.go -name=Yazid -n=谛听
// name: 谛听
func main() {
	var name string
	//绑定的参数指针，选项名称，默认值，帮助信息
	flag.StringVar(&name, "name", "默认值", "帮助信息")
	flag.StringVar(&name, "n", "默认值", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)
}
