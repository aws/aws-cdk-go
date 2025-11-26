package previewawscloudformationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudformation/previewawscloudformationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudFormation::HookDefaultVersion` resource specifies the default version of a Hook.
//
// The default version of the Hook is used in CloudFormation operations for this AWS account and AWS Region .
//
// For information about the CloudFormation registry, see [Managing extensions with the CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html) in the *CloudFormation User Guide* .
//
// This resource type is not compatible with Guard and Lambda Hooks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHookDefaultVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnHookDefaultVersionPropsMixin(&CfnHookDefaultVersionMixinProps{
//   	TypeName: jsii.String("typeName"),
//   	TypeVersionArn: jsii.String("typeVersionArn"),
//   	VersionId: jsii.String("versionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookdefaultversion.html
//
type CfnHookDefaultVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnHookDefaultVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHookDefaultVersionPropsMixin
type jsiiProxy_CfnHookDefaultVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnHookDefaultVersionPropsMixin) Props() *CfnHookDefaultVersionMixinProps {
	var returns *CfnHookDefaultVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookDefaultVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::HookDefaultVersion`.
func NewCfnHookDefaultVersionPropsMixin(props *CfnHookDefaultVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnHookDefaultVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHookDefaultVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHookDefaultVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnHookDefaultVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::HookDefaultVersion`.
func NewCfnHookDefaultVersionPropsMixin_Override(c CfnHookDefaultVersionPropsMixin, props *CfnHookDefaultVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnHookDefaultVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnHookDefaultVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHookDefaultVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnHookDefaultVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHookDefaultVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnHookDefaultVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHookDefaultVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnHookDefaultVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

