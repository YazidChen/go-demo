package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

// String 实现了flag的Value String接口
func (n *Name) String() string {
	return fmt.Sprint(*n)
}

// Set 实现了flag的Value Set接口
func (n *Name) Set(value string) error {
	if len(*n) > 0 {
		return errors.New("name flag already set")
	}
	*n = Name("输入信息：" + value)
	return nil
}

// > go run flag_value_set.go -name=Yazid
// name: 输入信息：Yazid
func main() {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)
}
