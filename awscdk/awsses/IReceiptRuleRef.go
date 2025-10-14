package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReceiptRule.
// Experimental.
type IReceiptRuleRef interface {
	constructs.IConstruct
	// A reference to a ReceiptRule resource.
	// Experimental.
	ReceiptRuleRef() *ReceiptRuleReference
}

// The jsii proxy for IReceiptRuleRef
type jsiiProxy_IReceiptRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReceiptRuleRef) ReceiptRuleRef() *ReceiptRuleReference {
	var returns *ReceiptRuleReference
	_jsii_.Get(
		j,
		"receiptRuleRef",
		&returns,
	)
	return returns
}

