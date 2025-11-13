package interfacesawslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Disk.
// Experimental.
type IDiskRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Disk resource.
	// Experimental.
	DiskRef() *DiskReference
}

// The jsii proxy for IDiskRef
type jsiiProxy_IDiskRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDiskRef) DiskRef() *DiskReference {
	var returns *DiskReference
	_jsii_.Get(
		j,
		"diskRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDiskRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDiskRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

