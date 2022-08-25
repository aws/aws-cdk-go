package awslambdadestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdadestinations/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
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
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("// your code")),
//   	// sqs queue for unsuccessful invocations
//   	onFailure: destinations.NewSqsDestination(deadLetterQueue),
//   })
//
// Experimental.
type SqsDestination interface {
	awslambda.IDestination
	// Returns a destination configuration.
	// Experimental.
	Bind(_scope awscdk.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig
}

// The jsii proxy struct for SqsDestination
type jsiiProxy_SqsDestination struct {
	internal.Type__awslambdaIDestination
}

// Experimental.
func NewSqsDestination(queue awssqs.IQueue) SqsDestination {
	_init_.Initialize()

	j := jsiiProxy_SqsDestination{}

	_jsii_.Create(
		"monocdk.aws_lambda_destinations.SqsDestination",
		[]interface{}{queue},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsDestination_Override(s SqsDestination, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_destinations.SqsDestination",
		[]interface{}{queue},
		s,
	)
}

func (s *jsiiProxy_SqsDestination) Bind(_scope awscdk.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) *awslambda.DestinationConfig {
	var returns *awslambda.DestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, fn, _options},
		&returns,
	)

	return returns
}

