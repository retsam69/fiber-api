package helper

import (
	"reflect"
	"strings"
)

func GetTagInStruct(s interface{}, tag_name string) map[string]string {
	st := reflect.TypeOf(s)
	tags := make(map[string]string)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		switch tag_name {
		case "gorm":
			tags[field.Name] = ParseTagGORM(field)
		case "json":
			tags[field.Name] = ParseTagJson(field)
		default:
			tags[field.Name] = ParseTagDefault(field, tag_name)
		}
	}
	return tags
}

func ParseTagDefault(field reflect.StructField, tag_name string) string {
	if tag, ok := field.Tag.Lookup(tag_name); ok {
		if idx := strings.Index(tag, ","); idx != -1 {
			return tag[:idx]
		}
	}
	return ToSnakeCase(field.Name)
}

func ParseTagJson(field reflect.StructField) string {
	if tag, ok := field.Tag.Lookup("json"); ok {
		if idx := strings.Index(tag, ","); idx != -1 {
			return tag[:idx]
		}
	}
	return ToSnakeCase(field.Name)
}

func ParseTagGORM(field reflect.StructField) string {
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
					return values[1]
				}
			}
		}
	}
	return ToSnakeCase(field.Name)
}
