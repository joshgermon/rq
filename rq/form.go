package rq

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/getkin/kin-openapi/openapi3"
)


func CallForm(apiSpec *openapi3.T) {
  var (
    endpoint string
    method string
    body string
  )

  form := huh.NewForm(
    huh.NewGroup(
      huh.NewSelect[string]().
        Value(&endpoint).
        Title("Endpoint").
        Options(huh.NewOptions(apiSpec.Paths.InMatchingOrder()...)...),
      huh.NewSelect[string]().
        Value(&method).
        Title("Method").
        OptionsFunc(func() []huh.Option[string] {
          opts := apiSpec.Paths.Find(endpoint).Operations()
          keys := make([]string, 0, len(opts))
          for k := range opts {
            keys = append(keys, k)
          }
          return huh.NewOptions(keys...)
        }, &endpoint),
    ),
  )

  err := form.Run()
  if err != nil {
    log.Fatal(err)
  }

  // Fallback example JSON body for now
  body = "{\"foo\": \"bar\"}"

  // Attempt to get an example body from the OpenAPI spec for the specified endpoint
  example := GetRequestBodyExample(apiSpec, endpoint, method)
  if example != nil {
    body = *example
  }

  // Create new form
  finalForm := huh.NewForm(
    huh.NewGroup(
      huh.NewText().
        Title("Request Body").
        Placeholder("Enter request body...").
        ShowLineNumbers(true).
        Value(&body),
    ).WithHideFunc(func() bool {
      return method != "POST"
    }),
  )

  err = finalForm.Run()
  if err != nil {
    log.Fatal(err)
  }

  log.Println("Calling endpoint:", endpoint)
  log.Println("With method:", method)
  if body != "" {
    log.Println("With body:", body)
  }
}
