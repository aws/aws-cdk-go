package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Use a self hosted Kafka installation as a streaming source for AWS Lambda.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // The secret that allows access to your self hosted Kafka cluster
//   var secret secret
//
//   var myFunction function
//
//
//   // The list of Kafka brokers
//   bootstrapServers := []*string{
//   	"kafka-broker:9092",
//   }
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   // (Optional) The consumer group id to use when connecting to the Kafka broker. If omitted the UUID of the event source mapping will be used.
//   consumerGroupId := "my-consumer-group-id"
//   myFunction.AddEventSource(awscdk.NewSelfManagedKafkaEventSource(&SelfManagedKafkaEventSourceProps{
//   	BootstrapServers: bootstrapServers,
//   	Topic: topic,
//   	ConsumerGroupId: consumerGroupId,
//   	Secret: secret,
//   	BatchSize: jsii.Number(100),
//   	 // default
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   }))
//
type SelfManagedKafkaEventSource interface {
	StreamEventSource
	Props() *StreamEventSourceProps
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for SelfManagedKafkaEventSource
type jsiiProxy_SelfManagedKafkaEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_SelfManagedKafkaEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewSelfManagedKafkaEventSource(props *SelfManagedKafkaEventSourceProps) SelfManagedKafkaEventSource {
	_init_.Initialize()

	if err := validateNewSelfManagedKafkaEventSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SelfManagedKafkaEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SelfManagedKafkaEventSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewSelfManagedKafkaEventSource_Override(s SelfManagedKafkaEventSource, props *SelfManagedKafkaEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SelfManagedKafkaEventSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SelfManagedKafkaEventSource) Bind(target awslambda.IFunction) {
	if err := s.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

func (s *jsiiProxy_SelfManagedKafkaEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	if err := s.validateEnrichMappingOptionsParameters(options); err != nil {
		panic(err)
	}
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		s,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

