package awsservicecatalogappregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicecatalogappregistry/internal"
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

