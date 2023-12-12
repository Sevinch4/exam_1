package store

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJSonFile(fileName string) (Store, error) {
	var s Store
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error is while reading a json file", err.Error())
	}
	err = json.Unmarshal(data, &s.Data)
	if err != nil {
		fmt.Println("error is while unmarshalling json file", err.Error())
	}
	//fmt.Println(string(data))
	return s, nil
}
