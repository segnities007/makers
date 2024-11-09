FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY logics logics
COPY models models

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /makers

EXPOSE 8080

CMD ["/makers"]
