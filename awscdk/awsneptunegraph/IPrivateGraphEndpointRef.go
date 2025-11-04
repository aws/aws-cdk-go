package awsneptunegraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptunegraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivateGraphEndpoint.
// Experimental.
type IPrivateGraphEndpointRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PrivateGraphEndpoint resource.
	// Experimental.
	PrivateGraphEndpointRef() *PrivateGraphEndpointReference
}

// The jsii proxy for IPrivateGraphEndpointRef
type jsiiProxy_IPrivateGraphEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPrivateGraphEndpointRef) PrivateGraphEndpointRef() *PrivateGraphEndpointReference {
	var returns *PrivateGraphEndpointReference
	_jsii_.Get(
		j,
		"privateGraphEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrivateGraphEndpointRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrivateGraphEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

