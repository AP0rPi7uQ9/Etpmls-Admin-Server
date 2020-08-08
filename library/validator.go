package library

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

//Validate
// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

// Example: https://github.com/go-playground/validator/blob/master/_examples/translations/main.go
func ValidateRegister(form interface{}) error {

	zh := zh.New()
	uni = ut.New(zh, zh)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("zh")

	validate = validator.New()
	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		return err
	}

	//Validate Customize Tag Error Message
	/*if err := RegTranslation(validate, trans, "unique-user", "用户名已存在!"); err != nil{
		return err
	}
	if err := RegTranslation(validate, trans, "eqfield", "密码不一致!"); err != nil{
		return err
	}*/

	//Validate Rule: unique-user
	/*if err := validate.RegisterValidation("unique-user", func(fl validator.FieldLevel) bool {
		var user database.User
		database.DB.Where(fl.Param() + " = ?", fl.Field().String()).First(&user)

		//Validate fails if user exists
		if user.ID != 0 {
			return false
		}
		return true
	}); err != nil {
		return err
	}*/

	//Validate Register Form
	if err := validate.Struct(form); err != nil {
		var errMsg string
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			errMsg += e.Translate(trans) + ";"
		}
		msg := strings.TrimRight(errMsg, ";")
		//{
		//    "error": "用户名已存在!;Password长度必须是2个字符"
		//}
		return errors.New(msg)
	}

	return nil
}

func ValidateZh(form interface{}) error {

	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	validate = validator.New()

	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		return err
	}

	if err := handleError(form, trans); err != nil {
		return err
	}

	return nil
}

func handleError(form interface{}, trans ut.Translator) error {
	if err := validate.Struct(form); err != nil {
		var errMsg string
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			errMsg += e.Translate(trans) + ";"
		}
		msg := strings.TrimRight(errMsg, ";")
		//{
		//    "error": "用户名已存在!;Password长度必须是2个字符"
		//}
		return errors.New(msg)
	}
	return nil
}

func RegTranslation(validate *validator.Validate, trans ut.Translator, tag string, text string) (err error) {
	err = validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, text, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
	return err
}