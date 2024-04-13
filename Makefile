default: roll

roll: tidy fmt generate vet cmd/roll/main.go
	go build -o bin/ github.com/jadefish/cata-rando/cmd/roll

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: generate
generate:
	go generate

.PHONY: vet
vet:
	go vet ./...
