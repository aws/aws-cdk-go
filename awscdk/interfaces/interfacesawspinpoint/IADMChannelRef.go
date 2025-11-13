package interfacesawspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ADMChannel.
// Experimental.
type IADMChannelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ADMChannel resource.
	// Experimental.
	AdmChannelRef() *ADMChannelReference
}

// The jsii proxy for IADMChannelRef
type jsiiProxy_IADMChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IADMChannelRef) AdmChannelRef() *ADMChannelReference {
	var returns *ADMChannelReference
	_jsii_.Get(
		j,
		"admChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IADMChannelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IADMChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

