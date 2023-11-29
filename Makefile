:PHONY: run

token := $(shell cat .token)

run:
	go run ./cmd/bot -t $(token)
