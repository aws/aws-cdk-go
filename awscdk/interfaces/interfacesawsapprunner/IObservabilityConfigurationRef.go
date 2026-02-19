package interfacesawsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ObservabilityConfiguration.
// Experimental.
type IObservabilityConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ObservabilityConfiguration resource.
	// Experimental.
	ObservabilityConfigurationRef() *ObservabilityConfigurationReference
}

// The jsii proxy for IObservabilityConfigurationRef
type jsiiProxy_IObservabilityConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IObservabilityConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IObservabilityConfigurationRef) ObservabilityConfigurationRef() *ObservabilityConfigurationReference {
	var returns *ObservabilityConfigurationReference
	_jsii_.Get(
		j,
		"observabilityConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IObservabilityConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IObservabilityConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

