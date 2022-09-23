// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/internal"
)

// A Service Catalog AppRegistry Application.
// Experimental.
type IApplication interface {
	awscdk.IResource
	// Associate thisapplication with an attribute group.
	// Experimental.
	AssociateAttributeGroup(attributeGroup IAttributeGroup)
	// Associate this application with a CloudFormation stack.
	// Experimental.
	AssociateStack(stack awscdk.Stack)
	// Share this application with other IAM entities, accounts, or OUs.
	// Experimental.
	ShareApplication(shareOptions *ShareOptions)
	// The ARN of the application.
	// Experimental.
	ApplicationArn() *string
	// The ID of the application.
	// Experimental.
	ApplicationId() *string
}

// The jsii proxy for IApplication
type jsiiProxy_IApplication struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApplication) AssociateAttributeGroup(attributeGroup IAttributeGroup) {
	if err := i.validateAssociateAttributeGroupParameters(attributeGroup); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"associateAttributeGroup",
		[]interface{}{attributeGroup},
	)
}

func (i *jsiiProxy_IApplication) AssociateStack(stack awscdk.Stack) {
	if err := i.validateAssociateStackParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"associateStack",
		[]interface{}{stack},
	)
}

func (i *jsiiProxy_IApplication) ShareApplication(shareOptions *ShareOptions) {
	if err := i.validateShareApplicationParameters(shareOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"shareApplication",
		[]interface{}{shareOptions},
	)
}

func (j *jsiiProxy_IApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

