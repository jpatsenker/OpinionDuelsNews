package config

import (
	"encoding/json"
	"os"
	"reflect"
)

// Get values stored in config files instead of hard coding
// Need to call InitConfig at startup
// Add a file with ReadFile(name, file), name will be id of those values
// To get a value use From(conf file name) to reference the file you want
// then use the getters to get the value you want
// TODO: add warnings, more gets

type configData map[string]interface{}

// Manages config files
func (c configData) Get(key string) interface{} {
	return c[key]
}

func (c configData) GetInt(key string) (int, bool) {
	tmp, ok := c.Get(key).(float64)
	return int(tmp), ok
}

func (c configData) GetBool(key string) (bool, bool) {
	tmp, ok := c.Get(key).(bool)
	return tmp, ok
}

func (c configData) GetArray(key string) ([]interface{}, bool) {
	tmp, ok := c.Get(key).([]interface{})
	return tmp, ok
}

// get a value from the file you want
func From(file string) configData {
	return values[file]
}

func (c configData) Nested(key string) configData {
	val := c.Get(key)

	// need to use reflection to convert {} interface to configData
	reflected := reflect.ValueOf(val)
	cReflected := reflect.ValueOf(c)
	if reflected.Kind() != cReflected.Kind() {
		// TODO: make this return nil and spit out a warning instead of panicing
		panic("ERROR: could not reflect properly on nested conf")
	}
	tmp := make(configData)
	for _, key := range reflected.MapKeys() {
		tmp[key.String()] = reflected.MapIndex(key).Interface()
	}
	return tmp
}

// private member to hold read variables
var values map[string]configData

// call when creating
func InitConfig() {
	values = make(map[string]configData)
}

func ReadFile(name string, file *os.File) error {
	dec := json.NewDecoder(file)
	c := make(configData)
	err := dec.Decode(&c)
	if err != nil {
		panic(err)
	}

	values[name] = c

	return nil
}
