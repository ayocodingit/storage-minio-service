# build
FROM --platform=linux/amd64 golang:1.17-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy -compat=1.17
RUN CGO_ENABLED=0 go build -o bin/services ./src/cmd

FROM --platform=linux/amd64 alpine:3.16.0
COPY --from=builder /app/bin/services /
COPY --from=builder /app/.env.example /.env
COPY --from=builder /app/src /app
EXPOSE 8080
CMD ["/services"]