package awsmanagedblockchain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmanagedblockchain/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Accessor.
// Experimental.
type IAccessorRef interface {
	constructs.IConstruct
	// A reference to a Accessor resource.
	// Experimental.
	AccessorRef() *AccessorReference
}

// The jsii proxy for IAccessorRef
type jsiiProxy_IAccessorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessorRef) AccessorRef() *AccessorReference {
	var returns *AccessorReference
	_jsii_.Get(
		j,
		"accessorRef",
		&returns,
	)
	return returns
}

