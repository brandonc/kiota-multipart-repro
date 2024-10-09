package main

import (
	"context"
	"log"

	"github.com/brandonc/kiota-multipart-repro/generated"
	"github.com/brandonc/kiota-multipart-repro/generated/api"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	bundle "github.com/microsoft/kiota-bundle-go"
)

func main() {
	adapter, _ := bundle.NewDefaultRequestAdapter(&authentication.AnonymousAuthenticationProvider{})
	client := generated.NewApiClient(adapter)

	contents := string([]byte("Hello, World!"))
	name := "file.txt"

	// Am I meant to use this somehow?
	// multipartBody := abstractions.NewMultipartBody()
	// multipartBody.AddOrReplacePart("name", "text/plain", name)
	// multipartBody.AddOrReplacePart("file", "application/octet-stream", bytes.NewBufferString(contents))

	// This appears to be what the Post method needs
	reqBody := api.NewEndpointPostRequestBody()
	reqBody.SetFile(&contents)
	reqBody.SetName(&name)

	err := client.Api().Endpoint().Post(context.TODO(), reqBody, nil)

	if err == nil {
		log.Fatalf("Expected error, got nil")
	}
	// This will print the error message: "only the serialization of multipart bodies is supported with MultipartSerializationWriter"
	log.Fatal(err)
}
