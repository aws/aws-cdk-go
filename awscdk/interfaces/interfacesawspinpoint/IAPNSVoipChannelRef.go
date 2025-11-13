package interfacesawspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSVoipChannel.
// Experimental.
type IAPNSVoipChannelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a APNSVoipChannel resource.
	// Experimental.
	ApnsVoipChannelRef() *APNSVoipChannelReference
}

// The jsii proxy for IAPNSVoipChannelRef
type jsiiProxy_IAPNSVoipChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAPNSVoipChannelRef) ApnsVoipChannelRef() *APNSVoipChannelReference {
	var returns *APNSVoipChannelReference
	_jsii_.Get(
		j,
		"apnsVoipChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAPNSVoipChannelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAPNSVoipChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

