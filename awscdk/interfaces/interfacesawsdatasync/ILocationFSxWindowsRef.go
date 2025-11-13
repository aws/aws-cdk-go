package interfacesawsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxWindows.
// Experimental.
type ILocationFSxWindowsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocationFSxWindows resource.
	// Experimental.
	LocationFSxWindowsRef() *LocationFSxWindowsReference
}

// The jsii proxy for ILocationFSxWindowsRef
type jsiiProxy_ILocationFSxWindowsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILocationFSxWindowsRef) LocationFSxWindowsRef() *LocationFSxWindowsReference {
	var returns *LocationFSxWindowsReference
	_jsii_.Get(
		j,
		"locationFSxWindowsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationFSxWindowsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationFSxWindowsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

