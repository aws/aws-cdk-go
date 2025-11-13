package interfacesawscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfiguredTableAssociation.
// Experimental.
type IConfiguredTableAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConfiguredTableAssociation resource.
	// Experimental.
	ConfiguredTableAssociationRef() *ConfiguredTableAssociationReference
}

// The jsii proxy for IConfiguredTableAssociationRef
type jsiiProxy_IConfiguredTableAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConfiguredTableAssociationRef) ConfiguredTableAssociationRef() *ConfiguredTableAssociationReference {
	var returns *ConfiguredTableAssociationReference
	_jsii_.Get(
		j,
		"configuredTableAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguredTableAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfiguredTableAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

