ENV ?= "dev"

install:
	@go get
	@go build

release:
	@./alfred -env=$(ENV) -test-mode=false

repos:
	@cp .alfred/env.sample.yml .alfred/$(ENV).yml
	@echo "Modify .alfred/$(ENV).yml to your liking"

setup:
	@cp .alfred/config.sample.yml .alfred/config.yml
	@echo "Modify .alfred/config.yml to your liking"

test:
	@go test -v

verify:
	@./alfred -env=$(ENV) -test-mode=true