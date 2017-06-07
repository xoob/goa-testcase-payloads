SHELL := PS4='\[\e[36m\]--->\[\e[m\] ' bash -x
gopackage := github.com/xoob/goa-testcase-payloads

all: clean generate test

generate:
	@goagen bootstrap --design="$(gopackage)/design"

run:
	@go run main.go greeting.go

test:
	@go test main_test.go || echo "static tests failed"

clean:
	@rm -rf main.go greeting.go extended_greeting.go app/ client/ swagger/ tool/

.PHONY: all generate run clean
