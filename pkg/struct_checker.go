package pkg

import (
	"fmt"
	"log"
	"reflect"
)

func StructChecker(s any) []fieldProperties {
	reflectStruct := reflect.TypeOf(s)
	fields := make([]fieldProperties, reflectStruct.NumField())

	for i := range fields {

		fieldName := reflectStruct.Field(i).Name
		fieldType := fmt.Sprintf("%v", reflectStruct.Field(i).Type)

		fieldRequired, err := checkRequired(string(reflectStruct.Field(i).Tag))
		log.Print(string(reflectStruct.Field(i).Tag))
		if err != nil {
			log.Panic(err)
		}

		fieldProperties := fieldProperties{
			Name:      fieldName,
			SnakeName: Case2SnakeUpper(fieldName),
			Type:      fieldType,
			Required:  &fieldRequired,
		}

		fields[i] = fieldProperties

	}

	return fields
}
