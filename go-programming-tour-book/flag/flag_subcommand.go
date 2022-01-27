package main

import (
	"flag"
	"log"
)

var name string

// > go run flag_subcommand.go go -name=Yazid
// name: Yazid
// > go run flag_subcommand.go java -n=Yazid
// name: Yazid
// > go run flag_subcommand.go java -name=Yazid
// flag provided but not defined: -name
// Usage of java:
//   -n string
//         帮助信息 (default "Java 语言")
// exit status 2
func main() {
	// 解析命令行
	flag.Parse()
	// 获取命令行参数
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	// 判断第一个参数
	switch args[0] {
	case "go":
		// 创建了新的命令集支持子命令，指定命令参数名称，及错误处理方式
		// // 返回错误描述
		// ContinueOnError ErrorHandling = iota
		// // 调用 os.Exit(2) 退出程序
		// ExitOnError
		// // 调用 painc 语句抛出错误异常
		// PanicOnError
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "java":
		javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
		javaCmd.StringVar(&name, "n", "Java 语言", "帮助信息")
		_ = javaCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}
