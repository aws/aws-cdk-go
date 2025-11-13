package interfacesawsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IpAccessSettings.
// Experimental.
type IIpAccessSettingsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IpAccessSettings resource.
	// Experimental.
	IpAccessSettingsRef() *IpAccessSettingsReference
}

// The jsii proxy for IIpAccessSettingsRef
type jsiiProxy_IIpAccessSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IIpAccessSettingsRef) IpAccessSettingsRef() *IpAccessSettingsReference {
	var returns *IpAccessSettingsReference
	_jsii_.Get(
		j,
		"ipAccessSettingsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpAccessSettingsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpAccessSettingsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

