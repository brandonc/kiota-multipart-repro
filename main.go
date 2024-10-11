package main

import (
	"bytes"
	"context"
	"log"

	"github.com/brandonc/kiota-multipart-repro/generated"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	bundle "github.com/microsoft/kiota-bundle-go"
)

func main() {
	adapter, _ := bundle.NewDefaultRequestAdapter(&authentication.AnonymousAuthenticationProvider{})
	client := generated.NewApiClient(adapter)

	contents := string([]byte("Hello, World!"))
	name := "file.txt"

	multipartBody := abstractions.NewMultipartBody()
	multipartBody.AddOrReplacePart("name", "text/plain", name)
	multipartBody.AddOrReplacePart("file", "application/octet-stream", bytes.NewBufferString(contents))

	// Fixed in 1.19.1
	err := client.Api().Endpoint().Post(context.TODO(), multipartBody, nil)

	if err == nil {
		log.Fatalf("Expected error, got nil")
	}
	// This will print the error message: "only the serialization of multipart bodies is supported with MultipartSerializationWriter"
	log.Fatal(err)
}
