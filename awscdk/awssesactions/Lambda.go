package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssesactions/internal"
)

// Calls an AWS Lambda function, and optionally, publishes a notification to Amazon SNS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var topic topic
//
//   lambda := awscdk.Aws_ses_actions.NewLambda(&LambdaProps{
//   	Function: function_,
//
//   	// the properties below are optional
//   	InvocationType: awscdk.*Aws_ses_actions.LambdaInvocationType_EVENT,
//   	Topic: topic,
//   })
//
type Lambda interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	Bind(rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Lambda
type jsiiProxy_Lambda struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewLambda(props *LambdaProps) Lambda {
	_init_.Initialize()

	if err := validateNewLambdaParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lambda{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Lambda",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLambda_Override(l Lambda, props *LambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Lambda",
		[]interface{}{props},
		l,
	)
}

func (l *jsiiProxy_Lambda) Bind(rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := l.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

