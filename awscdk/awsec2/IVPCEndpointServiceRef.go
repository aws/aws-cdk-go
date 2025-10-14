package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCEndpointService.
// Experimental.
type IVPCEndpointServiceRef interface {
	constructs.IConstruct
	// A reference to a VPCEndpointService resource.
	// Experimental.
	VpcEndpointServiceRef() *VPCEndpointServiceReference
}

// The jsii proxy for IVPCEndpointServiceRef
type jsiiProxy_IVPCEndpointServiceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCEndpointServiceRef) VpcEndpointServiceRef() *VPCEndpointServiceReference {
	var returns *VPCEndpointServiceReference
	_jsii_.Get(
		j,
		"vpcEndpointServiceRef",
		&returns,
	)
	return returns
}

