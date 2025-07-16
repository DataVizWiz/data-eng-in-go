package hello

import (
	"fmt"
	"log"

	"github.com/DataVizWiz/data-with-go/cmd/greetings"
)

func SayHello() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	const lang = "en"

	message, err := greetings.Hello("Metin", lang)

	if err != nil {
		var message, input string
		fmt.Print("Enter your name: ")
		fmt.Scanf("%s", &input)
		message, _ = greetings.Hello(input, lang)
		fmt.Println(message)
	}

	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range messages {
		fmt.Println(msg)
	}
}
