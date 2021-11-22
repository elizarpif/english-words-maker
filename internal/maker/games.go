package maker

import (
	"bufio"
	"fmt"
	"github.com/elizarpif/english-words-maker/internal/jsonwords"
	"os"
	"strconv"
	"strings"
)

func createRusEngGames(words [][]string, players int) error {
	for i := 0; i < players; i++ {
		records := shuflleRecords(words)

		var values []string
		var valuesEng []string

		// len(records) - 5 - чтобы покрыть большее число слов
		for j := 0; j < len(records)-5; j++ {
			values = append(values, fmt.Sprintf("%s%s", randSpace(), records[j][1]))
			valuesEng = append(valuesEng, fmt.Sprintf("%s%s", randSpace(), records[j][0]))
		}

		err := writeToFile(values, templateRandRusResult+strconv.Itoa(i+1)+txtExtension)
		if err != nil {
			return err
		}

		err = writeToFile(valuesEng, templateRandEndResult+strconv.Itoa(i+1)+txtExtension)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateLotoFiles creates files for players
func CreateLotoFiles(words [][]string, players int) error {
	for i := 0; i < players; i++ {
		records := shuflleRecords(words)

		var valuesEng []string

		// len(records) - 5 - чтобы покрыть большее число слов
		for j := 0; j < len(records)-5; j++ {
			valuesEng = append(valuesEng, fmt.Sprintf("%s%s", randSpace(), records[j][0]))
		}

		err := writeToFile(valuesEng, templateRandEndResult+strconv.Itoa(i+1)+txtExtension)
		if err != nil {
			return err
		}
	}

	return nil
}

const maxWordsForStoryTeller = 4

// StoryTeller randomises 4 words from list of words
func StoryTeller(peopleCount int) error {
	words, err := jsonwords.GetWords()
	if err != nil {
		return err
	}

	fmt.Printf("StoryTeller for %d people started\n\n", peopleCount)
	sc := bufio.NewScanner(os.Stdin)

	for j := 0; j < peopleCount; j++ {
		words = shuflleRecords(words)

		var generated []string

		for i := 0; i < maxWordsForStoryTeller; i++ {
			generated = append(generated, words[i][0])
		}

		fmt.Printf("generated: (%s)\n", strings.Join(generated, ", "))
		sc.Scan()
	}

	return nil
}

// Activity randomises words without duplicates
func Activity() error {
	words, err := jsonwords.GetWords()
	if err != nil {
		return err
	}

	// maybe shuffleRecords is not necessary
	words = shuflleRecords(words)

	mapW := GetMap(words)
	sc := bufio.NewScanner(os.Stdin)

	fmt.Printf("Activity started\n\n")

	for key, value := range mapW {
		fmt.Printf("generated \"%s\" - [%s]\n", key, value)
		sc.Scan()
	}

	return nil
}
