package awsbillingconductor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomLineItem.
// Experimental.
type ICustomLineItemRef interface {
	constructs.IConstruct
	// A reference to a CustomLineItem resource.
	// Experimental.
	CustomLineItemRef() *CustomLineItemReference
}

// The jsii proxy for ICustomLineItemRef
type jsiiProxy_ICustomLineItemRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomLineItemRef) CustomLineItemRef() *CustomLineItemReference {
	var returns *CustomLineItemReference
	_jsii_.Get(
		j,
		"customLineItemRef",
		&returns,
	)
	return returns
}

