package interfacesawsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConfigurationSetEventDestination.
// Experimental.
type IConfigurationSetEventDestinationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConfigurationSetEventDestination resource.
	// Experimental.
	ConfigurationSetEventDestinationRef() *ConfigurationSetEventDestinationReference
}

// The jsii proxy for IConfigurationSetEventDestinationRef
type jsiiProxy_IConfigurationSetEventDestinationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IConfigurationSetEventDestinationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IConfigurationSetEventDestinationRef) ConfigurationSetEventDestinationRef() *ConfigurationSetEventDestinationReference {
	var returns *ConfigurationSetEventDestinationReference
	_jsii_.Get(
		j,
		"configurationSetEventDestinationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationSetEventDestinationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationSetEventDestinationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

