package awsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::Bedrock::EnforcedGuardrailConfiguration Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnEnforcedGuardrailConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_bedrock.NewCfnEnforcedGuardrailConfigurationPropsMixin(&CfnEnforcedGuardrailConfigurationMixinProps{
//   	GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   	GuardrailVersion: jsii.String("guardrailVersion"),
//   	ModelEnforcement: &ModelEnforcementProperty{
//   		ExcludedModels: []*string{
//   			jsii.String("excludedModels"),
//   		},
//   		IncludedModels: []*string{
//   			jsii.String("includedModels"),
//   		},
//   	},
//   	SelectiveContentGuarding: &SelectiveContentGuardingProperty{
//   		Messages: jsii.String("messages"),
//   		System: jsii.String("system"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-enforcedguardrailconfiguration.html
//
type CfnEnforcedGuardrailConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnEnforcedGuardrailConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnforcedGuardrailConfigurationPropsMixin
type jsiiProxy_CfnEnforcedGuardrailConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEnforcedGuardrailConfigurationPropsMixin) Props() *CfnEnforcedGuardrailConfigurationMixinProps {
	var returns *CfnEnforcedGuardrailConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnforcedGuardrailConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Bedrock::EnforcedGuardrailConfiguration`.
func NewCfnEnforcedGuardrailConfigurationPropsMixin(props *CfnEnforcedGuardrailConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnEnforcedGuardrailConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnforcedGuardrailConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnforcedGuardrailConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnEnforcedGuardrailConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Bedrock::EnforcedGuardrailConfiguration`.
func NewCfnEnforcedGuardrailConfigurationPropsMixin_Override(c CfnEnforcedGuardrailConfigurationPropsMixin, props *CfnEnforcedGuardrailConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnEnforcedGuardrailConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEnforcedGuardrailConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnforcedGuardrailConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnEnforcedGuardrailConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnforcedGuardrailConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrock.CfnEnforcedGuardrailConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnforcedGuardrailConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEnforcedGuardrailConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

