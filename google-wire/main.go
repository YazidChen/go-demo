package main

import (
	"fmt"
	"os"
	"yazidchen.com/google-wire/wire"
)

/**
func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}
*/

func main() {
	e, err := wire.InitializeEvent("Init Message!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
