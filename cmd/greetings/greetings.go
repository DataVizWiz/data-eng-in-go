package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func Hello(name string, lang string) (string, error) {
	if len(name) < 1 {
		return "", errors.New("name cannot be empty")
	}

	var message string

	switch strings.ToLower(lang) {
	case "es":
		message = fmt.Sprintf("¡Hola!, %v. ¡bienvenido!", name)
	case "fr":
		message = fmt.Sprintf("Salut !, %v. bienvenue !", name)
	default:
		message = fmt.Sprintf(randomGreeting(), name)
	}
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name, "en")
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomGreeting() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
