FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8090

CMD ["./main"]

#FROM golang:latest
#
#WORKDIR /app
#
#COPY --from=builder /app/main .
#
#EXPOSE 8090
#
#CMD ["./main"]
