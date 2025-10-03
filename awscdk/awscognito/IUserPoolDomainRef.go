package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolDomain.
// Experimental.
type IUserPoolDomainRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolDomainRef
type jsiiProxy_IUserPoolDomainRef struct {
	internal.Type__constructsIConstruct
}

