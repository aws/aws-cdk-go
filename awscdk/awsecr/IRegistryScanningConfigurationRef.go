package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegistryScanningConfiguration.
// Experimental.
type IRegistryScanningConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RegistryScanningConfiguration resource.
	// Experimental.
	RegistryScanningConfigurationRef() *RegistryScanningConfigurationReference
}

// The jsii proxy for IRegistryScanningConfigurationRef
type jsiiProxy_IRegistryScanningConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRegistryScanningConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

