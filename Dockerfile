FROM golang:1.18 as builder
COPY . /app
WORKDIR /app/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
COPY --from=builder /app/src/app /app
COPY --from=builder /app/src/.env /.env
RUN chmod -R 777 /.env
CMD ["./app"]
