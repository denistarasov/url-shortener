default: run

run:
	@echo "Starting server"
	sudo docker-compose up --build

tests:
	@echo "Running tests"
	sudo docker-compose -f docker-compose-for-tests.yml up --build --abort-on-container-exit

stop:
	@echo "Stopping server"
	sudo docker-compose stop
