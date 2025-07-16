package main

import (
	"fmt"

	"github.com/DataVizWiz/data-with-go/cmd/hello"
	"github.com/DataVizWiz/data-with-go/cmd/words"
)

func main() {
	sentence := "hello metin my name is metin"
	counts, _ := words.CountWords(sentence)

	fmt.Print(counts)

	hello.SayHello()
}
