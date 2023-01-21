package helpers

import (
	"errors"
	"regexp"
	"strconv"
)

func ExtractIdFromPath(path string) (int, error) {
	r := regexp.MustCompile(`/([0-9]+)`)
	gr := r.FindAllStringSubmatch(path, -1)
	if len(gr) != 1 && len(gr[0]) != 2 {
		return -1, errors.New("invalid product ID")
	}

	idAsString := gr[0][1]
	return strconv.Atoi(idAsString)
}
