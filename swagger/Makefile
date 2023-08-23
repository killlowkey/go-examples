.PHONY: generate
generate:
	swag init -g main.go --output docs/swagger
    # swag init --output docs/swagger

.PHONY: run
run:
	go run main.go

.PHONY: tidy
tidy:
	go mod tidy