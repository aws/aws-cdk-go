package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdadestinations/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Lambda function as a Lambda destination.
//
// Example:
//   // Auto-extract response payload with a lambda destination
//   var destinationFn function
//
//
//   sourceFn := lambda.NewFunction(this, jsii.String("Source"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	// auto-extract on success
//   	onSuccess: destinations.NewLambdaDestination(destinationFn, &lambdaDestinationOptions{
//   		responseOnly: jsii.Boolean(true),
//   	}),
//   })
//
type LambdaDestination interface {
	awslambda.IDestination
	// Returns a destination configuration.
	Bind(scope constructs.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awslambdaIDestination
}

func NewLambdaDestination(fn awslambda.IFunction, options *LambdaDestinationOptions) LambdaDestination {
	_init_.Initialize()

	if err := validateNewLambdaDestinationParameters(fn, options); err != nil {
		panic(err)
	}
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

func (l *jsiiProxy_LambdaDestination) Bind(scope constructs.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
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

