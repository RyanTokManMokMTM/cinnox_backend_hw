FROM golang AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -o /app/server ./server.go

FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
WORKDIR /app
COPY --from=builder /app/server /app/server
COPY --from=builder /build/config /app/config

CMD ["./server","-f","config.yml"]