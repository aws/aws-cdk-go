// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog AppRegistry Application.
// Experimental.
type IApplication interface {
	awscdk.IResource
	// Create an attribute group and associate this application with the created attribute group.
	// Experimental.
	AddAttributeGroup(id *string, attributeGroupProps *AttributeGroupAssociationProps) IAttributeGroup
	// Associate this application with all stacks under the construct node.
	//
	// NOTE: This method won't automatically register stacks under pipeline stages,
	// and requires association of each pipeline stage by calling this method with stage Construct.
	// Experimental.
	AssociateAllStacksInScope(construct constructs.Construct)
	// Associate a Cloudformation statck with the application in the given stack.
	// Experimental.
	AssociateApplicationWithStack(stack awscdk.Stack)
	// Associate this application with an attribute group.
	// Experimental.
	AssociateAttributeGroup(attributeGroup IAttributeGroup)
	// Associate this application with a CloudFormation stack.
	// Deprecated: Use `associateApplicationWithStack` instead.
	AssociateStack(stack awscdk.Stack)
	// Share this application with other IAM entities, accounts, or OUs.
	// Experimental.
	ShareApplication(id *string, shareOptions *ShareOptions)
	// The ARN of the application.
	// Experimental.
	ApplicationArn() *string
	// The ID of the application.
	// Experimental.
	ApplicationId() *string
	// The name of the application.
	// Experimental.
	ApplicationName() *string
}

// The jsii proxy for IApplication
type jsiiProxy_IApplication struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApplication) AddAttributeGroup(id *string, attributeGroupProps *AttributeGroupAssociationProps) IAttributeGroup {
	if err := i.validateAddAttributeGroupParameters(id, attributeGroupProps); err != nil {
		panic(err)
	}
	var returns IAttributeGroup

	_jsii_.Invoke(
		i,
		"addAttributeGroup",
		[]interface{}{id, attributeGroupProps},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) AssociateAllStacksInScope(construct constructs.Construct) {
	if err := i.validateAssociateAllStacksInScopeParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"associateAllStacksInScope",
		[]interface{}{construct},
	)
}

func (i *jsiiProxy_IApplication) AssociateApplicationWithStack(stack awscdk.Stack) {
	if err := i.validateAssociateApplicationWithStackParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"associateApplicationWithStack",
		[]interface{}{stack},
	)
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

func (i *jsiiProxy_IApplication) ShareApplication(id *string, shareOptions *ShareOptions) {
	if err := i.validateShareApplicationParameters(id, shareOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"shareApplication",
		[]interface{}{id, shareOptions},
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

func (j *jsiiProxy_IApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

