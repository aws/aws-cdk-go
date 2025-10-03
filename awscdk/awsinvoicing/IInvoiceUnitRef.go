package awsinvoicing

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinvoicing/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InvoiceUnit.
// Experimental.
type IInvoiceUnitRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInvoiceUnitRef
type jsiiProxy_IInvoiceUnitRef struct {
	internal.Type__constructsIConstruct
}

