package awsservicecatalogappregistry

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AttributeGroup.
// Experimental.
type IAttributeGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAttributeGroupRef
type jsiiProxy_IAttributeGroupRef struct {
	internal.Type__constructsIConstruct
}

