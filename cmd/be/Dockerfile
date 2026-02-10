# Build stage
FROM golang:1.25-alpine AS build
WORKDIR /app
COPY main.go .
RUN go build -o server main.go

# Runtime stage
FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 80
ENTRYPOINT ["./server"]
CMD ["-port", ":80"]
