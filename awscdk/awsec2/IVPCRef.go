package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPC.
// Experimental.
type IVPCRef interface {
	constructs.IConstruct
	// A reference to a VPC resource.
	// Experimental.
	VpcRef() *VPCReference
}

// The jsii proxy for IVPCRef
type jsiiProxy_IVPCRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCRef) VpcRef() *VPCReference {
	var returns *VPCReference
	_jsii_.Get(
		j,
		"vpcRef",
		&returns,
	)
	return returns
}

