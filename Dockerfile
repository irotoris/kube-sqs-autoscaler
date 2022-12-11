FROM golang:1.19-alpine3.17 AS builder
WORKDIR /work
COPY . /work
RUN CGO_ENABLED=0 GOOS=linux go build -o kube-sqs-autoscaler

FROM alpine:3.17
RUN  apk add --no-cache --update ca-certificates
COPY --from=builder /work/kube-sqs-autoscaler ./
CMD ["/kube-sqs-autoscaler"]
