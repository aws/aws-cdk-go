package awslogsdestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogsdestinations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Kinesis stream as the destination for a log subscription.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import kinesis "github.com/aws/aws-cdk-go/awscdk/aws_kinesis"import awscdk "github.com/aws/aws-cdk-go/awscdk"import logs_destinations "github.com/aws/aws-cdk-go/awscdk/aws_logs_destinations"
//
//   var role role
//   var stream stream
//   kinesisDestination := logs_destinations.NewKinesisDestination(stream, &kinesisDestinationProps{
//   	role: role,
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
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{scope, _sourceLogGroup},
		&returns,
	)

	return returns
}

// Customize the Kinesis Logs Destination.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import logs_destinations "github.com/aws/aws-cdk-go/awscdk/aws_logs_destinations"
//
//   var role role
//   kinesisDestinationProps := &kinesisDestinationProps{
//   	role: role,
//   }
//
type KinesisDestinationProps struct {
	// The role to assume to write log events to the destination.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// Use a Lambda Function as the destination for a log subscription.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//   var logGroup logGroup
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &subscriptionFilterProps{
//   	logGroup: logGroup,
//   	destination: destinations.NewLambdaDestination(fn),
//   	filterPattern: logs.filterPattern.allTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   })
//
type LambdaDestination interface {
	awslogs.ILogSubscriptionDestination
	// Return the properties required to send subscription events to this destination.
	//
	// If necessary, the destination can use the properties of the SubscriptionFilter
	// object itself to configure its permissions to allow the subscription to write
	// to it.
	//
	// The destination may reconfigure its own permissions in response to this
	// function call.
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import logs_destinations "github.com/aws/aws-cdk-go/awscdk/aws_logs_destinations"
//   lambdaDestinationOptions := &lambdaDestinationOptions{
//   	addPermissions: jsii.Boolean(false),
//   }
//
type LambdaDestinationOptions struct {
	// Whether or not to add Lambda Permissions.
	AddPermissions *bool `json:"addPermissions" yaml:"addPermissions"`
}

