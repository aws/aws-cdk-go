package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Owner.
// Experimental.
type IOwnerRef interface {
	constructs.IConstruct
	// A reference to a Owner resource.
	// Experimental.
	OwnerRef() *OwnerReference
}

// The jsii proxy for IOwnerRef
type jsiiProxy_IOwnerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOwnerRef) OwnerRef() *OwnerReference {
	var returns *OwnerReference
	_jsii_.Get(
		j,
		"ownerRef",
		&returns,
	)
	return returns
}

