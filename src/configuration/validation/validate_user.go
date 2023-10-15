package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/go-server-updater/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		eng := en.New()
		unt := ut.New(eng, eng)
		transl, _ = unt.GetTranslator("en")
		err := entranslations.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}

}
