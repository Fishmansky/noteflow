FROM golang:1.21

WORKDIR /app

COPY . /app

RUN go mod download && \
go build -o bin .
EXPOSE 80/tcp

ENTRYPOINT [ "/app/bin" ]
