package interfacesawsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OdbNetwork.
// Experimental.
type IOdbNetworkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OdbNetwork resource.
	// Experimental.
	OdbNetworkRef() *OdbNetworkReference
}

// The jsii proxy for IOdbNetworkRef
type jsiiProxy_IOdbNetworkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IOdbNetworkRef) OdbNetworkRef() *OdbNetworkReference {
	var returns *OdbNetworkReference
	_jsii_.Get(
		j,
		"odbNetworkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOdbNetworkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOdbNetworkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

