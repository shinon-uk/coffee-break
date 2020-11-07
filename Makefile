setup:
	npm install
	npm run webpack
	docker-compose build --no-cache
	docker-compose up

clean:
	sudo docker-compose down --rmi all --volumes --remove-orphans
	rm -rf node_modules/
	rm -rf dist/
	rm -rf mysql/data/
	rm -rf mysql/logs/
	rm -rf coffee-break
