package validation

import (
	"regexp"

	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/gimlet/validators"
)

func IsValidPhone(ctx *validators.ValidationCtx) (bool, error) {
	val := ctx.Value.String()
	ok, err := regexp.Match("^(\\+\\d{1,2}\\s)?\\(?\\d{3}\\)?[\\s.-]\\d{3}[\\s.-]\\d{4}$", []byte(val))

	if !ok || err != nil {
		return false, responses.NewError("Le numéro de téléphone doit être du format (xxx)-xxx-xxxx ou xxx-xxx-xxxx")
	}

	return true, nil
}
