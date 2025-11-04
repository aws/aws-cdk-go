package awslogsdestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogsdestinations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Lambda Function as the destination for a log subscription.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn Function
//   var logGroup LogGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &SubscriptionFilterProps{
//   	LogGroup: LogGroup,
//   	Destination: destinations.NewLambdaDestination(fn),
//   	FilterPattern: logs.FilterPattern_AllTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   	FilterName: jsii.String("ErrorInMainThread"),
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
	Bind(scope constructs.Construct, sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awslogsILogSubscriptionDestination
}

// LambdaDestinationOptions.
func NewLambdaDestination(fn awslambda.IFunction, options *LambdaDestinationOptions) LambdaDestination {
	_init_.Initialize()

	if err := validateNewLambdaDestinationParameters(fn, options); err != nil {
		panic(err)
	}
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

func (l *jsiiProxy_LambdaDestination) Bind(scope constructs.Construct, sourceLogGroup awslogs.ILogGroup) *awslogs.LogSubscriptionDestinationConfig {
	if err := l.validateBindParameters(scope, sourceLogGroup); err != nil {
		panic(err)
	}
	var returns *awslogs.LogSubscriptionDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, sourceLogGroup},
		&returns,
	)

	return returns
}

