FROM golang:1.15-alpine3.14 AS builder
WORKDIR /work
COPY . /work
RUN CGO_ENABLED=0 GOOS=linux go build -o kube-sqs-autoscaler

FROM alpine:3.14
RUN  apk add --no-cache --update ca-certificates
COPY --from=builder /work/kube-sqs-autoscaler ./
CMD ["/kube-sqs-autoscaler"]
