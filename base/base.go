package main

import (
	"log"
)

func array() {
	arr1 := [3]int{1, 2, 3}
	log.Println(arr1)
	arr2 := [...]int{1, 2}
	log.Println(arr2)
}

func slice() {
	arr1 := [3]int{1, 2, 3}
	// 利用下标的方式取切片，只会创建指向原数组的切片结构体，并不会拷贝，修改切片数据也会改动到原数组
	slice := arr1[0:1]
	log.Println(slice)

	slice1 := []int{4, 5, 6}
	log.Println(slice1)

	slice2 := make([]int, 10)
	slice2[0] = 1
	log.Println(slice2)

	slice1 = append(slice1, 7)
	log.Println(slice1)

	copy(slice2, slice1)
	println(slice1[0])
	println(slice2[0])
}

func hash() {
	hash1 := map[string]int{
		"1": 1,
		"2": 2,
	}
	log.Println(hash1)

	hash2 := make(map[string]int, 2)
	hash2["1"] = 1
	hash2["2"] = 2
	log.Println(hash2)

	delete(hash2, "1")
	log.Println(hash2)
}

func stringDemo() {
	json1 := "{\"a\":1}"
	log.Println(json1)

	json2 := `{"a":1,
"b":2}`
	log.Println(json2)
}

func main() {
	stringDemo()
}
