FROM golang:latest AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY ./ ./
RUN cd cmd/server && CGO_ENABLED=0 GOOS=linux go build .

FROM scratch AS final
WORKDIR /app
COPY --from=builder /app/cmd/server/server .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD [ "./server" ]

