package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolUser.
// Experimental.
type IUserPoolUserRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolUserRef
type jsiiProxy_IUserPoolUserRef struct {
	internal.Type__constructsIConstruct
}

