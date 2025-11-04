package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCEndpoint.
// Experimental.
type IVPCEndpointRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VPCEndpoint resource.
	// Experimental.
	VpcEndpointRef() *VPCEndpointReference
}

// The jsii proxy for IVPCEndpointRef
type jsiiProxy_IVPCEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVPCEndpointRef) VpcEndpointRef() *VPCEndpointReference {
	var returns *VPCEndpointReference
	_jsii_.Get(
		j,
		"vpcEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCEndpointRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

