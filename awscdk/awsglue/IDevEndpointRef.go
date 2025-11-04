package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DevEndpoint.
// Experimental.
type IDevEndpointRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DevEndpoint resource.
	// Experimental.
	DevEndpointRef() *DevEndpointReference
}

// The jsii proxy for IDevEndpointRef
type jsiiProxy_IDevEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDevEndpointRef) DevEndpointRef() *DevEndpointReference {
	var returns *DevEndpointReference
	_jsii_.Get(
		j,
		"devEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDevEndpointRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDevEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

