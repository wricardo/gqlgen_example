run-local:
	go run ./cmd/gqlgen_example

install-tools:
	go install github.com/vektra/mockery/v2@latest
	go install github.com/99designs/gqlgen@latest

test: 
	go test ./...

generate: generate-graphql

generate-mocks:
	rm -rf mocks
	mockery --config mockery.yaml --name ".*" --keeptree --outpkg mock_gqlgen_example --dir internal/gqlgen_example/generated --output ./mocks/internal/gqlgen_example/generated

generate-graphql:
	go run github.com/99designs/gqlgen generate
