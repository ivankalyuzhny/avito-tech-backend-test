FROM golang:1.21

WORKDIR /app

COPY ../go.mod  ./
COPY ../go.sum  ./
RUN go mod download

COPY .. .

RUN go build -o user-segmentation-service ./cmd/user-segmentation-service

CMD ["./user-segmentation-service"]
