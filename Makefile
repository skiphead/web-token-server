.SILENT:

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./app/main ./cmd/main.go

