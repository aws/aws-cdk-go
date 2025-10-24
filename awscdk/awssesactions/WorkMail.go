package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssesactions/internal"
)

// Integrates an Amazon WorkMail action into a receipt rule set, and optionally publishes a notification to Amazon SNS.
//
// Beware that WorkMail should already set up an active `INBOUND_MAIL` rule set
// that includes a WorkMail rule action for this exact purpose.
// This action should be used to customize the behavior of replacement rule set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic Topic
//
//   workMail := awscdk.Aws_ses_actions.NewWorkMail(&WorkMailProps{
//   	OrganizationArn: jsii.String("organizationArn"),
//
//   	// the properties below are optional
//   	Topic: topic,
//   })
//
// See: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-action-workmail.html
//
type WorkMail interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for WorkMail
type jsiiProxy_WorkMail struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewWorkMail(props *WorkMailProps) WorkMail {
	_init_.Initialize()

	if err := validateNewWorkMailParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkMail{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.WorkMail",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewWorkMail_Override(w WorkMail, props *WorkMailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.WorkMail",
		[]interface{}{props},
		w,
	)
}

func (w *jsiiProxy_WorkMail) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := w.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}

