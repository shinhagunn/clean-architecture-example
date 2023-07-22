build:
	docker build -t mygo .

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f mygo
