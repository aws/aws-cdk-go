package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReceiptRule.
// Experimental.
type IReceiptRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReceiptRuleRef
type jsiiProxy_IReceiptRuleRef struct {
	internal.Type__constructsIConstruct
}

