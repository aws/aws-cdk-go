package previewawscloudformationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudformation/previewawscloudformationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudFormation::HookTypeConfig` resource specifies the configuration of an activated Hook.
//
// For information about the CloudFormation registry, see [Managing extensions with the CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html) in the *CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHookTypeConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnHookTypeConfigPropsMixin(&CfnHookTypeConfigMixinProps{
//   	Configuration: jsii.String("configuration"),
//   	ConfigurationAlias: jsii.String("configurationAlias"),
//   	TypeArn: jsii.String("typeArn"),
//   	TypeName: jsii.String("typeName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hooktypeconfig.html
//
type CfnHookTypeConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnHookTypeConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHookTypeConfigPropsMixin
type jsiiProxy_CfnHookTypeConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnHookTypeConfigPropsMixin) Props() *CfnHookTypeConfigMixinProps {
	var returns *CfnHookTypeConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHookTypeConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::HookTypeConfig`.
func NewCfnHookTypeConfigPropsMixin(props *CfnHookTypeConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnHookTypeConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHookTypeConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHookTypeConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnHookTypeConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::HookTypeConfig`.
func NewCfnHookTypeConfigPropsMixin_Override(c CfnHookTypeConfigPropsMixin, props *CfnHookTypeConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnHookTypeConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnHookTypeConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHookTypeConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnHookTypeConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHookTypeConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnHookTypeConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHookTypeConfigPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnHookTypeConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

