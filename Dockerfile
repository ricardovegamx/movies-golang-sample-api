FROM golang

WORKDIR /go

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["movies"]