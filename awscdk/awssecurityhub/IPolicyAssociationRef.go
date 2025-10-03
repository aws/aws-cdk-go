package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyAssociation.
// Experimental.
type IPolicyAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPolicyAssociationRef
type jsiiProxy_IPolicyAssociationRef struct {
	internal.Type__constructsIConstruct
}

