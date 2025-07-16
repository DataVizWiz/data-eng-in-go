package words

import (
	"errors"
	"strings"
)

func CountWords(sentence string) (map[string]int, error) {
	if len(sentence) < 1 {
		return make(map[string]int),
			errors.New("name cannot be empty")
	}

	words := strings.Fields(sentence)
	counts := make(map[string]int)

	for _, word := range words {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts, nil
}
