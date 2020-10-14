FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go install -mod=vendor -v ./...

WORKDIR /go/src/app
EXPOSE 80
CMD ["app"]