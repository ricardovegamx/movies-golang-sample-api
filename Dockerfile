FROM golang

WORKDIR /go/src/app

COPY src/api/ .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["app"]