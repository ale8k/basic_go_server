PWD = $(shell pwd)

air:
	docker run -it --rm \
		-w "/go/src" \
		-v $(PWD):/go/src \
		-p 8000:8001 \
		cosmtrek/air

build:
	go build -o ./tmp/main .