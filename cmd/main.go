package main

import (
	"fmt"
	"github.com/elizarpif/english-words-maker/internal/json_to_csv"
	"github.com/elizarpif/english-words-maker/internal/maker"
	"math/rand"
)

func main() {
	// json()

	c := json_to_csv.NewCsvMaker()
	words := c.GetWords()

	mapW := maker.GetMap(words)

	for len(mapW) > 0 {
		key := words[rand.Intn(len(mapW))][0]
		//word := mapW[key]

		fmt.Println("generated ", key)
		var a string
		fmt.Scan(&a)
		delete(mapW, key)
	}
	//err = maker.CreateGames(words, 5)
	//if err != nil {
	//	panic(err)
	//}
}

func json() {
	c := json_to_csv.NewCsvMaker()
	words := c.GetWords()

	err := maker.CreateGames(words, 5)
	if err != nil {
		panic(err)
	}
}

func makeFiles() error {
	m := maker.NewMaker(1, true)
	defer m.Close()

	// m.AllFromToeflVocabulary()

	err := m.RandFromToeflVocabulary(3)
	if err != nil {
		return err
	}

	return nil
}