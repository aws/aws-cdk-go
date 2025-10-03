package awsservicecatalogappregistry

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceAssociation.
// Experimental.
type IResourceAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceAssociationRef
type jsiiProxy_IResourceAssociationRef struct {
	internal.Type__constructsIConstruct
}

