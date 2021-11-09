package helper

// import "github.com/go-playground/validator/v10"

// type ErrorValue struct {
// 	FailedField string `json:"failed_field"`
// 	Tag         string `json:"tag"`
// 	Value       string `json:"value"`
// }

// func ValidateStruct(str interface{}) []*ErrorValue {
// 	var errors []*ErrorValue
// 	validate := validator.New()
// 	err := validate.Struct(str)
// 	nameStruct := GetTagInStruct(str, "json")
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			var element ErrorValue
// 			if n, ok := nameStruct[err.Field()]; ok {
// 				element.FailedField = n
// 			} else {
// 				element.FailedField = err.Field()
// 			}
// 			element.Tag = err.Tag()
// 			element.Value = err.Param()
// 			errors = append(errors, &element)
// 		}
// 	}
// 	return errors
// }
