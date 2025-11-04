package awsservicecatalogappregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AttributeGroupAssociation.
// Experimental.
type IAttributeGroupAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AttributeGroupAssociation resource.
	// Experimental.
	AttributeGroupAssociationRef() *AttributeGroupAssociationReference
}

// The jsii proxy for IAttributeGroupAssociationRef
type jsiiProxy_IAttributeGroupAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAttributeGroupAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAttributeGroupAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

