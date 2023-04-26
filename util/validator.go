package util

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	english  = en.New()
	uni      = ut.New(english, english)
	trans, _ = uni.GetTranslator("en")
	_        = enTranslations.RegisterDefaultTranslations(validate, trans)
)

func Validate(obj interface{}) error {
	err := validate.Struct(obj)
	if err != nil {
		return translateError(err)
	}
	return nil
}

func translateError(err error) error {
	validatorErrs := err.(validator.ValidationErrors)
	translatedErr := fmt.Errorf(validatorErrs[0].Translate(trans))

	return translatedErr
}
