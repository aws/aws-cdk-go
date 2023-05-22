package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsses"
	"github.com/aws/aws-cdk-go/awscdk/awssesactions/internal"
)

// Terminates the evaluation of the receipt rule set and optionally publishes a notification to Amazon SNS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   stop := awscdk.Aws_ses_actions.NewStop(&StopProps{
//   	Topic: topic,
//   })
//
// Experimental.
type Stop interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	// Experimental.
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Stop
type jsiiProxy_Stop struct {
	internal.Type__awssesIReceiptRuleAction
}

// Experimental.
func NewStop(props *StopProps) Stop {
	_init_.Initialize()

	if err := validateNewStopParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Stop{}

	_jsii_.Create(
		"monocdk.aws_ses_actions.Stop",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewStop_Override(s Stop, props *StopProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ses_actions.Stop",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_Stop) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := s.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}

