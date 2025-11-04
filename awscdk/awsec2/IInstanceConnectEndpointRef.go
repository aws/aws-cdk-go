package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceConnectEndpoint.
// Experimental.
type IInstanceConnectEndpointRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InstanceConnectEndpoint resource.
	// Experimental.
	InstanceConnectEndpointRef() *InstanceConnectEndpointReference
}

// The jsii proxy for IInstanceConnectEndpointRef
type jsiiProxy_IInstanceConnectEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IInstanceConnectEndpointRef) InstanceConnectEndpointRef() *InstanceConnectEndpointReference {
	var returns *InstanceConnectEndpointReference
	_jsii_.Get(
		j,
		"instanceConnectEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceConnectEndpointRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceConnectEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

