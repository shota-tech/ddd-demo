.PHONY: build
build:
	docker-compose -f ./build/docker-compose.yml build --no-cache

.PHONY: run
run:
	docker-compose -f ./build/docker-compose.yml up -d

.PHONY: stop
stop:
	docker-compose -f ./build/docker-compose.yml down -v