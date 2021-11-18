package jsonwords

import (
	"encoding/json"
	"io/ioutil"
)

const (
	jsonName = "assets/data.json"
)

func GetWords() ([][]string, error) {
	plan, err := ioutil.ReadFile(jsonName)
	if err != nil {
		return nil, err
	}

	var data map[string]string

	err = json.Unmarshal(plan, &data)
	if err != nil {
		return nil, err
	}

	var res [][]string

	for k, v := range data {
		t := [][]string{{k, v}}

		res = append(res, t...)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
