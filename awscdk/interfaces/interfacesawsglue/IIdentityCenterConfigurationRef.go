package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityCenterConfiguration.
// Experimental.
type IIdentityCenterConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IdentityCenterConfiguration resource.
	// Experimental.
	IdentityCenterConfigurationRef() *IdentityCenterConfigurationReference
}

// The jsii proxy for IIdentityCenterConfigurationRef
type jsiiProxy_IIdentityCenterConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IIdentityCenterConfigurationRef) IdentityCenterConfigurationRef() *IdentityCenterConfigurationReference {
	var returns *IdentityCenterConfigurationReference
	_jsii_.Get(
		j,
		"identityCenterConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityCenterConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityCenterConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

