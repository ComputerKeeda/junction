package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/airchains-network/junction/x/trackgate/types"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// ValidateTrackID checks if the given trackId follows the pattern TRK-yyyyMMdd-nnnn
func ValidateTrackID(trackId string) bool {
	// Create a regular expression to match the track ID format
	date := time.Now().Format("20060102")
	regexPattern := fmt.Sprintf("^TRK-%s-[0-9]{4}$", date)
	match, _ := regexp.MatchString(regexPattern, trackId)
	return match
}

func DynamicUnmarshal(structDef types.StructDef, jsonData string) (string, error) {
	// Function to recursively create fields
	var createFields func(fields map[string]interface{}) ([]reflect.StructField, error)
	createFields = func(fields map[string]interface{}) ([]reflect.StructField, error) {
		var result []reflect.StructField
		for name, field := range fields {
			switch f := field.(type) {
			case string: // Base types
				var typ reflect.Type
				switch f {
				case "string":
					typ = reflect.TypeOf("")
				case "int":
					typ = reflect.TypeOf(0)
				case "uint":
					typ = reflect.TypeOf(uint(0))
				case "bytes":
					typ = reflect.TypeOf([]byte{})
				default:
					return nil, fmt.Errorf("unsupported type: %s", f)
				}
				result = append(result, reflect.StructField{
					Name: strings.Title(name),
					Type: typ,
					Tag:  reflect.StructTag(`json:"` + name + `"`),
				})
			case map[string]interface{}: // Nested structs
				nestedFields, err := createFields(f)
				if err != nil {
					return nil, err
				}
				structType := reflect.StructOf(nestedFields)
				result = append(result, reflect.StructField{
					Name: strings.Title(name),
					Type: structType,
					Tag:  reflect.StructTag(`json:"` + name + `"`),
				})
			default:
				return nil, fmt.Errorf("invalid field definition: %v", field)
			}
		}
		return result, nil
	}

	// Create struct type from fields
	fields, err := createFields(structDef.Fields)
	if err != nil {
		return "", err
	}
	structType := reflect.StructOf(fields)
	structInstance := reflect.New(structType).Elem()

	// Unmarshal JSON into the dynamically created struct
	if err := json.Unmarshal([]byte(jsonData), structInstance.Addr().Interface()); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Marshal the struct back to JSON with indentation
	formattedJson, err := json.MarshalIndent(structInstance.Interface(), "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}

	return string(formattedJson), nil
}
