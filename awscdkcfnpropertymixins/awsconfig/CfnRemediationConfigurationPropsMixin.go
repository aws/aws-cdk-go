package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An object that represents the details about the remediation configuration that includes the remediation action, parameters, and data to execute the action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var parameters interface{}
//
//   cfnRemediationConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_config.NewCfnRemediationConfigurationPropsMixin(&CfnRemediationConfigurationMixinProps{
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-remediationconfiguration.html
//
type CfnRemediationConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRemediationConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRemediationConfigurationPropsMixin
type jsiiProxy_CfnRemediationConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnRemediationConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Config::RemediationConfiguration`.
func NewCfnRemediationConfigurationPropsMixin(props *CfnRemediationConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRemediationConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRemediationConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRemediationConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Config::RemediationConfiguration`.
func NewCfnRemediationConfigurationPropsMixin_Override(c CfnRemediationConfigurationPropsMixin, props *CfnRemediationConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRemediationConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRemediationConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_config.CfnRemediationConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRemediationConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

