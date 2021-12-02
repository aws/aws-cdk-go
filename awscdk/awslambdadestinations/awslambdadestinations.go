package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdadestinations/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an Event Bridge event bus as a Lambda destination.
//
// If no event bus is specified, the default event bus is used.
//
// TODO: EXAMPLE
//
type EventBridgeDestination interface {
	awslambda.IDestination
	Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for EventBridgeDestination
type jsiiProxy_EventBridgeDestination struct {
	internal.Type__awslambdaIDestination
}

func NewEventBridgeDestination(eventBus awsevents.IEventBus) EventBridgeDestination {
	_init_.Initialize()

	j := jsiiProxy_EventBridgeDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.EventBridgeDestination",
		[]interface{}{eventBus},
		&j,
	)

	return &j
}

func NewEventBridgeDestination_Override(e EventBridgeDestination, eventBus awsevents.IEventBus) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.EventBridgeDestination",
		[]interface{}{eventBus},
		e,
	)
}

// Returns a destination configuration.
func (e *jsiiProxy_EventBridgeDestination) Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope, fn, _options},
		&returns,
	)

	return returns
}

// Use a Lambda function as a Lambda destination.
//
// TODO: EXAMPLE
//
type LambdaDestination interface {
	awslambda.IDestination
	Bind(scope constructs.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awslambdaIDestination
}

func NewLambdaDestination(fn awslambda.IFunction, options *LambdaDestinationOptions) LambdaDestination {
	_init_.Initialize()

	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.LambdaDestination",
		[]interface{}{fn, options},
		&j,
	)

	return &j
}

func NewLambdaDestination_Override(l LambdaDestination, fn awslambda.IFunction, options *LambdaDestinationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.LambdaDestination",
		[]interface{}{fn, options},
		l,
	)
}

// Returns a destination configuration.
func (l *jsiiProxy_LambdaDestination) Bind(scope constructs.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, fn, options},
		&returns,
	)

	return returns
}

// Options for a Lambda destination.
//
// TODO: EXAMPLE
//
type LambdaDestinationOptions struct {
	// Whether the destination function receives only the `responsePayload` of the source function.
	//
	// When set to `true` and used as `onSuccess` destination, the destination
	// function will be invoked with the payload returned by the source function.
	//
	// When set to `true` and used as `onFailure` destination, the destination
	// function will be invoked with the error object returned by source function.
	//
	// See the README of this module to see a full explanation of this option.
	ResponseOnly *bool `json:"responseOnly"`
}

// Use a SNS topic as a Lambda destination.
//
// TODO: EXAMPLE
//
type SnsDestination interface {
	awslambda.IDestination
	Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for SnsDestination
type jsiiProxy_SnsDestination struct {
	internal.Type__awslambdaIDestination
}

func NewSnsDestination(topic awssns.ITopic) SnsDestination {
	_init_.Initialize()

	j := jsiiProxy_SnsDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.SnsDestination",
		[]interface{}{topic},
		&j,
	)

	return &j
}

func NewSnsDestination_Override(s SnsDestination, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.SnsDestination",
		[]interface{}{topic},
		s,
	)
}

// Returns a destination configuration.
func (s *jsiiProxy_SnsDestination) Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, fn, _options},
		&returns,
	)

	return returns
}

// Use a SQS queue as a Lambda destination.
//
// TODO: EXAMPLE
//
type SqsDestination interface {
	awslambda.IDestination
	Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for SqsDestination
type jsiiProxy_SqsDestination struct {
	internal.Type__awslambdaIDestination
}

func NewSqsDestination(queue awssqs.IQueue) SqsDestination {
	_init_.Initialize()

	j := jsiiProxy_SqsDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.SqsDestination",
		[]interface{}{queue},
		&j,
	)

	return &j
}

func NewSqsDestination_Override(s SqsDestination, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.SqsDestination",
		[]interface{}{queue},
		s,
	)
}

// Returns a destination configuration.
func (s *jsiiProxy_SqsDestination) Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, fn, _options},
		&returns,
	)

	return returns
}

