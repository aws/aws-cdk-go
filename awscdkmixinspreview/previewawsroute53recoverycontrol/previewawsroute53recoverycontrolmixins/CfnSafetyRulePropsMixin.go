package previewawsroute53recoverycontrolmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53recoverycontrol/previewawsroute53recoverycontrolmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a safety rule in a control panel in Amazon Route 53 Application Recovery Controller.
//
// Safety rules in Amazon Route 53 Application Recovery Controller let you add safeguards around changing routing control states, and enabling and disabling routing controls, to help prevent unwanted outcomes. Note that the name of a safety rule must be unique within a control panel.
//
// There are two types of safety rules in Route 53 ARC: assertion rules and gating rules.
//
// Assertion rule: An assertion rule enforces that, when you change a routing control state, certain criteria are met. For example, the criteria might be that at least one routing control state is `On` after the transaction completes so that traffic continues to be directed to at least one cell for the application. This prevents a fail-open scenario.
//
// Gating rule: A gating rule lets you configure a gating routing control as an overall on-off switch for a group of routing controls. Or, you can configure more complex gating scenarios, for example, by configuring multiple gating routing controls.
//
// For more information, see [Safety rules](https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.safety-rules.html) in the Amazon Route 53 Application Recovery Controller Developer Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSafetyRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnSafetyRulePropsMixin(&CfnSafetyRuleMixinProps{
//   	AssertionRule: &AssertionRuleProperty{
//   		AssertedControls: []*string{
//   			jsii.String("assertedControls"),
//   		},
//   		WaitPeriodMs: jsii.Number(123),
//   	},
//   	ControlPanelArn: jsii.String("controlPanelArn"),
//   	GatingRule: &GatingRuleProperty{
//   		GatingControls: []*string{
//   			jsii.String("gatingControls"),
//   		},
//   		TargetControls: []*string{
//   			jsii.String("targetControls"),
//   		},
//   		WaitPeriodMs: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	RuleConfig: &RuleConfigProperty{
//   		Inverted: jsii.Boolean(false),
//   		Threshold: jsii.Number(123),
//   		Type: jsii.String("type"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-safetyrule.html
//
type CfnSafetyRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSafetyRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSafetyRulePropsMixin
type jsiiProxy_CfnSafetyRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSafetyRulePropsMixin) Props() *CfnSafetyRuleMixinProps {
	var returns *CfnSafetyRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSafetyRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53RecoveryControl::SafetyRule`.
func NewCfnSafetyRulePropsMixin(props *CfnSafetyRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSafetyRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSafetyRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSafetyRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoverycontrol.mixins.CfnSafetyRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53RecoveryControl::SafetyRule`.
func NewCfnSafetyRulePropsMixin_Override(c CfnSafetyRulePropsMixin, props *CfnSafetyRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoverycontrol.mixins.CfnSafetyRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSafetyRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSafetyRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoverycontrol.mixins.CfnSafetyRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSafetyRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53recoverycontrol.mixins.CfnSafetyRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSafetyRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSafetyRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

