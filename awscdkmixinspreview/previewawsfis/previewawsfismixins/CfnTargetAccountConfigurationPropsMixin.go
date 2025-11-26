package previewawsfismixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsfis/previewawsfismixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a target account configuration for the experiment template.
//
// A target account configuration is required when `accountTargeting` of `experimentOptions` is set to `multi-account` . For more information, see [experiment options](https://docs.aws.amazon.com/fis/latest/userguide/experiment-options.html) in the *AWS Fault Injection Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTargetAccountConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnTargetAccountConfigurationPropsMixin(&CfnTargetAccountConfigurationMixinProps{
//   	AccountId: jsii.String("accountId"),
//   	Description: jsii.String("description"),
//   	ExperimentTemplateId: jsii.String("experimentTemplateId"),
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fis-targetaccountconfiguration.html
//
type CfnTargetAccountConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTargetAccountConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTargetAccountConfigurationPropsMixin
type jsiiProxy_CfnTargetAccountConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTargetAccountConfigurationPropsMixin) Props() *CfnTargetAccountConfigurationMixinProps {
	var returns *CfnTargetAccountConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTargetAccountConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::FIS::TargetAccountConfiguration`.
func NewCfnTargetAccountConfigurationPropsMixin(props *CfnTargetAccountConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTargetAccountConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTargetAccountConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTargetAccountConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fis.mixins.CfnTargetAccountConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::FIS::TargetAccountConfiguration`.
func NewCfnTargetAccountConfigurationPropsMixin_Override(c CfnTargetAccountConfigurationPropsMixin, props *CfnTargetAccountConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fis.mixins.CfnTargetAccountConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTargetAccountConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTargetAccountConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_fis.mixins.CfnTargetAccountConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTargetAccountConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_fis.mixins.CfnTargetAccountConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTargetAccountConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTargetAccountConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

