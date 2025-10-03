package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolResourceServer.
// Experimental.
type IUserPoolResourceServerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolResourceServerRef
type jsiiProxy_IUserPoolResourceServerRef struct {
	internal.Type__constructsIConstruct
}

