run:
	MONGO_URL=mongodb://nature:nature@localhost:27017 PORT=3333 REDIS_URL=localhost:6379 go run cmd/api/main.go