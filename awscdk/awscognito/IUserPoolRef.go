package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPool.
// Experimental.
type IUserPoolRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolRef
type jsiiProxy_IUserPoolRef struct {
	internal.Type__constructsIConstruct
}

