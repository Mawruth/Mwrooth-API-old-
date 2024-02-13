SHELL := /bin/bash

.PHONY: build, run*, kill

build:
	docker build -t mwruth .
run:
	docker run -p 8080:8080 --name mwruth mwruth
run-bg:
	docker rm -f mwruth
	docker run -p 8080:8080 -d --name mwruth mwruth
kill:
	docker rm -f mwruth
