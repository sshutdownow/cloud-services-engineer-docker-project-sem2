FROM golang:1.17-alpine AS builder
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

# Building the application, statically compile with all libraries built in
RUN CGO_ENABLED=0 \
  go build -v -o appd cmd/api


#  https://hub.docker.com/_/scratch/ - это пустой образ в докере, размер которого - 0 MB
FROM scratch

# копируем собранное приложение
COPY --from=builder app/appd /backend

# https://docs.docker.com/reference/dockerfile/#user
USER 1001:1001

# из приложения
EXPOSE 8081

ENTRYPOINT ["/backend"]
