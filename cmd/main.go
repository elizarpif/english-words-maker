package main

import (
	"github.com/elizarpif/english-words-maker/internal/json_to_csv"
	"github.com/elizarpif/english-words-maker/internal/maker"
)

func main() {
	// json()
	make()
}

func json() {
	c := json_to_csv.NewCsvMaker()
	c.Do()
}

func make() {
	m := maker.NewMaker(1, true)

	m.AllFromToeflVocabulary()
	m.RandFromToeflVocabulary(3)

	m.Close()
}