package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityProviderConfig.
// Experimental.
type IIdentityProviderConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdentityProviderConfigRef
type jsiiProxy_IIdentityProviderConfigRef struct {
	internal.Type__constructsIConstruct
}

