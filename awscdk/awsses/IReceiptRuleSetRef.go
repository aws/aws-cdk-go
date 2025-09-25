package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReceiptRuleSet.
// Experimental.
type IReceiptRuleSetRef interface {
	constructs.IConstruct
	// A reference to a ReceiptRuleSet resource.
	// Experimental.
	ReceiptRuleSetRef() *ReceiptRuleSetReference
}

// The jsii proxy for IReceiptRuleSetRef
type jsiiProxy_IReceiptRuleSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReceiptRuleSetRef) ReceiptRuleSetRef() *ReceiptRuleSetReference {
	var returns *ReceiptRuleSetReference
	_jsii_.Get(
		j,
		"receiptRuleSetRef",
		&returns,
	)
	return returns
}

