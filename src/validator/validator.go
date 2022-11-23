package validator

import (
	"strings"

	"github.com/ayocodingit/storage-minio-service/src/lang"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validate *validator.Validate
	Lang     lang.Lang
}

func New(lang lang.Lang) Validator {
	return Validator{validator.New(), lang}
}

func (v Validator) Validation(args interface{}) map[string]string {
	errors := map[string]string{}

	if err := v.Validate.Struct(args); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := v.Lang.GetMessage(err.Tag(), map[string]interface{}{
				"Field": err.Field(),
			})

			key := strings.ToLower(err.StructNamespace())
			errors[key] = message
		}
	}

	if len(errors) == 0 {
		return nil
	}

	return errors
}
