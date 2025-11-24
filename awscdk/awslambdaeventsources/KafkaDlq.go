package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
)

// A Kafka topic dead letter queue destination configuration for a Lambda event source.
//
// This destination can only be used with Kafka-based event sources (MSK and self-managed Kafka).
// When used with other event source types, a validation error will be thrown.
//
// ## Kafka URI Format
//
// new KafkaDlq('my-topic');
//
// ## Topic Naming Requirements
//
// Kafka topic names must follow these rules:
// - Only alphanumeric characters, dots (.), underscores (_), and hyphens (-) are allowed
// - Cannot be empty
// - Must be a valid Kafka topic name.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction Function
//
//
//   // Your MSK cluster arn
//   clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   // Create a Kafka DLQ destination
//   kafkaDlq := awscdk.NewKafkaDlq(jsii.String("failure-topic"))
//
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: jsii.String(Topic),
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   	OnFailure: kafkaDlq,
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(1),
//   		MaximumPollers: jsii.Number(1),
//   	},
//   }))
//
type KafkaDlq interface {
	awslambda.IEventSourceDlq
	// Returns a destination configuration for the DLQ.
	//
	// The returned configuration is used in the AWS Lambda EventSourceMapping's DestinationConfig
	// to specify where failed records should be sent.
	//
	// Returns: The DLQ destination configuration with the properly formatted Kafka URI.
	Bind(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig
}

// The jsii proxy struct for KafkaDlq
type jsiiProxy_KafkaDlq struct {
	internal.Type__awslambdaIEventSourceDlq
}

// Creates a new Kafka DLQ destination.
func NewKafkaDlq(topicName *string) KafkaDlq {
	_init_.Initialize()

	if err := validateNewKafkaDlqParameters(topicName); err != nil {
		panic(err)
	}
	j := jsiiProxy_KafkaDlq{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.KafkaDlq",
		[]interface{}{topicName},
		&j,
	)

	return &j
}

// Creates a new Kafka DLQ destination.
func NewKafkaDlq_Override(k KafkaDlq, topicName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.KafkaDlq",
		[]interface{}{topicName},
		k,
	)
}

func (k *jsiiProxy_KafkaDlq) Bind(target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig {
	if err := k.validateBindParameters(target, targetHandler); err != nil {
		panic(err)
	}
	var returns *awslambda.DlqDestinationConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{target, targetHandler},
		&returns,
	)

	return returns
}

