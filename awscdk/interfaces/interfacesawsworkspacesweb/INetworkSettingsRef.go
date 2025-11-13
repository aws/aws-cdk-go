package interfacesawsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkSettings.
// Experimental.
type INetworkSettingsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NetworkSettings resource.
	// Experimental.
	NetworkSettingsRef() *NetworkSettingsReference
}

// The jsii proxy for INetworkSettingsRef
type jsiiProxy_INetworkSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_INetworkSettingsRef) NetworkSettingsRef() *NetworkSettingsReference {
	var returns *NetworkSettingsReference
	_jsii_.Get(
		j,
		"networkSettingsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkSettingsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkSettingsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

