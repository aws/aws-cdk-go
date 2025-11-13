package interfacesawsamazonmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsamazonmq/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationAssociation.
// Experimental.
type IConfigurationAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConfigurationAssociation resource.
	// Experimental.
	ConfigurationAssociationRef() *ConfigurationAssociationReference
}

// The jsii proxy for IConfigurationAssociationRef
type jsiiProxy_IConfigurationAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConfigurationAssociationRef) ConfigurationAssociationRef() *ConfigurationAssociationReference {
	var returns *ConfigurationAssociationReference
	_jsii_.Get(
		j,
		"configurationAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

