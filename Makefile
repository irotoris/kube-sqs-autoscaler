.PHONY: test clean compile build push

IMAGE=irotoris/kube-sqs-autoscaler
VERSION=v2.2.1

test:
	go test ./...

clean:
	rm -f kube-sqs-autoscaler

compile: clean
	GOOS=linux go build .

build: clean
	docker build -t $(IMAGE):$(VERSION) .

push: build
	docker push $(IMAGE):$(VERSION)
