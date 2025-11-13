package interfacesawseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityProviderConfig.
// Experimental.
type IIdentityProviderConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IdentityProviderConfig resource.
	// Experimental.
	IdentityProviderConfigRef() *IdentityProviderConfigReference
}

// The jsii proxy for IIdentityProviderConfigRef
type jsiiProxy_IIdentityProviderConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IIdentityProviderConfigRef) IdentityProviderConfigRef() *IdentityProviderConfigReference {
	var returns *IdentityProviderConfigReference
	_jsii_.Get(
		j,
		"identityProviderConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityProviderConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityProviderConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

