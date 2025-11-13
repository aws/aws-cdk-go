package interfacesawsevidently

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Launch.
// Experimental.
type ILaunchRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Launch resource.
	// Experimental.
	LaunchRef() *LaunchReference
}

// The jsii proxy for ILaunchRef
type jsiiProxy_ILaunchRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILaunchRef) LaunchRef() *LaunchReference {
	var returns *LaunchReference
	_jsii_.Get(
		j,
		"launchRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

