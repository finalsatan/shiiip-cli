build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
	docker build -t shiiip-cli .

run:
	docker run shiiip-cli