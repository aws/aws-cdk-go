package interfacesawsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityProvider.
// Experimental.
type IIdentityProviderRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IdentityProvider resource.
	// Experimental.
	IdentityProviderRef() *IdentityProviderReference
}

// The jsii proxy for IIdentityProviderRef
type jsiiProxy_IIdentityProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIdentityProviderRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIdentityProviderRef) IdentityProviderRef() *IdentityProviderReference {
	var returns *IdentityProviderReference
	_jsii_.Get(
		j,
		"identityProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityProviderRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

