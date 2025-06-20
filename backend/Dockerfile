FROM golang:1.17-alpine AS builder
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

# Build the Go application as a static binary
# CGO_ENABLED=0 disables CGO, which is crucial for static binaries
# -a ensures all packages are rebuilt,
# -installsuffix cgo prevents CGO conflicts
# -ldflags="-s -w" removes debug information and symbol table for smaller size
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -v -o backend cmd/api

#  https://hub.docker.com/_/scratch/ - это пустой образ в докере, размер которого - 0 MB
FROM scratch

# копируем собранное приложение
COPY --from=builder /app/backend /backend

# https://docs.docker.com/reference/dockerfile/#user
USER 1001:1001

# из приложения
EXPOSE 8081

ENTRYPOINT ["/backend"]
