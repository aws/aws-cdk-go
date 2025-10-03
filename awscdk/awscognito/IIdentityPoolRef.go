package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityPool.
// Experimental.
type IIdentityPoolRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdentityPoolRef
type jsiiProxy_IIdentityPoolRef struct {
	internal.Type__constructsIConstruct
}

