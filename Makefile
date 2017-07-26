.PHONY: build run
build:
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t $(image) .

run:
	docker run -it --rm $(image)
