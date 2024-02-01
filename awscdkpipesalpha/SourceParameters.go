package awscdkpipesalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspipes"
)

// Source properties.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//
//
//   sourceFilter := pipes.NewFilter([]iFilterPattern{
//   	pipes.FilterPattern_FromObject(map[string]interface{}{
//   		"body": map[string][]*string{
//   			// only forward events with customerType B2B or B2C
//   			"customerType": []*string{
//   				jsii.String("B2B"),
//   				jsii.String("B2C"),
//   			},
//   		},
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: NewSqsSource(sourceQueue),
//   	Target: NewSqsTarget(targetQueue),
//   	Filter: sourceFilter,
//   })
//
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-source.html
//
// Experimental.
type SourceParameters struct {
	// ActiveMQBroker configuration parameters.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-mq.html
	//
	// Default: - none.
	//
	// Experimental.
	ActiveMqBrokerParameters *awspipes.CfnPipe_PipeSourceActiveMQBrokerParametersProperty `field:"optional" json:"activeMqBrokerParameters" yaml:"activeMqBrokerParameters"`
	// DynamoDB stream configuration parameters.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-dynamodb.html
	//
	// Default: - none.
	//
	// Experimental.
	DynamoDbStreamParameters *awspipes.CfnPipe_PipeSourceDynamoDBStreamParametersProperty `field:"optional" json:"dynamoDbStreamParameters" yaml:"dynamoDbStreamParameters"`
	// Kinesis stream configuration parameters.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-kinesis.html
	//
	// Default: - none.
	//
	// Experimental.
	KinesisStreamParameters *awspipes.CfnPipe_PipeSourceKinesisStreamParametersProperty `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// Managed streaming Kafka configuration parameters.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-msk.html
	//
	// Default: - none.
	//
	// Experimental.
	ManagedStreamingKafkaParameters *awspipes.CfnPipe_PipeSourceManagedStreamingKafkaParametersProperty `field:"optional" json:"managedStreamingKafkaParameters" yaml:"managedStreamingKafkaParameters"`
	// RabbitMQ broker configuration parameters.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-mq.html
	//
	// Default: - none.
	//
	// Experimental.
	RabbitMqBrokerParameters *awspipes.CfnPipe_PipeSourceRabbitMQBrokerParametersProperty `field:"optional" json:"rabbitMqBrokerParameters" yaml:"rabbitMqBrokerParameters"`
	// Self-managed Kafka configuration parameters.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-kafka.html
	//
	// Default: - none.
	//
	// Experimental.
	SelfManagedKafkaParameters *awspipes.CfnPipe_PipeSourceSelfManagedKafkaParametersProperty `field:"optional" json:"selfManagedKafkaParameters" yaml:"selfManagedKafkaParameters"`
	// SQS queue configuration parameters.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-sqs.html
	//
	// Default: - none.
	//
	// Experimental.
	SqsQueueParameters *awspipes.CfnPipe_PipeSourceSqsQueueParametersProperty `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
}

