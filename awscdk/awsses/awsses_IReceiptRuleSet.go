package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
)

// A receipt rule set.
type IReceiptRuleSet interface {
	awscdk.IResource
	// Adds a new receipt rule in this rule set.
	//
	// The new rule is added after
	// the last added rule unless `after` is specified.
	AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule
	// The receipt rule set name.
	ReceiptRuleSetName() *string
}

// The jsii proxy for IReceiptRuleSet
type jsiiProxy_IReceiptRuleSet struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IReceiptRuleSet) AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule {
	var returns ReceiptRule

	_jsii_.Invoke(
		i,
		"addRule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IReceiptRuleSet) ReceiptRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleSetName",
		&returns,
	)
	return returns
}

