package interfacesawspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GCMChannel.
// Experimental.
type IGCMChannelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GCMChannel resource.
	// Experimental.
	GcmChannelRef() *GCMChannelReference
}

// The jsii proxy for IGCMChannelRef
type jsiiProxy_IGCMChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IGCMChannelRef) GcmChannelRef() *GCMChannelReference {
	var returns *GCMChannelReference
	_jsii_.Get(
		j,
		"gcmChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGCMChannelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGCMChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

