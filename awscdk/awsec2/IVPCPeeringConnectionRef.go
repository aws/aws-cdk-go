package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCPeeringConnection.
// Experimental.
type IVPCPeeringConnectionRef interface {
	constructs.IConstruct
	// A reference to a VPCPeeringConnection resource.
	// Experimental.
	VpcPeeringConnectionRef() *VPCPeeringConnectionReference
}

// The jsii proxy for IVPCPeeringConnectionRef
type jsiiProxy_IVPCPeeringConnectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCPeeringConnectionRef) VpcPeeringConnectionRef() *VPCPeeringConnectionReference {
	var returns *VPCPeeringConnectionReference
	_jsii_.Get(
		j,
		"vpcPeeringConnectionRef",
		&returns,
	)
	return returns
}

