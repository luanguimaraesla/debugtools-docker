VERSION := 0.2.0

PHONY: build
build:
	sudo docker build -t luanguimaraesla/debugtools:${VERSION} .
	sudo docker tag luanguimaraesla/debugtools:${VERSION} luanguimaraesla/debugtools:latest 

PHONY: publish
publish: build
	sudo docker push luanguimaraesla/debugtools:${VERSION}
	sudo docker push luanguimaraesla/debugtools:latest

PHONY: run
run:
	sudo docker run -it --rm -v "${PWD}/warzone":/warzone luanguimaraesla/debugtools:${VERSION} bash
