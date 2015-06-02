package config

import (
	"encoding/json"
	"io/ioutil"
)

func loadConfigIntoStruct(filepath string, target interface{}) error {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, target)
	if err != nil {
		return err
	}

	return nil
}
