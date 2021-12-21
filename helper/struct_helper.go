package helper

import (
	"reflect"
	"strings"
	// 	"github.com/leebenson/conform" // trim data struct
)

type StructTagDetail struct {
	Index     int
	FielName  string
	TagName   string
	TagOption string
}

func GetTagInStruct(s interface{}, tag_name string) []StructTagDetail {
	st := reflect.TypeOf(s)
	tags := []StructTagDetail{}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		tag := StructTagDetail{Index: i, FielName: field.Name}
		switch tag_name {
		case "gorm":
			tag.TagName, tag.TagOption = parseTagGORM(field)
		default:
			tag.TagName, tag.TagOption = parseTagDefault(field, tag_name)
		}
	}
	return tags
}

func parseTagDefault(field reflect.StructField, tag_name string) (string, string) {
	if tag, ok := field.Tag.Lookup(tag_name); ok {
		if idx := strings.SplitN(tag, ",", 1); len(idx) > 1 {
			return idx[0], idx[1]
		} else {
			return idx[0], ""
		}
	}
	return ToSnakeCase(field.Name), ""
}

func parseTagGORM(field reflect.StructField) (string, string) {
	sep := ";"
	if tag, ok := field.Tag.Lookup("gorm"); ok {
		names := strings.Split(tag, sep)
		for i := 0; i < len(names); i++ {
			j := i
			if len(names[j]) > 0 {
				for {
					if names[j][len(names[j])-1] == '\\' {
						i++
						names[j] = names[j][0:len(names[j])-1] + sep + names[i]
						names[i] = ""
					} else {
						break
					}
				}
			}
			values := strings.Split(names[j], ":")
			k := strings.TrimSpace(strings.ToUpper(values[0]))
			if k == "COLUMN" {
				if len(values) >= 2 {
					return values[1], tag
				}
			}
		}
	}
	return ToSnakeCase(field.Name), ""
}
