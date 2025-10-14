package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityProvider.
// Experimental.
type IIdentityProviderRef interface {
	constructs.IConstruct
	// A reference to a IdentityProvider resource.
	// Experimental.
	IdentityProviderRef() *IdentityProviderReference
}

// The jsii proxy for IIdentityProviderRef
type jsiiProxy_IIdentityProviderRef struct {
	internal.Type__constructsIConstruct
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

