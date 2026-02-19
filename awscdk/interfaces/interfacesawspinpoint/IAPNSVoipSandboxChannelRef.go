package interfacesawspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a APNSVoipSandboxChannel.
// Experimental.
type IAPNSVoipSandboxChannelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a APNSVoipSandboxChannel resource.
	// Experimental.
	ApnsVoipSandboxChannelRef() *APNSVoipSandboxChannelReference
}

// The jsii proxy for IAPNSVoipSandboxChannelRef
type jsiiProxy_IAPNSVoipSandboxChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAPNSVoipSandboxChannelRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAPNSVoipSandboxChannelRef) ApnsVoipSandboxChannelRef() *APNSVoipSandboxChannelReference {
	var returns *APNSVoipSandboxChannelReference
	_jsii_.Get(
		j,
		"apnsVoipSandboxChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAPNSVoipSandboxChannelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAPNSVoipSandboxChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

