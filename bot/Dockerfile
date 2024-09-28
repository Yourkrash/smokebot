FROM golang

WORKDIR /bot
COPY . .
COPY ../go.mod .
COPY ../dbservice/proto ../dbservice/proto
RUN go mod download
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Запускаем приложение
CMD ["./main"]