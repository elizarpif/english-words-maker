package json_to_csv

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	jsonName = "assets/data.json"
	csvName  = "assets/data.csv"
)

type CsvMaker struct {
}

func NewCsvMaker() *CsvMaker {
	return &CsvMaker{}
}

func (m *CsvMaker) Do() {
	plan, err := ioutil.ReadFile(jsonName)
	if err != nil {
		panic(err)
	}

	var data map[string]string

	err = json.Unmarshal(plan, &data)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(csvName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	for k, v := range data{
		err = w.Write([]string{k, v})
		if err != nil {
			panic(err)
		}
	}

	w.Flush()
}
