package utility

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func msgForTag(f validator.FieldError) string {
	switch f.Tag() {
	case "required":
		return "Ce champ est obligatoire"
	case "email":
		return "Email invalide"
	case "alpha":
		return "Ce champ ne doit contenir exclusivement des letrres"
	}
	return f.Error() // default error
}

func BeautifulError(err error) map[string]string {
	errMsg := make(map[string]string)

	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		for _, f := range verr {
			errMsg[f.Field()] = msgForTag(f)
		}
	}

	return errMsg
}
