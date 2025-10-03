package awsservicecatalogappregistry

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AttributeGroupAssociation.
// Experimental.
type IAttributeGroupAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAttributeGroupAssociationRef
type jsiiProxy_IAttributeGroupAssociationRef struct {
	internal.Type__constructsIConstruct
}

