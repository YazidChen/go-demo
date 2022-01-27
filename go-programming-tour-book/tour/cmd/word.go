package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"yazidchen.com/go-programming-tour-book/tour/internal/word"
)

const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var desc = strings.Join([]string{
	"该子命令支持的单词转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果：%s", content)
	},
}

// 待转换的字符串
var str string

// 转换模式
var mode int8

func init() {
	// 在 VarP 函数中：
	// p: 需绑定的变量
	// name: 完整命令标志
	// shorthand: 短标志
	// value: 默认值
	// usage: 使用说明
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入待转换字符串")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入转换模式序号")
}
