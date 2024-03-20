FROM golang:1.21.4-alpine3.18 as builder

WORKDIR /k8s_learn

COPY . .

RUN go mod tidy

RUN go build -o main .

FROM alpine:3.15.0-rc.4

WORKDIR /

COPY --from=builder /k8s_learn .

EXPOSE 8080

ENTRYPOINT ["./main"]