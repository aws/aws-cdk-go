package awsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcEndpoint.
// Experimental.
type IVpcEndpointRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VpcEndpoint resource.
	// Experimental.
	VpcEndpointRef() *VpcEndpointReference
}

// The jsii proxy for IVpcEndpointRef
type jsiiProxy_IVpcEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVpcEndpointRef) VpcEndpointRef() *VpcEndpointReference {
	var returns *VpcEndpointReference
	_jsii_.Get(
		j,
		"vpcEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

