# Builder Stage
FROM golang:1.13.0-alpine AS builder
WORKDIR /src
COPY ["*.go", "./"]
COPY ["go.mod", "go.sum", "./"]
RUN go build -o bin/bikram-sambat-progress .

# Final Stage
FROM alpine:3.10.2
WORKDIR /bin
COPY --from=builder /src/bin /bin
CMD /bin/bikram-sambat-progress
