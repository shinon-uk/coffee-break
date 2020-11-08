# Go パラメータ
GOCMD=go
GOCLEAN=$(GOCMD) clean
BINARY_NAME=coffee-break
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: clean
clean:
	sudo docker-compose down --rmi all --volumes --remove-orphans
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -rf node_modules/
	rm -rf dist/
	rm -rf mysql/data/
	rm -rf mysql/logs/
	rm -rf coffee-break
.PHONY: setup
setup:
	docker-compose build --no-cache
	npm install
	npm run webpack
	docker-compose up -d
.PHONY: migrate
migrate:
	bee migrate
