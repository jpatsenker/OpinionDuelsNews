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
	err = config.ReadFile("default", f)

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
		got := config.From("default").Get(c.key)
		if got != c.want {
			t.Errorf("config[%q] == %v, want %v", c.key, got, c.want)
		}
	}
}

func TestGetInt(t *testing.T) {
	val, _ := config.From("default").GetInt("one")
	val = val * 2
	if val != 2 {
		t.Errorf("muling failed: %v", val)
	}

	boo, _ := config.From("default").GetBool("true")
	if boo == false {
		t.Errorf("messed up getting bool")
	}
}

func TestGetArr(t *testing.T) {
	val, _ := config.From("default").GetArray("array")
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
	cases := []struct {
		key  string
		want interface{}
	}{
		{"num", float64(1)},
		{"letter", "a"},
	}

	for _, c := range cases {
		got := config.From("default").Nested("nest").Get(c.key)
		if got != c.want {
			t.Errorf("config[%q] == %v, want %v", c.key, got, c.want)
		}
	}
}

func TestDoubleNested(t *testing.T) {
	val := config.From("default").Nested("double_nest").Nested("inner").Get("val")
	if val != "a" {
		t.Errorf("expected %v, got %v", "a", val)
	}
}
func TestBadNested(t *testing.T) {
	t.Skip("skipping until panic is updated to be less panic-y")
	val := config.From("default").Nested("one")
	if val != nil {
		t.Errorf("failed to handle val properly")
	}

}
func TestMultiRead(t *testing.T) {
	f, err := os.Open("test_config2.json")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	err = config.ReadFile("other", f)

	if err != nil {
		panic(err)
	}

	cases := []struct {
		key  string
		want interface{}
	}{
		{"three", float64(3)},
		{"four", "four"},
	}

	for _, c := range cases {
		got := config.From("other").Get(c.key)
		if got != c.want {
			t.Errorf("config[%q] == %v, want %v", c.key, got, c.want)
		}
	}
}

func TestGetMissing(t *testing.T) {
	val := config.From("other").Get("ten")
	if val != nil {
		t.Errorf("config[ten] == %v, want nil", val)
	}
}
