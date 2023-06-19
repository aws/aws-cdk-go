package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdadestinations/internal"
)

// Use a Lambda function as a Lambda destination.
//
// Example:
//   // Auto-extract response payload with a lambda destination
//   var destinationFn function
//
//
//   sourceFn := lambda.NewFunction(this, jsii.String("Source"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	// auto-extract on success
//   	OnSuccess: destinations.NewLambdaDestination(destinationFn, &LambdaDestinationOptions{
//   		ResponseOnly: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type LambdaDestination interface {
	awslambda.IDestination
	// Returns a destination configuration.
	// Experimental.
	Bind(scope awscdk.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awslambdaIDestination
}

// Experimental.
func NewLambdaDestination(fn awslambda.IFunction, options *LambdaDestinationOptions) LambdaDestination {
	_init_.Initialize()

	if err := validateNewLambdaDestinationParameters(fn, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"monocdk.aws_lambda_destinations.LambdaDestination",
		[]interface{}{fn, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaDestination_Override(l LambdaDestination, fn awslambda.IFunction, options *LambdaDestinationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_destinations.LambdaDestination",
		[]interface{}{fn, options},
		l,
	)
}

func (l *jsiiProxy_LambdaDestination) Bind(scope awscdk.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	if err := l.validateBindParameters(scope, fn, options); err != nil {
		panic(err)
	}
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, fn, options},
		&returns,
	)

	return returns
}

