FROM golang:latest
ENV POSTGRESURI="postgresql://test:test@postgres:5432/ledgerDB?sslmode=disable"
WORKDIR /opt
COPY . .
#RUN go mod tidy
RUN go build -mod vendor ./cmd/ledger/...
EXPOSE 8080
CMD ["./ledger"]
