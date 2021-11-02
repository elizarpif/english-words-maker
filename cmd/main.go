package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("assets/toefl_words.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	englishWordsFile, err := os.Create("assets/result.txt")
	if err != nil {
		panic(err)
	}
	defer englishWordsFile.Close()

	newWordsFile1, err := os.Create("assets/result1.txt")
	if err != nil {
		panic(err)
	}
	defer newWordsFile1.Close()

	newWordsFile2, err := os.Create("assets/result2.txt")
	if err != nil {
		panic(err)
	}
	defer newWordsFile2.Close()

	reader := csv.NewReader(file)

	max := 20

	records, err := reader.ReadAll()
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	for i, v := range records {
		if i == max {
			break
		}

		str := strings.Trim(v[2], "\"")+"\n"

		_, err := englishWordsFile.WriteString(fmt.Sprintf("%s - %s - %s\n",
			strings.Trim(v[0], "\""),
			strings.Trim(v[1], "\""),
			strings.Trim(v[2], "\"")))
		if err != nil {
			panic(err)
		}

		if i%2 == 1 {
			// надя
			_, err = newWordsFile1.WriteString(fmt.Sprintf("%s%s", randSpace(), str))
		} else {
			// лиза
			_, err = newWordsFile2.WriteString(fmt.Sprintf("%s%s", randSpace(), str))
		}

		if err != nil {
			panic(err)
		}
	}
}

func randSpace() string {
	n := rand.Intn(150)

	var res []string
	for i := 0; i < n; i++ {
		res = append(res, " ")
	}

	return strings.Join(res, "")
}
