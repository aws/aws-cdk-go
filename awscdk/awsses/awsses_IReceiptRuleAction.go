package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An abstract action for a receipt rule.
type IReceiptRuleAction interface {
	// Returns the receipt rule action specification.
	Bind(receiptRule IReceiptRule) *ReceiptRuleActionConfig
}

// The jsii proxy for IReceiptRuleAction
type jsiiProxy_IReceiptRuleAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IReceiptRuleAction) Bind(receiptRule IReceiptRule) *ReceiptRuleActionConfig {
	if err := i.validateBindParameters(receiptRule); err != nil {
		panic(err)
	}
	var returns *ReceiptRuleActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{receiptRule},
		&returns,
	)

	return returns
}

