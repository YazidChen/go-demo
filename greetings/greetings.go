package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi,%v,Welcome!", name)
	return message, nil
}

func RandomHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// make(map[key-type]value-type)
	messages := make(map[string]string)
	// _ 空白标识符，range返回索引值和数据值，此处不需要索引值，用_替代
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi,%v,Welcome",
		"Great to see you,%v!",
		"Hail,%v! Well met",
	}
	return formats[rand.Intn(len(formats))]
}
