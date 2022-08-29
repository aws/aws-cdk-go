package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Use an Amazon SNS topic as an event source for AWS Lambda.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   var fn function
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.addEventSource(awscdk.NewSnsEventSource(topic, &snsEventSourceProps{
//   	filterPolicy: map[string]interface{}{
//   	},
//   	deadLetterQueue: deadLetterQueue,
//   }))
//
type SnsEventSource interface {
	awslambda.IEventSource
	Topic() awssns.ITopic
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
}

// The jsii proxy struct for SnsEventSource
type jsiiProxy_SnsEventSource struct {
	internal.Type__awslambdaIEventSource
}

func (j *jsiiProxy_SnsEventSource) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


func NewSnsEventSource(topic awssns.ITopic, props *SnsEventSourceProps) SnsEventSource {
	_init_.Initialize()

	j := jsiiProxy_SnsEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsEventSource",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

func NewSnsEventSource_Override(s SnsEventSource, topic awssns.ITopic, props *SnsEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsEventSource",
		[]interface{}{topic, props},
		s,
	)
}

func (s *jsiiProxy_SnsEventSource) Bind(target awslambda.IFunction) {
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{target},
	)
}

