FROM golang:1.18.1


RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go build -o /app ./cmd/backend
CMD ["/app"]
