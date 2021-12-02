package main

import (
	"github.com/elizarpif/english-words-maker/internal/jsonwords"
	"github.com/elizarpif/english-words-maker/internal/maker"
	"log"
)

func main() {
	files()

	//err := maker.StoryTeller()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err := maker.Activity()
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func files() {
	words, err := jsonwords.GetWords()
	if err != nil {
		log.Fatal(err)
	}

	err = maker.Vocabulary(words)
	if err != nil {
		log.Fatal(err)
	}

	err = maker.CreateLotoFiles(words, 4)
	if err != nil {
		log.Fatal(err)
	}
}

