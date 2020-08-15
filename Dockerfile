FROM golang:alpine as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o main .

FROM alpine
RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]