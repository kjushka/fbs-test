install:
	docker build . -f ./server-d -t server-d

run:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down

logs:
	docker-compose logs -f

reload:
	docker-compose down
	docker build -t fib_grpc:multistage -f Dockerfile.multistage .
	docker-compose up -d
	docker-compose logs -f