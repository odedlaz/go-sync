FROM bitnami/golang:1.19 as builder

ARG VERSION
ARG BUILD_TIME
ARG COMMIT

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X github.com/odedlaz/go-sync/server.version=${VERSION} -X github.com/odedlaz/go-sync/server.buildTime=${BUILD_TIME} -X github.com/odedlaz/go-sync/server.commit=${COMMIT}" \
    -o main .

FROM alpine:3.6 as artifact
RUN apk add --update ca-certificates # Certificates for SSL
COPY --from=builder /src/main main

EXPOSE 8295

CMD ["./main"]
