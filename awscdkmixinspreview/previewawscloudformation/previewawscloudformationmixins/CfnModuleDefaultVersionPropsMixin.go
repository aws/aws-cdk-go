package previewawscloudformationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudformation/previewawscloudformationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the default version of a module.
//
// The default version of the module will be used in CloudFormation operations for this account and Region.
//
// For more information, see [Create reusable resource configurations that can be included across templates with CloudFormation modules](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/modules.html) in the *CloudFormation User Guide* .
//
// For information about the CloudFormation registry, see [Managing extensions with the CloudFormation registry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html) in the *CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnModuleDefaultVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnModuleDefaultVersionPropsMixin(&CfnModuleDefaultVersionMixinProps{
//   	Arn: jsii.String("arn"),
//   	ModuleName: jsii.String("moduleName"),
//   	VersionId: jsii.String("versionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduledefaultversion.html
//
type CfnModuleDefaultVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnModuleDefaultVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnModuleDefaultVersionPropsMixin
type jsiiProxy_CfnModuleDefaultVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnModuleDefaultVersionPropsMixin) Props() *CfnModuleDefaultVersionMixinProps {
	var returns *CfnModuleDefaultVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModuleDefaultVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFormation::ModuleDefaultVersion`.
func NewCfnModuleDefaultVersionPropsMixin(props *CfnModuleDefaultVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnModuleDefaultVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnModuleDefaultVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModuleDefaultVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnModuleDefaultVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFormation::ModuleDefaultVersion`.
func NewCfnModuleDefaultVersionPropsMixin_Override(c CfnModuleDefaultVersionPropsMixin, props *CfnModuleDefaultVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnModuleDefaultVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnModuleDefaultVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModuleDefaultVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnModuleDefaultVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModuleDefaultVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudformation.mixins.CfnModuleDefaultVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModuleDefaultVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnModuleDefaultVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

