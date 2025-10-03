package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityPoolPrincipalTag.
// Experimental.
type IIdentityPoolPrincipalTagRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdentityPoolPrincipalTagRef
type jsiiProxy_IIdentityPoolPrincipalTagRef struct {
	internal.Type__constructsIConstruct
}

