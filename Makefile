run:
	docker-compose down --volumes
	docker build -f ./docker/base/Dockerfile . --tag app
	docker-compose up --build

build:
	go build -o app.out -v ./cmd/app/main.go

clean:
	rm -rf *.out *.exe