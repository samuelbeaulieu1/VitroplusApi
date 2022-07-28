package validation

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/gimlet/validators"
	"github.com/samuelbeaulieu1/vitroplus-api/src/entities"
	"github.com/samuelbeaulieu1/vitroplus-api/src/models"
)

func IsUniquePin(ctx *validators.ValidationCtx) (bool, error) {
	val := ctx.Value.String()
	if len(val) == models.PinLength && entities.NewEmployee().PinExists(val) {
		return false, responses.NewError("Un employé existe déjà avec le même pin")
	}

	return true, nil
}

func IsValidPin(ctx *validators.ValidationCtx) (bool, error) {
	val := ctx.Value.String()
	if _, err := strconv.Atoi(val); len(val) != 0 && (len(val) != models.PinLength || err != nil) {
		return false, responses.NewError(fmt.Sprintf("Le pin doit être composé de %d chiffres", models.PinLength))
	}

	return true, nil
}

func IsValidBranch(ctx *validators.ValidationCtx) (bool, error) {
	val := ctx.Value.String()
	if !entities.NewBranch().Exists(val) {
		return false, responses.NewError("La succursale est inexistante")
	}

	return true, nil
}

func IsValidEmail(ctx *validators.ValidationCtx) (bool, error) {
	val := ctx.Value.String()
	ok, err := regexp.Match("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", []byte(val))

	if !ok || err != nil {
		return false, responses.NewError("Le format du courriel est invalide")
	}

	return true, nil
}
