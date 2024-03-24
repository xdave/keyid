package util

import (
	"regexp"
	"strconv"
)

func ParseEnergy(comment string) int {
	r := regexp.MustCompile(`Energy ([0-9]+)`)
	matches := r.FindAllStringSubmatch(comment, -1)
	if len(matches) == 0 || len(matches[0]) == 0 {
		return 0
	}
	value, _ := strconv.Atoi(matches[0][1])
	return value
}
