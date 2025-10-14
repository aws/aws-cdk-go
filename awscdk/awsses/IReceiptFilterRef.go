package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReceiptFilter.
// Experimental.
type IReceiptFilterRef interface {
	constructs.IConstruct
	// A reference to a ReceiptFilter resource.
	// Experimental.
	ReceiptFilterRef() *ReceiptFilterReference
}

// The jsii proxy for IReceiptFilterRef
type jsiiProxy_IReceiptFilterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReceiptFilterRef) ReceiptFilterRef() *ReceiptFilterReference {
	var returns *ReceiptFilterReference
	_jsii_.Get(
		j,
		"receiptFilterRef",
		&returns,
	)
	return returns
}

