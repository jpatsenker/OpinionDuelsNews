package config

import (
	"encoding/json"
	"os"
)

// TODO: add warnings, more gets

// to use: call init when creating array
// pass ReadFile a file to read and insert into the config map
// get values from config map using built in function

// Manages config files
func Get(key string) interface{} {
	return values[key]
}

func GetInt(key string) (int, bool) {
	tmp, ok := Get(key).(float64)
	return int(tmp), ok
}

func GetBool(key string) (bool, bool) {
	tmp, ok := Get(key).(bool)
	return tmp, ok
}

func GetArray(key string) ([]interface{}, bool) {
	tmp, ok := Get(key).([]interface{})
	return tmp, ok
}

// private member to hold read variables
var values map[string]interface{}

// call when creating
func InitConfig() {
	values = make(map[string]interface{})
}

func ReadFile(file *os.File) error {
	// NOTE: will not warn for overlap
	dec := json.NewDecoder(file)
	err := dec.Decode(&values)
	if err != nil {
		panic(err)
	}

	return nil
}
