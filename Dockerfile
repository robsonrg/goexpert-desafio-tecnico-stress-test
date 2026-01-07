FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o stresstest main.go

# FROM alpine:latest
# WORKDIR /app

# COPY --from=builder /app/stresstest .

FROM scratch
WORKDIR /app

# Copia os certificados CA do sistema Alpine da fase de build, pq a imagem scratch não tem
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/stresstest .

ENTRYPOINT ["./stresstest"]