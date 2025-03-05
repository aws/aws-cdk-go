package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdadestinations/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
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
//   myFn := lambda.NewFunction(this, jsii.String("Fn"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	// sns topic for successful invocations
//   	OnSuccess: destinations.NewSnsDestination(myTopic),
//   })
//
type SnsDestination interface {
	awslambda.IDestination
	// Returns a destination configuration.
	Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for SnsDestination
type jsiiProxy_SnsDestination struct {
	internal.Type__awslambdaIDestination
}

func NewSnsDestination(topic awssns.ITopic) SnsDestination {
	_init_.Initialize()

	if err := validateNewSnsDestinationParameters(topic); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.SnsDestination",
		[]interface{}{topic},
		&j,
	)

	return &j
}

func NewSnsDestination_Override(s SnsDestination, topic awssns.ITopic) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.SnsDestination",
		[]interface{}{topic},
		s,
	)
}

func (s *jsiiProxy_SnsDestination) Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
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

