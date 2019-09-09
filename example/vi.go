package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

type User struct {
	FirstName string `validate:"required"`
	LastName string `validate:"required"`
	Age uint8 `validate:"gte=0,lte=130"`
	Email string `validate:"required,email"`
}

func main() {
	user := &User{
		FirstName: "firstName",
		LastName: "lastName",
		Age: 136,
		Email: "fl163.com",
	}

	validate := validator.New()
	cn := zh.New()
	uni := ut.New(cn, cn)
	transtator, found := uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, transtator)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("not found")
	}
	err := validate.Struct(user)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			fmt.Println(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				fmt.Println(err.Translate(transtator))

			}
		}
	}
}