package interfacesawsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentBlueprintConfiguration.
// Experimental.
type IEnvironmentBlueprintConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EnvironmentBlueprintConfiguration resource.
	// Experimental.
	EnvironmentBlueprintConfigurationRef() *EnvironmentBlueprintConfigurationReference
}

// The jsii proxy for IEnvironmentBlueprintConfigurationRef
type jsiiProxy_IEnvironmentBlueprintConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEnvironmentBlueprintConfigurationRef) EnvironmentBlueprintConfigurationRef() *EnvironmentBlueprintConfigurationReference {
	var returns *EnvironmentBlueprintConfigurationReference
	_jsii_.Get(
		j,
		"environmentBlueprintConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentBlueprintConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentBlueprintConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

