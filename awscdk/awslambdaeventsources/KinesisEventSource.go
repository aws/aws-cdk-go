package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Use an Amazon Kinesis stream as an event source for AWS Lambda.
//
// Example:
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction function
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//   myFunction.AddEventSource(awscdk.NewKinesisEventSource(stream, &KinesisEventSourceProps{
//   	BatchSize: jsii.Number(100),
//   	 // default
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   }))
//
// Experimental.
type KinesisEventSource interface {
	StreamEventSource
	// The identifier for this EventSourceMapping.
	// Experimental.
	EventSourceMappingId() *string
	// Experimental.
	Props() *StreamEventSourceProps
	// Experimental.
	Stream() awskinesis.IStream
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	// Experimental.
	Bind(target awslambda.IFunction)
	// Experimental.
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for KinesisEventSource
type jsiiProxy_KinesisEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_KinesisEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisEventSource) Stream() awskinesis.IStream {
	var returns awskinesis.IStream
	_jsii_.Get(
		j,
		"stream",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisEventSource(stream awskinesis.IStream, props *KinesisEventSourceProps) KinesisEventSource {
	_init_.Initialize()

	if err := validateNewKinesisEventSourceParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisEventSource{}

	_jsii_.Create(
		"monocdk.aws_lambda_event_sources.KinesisEventSource",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisEventSource_Override(k KinesisEventSource, stream awskinesis.IStream, props *KinesisEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_event_sources.KinesisEventSource",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisEventSource) Bind(target awslambda.IFunction) {
	if err := k.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"bind",
		[]interface{}{target},
	)
}

func (k *jsiiProxy_KinesisEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
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

