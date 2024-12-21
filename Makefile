.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: mock
mock:
	mockgen -source=internal/service/user.go -destination test/mocks/service/user.go
	mockgen -source=internal/repository/user.go -destination test/mocks/repository/user.go
	mockgen -source=internal/repository/repository.go -destination test/mocks/repository/repository.go

.PHONY: test
test:
	go test -coverpkg=./internal/handler,./internal/service,./internal/repository -coverprofile=./test/coverage.out ./test/server/...
	go tool cover -html=./coverage.out -o ./test/coverage.html

.PHONY: build
build:
	go build -ldflags="-s -w" -o ./deploy/bin/server ./cmd/server

.PHONY: docker
docker:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./deploy/bin/server ./cmd/server
	cd ./deploy && docker compose build app && docker compose up -d

.PHONY: swag
swag:
	swag init  -g cmd/server/main.go -o ./docs --parseDependency

.PHONY: gorm
gorm:
	go run cmd/gorm/main.go