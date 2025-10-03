package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolIdentityProvider.
// Experimental.
type IUserPoolIdentityProviderRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolIdentityProviderRef
type jsiiProxy_IUserPoolIdentityProviderRef struct {
	internal.Type__constructsIConstruct
}

