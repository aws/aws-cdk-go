package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssesactions/internal"
)

// Rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon SNS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bounceTemplate bounceTemplate
//   var topic topic
//
//   bounce := awscdk.Aws_ses_actions.NewBounce(&BounceProps{
//   	Sender: jsii.String("sender"),
//   	Template: bounceTemplate,
//
//   	// the properties below are optional
//   	Topic: topic,
//   })
//
type Bounce interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Bounce
type jsiiProxy_Bounce struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewBounce(props *BounceProps) Bounce {
	_init_.Initialize()

	if err := validateNewBounceParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Bounce{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Bounce",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBounce_Override(b Bounce, props *BounceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Bounce",
		[]interface{}{props},
		b,
	)
}

func (b *jsiiProxy_Bounce) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := b.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}

