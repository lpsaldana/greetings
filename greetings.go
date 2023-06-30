package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("nombre vacío")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	if names == nil {
		return nil, errors.New("no hay nombres")
	}
	greets := make(map[string]string)
	for _, name := range names {
		greet, err := Hello(name)
		if err != nil {
			return nil, err
		}
		greets[name] = greet
	}
	return greets, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
