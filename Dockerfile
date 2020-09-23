FROM golang:1.14 AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v ./cmd/...

FROM alpine:3.12
COPY --from=builder /go/bin/* /usr/local/bin/
EXPOSE 8080
USER 1000:1000
ENV PORT=8080
CMD ["app"]
