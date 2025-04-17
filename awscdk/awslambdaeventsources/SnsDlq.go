package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// An SNS dead letter queue destination configuration for a Lambda event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   snsDlq := awscdk.Aws_lambda_event_sources.NewSnsDlq(topic)
//
type SnsDlq interface {
	awslambda.IEventSourceDlq
	// Returns a destination configuration for the DLQ.
	Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig
}

// The jsii proxy struct for SnsDlq
type jsiiProxy_SnsDlq struct {
	internal.Type__awslambdaIEventSourceDlq
}

func NewSnsDlq(topic awssns.ITopic) SnsDlq {
	_init_.Initialize()

	if err := validateNewSnsDlqParameters(topic); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsDlq{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsDlq",
		[]interface{}{topic},
		&j,
	)

	return &j
}

func NewSnsDlq_Override(s SnsDlq, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SnsDlq",
		[]interface{}{topic},
		s,
	)
}

func (s *jsiiProxy_SnsDlq) Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig {
	if err := s.validateBindParameters(_target, targetHandler); err != nil {
		panic(err)
	}
	var returns *awslambda.DlqDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_target, targetHandler},
		&returns,
	)

	return returns
}

