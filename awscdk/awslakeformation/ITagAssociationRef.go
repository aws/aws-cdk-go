package awslakeformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagAssociation.
// Experimental.
type ITagAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITagAssociationRef
type jsiiProxy_ITagAssociationRef struct {
	internal.Type__constructsIConstruct
}

