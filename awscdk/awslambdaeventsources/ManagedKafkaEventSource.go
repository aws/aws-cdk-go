package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Use a MSK cluster as a streaming source for AWS Lambda.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // With provisioned pollers and poller group for cost optimization
//   var myFunction Function
//
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String("arn:aws:kafka:us-east-1:123456789012:cluster/my-cluster/abcd1234-abcd-cafe-abab-9876543210ab-4"),
//   	Topic: jsii.String("orders-topic"),
//   	StartingPosition: awscdk.StartingPosition_LATEST,
//   	ProvisionedPollerConfig: &ProvisionedPollerConfig{
//   		MinimumPollers: jsii.Number(2),
//   		MaximumPollers: jsii.Number(10),
//   		PollerGroupName: jsii.String("shared-kafka-pollers"),
//   	},
//   }))
//
type ManagedKafkaEventSource interface {
	StreamEventSource
	// The ARN for this EventSourceMapping.
	EventSourceMappingArn() *string
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for ManagedKafkaEventSource
type jsiiProxy_ManagedKafkaEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_ManagedKafkaEventSource) EventSourceMappingArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKafkaEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewManagedKafkaEventSource(props *ManagedKafkaEventSourceProps) ManagedKafkaEventSource {
	_init_.Initialize()

	if err := validateNewManagedKafkaEventSourceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedKafkaEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ManagedKafkaEventSource",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewManagedKafkaEventSource_Override(m ManagedKafkaEventSource, props *ManagedKafkaEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.ManagedKafkaEventSource",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_ManagedKafkaEventSource) Bind(target awslambda.IFunction) {
	if err := m.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"bind",
		[]interface{}{target},
	)
}

func (m *jsiiProxy_ManagedKafkaEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	if err := m.validateEnrichMappingOptionsParameters(options); err != nil {
		panic(err)
	}
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		m,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

