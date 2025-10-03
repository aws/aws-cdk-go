package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomLineItem.
// Experimental.
type ICustomLineItemRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomLineItemRef
type jsiiProxy_ICustomLineItemRef struct {
	internal.Type__constructsIConstruct
}

