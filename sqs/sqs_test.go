package sqs

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
)

func TestNumMessages(t *testing.T) {
	s := NewMockSqsClient()

	num, err := s.NumMessages()
	assert.Equal(t, 30, num)
	assert.Nil(t, err)
}

type MockSQS struct {
	QueueAttributes *sqs.GetQueueAttributesOutput
}

func (m *MockSQS) GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	return m.QueueAttributes, nil
}

func (m *MockSQS) SetQueueAttributes(input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
	m.QueueAttributes = &sqs.GetQueueAttributesOutput{
		Attributes: input.Attributes,
	}

	return &sqs.SetQueueAttributesOutput{}, nil
}

func NewMockSqsClient() *SqsClient {
	Attributes := map[string]*string{
		"ApproximateNumberOfMessages":           aws.String("10"),
		"ApproximateNumberOfMessagesDelayed":    aws.String("10"),
		"ApproximateNumberOfMessagesNotVisible": aws.String("10"),
	}

	return &SqsClient{
		Client: &MockSQS{
			QueueAttributes: &sqs.GetQueueAttributesOutput{
				Attributes: Attributes,
			},
		},
		QueueUrl: "example.com",
	}
}
