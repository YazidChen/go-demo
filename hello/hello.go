package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"log"
	"rsc.io/quote"
	"yazid.com/greetings"
	"yazid.com/hello/morestrings"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	fmt.Println("Hello World!")
	fmt.Println(quote.Go())

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	names := []string{"Yazid", "Toby", "Tom"}
	message3, err3 := greetings.Hellos(names)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(message3)

	message, nil := greetings.Hello("Yazid")
	fmt.Println(message)

	message2, err2 := greetings.RandomHello("Yazid")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(message2)

	message1, err1 := greetings.Hello("")
	if err1 != nil {
		log.Fatal(err1) //后面代码不执行
	}
	fmt.Println(message1)

}
