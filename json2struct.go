package json2struct

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func capitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}

	// Convert the first character to uppercase
	firstChar := string(input[0])
	upperFirstChar := string([]rune(firstChar)[0] - ('a' - 'A'))

	// Concatenate the modified first character with the rest of the string
	return upperFirstChar + input[1:]
}

func getFirstValue(myInterface interface{}) (interface{}, error) {
	value := reflect.ValueOf(myInterface)

	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		if value.Len() > 0 {
			return value.Index(0).Interface(), nil
		} else {
			return "", nil
		}
	case reflect.Map:
		key := value.MapKeys()[0]
		return value.MapIndex(key).Interface(), nil
	case reflect.String:
		if value.Len() > 0 {
			return value.Index(0).Interface(), nil
		} else {
			return "", nil
		}
	default:
		return "", nil
	}
}

// JSONToStruct converts a JSON object to a Go struct definition string.
func JSONToStruct(key string, jsonData string) (string, error) {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return "", err
	}

	structString := "type " + capitalizeFirstLetter(key) + " struct {\n"

	for key, value := range data {
		fieldType := getFieldType(value)

		if fieldType == "struct{}" {
			jsonStr, _ := json.Marshal(value)
			tempStructString, _ := JSONToStruct(key, string(jsonStr))
			structString = tempStructString + structString
			key = capitalizeFirstLetter(key)
			structString += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", key, key, key)
		} else if fieldType == "[]interface{}" {
			arrValue, _ := getFirstValue(value)
			if arrValue != "" {
				jsonStr, _ := json.Marshal(arrValue)
				tempStructString, _ := JSONToStruct(key, string(jsonStr))
				structString = tempStructString + structString
				key = capitalizeFirstLetter(key)
				structString += fmt.Sprintf("\t%s []%s `json:\"%s\"`\n", key, key, key)
			} else {
				structString += fmt.Sprintf("\t%s [] `json:\"%s\"`\n", key, key)
			}
		} else {
			keyCap := capitalizeFirstLetter(key)
			structString += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", keyCap, fieldType, key)
		}
	}

	structString += "}\n"

	return structString, nil
}

// getFieldType returns the Go type for a given JSON value.
func getFieldType(value interface{}) string {
	switch value.(type) {
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case map[string]interface{}:
		return "struct{}"
	case []interface{}:
		return "[]interface{}"
	default:
		return "interface{}"
	}
}
