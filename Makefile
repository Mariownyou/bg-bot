:PHONY: run

token := $(shell cat .token)

run:
	go run . -t $(token)
