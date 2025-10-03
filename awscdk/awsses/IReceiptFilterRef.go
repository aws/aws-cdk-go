package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReceiptFilter.
// Experimental.
type IReceiptFilterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IReceiptFilterRef
type jsiiProxy_IReceiptFilterRef struct {
	internal.Type__constructsIConstruct
}

