package helper

import "github.com/go-playground/validator/v10"

type ErrorValue struct {
	FailedField string      `json:"field"`
	Tag         string      `json:"tag_error"`
	Value       interface{} `json:"value"`
}

func ValidateStruct(str interface{}) []*ErrorValue {
	var errors []*ErrorValue
	validate := validator.New()
	err := validate.Struct(str)
	nameStruct := GetTagInStruct(str, "json")
	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValue
			element.FailedField = ToSnakeCase(err.Field())
			for _, ns := range nameStruct {
				if ns.FielName == err.Field() {
					element.FailedField = ns.TagName
				}
			}
			element.Tag = err.Tag()
			element.Value = err.Value()
			errors = append(errors, &element)
		}
	}
	return errors
}
