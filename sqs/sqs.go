package sqs

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/pkg/errors"
)

type SQS interface {
	GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)
	// SetQueueAttributes is only implemented on unit tests
	SetQueueAttributes(*sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error)
}

type SqsClient struct {
	Client         SQS
	QueueUrl       string
	AttributeNames []*string
}

var (
	// DefaultAttributeNames define common default attribute names to pass to NewSqsClient.
	DefaultAttributeNames []*string = []*string{
		aws.String("ApproximateNumberOfMessages"),
		aws.String("ApproximateNumberOfMessagesDelayed"),
		aws.String("ApproximateNumberOfMessagesNotVisible"),
	}
)

func NewSqsClient(queue string, region string, attributeNames []*string) *SqsClient {
	svc := sqs.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(region))

	return &SqsClient{
		svc,
		queue,
		attributeNames,
	}
}

func (s *SqsClient) NumMessages() (int, error) {
	params := &sqs.GetQueueAttributesInput{
		AttributeNames: s.AttributeNames,
		QueueUrl:       aws.String(s.QueueUrl),
	}

	out, err := s.Client.GetQueueAttributes(params)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to get messages in SQS")
	}

	var messages int
	for _, attr := range s.AttributeNames {
		approximateNumberOfMessages, err := strconv.Atoi(*out.Attributes[*attr])
		if err != nil {
			return 0, errors.Wrap(err, fmt.Sprintf("Failed to get '%s' number of messages in queue", *attr))
		}

		messages = messages + approximateNumberOfMessages
	}

	return messages, nil
}
