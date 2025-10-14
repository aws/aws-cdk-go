package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCConnection.
// Experimental.
type IVPCConnectionRef interface {
	constructs.IConstruct
	// A reference to a VPCConnection resource.
	// Experimental.
	VpcConnectionRef() *VPCConnectionReference
}

// The jsii proxy for IVPCConnectionRef
type jsiiProxy_IVPCConnectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCConnectionRef) VpcConnectionRef() *VPCConnectionReference {
	var returns *VPCConnectionReference
	_jsii_.Get(
		j,
		"vpcConnectionRef",
		&returns,
	)
	return returns
}

