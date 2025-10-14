package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityProviderConfig.
// Experimental.
type IIdentityProviderConfigRef interface {
	constructs.IConstruct
	// A reference to a IdentityProviderConfig resource.
	// Experimental.
	IdentityProviderConfigRef() *IdentityProviderConfigReference
}

// The jsii proxy for IIdentityProviderConfigRef
type jsiiProxy_IIdentityProviderConfigRef struct {
	internal.Type__constructsIConstruct
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

