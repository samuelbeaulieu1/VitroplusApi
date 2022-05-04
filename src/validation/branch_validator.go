package validation

import (
	"reflect"
	"regexp"

	"github.com/samuelbeaulieu1/gimlet/actions"
	"github.com/samuelbeaulieu1/gimlet/responses"
)

func IsValidPhone(action actions.Action, value reflect.Value, field reflect.StructField) (bool, error) {
	val := value.String()
	ok, err := regexp.Match("^(\\+\\d{1,2}\\s)?\\(?\\d{3}\\)?[\\s.-]\\d{3}[\\s.-]\\d{4}$", []byte(val))

	if !ok || err != nil {
		return false, responses.NewError("Le format du numéro de téléphone est invalide")
	}

	return true, nil
}