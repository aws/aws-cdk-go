// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/internal"
)

// A Service Catalog AppRegistry Attribute Group.
// Experimental.
type IAttributeGroup interface {
	awscdk.IResource
	// Share the attribute group resource with other IAM entities, accounts, or OUs.
	// Experimental.
	ShareAttributeGroup(shareOptions *ShareOptions)
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

func (i *jsiiProxy_IAttributeGroup) ShareAttributeGroup(shareOptions *ShareOptions) {
	if err := i.validateShareAttributeGroupParameters(shareOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"shareAttributeGroup",
		[]interface{}{shareOptions},
	)
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

