package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReceiptRuleSet.
// Experimental.
type IReceiptRuleSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReceiptRuleSetRef
type jsiiProxy_IReceiptRuleSetRef struct {
	internal.Type__constructsIConstruct
}

