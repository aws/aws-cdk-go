package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolIdentityProvider.
// Experimental.
type IUserPoolIdentityProviderRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a UserPoolIdentityProvider resource.
	// Experimental.
	UserPoolIdentityProviderRef() *UserPoolIdentityProviderReference
}

// The jsii proxy for IUserPoolIdentityProviderRef
type jsiiProxy_IUserPoolIdentityProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IUserPoolIdentityProviderRef) UserPoolIdentityProviderRef() *UserPoolIdentityProviderReference {
	var returns *UserPoolIdentityProviderReference
	_jsii_.Get(
		j,
		"userPoolIdentityProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolIdentityProviderRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolIdentityProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

