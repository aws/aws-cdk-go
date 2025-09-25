package awsservicecatalogappregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AttributeGroupAssociation.
// Experimental.
type IAttributeGroupAssociationRef interface {
	constructs.IConstruct
	// A reference to a AttributeGroupAssociation resource.
	// Experimental.
	AttributeGroupAssociationRef() *AttributeGroupAssociationReference
}

// The jsii proxy for IAttributeGroupAssociationRef
type jsiiProxy_IAttributeGroupAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAttributeGroupAssociationRef) AttributeGroupAssociationRef() *AttributeGroupAssociationReference {
	var returns *AttributeGroupAssociationReference
	_jsii_.Get(
		j,
		"attributeGroupAssociationRef",
		&returns,
	)
	return returns
}

