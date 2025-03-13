package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Use an Amazon Kinesis stream consumer as an event source for AWS Lambda.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction function
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//   streamConsumer := kinesis.NewStreamConsumer(this, jsii.String("MyStreamConsumer"), &StreamConsumerProps{
//   	Stream: Stream,
//   	StreamConsumerName: jsii.String("MyStreamConsumer"),
//   })
//   myFunction.AddEventSource(awscdk.NewKinesisConsumerEventSource(streamConsumer, &KinesisEventSourceProps{
//   	BatchSize: jsii.Number(100),
//   	 // default
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   }))
//
type KinesisConsumerEventSource interface {
	StreamEventSource
	// The ARN for this EventSourceMapping.
	EventSourceMappingArn() *string
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	StreamConsumer() awskinesis.IStreamConsumer
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for KinesisConsumerEventSource
type jsiiProxy_KinesisConsumerEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_KinesisConsumerEventSource) EventSourceMappingArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisConsumerEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisConsumerEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisConsumerEventSource) StreamConsumer() awskinesis.IStreamConsumer {
	var returns awskinesis.IStreamConsumer
	_jsii_.Get(
		j,
		"streamConsumer",
		&returns,
	)
	return returns
}


func NewKinesisConsumerEventSource(streamConsumer awskinesis.IStreamConsumer, props *KinesisEventSourceProps) KinesisConsumerEventSource {
	_init_.Initialize()

	if err := validateNewKinesisConsumerEventSourceParameters(streamConsumer, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisConsumerEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.KinesisConsumerEventSource",
		[]interface{}{streamConsumer, props},
		&j,
	)

	return &j
}

func NewKinesisConsumerEventSource_Override(k KinesisConsumerEventSource, streamConsumer awskinesis.IStreamConsumer, props *KinesisEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.KinesisConsumerEventSource",
		[]interface{}{streamConsumer, props},
		k,
	)
}

func (k *jsiiProxy_KinesisConsumerEventSource) Bind(target awslambda.IFunction) {
	if err := k.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"bind",
		[]interface{}{target},
	)
}

func (k *jsiiProxy_KinesisConsumerEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	if err := k.validateEnrichMappingOptionsParameters(options); err != nil {
		panic(err)
	}
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		k,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

