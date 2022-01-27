package word

import (
	"strings"
	"unicode"
)

// ToUpper 全部转为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 全部转为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线转大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	// 先将下划线替换为空格
	s = strings.Replace(s, "_", " ", -1)
	// 将每个单词首字母大写
	s = strings.Title(s)
	// 移除空格
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线转小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	// 得到大写驼峰
	s = UnderscoreToUpperCamelCase(s)
	// 将首字母转为小写
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线
// 遍历所有字符，如果是大写，则前面加下划线，并且字符转为小写
func CamelCaseToUnderscore(s string) string {
	var out []rune
	for i, r := range s {
		if i == 0 {
			out = append(out, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(r))
	}
	return string(out)
}
