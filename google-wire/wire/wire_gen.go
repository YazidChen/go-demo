// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"yazidchen.com/google-wire/internal"
)

// Injectors from wire.go:

//go get github.com/google/wire/cmd/wire
func InitializeEvent(phrase string) (internal.Event, error) {
	message := internal.NewMessage(phrase)
	greeter := internal.NewGreeter(message)
	event, err := internal.NewEvent(greeter)
	if err != nil {
		return internal.Event{}, err
	}
	return event, nil
}
