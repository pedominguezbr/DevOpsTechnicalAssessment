build:
	swag init -g main.go
	go build main.go

run:
	swag init -g main.go
	go run main.go

test: generate-grpc ## generate grpc code and run short tests
	go test -v ./... -short

test-it: generate-grpc generate-mocks   ## generate grpc code and mocks and run all tests
	go test -v ./...

test-bench: ## run benchmark tests
	go test -bench ./...

swagger-ui:     ## Run a local swagger-ui to view the generated swagger docs
	sleep 1 && open http://localhost/ &
	docker run --rm -p 80:8080 -e SWAGGER_JSON=/swagger/user.swagger.json -v ${CURDIR}/swagger:/swagger swaggerapi/swagger-ui