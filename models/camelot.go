package models

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/xdave/keyid/interfaces"
)

var CamelotKeys map[string]interfaces.Scale = map[string]interfaces.Scale{
	"1A":  NewKey("1A"),
	"1B":  NewKey("1B"),
	"2A":  NewKey("2A"),
	"2B":  NewKey("2B"),
	"3A":  NewKey("3A"),
	"3B":  NewKey("3B"),
	"4A":  NewKey("4A"),
	"4B":  NewKey("4B"),
	"5A":  NewKey("5A"),
	"5B":  NewKey("5B"),
	"6A":  NewKey("6A"),
	"6B":  NewKey("6B"),
	"7A":  NewKey("7A"),
	"7B":  NewKey("7B"),
	"8A":  NewKey("8A"),
	"8B":  NewKey("8B"),
	"9A":  NewKey("9A"),
	"9B":  NewKey("9B"),
	"10A": NewKey("10A"),
	"10B": NewKey("10B"),
	"11A": NewKey("11A"),
	"11B": NewKey("11B"),
	"12A": NewKey("12A"),
	"12B": NewKey("12B"),
}

func IsCamelotKey(key string) bool {
	r := regexp.MustCompile("[0-9]{1,2}[AB]")
	return r.Match([]byte(key))
}

func ParseCamelotKey(key string) *CamelotScale {
	camelotKey := &CamelotScale{}

	if !IsCamelotKey(key) {
		panic(fmt.Errorf("%v is not a camelot key", key))
	}

	r := regexp.MustCompile(`(?P<Number>\d{1,2})(?P<Letter>[AB]{1})`)
	match := r.FindStringSubmatch(key)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			switch name {
			case "Number":
				{
					n, _ := strconv.Atoi(match[i])
					camelotKey.Index = n
				}
			case "Letter":
				{
					camelotKey.Kind = interfaces.ScaleKind(rune([]byte(match[i])[0]))
				}
			}
		}
	}

	return camelotKey
}
