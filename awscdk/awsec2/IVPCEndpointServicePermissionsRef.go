package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCEndpointServicePermissions.
// Experimental.
type IVPCEndpointServicePermissionsRef interface {
	constructs.IConstruct
	// A reference to a VPCEndpointServicePermissions resource.
	// Experimental.
	VpcEndpointServicePermissionsRef() *VPCEndpointServicePermissionsReference
}

// The jsii proxy for IVPCEndpointServicePermissionsRef
type jsiiProxy_IVPCEndpointServicePermissionsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCEndpointServicePermissionsRef) VpcEndpointServicePermissionsRef() *VPCEndpointServicePermissionsReference {
	var returns *VPCEndpointServicePermissionsReference
	_jsii_.Get(
		j,
		"vpcEndpointServicePermissionsRef",
		&returns,
	)
	return returns
}

