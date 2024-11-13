# Stage 1: Build App
FROM golang:1.23-bookworm AS build

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN make deps

COPY . .

RUN make build

# Stage 2: Run App
FROM debian:bookworm-slim AS runtime

WORKDIR /app

COPY --from=build /app/dist/app .

ENTRYPOINT ["./app"]
