package mixinsawscloudformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscloudformation/mixinsawscloudformation/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudFormation::Stack` resource nests a stack as a resource in a top-level template.
//
// For more information, see [Nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) in the *CloudFormation User Guide* .
//
// You can add output values from a nested stack within the containing template. You use the [GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-getatt.html) function with the nested stack's logical name and the name of the output value in the nested stack in the format `Outputs. *NestedStackOutputName*` .
//
// We strongly recommend that updates to nested stacks are run from the parent stack.
//
// When you apply template changes to update a top-level stack, CloudFormation updates the top-level stack and initiates an update to its nested stacks. CloudFormation updates the resources of modified nested stacks, but doesn't update the resources of unmodified nested stacks.
//
// For stacks that contain IAM resources, you must acknowledge IAM capabilities. Also, make sure that you have cancel update stack permissions, which are required if an update rolls back. For more information about IAM and CloudFormation , see [Controlling access with AWS Identity and Access Management](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/control-access-with-iam.html) in the *CloudFormation User Guide* .
//
// > A subset of `AWS::CloudFormation::Stack` resource type properties listed below are available to customers using CloudFormation , AWS CDK , and Cloud Control  to configure.
// >
// > - `NotificationARNs`
// > - `Parameters`
// > - `Tags`
// > - `TemplateURL`
// > - `TimeoutInMinutes`
// >
// > These properties can be configured only when using Cloud Control  . This is because the below properties are set by the parent stack, and thus cannot be configured using CloudFormation or AWS CDK but only Cloud Control  .
// >
// > - `Capabilities`
// > - `Description`
// > - `DisableRollback`
// > - `EnableTerminationProtection`
// > - `RoleARN`
// > - `StackName`
// > - `StackPolicyBody`
// > - `StackPolicyURL`
// > - `StackStatusReason`
// > - `TemplateBody`
// >
// > Customers that configure `AWS::CloudFormation::Stack` using CloudFormation and AWS CDK can do so for nesting a CloudFormation stack as a resource in their top-level template.
// >
// > These read-only properties can be accessed only when using Cloud Control  .
// >
// > - `ChangeSetId`
// > - `CreationTime`
// > - `LastUpdateTime`
// > - `Outputs`
// > - `ParentId`
// > - `RootId`
// > - `StackId`
// > - `StackStatus`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStackPropsMixin := awscdkmixinspreview.Mixins.NewCfnStackPropsMixin(&CfnStackMixinProps{
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateUrl: jsii.String("templateUrl"),
//   	TimeoutInMinutes: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-stack.html
//
type CfnStackPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStackMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStackPropsMixin
type jsiiProxy_CfnStackPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStackPropsMixin) Props() *CfnStackMixinProps {
	var returns *CfnStackMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::Stack`.
func NewCfnStackPropsMixin(props *CfnStackMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStackPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStackPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStackPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnStackPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::Stack`.
func NewCfnStackPropsMixin_Override(c CfnStackPropsMixin, props *CfnStackMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnStackPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStackPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStackPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnStackPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStackPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnStackPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStackPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

