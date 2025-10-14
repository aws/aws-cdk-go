package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcConnection.
// Experimental.
type IVpcConnectionRef interface {
	constructs.IConstruct
	// A reference to a VpcConnection resource.
	// Experimental.
	VpcConnectionRef() *VpcConnectionReference
}

// The jsii proxy for IVpcConnectionRef
type jsiiProxy_IVpcConnectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVpcConnectionRef) VpcConnectionRef() *VpcConnectionReference {
	var returns *VpcConnectionReference
	_jsii_.Get(
		j,
		"vpcConnectionRef",
		&returns,
	)
	return returns
}

