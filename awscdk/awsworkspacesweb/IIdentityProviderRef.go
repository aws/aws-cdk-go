package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityProvider.
// Experimental.
type IIdentityProviderRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdentityProviderRef
type jsiiProxy_IIdentityProviderRef struct {
	internal.Type__constructsIConstruct
}

