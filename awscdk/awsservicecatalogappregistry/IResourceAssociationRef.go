package awsservicecatalogappregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceAssociation.
// Experimental.
type IResourceAssociationRef interface {
	constructs.IConstruct
	// A reference to a ResourceAssociation resource.
	// Experimental.
	ResourceAssociationRef() *ResourceAssociationReference
}

// The jsii proxy for IResourceAssociationRef
type jsiiProxy_IResourceAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceAssociationRef) ResourceAssociationRef() *ResourceAssociationReference {
	var returns *ResourceAssociationReference
	_jsii_.Get(
		j,
		"resourceAssociationRef",
		&returns,
	)
	return returns
}

