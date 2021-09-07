FROM golang:alpine
RUN apk add git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main ./cmd/productmanager/
CMD ["/app/main"]