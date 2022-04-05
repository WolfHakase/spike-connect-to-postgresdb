connect_database_up:
	docker build -f Dockerfile -t connect_database .
	docker run -d --name connect_database -p 5432:5432 connect_database
