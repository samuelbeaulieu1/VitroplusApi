package validation

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/samuelbeaulieu1/gimlet/actions"
	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/vitroplus-api/src/entities"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

func IsUniquePin(action actions.Action, value reflect.Value, field reflect.StructField) (bool, error) {
	val := value.String()
	if len(val) == models.PinLength && entities.NewEmployee().PinExists(val) {
		return false, responses.NewError("Un employé existe déjà avec le même pin")
	}

	return true, nil
}

func IsValidPin(action actions.Action, value reflect.Value, field reflect.StructField) (bool, error) {
	val := value.String()
	if len(val) != 0 && len(val) < models.PinLength {
		return false, responses.NewError(fmt.Sprintf("Le pin doit être composé de %d chiffres", models.PinLength))
	}

	return true, nil
}

func IsValidBranch(action actions.Action, value reflect.Value, field reflect.StructField) (bool, error) {
	val := value.String()
	if !entities.NewBranch().Exists(val) {
		return false, responses.NewError("La succursale est inexistante")
	}

	return true, nil
}

func IsValidEmail(action actions.Action, value reflect.Value, field reflect.StructField) (bool, error) {
	val := value.String()
	ok, err := regexp.Match("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", []byte(val))

	if !ok || err != nil {
		return false, responses.NewError("Le format du courriel est invalide")
	}

	return true, nil
}
