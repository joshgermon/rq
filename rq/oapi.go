package rq

import (
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
