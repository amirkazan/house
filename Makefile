
build:
	docker-compose build

run:
	docker-compose up

stop:
	docker-compose down

clean:
	rm -f main
	docker-compose down -v

rebuild: clean build run
