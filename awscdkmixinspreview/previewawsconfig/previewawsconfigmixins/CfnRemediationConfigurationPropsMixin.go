package previewawsconfigmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconfig/previewawsconfigmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An object that represents the details about the remediation configuration that includes the remediation action, parameters, and data to execute the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//
//   cfnRemediationConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnRemediationConfigurationPropsMixin(&CfnRemediationConfigurationMixinProps{
//   	Automatic: jsii.Boolean(false),
//   	ConfigRuleName: jsii.String("configRuleName"),
//   	ExecutionControls: &ExecutionControlsProperty{
//   		SsmControls: &SsmControlsProperty{
//   			ConcurrentExecutionRatePercentage: jsii.Number(123),
//   			ErrorPercentage: jsii.Number(123),
//   		},
//   	},
//   	MaximumAutomaticAttempts: jsii.Number(123),
//   	Parameters: parameters,
//   	ResourceType: jsii.String("resourceType"),
//   	RetryAttemptSeconds: jsii.Number(123),
//   	TargetId: jsii.String("targetId"),
//   	TargetType: jsii.String("targetType"),
//   	TargetVersion: jsii.String("targetVersion"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-remediationconfiguration.html
//
type CfnRemediationConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRemediationConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRemediationConfigurationPropsMixin
type jsiiProxy_CfnRemediationConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRemediationConfigurationPropsMixin) Props() *CfnRemediationConfigurationMixinProps {
	var returns *CfnRemediationConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRemediationConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Config::RemediationConfiguration`.
func NewCfnRemediationConfigurationPropsMixin(props *CfnRemediationConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRemediationConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRemediationConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRemediationConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnRemediationConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Config::RemediationConfiguration`.
func NewCfnRemediationConfigurationPropsMixin_Override(c CfnRemediationConfigurationPropsMixin, props *CfnRemediationConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnRemediationConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRemediationConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRemediationConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnRemediationConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRemediationConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnRemediationConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRemediationConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRemediationConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

