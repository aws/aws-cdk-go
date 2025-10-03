package awswafregional

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafregional/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebACLAssociation.
// Experimental.
type IWebACLAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWebACLAssociationRef
type jsiiProxy_IWebACLAssociationRef struct {
	internal.Type__constructsIConstruct
}

