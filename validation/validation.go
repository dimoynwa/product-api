package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateSku(fl validator.FieldLevel) bool {
	if fl.Field().String() == "invalid" {
		return false
	}

	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	gr := re.FindAllString(fl.Field().String(), -1)

	return len(gr) == 1
}
