FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
ADD shiiip-cli /app/shiiip-cli

CMD ["./shiiip-cli"]