package interfacesawsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InfrastructureConfiguration.
// Experimental.
type IInfrastructureConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InfrastructureConfiguration resource.
	// Experimental.
	InfrastructureConfigurationRef() *InfrastructureConfigurationReference
}

// The jsii proxy for IInfrastructureConfigurationRef
type jsiiProxy_IInfrastructureConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IInfrastructureConfigurationRef) InfrastructureConfigurationRef() *InfrastructureConfigurationReference {
	var returns *InfrastructureConfigurationReference
	_jsii_.Get(
		j,
		"infrastructureConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInfrastructureConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInfrastructureConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

