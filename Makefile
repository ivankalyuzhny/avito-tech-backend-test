BINARY_NAME=avito-tech-backend-test

build:
	go build -o ${BINARY_NAME} ./cmd/user-segmentation-service/main.go
.PHONY: build

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
.PHONY: build

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run ./...

up:
	docker-compose -f build/docker-compose.yml -p ${BINARY_NAME} up --force-recreate --build -d

down:
	docker-compose -f build/docker-compose.yml -p ${BINARY_NAME} down

