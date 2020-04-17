package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"
)

// isInternalError determines if a given error message is related to the schema requirements itself
func isInternalError(errorType string) bool {
	switch errorType {
	case
		"condition_else",
		"condition_then",
		"number_any_of",
		"number_one_of",
		"number_all_of",
		"number_not":
		return true
	default:
		return false
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Missing argument. Set the file out want to validate as first command argument.")
	}
	var fileName = os.Args[1]
	var err error

	binDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + filepath.Join(binDir, "..", "schema", "project-configuration.schema.json"))

	configData, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("read config file: " + err.Error())
	}

	// YAML to JSON
	configJson, err := yaml.YAMLToJSON(configData)
	if err != nil {
		panic("yaml to json: " + err.Error())
	}
	documentLoader := gojsonschema.NewBytesLoader(configJson)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors:\n")
		for _, desc := range result.Errors() {
			if isInternalError(desc.Type()) {
				continue
			}
			fmt.Printf("- %s\n", desc)
		}
	}
}
