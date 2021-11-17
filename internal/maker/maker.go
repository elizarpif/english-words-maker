package maker

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Maker struct {
	vocabulary *os.File
	result     *os.File
	res1       *os.File
	res2       *os.File
	isLesson   bool

	records [][]string
}

func NewMaker(day int, isLesson bool) *Maker {
	if day < 1 {
		panic("day < 1")
	}

	m := &Maker{
		isLesson: isLesson,
	}

	vocabularyName := vocabulary
	if isLesson {
		vocabularyName = lessonVocabulary
	}

	m.openVocabulary(vocabularyName)
	m.createResults()
	m.readVocabulary(day)

	return m
}

const (
	vocabulary               = "assets/toefl_words.csv"
	lessonVocabulary         = "assets/data.csv"
	resultWords              = "assets/result.txt"
	result1                  = "assets/result1.txt"
	result2                  = "assets/result2.txt"
	templateRandRusResult    = "assets/rand_rus"
	templateRandEndResult    = "assets/rand_eng"
	templateVocabularyResult = "assets/vocabulary"
	max                      = 30
	maxForPerson             = 20
)

func (m *Maker) readVocabulary(day int) {
	reader := csv.NewReader(m.vocabulary)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	if !m.isLesson {
		records = shuflleRecords(records)
	}

	n := day - 1 // 0
	m.records = records[n*max : day*max]
}

func shuflleRecords(records [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	return records
}

func (m *Maker) openVocabulary(vocabularyName string) {
	file, err := os.Open(vocabularyName)
	if err != nil {
		panic(err)
	}

	m.vocabulary = file
}

func (m *Maker) createResults() {
	englishWordsFile, err := os.Create(resultWords)
	if err != nil {
		panic(err)
	}
	m.result = englishWordsFile

	newWordsFile1, err := os.Create(result1)
	if err != nil {
		panic(err)
	}
	m.res1 = newWordsFile1

	newWordsFile2, err := os.Create(result2)
	if err != nil {
		panic(err)
	}
	m.res2 = newWordsFile2
}

func (m *Maker) Close() {
	err := m.vocabulary.Close()
	if err != nil {
		panic(err)
	}

	err = m.result.Close()
	if err != nil {
		panic(err)
	}

	err = m.res1.Close()
	if err != nil {
		panic(err)
	}

	err = m.res2.Close()
	if err != nil {
		panic(err)
	}
}

func (m *Maker) AllFromToeflVocabulary() {
	records := m.records

	for i, v := range records {
		if i == max {
			break
		}

		var err error
		var str string

		if m.isLesson {
			str = strings.Trim(v[1], "\"") + "\n"

			_, err = m.result.WriteString(fmt.Sprintf("%s - %s \n",
				strings.Trim(v[0], "\""),
				strings.Trim(v[1], "\"")))
		} else {
			str = strings.Trim(v[2], "\"") + "\n"

			_, err = m.result.WriteString(fmt.Sprintf("%s - %s - %s\n",
				strings.Trim(v[0], "\""),
				strings.Trim(v[1], "\""),
				strings.Trim(v[2], "\"")))
		}

		if err != nil {
			panic(err)
		}

		if i%2 == 1 {
			// надя
			_, err = m.res1.WriteString(fmt.Sprintf("%s%s", randSpace(), str))
		} else {
			// лиза
			_, err = m.res2.WriteString(fmt.Sprintf("%s%s", randSpace(), str))
		}

		if err != nil {
			panic(err)
		}
	}
}

// randFromToeflVocabulary из 20 слов берет рандомные 10 на листик
func (m *Maker) randFromToeflVocabulary(fileName string) error {
	records := m.records
	records = shuflleRecords(records)

	values := make([]string, 0, len(records))

	for i, v := range records {
		if i == maxForPerson {
			break
		}

		str := strings.Trim(v[1], "\"") + "\n"
		if !m.isLesson {
			str = strings.Trim(v[2], "\"") + "\n"
		}

		values = append(values, fmt.Sprintf("%s%s", randSpace(), str))
	}

	err := toOdf(fileName, values)
	return err
}

func (m *Maker) RandFromToeflVocabulary(listCount int) error {
	var err error

	for i := 0; i < listCount; i++ {
		err = m.randFromToeflVocabulary(templateRandRusResult + strconv.Itoa(i+1) + ".txt")
		if err != nil {
			break
		}
	}

	return err
}

func GetMap(words [][]string) map[string]string {
	mapW := make(map[string]string)
	for _, w := range words {
		mapW[w[0]] = w[1]
	}

	return mapW
}

func Vocabulary(words [][]string) error {
	records := shuflleRecords(words)

	newWordsFile, err := os.Create(templateVocabularyResult + ".txt")
	if err != nil {
		return err
	}
	defer newWordsFile.Close()

	for j := 0; j < len(records); j++ {
		_, err = newWordsFile.WriteString(fmt.Sprintf("%s - %s\n", records[j][0], records[j][1]))
		if err != nil {
			return err
		}
	}

	return nil
}
func CreateGames(words [][]string, players int) error {
	for i := 0; i < players; i++ {
		records := shuflleRecords(words)

		var values []string
		var valuesEng []string
		for j := 0; j < len(records)-5; j++ {
			values = append(values, fmt.Sprintf("%s%s", randSpace(), records[j][1]))
			valuesEng = append(valuesEng, fmt.Sprintf("%s%s", randSpace(), records[j][0]))
		}

		newWordsFile, err := os.Create(templateRandRusResult + strconv.Itoa(i+1) + ".txt")
		if err != nil {
			return err
		}
		defer newWordsFile.Close()

		_, err = newWordsFile.WriteString(strings.Join(values, "\n"))
		if err != nil {
			return err
		}

		newWordsFileEng, err := os.Create(templateRandEndResult + strconv.Itoa(i+1) + ".txt")
		if err != nil {
			return err
		}
		defer newWordsFileEng.Close()

		_, err = newWordsFileEng.WriteString(strings.Join(valuesEng, "\n"))
		if err != nil {
			return err
		}
	}

	return nil
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
