package awslogsdestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awslogsdestinations/internal"
)

// Use a Lambda Function as the destination for a log subscription.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//   var logGroup logGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &SubscriptionFilterProps{
//   	LogGroup: LogGroup,
//   	Destination: destinations.NewLambdaDestination(fn),
//   	FilterPattern: logs.FilterPattern_AllTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   })
//
// Experimental.
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
	// Experimental.
	Bind(scope awscdk.Construct, logGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

// LambdaDestinationOptions.
// Experimental.
func NewLambdaDestination(fn awslambda.IFunction, options *LambdaDestinationOptions) LambdaDestination {
	_init_.Initialize()

	if err := validateNewLambdaDestinationParameters(fn, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"monocdk.aws_logs_destinations.LambdaDestination",
		[]interface{}{fn, options},
		&j,
	)

	return &j
}

// LambdaDestinationOptions.
// Experimental.
func NewLambdaDestination_Override(l LambdaDestination, fn awslambda.IFunction, options *LambdaDestinationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_logs_destinations.LambdaDestination",
		[]interface{}{fn, options},
		l,
	)
}

func (l *jsiiProxy_LambdaDestination) Bind(scope awscdk.Construct, logGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
	if err := l.validateBindParameters(scope, logGroup); err != nil {
		panic(err)
	}
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, logGroup},
		&returns,
	)

	return returns
}

