FROM golang:alpine
WORKDIR /app
COPY . /app
RUN go build -o text-in-text
ENTRYPOINT ["./text-in-text"]