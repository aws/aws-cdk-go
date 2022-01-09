package awslogsdestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogsdestinations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Kinesis stream as the destination for a log subscription.
//
// TODO: EXAMPLE
//
type KinesisDestination interface {
	awslogs.ILogSubscriptionDestination
	Bind(scope constructs.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for KinesisDestination
type jsiiProxy_KinesisDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

func NewKinesisDestination(stream awskinesis.IStream) KinesisDestination {
	_init_.Initialize()

	j := jsiiProxy_KinesisDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs_destinations.KinesisDestination",
		[]interface{}{stream},
		&j,
	)

	return &j
}

func NewKinesisDestination_Override(k KinesisDestination, stream awskinesis.IStream) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs_destinations.KinesisDestination",
		[]interface{}{stream},
		k,
	)
}

// Return the properties required to send subscription events to this destination.
//
// If necessary, the destination can use the properties of the SubscriptionFilter
// object itself to configure its permissions to allow the subscription to write
// to it.
//
// The destination may reconfigure its own permissions in response to this
// function call.
func (k *jsiiProxy_KinesisDestination) Bind(scope constructs.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{scope, _sourceLogGroup},
		&returns,
	)

	return returns
}

// Use a Lambda Function as the destination for a log subscription.
//
// TODO: EXAMPLE
//
type LambdaDestination interface {
	awslogs.ILogSubscriptionDestination
	Bind(scope constructs.Construct, logGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

// LambdaDestinationOptions.
func NewLambdaDestination(fn awslambda.IFunction, options *LambdaDestinationOptions) LambdaDestination {
	_init_.Initialize()

	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs_destinations.LambdaDestination",
		[]interface{}{fn, options},
		&j,
	)

	return &j
}

// LambdaDestinationOptions.
func NewLambdaDestination_Override(l LambdaDestination, fn awslambda.IFunction, options *LambdaDestinationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs_destinations.LambdaDestination",
		[]interface{}{fn, options},
		l,
	)
}

// Return the properties required to send subscription events to this destination.
//
// If necessary, the destination can use the properties of the SubscriptionFilter
// object itself to configure its permissions to allow the subscription to write
// to it.
//
// The destination may reconfigure its own permissions in response to this
// function call.
func (l *jsiiProxy_LambdaDestination) Bind(scope constructs.Construct, logGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, logGroup},
		&returns,
	)

	return returns
}

// Options that may be provided to LambdaDestination.
//
// TODO: EXAMPLE
//
type LambdaDestinationOptions struct {
	// Whether or not to add Lambda Permissions.
	AddPermissions *bool `json:"addPermissions"`
}

