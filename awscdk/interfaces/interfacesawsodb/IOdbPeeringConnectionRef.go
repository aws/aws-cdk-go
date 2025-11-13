package interfacesawsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OdbPeeringConnection.
// Experimental.
type IOdbPeeringConnectionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OdbPeeringConnection resource.
	// Experimental.
	OdbPeeringConnectionRef() *OdbPeeringConnectionReference
}

// The jsii proxy for IOdbPeeringConnectionRef
type jsiiProxy_IOdbPeeringConnectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IOdbPeeringConnectionRef) OdbPeeringConnectionRef() *OdbPeeringConnectionReference {
	var returns *OdbPeeringConnectionReference
	_jsii_.Get(
		j,
		"odbPeeringConnectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOdbPeeringConnectionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOdbPeeringConnectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

