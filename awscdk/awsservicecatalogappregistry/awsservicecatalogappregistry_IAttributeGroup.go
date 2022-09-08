package awsservicecatalogappregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicecatalogappregistry/internal"
)

// A Service Catalog AppRegistry Attribute Group.
// Experimental.
type IAttributeGroup interface {
	awscdk.IResource
	// The ARN of the attribute group.
	// Experimental.
	AttributeGroupArn() *string
	// The ID of the attribute group.
	// Experimental.
	AttributeGroupId() *string
}

// The jsii proxy for IAttributeGroup
type jsiiProxy_IAttributeGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAttributeGroup) AttributeGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAttributeGroup) AttributeGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeGroupId",
		&returns,
	)
	return returns
}

