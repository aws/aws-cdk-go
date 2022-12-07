package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdadestinations/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a SQS queue as a Lambda destination.
//
// Example:
//   // An sqs queue for unsuccessful invocations of a lambda function
//   import sqs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   myFn := lambda.NewFunction(this, jsii.String("Fn"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("// your code")),
//   	// sqs queue for unsuccessful invocations
//   	onFailure: destinations.NewSqsDestination(deadLetterQueue),
//   })
//
type SqsDestination interface {
	awslambda.IDestination
	// Returns a destination configuration.
	Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for SqsDestination
type jsiiProxy_SqsDestination struct {
	internal.Type__awslambdaIDestination
}

func NewSqsDestination(queue awssqs.IQueue) SqsDestination {
	_init_.Initialize()

	if err := validateNewSqsDestinationParameters(queue); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.SqsDestination",
		[]interface{}{queue},
		&j,
	)

	return &j
}

func NewSqsDestination_Override(s SqsDestination, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_destinations.SqsDestination",
		[]interface{}{queue},
		s,
	)
}

func (s *jsiiProxy_SqsDestination) Bind(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
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

