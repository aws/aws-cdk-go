package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCEndpointService.
// Experimental.
type IVPCEndpointServiceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VPCEndpointService resource.
	// Experimental.
	VpcEndpointServiceRef() *VPCEndpointServiceReference
}

// The jsii proxy for IVPCEndpointServiceRef
type jsiiProxy_IVPCEndpointServiceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IVPCEndpointServiceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCEndpointServiceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

