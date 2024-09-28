FROM golang


WORKDIR /go
COPY . go

RUN go mod download
RUN go mod tidy

RUN go build -o dbservice/main.go .

CMD ["./main.go"]