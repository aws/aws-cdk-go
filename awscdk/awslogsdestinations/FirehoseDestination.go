package awslogsdestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogsdestinations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Data Firehose delivery stream as the destination for a log subscription.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryStream iDeliveryStream
//   var logGroup logGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &SubscriptionFilterProps{
//   	LogGroup: LogGroup,
//   	Destination: destinations.NewFirehoseDestination(deliveryStream),
//   	FilterPattern: logs.FilterPattern_AllEvents(),
//   })
//
type FirehoseDestination interface {
	awslogs.ILogSubscriptionDestination
	// Return the properties required to send subscription events to this destination.
	//
	// If necessary, the destination can use the properties of the SubscriptionFilter
	// object itself to configure its permissions to allow the subscription to write
	// to it.
	//
	// The destination may reconfigure its own permissions in response to this
	// function call.
	Bind(scope constructs.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for FirehoseDestination
type jsiiProxy_FirehoseDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

func NewFirehoseDestination(stream awskinesisfirehose.IDeliveryStream, props *FirehoseDestinationProps) FirehoseDestination {
	_init_.Initialize()

	if err := validateNewFirehoseDestinationParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehoseDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs_destinations.FirehoseDestination",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

func NewFirehoseDestination_Override(f FirehoseDestination, stream awskinesisfirehose.IDeliveryStream, props *FirehoseDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs_destinations.FirehoseDestination",
		[]interface{}{stream, props},
		f,
	)
}

func (f *jsiiProxy_FirehoseDestination) Bind(scope constructs.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
	if err := f.validateBindParameters(scope, _sourceLogGroup); err != nil {
		panic(err)
	}
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope, _sourceLogGroup},
		&returns,
	)

	return returns
}

