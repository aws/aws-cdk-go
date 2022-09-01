package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Use an Amazon SQS queue as an event source for AWS Lambda.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"))
//   eventSource := awscdk.NewSqsEventSource(queue)
//   fn.addEventSource(eventSource)
//
//   eventSourceId := eventSource.eventSourceMappingId
//
type SqsEventSource interface {
	awslambda.IEventSource
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
	Queue() awssqs.IQueue
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for SqsEventSource
type jsiiProxy_SqsEventSource struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_SqsEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsEventSource) Queue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}


func NewSqsEventSource(queue awssqs.IQueue, props *SqsEventSourceProps) SqsEventSource {
	_init_.Initialize()

	if err := validateNewSqsEventSourceParameters(queue, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsEventSource",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

func NewSqsEventSource_Override(s SqsEventSource, queue awssqs.IQueue, props *SqsEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsEventSource",
		[]interface{}{queue, props},
		s,
	)
}

func (s *jsiiProxy_SqsEventSource) Bind(target awslambda.IFunction) {
	if err := s.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

