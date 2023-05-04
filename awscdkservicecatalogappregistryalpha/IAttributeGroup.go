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
	// Associate an application with attribute group If the attribute group is already associated, it will ignore duplicate request.
	// Experimental.
	AssociateWith(application IApplication)
	// Share the attribute group resource with other IAM entities, accounts, or OUs.
	// Experimental.
	ShareAttributeGroup(id *string, shareOptions *ShareOptions)
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

func (i *jsiiProxy_IAttributeGroup) AssociateWith(application IApplication) {
	if err := i.validateAssociateWithParameters(application); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"associateWith",
		[]interface{}{application},
	)
}

func (i *jsiiProxy_IAttributeGroup) ShareAttributeGroup(id *string, shareOptions *ShareOptions) {
	if err := i.validateShareAttributeGroupParameters(id, shareOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"shareAttributeGroup",
		[]interface{}{id, shareOptions},
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

