package awslogsdestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awslogsdestinations/internal"
)

// Use a Kinesis stream as the destination for a log subscription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var stream stream
//
//   kinesisDestination := awscdk.Aws_logs_destinations.NewKinesisDestination(stream, &KinesisDestinationProps{
//   	Role: role,
//   })
//
// Experimental.
type KinesisDestination interface {
	awslogs.ILogSubscriptionDestination
	// Return the properties required to send subscription events to this destination.
	//
	// If necessary, the destination can use the properties of the SubscriptionFilter
	// object itself to configure its permissions to allow the subscription to write
	// to it.
	//
	// The destination may reconfigure its own permissions in response to this
	// function call.
	// Experimental.
	Bind(scope awscdk.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for KinesisDestination
type jsiiProxy_KinesisDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

// Experimental.
func NewKinesisDestination(stream awskinesis.IStream, props *KinesisDestinationProps) KinesisDestination {
	_init_.Initialize()

	if err := validateNewKinesisDestinationParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisDestination{}

	_jsii_.Create(
		"monocdk.aws_logs_destinations.KinesisDestination",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisDestination_Override(k KinesisDestination, stream awskinesis.IStream, props *KinesisDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs_destinations.KinesisDestination",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisDestination) Bind(scope awscdk.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
	if err := k.validateBindParameters(scope, _sourceLogGroup); err != nil {
		panic(err)
	}
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{scope, _sourceLogGroup},
		&returns,
	)

	return returns
}

