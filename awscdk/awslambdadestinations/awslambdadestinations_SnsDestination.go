package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdadestinations/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Use a SNS topic as a Lambda destination.
//
// Example:
//   // An sns topic for successful invocations of a lambda function
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//
//   myFn := lambda.NewFunction(this, jsii.String("Fn"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	// sns topic for successful invocations
//   	onSuccess: destinations.NewSnsDestination(myTopic),
//   })
//
// Experimental.
type SnsDestination interface {
	awslambda.IDestination
	// Returns a destination configuration.
	// Experimental.
	Bind(_scope awscdk.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for SnsDestination
type jsiiProxy_SnsDestination struct {
	internal.Type__awslambdaIDestination
}

// Experimental.
func NewSnsDestination(topic awssns.ITopic) SnsDestination {
	_init_.Initialize()

	if err := validateNewSnsDestinationParameters(topic); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsDestination{}

	_jsii_.Create(
		"monocdk.aws_lambda_destinations.SnsDestination",
		[]interface{}{topic},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsDestination_Override(s SnsDestination, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_destinations.SnsDestination",
		[]interface{}{topic},
		s,
	)
}

func (s *jsiiProxy_SnsDestination) Bind(_scope awscdk.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	if err := s.validateBindParameters(_scope, fn, _options); err != nil {
		panic(err)
	}
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, fn, _options},
		&returns,
	)

	return returns
}

