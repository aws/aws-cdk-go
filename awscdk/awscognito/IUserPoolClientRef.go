package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolClient.
// Experimental.
type IUserPoolClientRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolClientRef
type jsiiProxy_IUserPoolClientRef struct {
	internal.Type__constructsIConstruct
}

