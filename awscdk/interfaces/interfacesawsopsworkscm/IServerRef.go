package interfacesawsopsworkscm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopsworkscm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Server.
// Experimental.
type IServerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Server resource.
	// Experimental.
	ServerRef() *ServerReference
}

// The jsii proxy for IServerRef
type jsiiProxy_IServerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IServerRef) ServerRef() *ServerReference {
	var returns *ServerReference
	_jsii_.Get(
		j,
		"serverRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

