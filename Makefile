.PHONY: test clean compile build push

IMAGE=irotoris/kube-sqs-autoscaler
TAG := latest

test:
	go test ./...

clean:
	rm -f kube-sqs-autoscaler

compile: clean
	GOOS=linux go build .

build: clean
	docker build -t $(IMAGE):$(TAG) .

push: build
	docker push $(IMAGE):$(TAG)
