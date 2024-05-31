gen:
	protoc --go_out=. --go-grpc_out=. \
		-I ./proto ./proto/cv.proto
	wire ./internal/app

migrate.up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@10.244.0.2:5436/cvs?sslmode=disable' up

migrate.down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@10.244.0.2:5436/cvs?sslmode=disable' down
	
conn:
	psql -h 10.244.0.2 -p 5436 -U postgres -d cvs
