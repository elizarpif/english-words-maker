package main

import "github.com/elizarpif/english-words-maker/internal/maker"

func main() {
	m := maker.NewMaker(1)

	m.AllFromToeflVocabulary()
	m.RandFromToeflVocabulary(3)

	m.Close()
}
