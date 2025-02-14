build:
	go build -o lido-no-helper-cli ./cmd

run:
	go run ./cmd/main.go

test:
	go test ./...

clean:
	rm -f lido-no-helper-cli
