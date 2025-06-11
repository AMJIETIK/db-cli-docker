FROM golang:1.23.6
LABEL authors="stormside7"

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]