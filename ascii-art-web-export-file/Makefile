run: 
	go run cmd/main.go

docker:
	docker build -t ascii-docker .
	docker rmi `docker images --filter label=build=temp -q`
	docker rmi alpine:3.15
	docker rmi golang:1.16-alpine3.15
	docker run -d -p 8080:8080 --name ascii-cont ascii-docker
	docker ps -a
	docker exec -it ascii-cont /bin/sh

remove:
	docker stop ascii-cont
	docker rm ascii-cont
	docker rmi ascii-docker:latest
	docker ps -a
	docker images