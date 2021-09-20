.PHONY: build
build:
	docker build -t layered-architecture-demo -f ./build/Dockerfile .

.PHONY: run
run:
	docker run -d -p 8080:8080 --rm --name layered-architecture-demo layered-architecture-demo

.PHONY: stop
stop:
	docker stop layered-architecture-demo