package maker

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	templateRandRusResult    = "assets/rand_rus"
	templateRandEndResult    = "assets/rand_eng"
	templateVocabularyResult = "assets/vocabulary"
	maxForPerson             = 20
	txtExtension             = ".txt"
)

// shuffleRecords shuffles records
func shuflleRecords(records [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	return records
}

// GetMap returns map of words where key is eng and value is rus
func GetMap(words [][]string) map[string]string {
	return getMap(words, false)
}

// GetSmallMap returns small map (with len(words) = 20)
func GetSmallMap(words [][]string) map[string]string {
	return getMap(words, true)
}

func getMap(words [][]string, isDecrease bool) map[string]string {
	mapW := make(map[string]string)
	words = shuflleRecords(words)

	if isDecrease {
		words = words[:maxForPerson]
	}

	for _, w := range words {
		mapW[w[0]] = w[1]
	}

	return mapW
}

// Vocabulary creates vocabulary file
func Vocabulary(words [][]string) error {
	records := shuflleRecords(words)

	newWordsFile, err := os.Create(templateVocabularyResult + ".txt")
	if err != nil {
		return err
	}
	defer newWordsFile.Close()

	for j := 0; j < len(records); j++ {
		_, err = newWordsFile.WriteString(fmt.Sprintf("%d) %s - %s\n", j, records[j][0], records[j][1]))
		if err != nil {
			return err
		}
	}

	return nil
}

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

// CreateGames creates files for players
func CreateGames(words [][]string, players int) error {
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

// writeToFile writes []string to file
func writeToFile(values []string, fileName string) error {
	newWordsFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer newWordsFile.Close()

	_, err = newWordsFile.WriteString(strings.Join(values, "\n"))
	return err
}

// randSpace - генерит рандомное количество пробелов для размещения слов в любом месте строки
func randSpace() string {
	n := rand.Intn(150)

	var res []string
	for i := 0; i < n; i++ {
		res = append(res, " ")
	}

	return strings.Join(res, "")
}
