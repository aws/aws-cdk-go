package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdaeventsources/internal"
)

// Use an stream as an event source for AWS Lambda.
// Experimental.
type StreamEventSource interface {
	awslambda.IEventSource
	// Experimental.
	Props() *StreamEventSourceProps
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	// Experimental.
	Bind(_target awslambda.IFunction)
	// Experimental.
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for StreamEventSource
type jsiiProxy_StreamEventSource struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_StreamEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewStreamEventSource_Override(s StreamEventSource, props *StreamEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_event_sources.StreamEventSource",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StreamEventSource) Bind(_target awslambda.IFunction) {
	if err := s.validateBindParameters(_target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{_target},
	)
}

func (s *jsiiProxy_StreamEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
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

