build: 
	docker build -t forum:1.0 .

run:
	docker run -d --name forum-app -p 8080:8080 forum:1.0 && echo "server started at http://localhost:8080/"