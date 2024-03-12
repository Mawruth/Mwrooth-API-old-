SHELL := /bin/bash

.PHONY: build, run*, kill

build:
	docker image rm -f mwruth
	docker build -t mwruth .
run:
	docker run -p 3000:3000 --name mwruth mwruth
run-bg:
	docker rm -f mwruth
	docker run -p 3000:3000 -d --name mwruth mwruth
kill:
	docker rm -f mwruth
