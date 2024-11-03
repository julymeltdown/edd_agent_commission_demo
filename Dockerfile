FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN go build -o main . # 현재 디렉토리에서 빌드

WORKDIR /app

CMD ["./cmd/main"]