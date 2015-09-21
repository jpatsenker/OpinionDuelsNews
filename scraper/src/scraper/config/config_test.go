package config_test

import (
	"os"
	"scraper/config"
	"testing"
)

func init() {
	f, err := os.Open("test_config.json")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	config.InitConfig()
	err = config.ReadFile(f)

	if err != nil {
		panic(err)
	}
}

func TestRead(t *testing.T) {

	cases := []struct {
		key  string
		want interface{}
	}{
		{"one", float64(1)},
		{"two", "2"},
		{"true", true},
	}

	for _, c := range cases {
		got := config.Get(c.key)
		if got != c.want {
			t.Errorf("config[%q] == %v, want %v", c.key, got, c.want)
		}
	}
}

func TestGetInt(t *testing.T) {
	val, _ := config.GetInt("one")
	val = val * 2
	if val != 2 {
		t.Errorf("muling failed: %v", val)
	}

	boo, _ := config.GetBool("true")
	if boo == false {
		t.Errorf("messed up getting bool")
	}
}

func TestGetArr(t *testing.T) {
	val, _ := config.GetArray("array")
	cases := []struct {
		index int
		value string
	}{
		{0, "a"},
		{1, "b"},
		{2, "c"},
	}

	for _, c := range cases {
		if val[c.index] != c.value {
			t.Errorf("config_array[%d] == %v, want %v", c.index, val[c.index], c.value)
		}
	}
}

func TestGetNested(t *testing.T) {
	val, _ := config.Get("nest").(map[string]interface{})
	cases := []struct {
		key  string
		want interface{}
	}{
		{"num", float64(1)},
		{"letter", "a"},
	}

	for _, c := range cases {
		got := val[c.key]
		if got != c.want {
			t.Errorf("config[%q] == %v, want %v", c.key, got, c.want)
		}
	}
}

func TestMultiRead(t *testing.T) {
	f, err := os.Open("test_config2.json")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	err = config.ReadFile(f)

	if err != nil {
		panic(err)
	}

	cases := []struct {
		key  string
		want interface{}
	}{
		{"one", float64(1)},
		{"two", "2"},
		{"true", true},
		{"three", float64(3)},
		{"four", "four"},
	}

	for _, c := range cases {
		got := config.Get(c.key)
		if got != c.want {
			t.Errorf("config[%q] == %v, want %v", c.key, got, c.want)
		}
	}
}

func TestGetMissing(t *testing.T) {
	val := config.Get("ten")
	if val != nil {
		t.Errorf("config[ten] == %v, want nil", val)
	}
}
