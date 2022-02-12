//+build wireinject

package wire

import (
	"github.com/google/wire"
	w "yazidchen.com/google-wire/internal"
)

//go get github.com/google/wire/cmd/wire
func InitializeEvent(phrase string) (w.Event, error) {
	wire.Build(w.NewEvent, w.NewGreeter, w.NewMessage)
	return w.Event{}, nil
}
