package interfacesawschime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawschime/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VoiceConnector.
// Experimental.
type IVoiceConnectorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VoiceConnector resource.
	// Experimental.
	VoiceConnectorRef() *VoiceConnectorReference
}

// The jsii proxy for IVoiceConnectorRef
type jsiiProxy_IVoiceConnectorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVoiceConnectorRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVoiceConnectorRef) VoiceConnectorRef() *VoiceConnectorReference {
	var returns *VoiceConnectorReference
	_jsii_.Get(
		j,
		"voiceConnectorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVoiceConnectorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVoiceConnectorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

