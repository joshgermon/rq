package rq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/getkin/kin-openapi/openapi3"
)

// ParseOpenApiSpec parses an OpenAPI spec from a YAML file
// and outputs the parsed structure
func ParseOpenApiSpec(openApiSpecFilePath string) *openapi3.T {
	// Perform checks on file validity
  loader := openapi3.NewLoader()
  apiSpec, err := loader.LoadFromFile(openApiSpecFilePath)
	if err != nil {
		log.Fatalf("Error loading OpenAPI file: %v", err)
	}

  return apiSpec
}

// GetRequestBodyExample retrieves the request body example for a specific path and method.
// Returns a JSON string of the example if found, otherwise returns nil.
func GetRequestBodyExample(apiSpec *openapi3.T, path string, method string) *string {
	// Find the path item for the specified path
	pathItem := apiSpec.Paths.Find(path)
	if pathItem == nil {
		log.Printf("Path %s not found in the OpenAPI spec", path)
		return nil
	}

	// Get the operation for the specified method
	operation := pathItem.GetOperation(method)
	if operation == nil {
		log.Printf("Method %s not found for path %s in the OpenAPI spec", method, path)
		return nil
	}

	// Check if the operation has a request body
	if operation.RequestBody != nil && operation.RequestBody.Value != nil {
		requestBody := operation.RequestBody.Value

		// Iterate through the content types in the request body
		for _, mediaType := range requestBody.Content {
			// Check for a single example
			if mediaType.Example != nil {
				jsonData, err := json.Marshal(mediaType.Example)
				if err != nil {
					log.Printf("Error marshaling example to JSON: %v", err)
					return nil
				}
				jsonString := string(jsonData)
				return &jsonString
			}

			// Check for multiple examples and return the first one found
			if mediaType.Examples != nil {
				for _, exampleRef := range mediaType.Examples {
					if exampleRef.Value != nil {
						jsonData, err := json.Marshal(exampleRef.Value.Value)
						if err != nil {
							log.Printf("Error marshaling example to JSON: %v", err)
							return nil
						}
						jsonString := string(jsonData)
						return &jsonString
					}
				}
			}
		}
	}

	// No example found for the given path and method
	return nil
}

// Function just to sanity check the OpenAPI spec remove later
func debugPrintSpec(apiSpec *openapi3.T) {
	// Print some information about the spec
	fmt.Println("OpenAPI Spec Info:")
	fmt.Printf("Title: %s\n", apiSpec.Info.Title)
	fmt.Printf("Version: %s\n", apiSpec.Info.Version)
	fmt.Printf("Description: %s\n", apiSpec.Info.Description)

	// Iterate over paths and methods in the OpenAPI spec
	for _, path := range apiSpec.Paths.InMatchingOrder() {
		fmt.Printf("\nPath: %s\n", path)
    pathItem := apiSpec.Paths.Find(path)

		for method, operation := range pathItem.Operations() {
      fmt.Println("==========")
			fmt.Printf("  Method: %s\n", method)
			fmt.Printf("    Summary: %s\n", operation.Summary)
			fmt.Printf("    Description: %s\n", operation.Description)
		}
	}
}
