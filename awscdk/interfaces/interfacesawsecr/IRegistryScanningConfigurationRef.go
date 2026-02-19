package interfacesawsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegistryScanningConfiguration.
// Experimental.
type IRegistryScanningConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RegistryScanningConfiguration resource.
	// Experimental.
	RegistryScanningConfigurationRef() *RegistryScanningConfigurationReference
}

// The jsii proxy for IRegistryScanningConfigurationRef
type jsiiProxy_IRegistryScanningConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRegistryScanningConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRegistryScanningConfigurationRef) RegistryScanningConfigurationRef() *RegistryScanningConfigurationReference {
	var returns *RegistryScanningConfigurationReference
	_jsii_.Get(
		j,
		"registryScanningConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRegistryScanningConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRegistryScanningConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

