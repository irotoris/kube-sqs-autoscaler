# kube-sqs-autoscaler

Kubernetes pod autoscaler based on queue size in AWS SQS. It periodically retrieves the number of messages in your queue and scales pods accordingly.

Forked https://github.com/Wattpad/kube-sqs-autoscaler

## Setting up

Setting up kube-sqs-autoscaler requires two steps:

1) Deploying it as an incluster service in your cluster
2) Adding AWS permissions so it can read the number of messages in your queues.

### Deploying kube-sqs-autoscaler

Deploy in kube-sqs-autoscaler, see [./deploykube-sqs-autoscaler.yaml](./deploy/kube-sqs-autoscaler.yaml) and set your parameters.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kube-sqs-autoscaler
  labels:
    app: kube-sqs-autoscaler
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kube-sqs-autoscaler
  template:
    metadata:
      labels:
        app: kube-sqs-autoscaler
    spec:
      serviceAccountName: kube-sqs-autoscaler
      containers:
      - name: kube-sqs-autoscaler
        image: irotoris/kube-sqs-autoscaler:latest
        command:
          - /kube-sqs-autoscaler
          - --sqs-queue-url=https://sqs.your_aws_region.amazonaws.com/your_aws_account_number/your_queue_name  # required
          - --kubernetes-deployment=your-kubernetes-deployment-name # required
          - --kubernetes-namespace=$(POD_NAMESPACE) # optional
          - --aws-region=your_aws_region  #required
          - --poll-period=5s # optional
          - --scale-down-cool-down=30s # optional
          - --scale-up-cool-down=5m # optional
          - --scale-up-messages=100 # optional
          - --scale-down-messages=10 # optional
          - --scale-up-pods=1 # optional
          - --scale-down-pods=1 # optional
          - --max-pods=5 # optional
          - --min-pods=1 # optional
          - --attribute-names=ApproximateNumberOfMessages # optional; defaults to ApproximateNumberOfMessages,ApproximateNumberOfMessagesDelayed,ApproximateNumberOfMessagesNotVisible
        env:
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                fieldPath: metadata.namespace
        resources:
          requests:
            memory: "200Mi"
            cpu: "100m"
          limits:
            memory: "200Mi"
            cpu: "100m"
```

Set base64 encrypted `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` to secret resource if you use IAM User and access keys.

### Permissions

Next you want to attach this policy so kube-sqs-autoscaler can retreive SQS attributes:

```json
{
    "Version": "2012-10-17",
    "Statement": [{
        "Effect": "Allow",
        "Action": "sqs:GetQueueAttributes",
        "Resource": "arn:aws:sqs:your_aws_account_number:your_region:your_sqs_queue"
    }]
}
```

Set base64 encrypted `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` to secret resource if you use IAM User and access keys.
