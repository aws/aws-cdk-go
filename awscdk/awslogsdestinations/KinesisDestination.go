package awslogsdestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogsdestinations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Kinesis stream as the destination for a log subscription.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//   var stream Stream
//   var logGroup LogGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &SubscriptionFilterProps{
//   	LogGroup: LogGroup,
//   	Destination: destinations.NewKinesisDestination(stream),
//   	FilterPattern: logs.FilterPattern_AllTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   	FilterName: jsii.String("ErrorInMainThread"),
//   	Distribution: logs.Distribution_RANDOM,
//   })
//
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
	Bind(scope constructs.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for KinesisDestination
type jsiiProxy_KinesisDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

func NewKinesisDestination(stream awskinesis.IStream, props *KinesisDestinationProps) KinesisDestination {
	_init_.Initialize()

	if err := validateNewKinesisDestinationParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs_destinations.KinesisDestination",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

func NewKinesisDestination_Override(k KinesisDestination, stream awskinesis.IStream, props *KinesisDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs_destinations.KinesisDestination",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisDestination) Bind(scope constructs.Construct, _sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
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

