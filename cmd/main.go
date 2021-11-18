package main

import (
	"bufio"
	"fmt"
	"github.com/elizarpif/english-words-maker/internal/jsonwords"
	"github.com/elizarpif/english-words-maker/internal/maker"
	"log"
	"os"
)

func main() {
	words, err := jsonwords.GetWords()
	if err != nil {
		log.Fatal(err)
	}

	err = maker.Vocabulary(words)
	if err != nil {
		log.Fatal(err)
	}

	err = maker.CreateGames(words, 3)
	if err != nil {
		log.Fatal(err)
	}
}

// Activity randomises words without duplicates
func Activity() {
	words, err := jsonwords.GetWords()
	if err != nil {
		log.Fatal(err)
	}

	mapW := maker.GetMap(words)

	sc := bufio.NewScanner(os.Stdin)
	for key, value := range mapW {
		fmt.Printf("generated %s - %s\n", key, value)
		sc.Scan()
	}
}
