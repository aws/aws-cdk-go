package awswafv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2/internal"
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

