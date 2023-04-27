FROM cosmtrek/air:latest

COPY . /app

WORKDIR /app

RUN go mod download
