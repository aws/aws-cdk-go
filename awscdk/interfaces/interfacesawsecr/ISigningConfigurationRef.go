package interfacesawsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SigningConfiguration.
// Experimental.
type ISigningConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SigningConfiguration resource.
	// Experimental.
	SigningConfigurationRef() *SigningConfigurationReference
}

// The jsii proxy for ISigningConfigurationRef
type jsiiProxy_ISigningConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISigningConfigurationRef) SigningConfigurationRef() *SigningConfigurationReference {
	var returns *SigningConfigurationReference
	_jsii_.Get(
		j,
		"signingConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

