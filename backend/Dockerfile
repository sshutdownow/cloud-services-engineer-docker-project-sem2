FROM golang:1.24 AS builder
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application as a static binary
# CGO_ENABLED=0 disables CGO, which is crucial for static binaries
# -a ensures all packages are rebuilt,
# -ldflags="-s -w" removes debug information and symbol table for smaller size
RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -v  -o /app ./cmd/api ./cmd/healthcheck 

#  https://hub.docker.com/_/scratch/ - это пустой образ в докере, размер которого - 0 MB
FROM scratch

# копируем собранные приложения
COPY --from=builder /app/api /app/healthcheck /

# https://docs.docker.com/reference/dockerfile/#user
USER 1001:1001

# из приложения
EXPOSE 8081

ENTRYPOINT ["/api"]
