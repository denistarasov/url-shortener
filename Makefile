default: run

run:
	@echo "Starting server"
	sudo docker-compose up --build

stop:
	@echo "Stopping server"
	sudo docker-compose stop
