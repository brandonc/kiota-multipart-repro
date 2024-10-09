default: generate

generate:
	mkdir -p generated
	docker run -v ./generated:/app/output -v ./spec.yml:/app/openapi.yaml mcr.microsoft.com/openapi/kiota generate --language go -n github.com/brandonc/kiota-multipart-repro/generated

.PHONY: generate
