package interfacesawsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RemediationConfiguration.
// Experimental.
type IRemediationConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RemediationConfiguration resource.
	// Experimental.
	RemediationConfigurationRef() *RemediationConfigurationReference
}

// The jsii proxy for IRemediationConfigurationRef
type jsiiProxy_IRemediationConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRemediationConfigurationRef) RemediationConfigurationRef() *RemediationConfigurationReference {
	var returns *RemediationConfigurationReference
	_jsii_.Get(
		j,
		"remediationConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemediationConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRemediationConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

