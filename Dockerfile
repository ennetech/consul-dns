FROM golang:1.9.3

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 53
EXPOSE 4367

CMD ["consul-dns"]