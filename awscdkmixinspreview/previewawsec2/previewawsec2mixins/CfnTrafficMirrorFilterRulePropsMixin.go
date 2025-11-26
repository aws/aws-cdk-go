package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a Traffic Mirror filter rule.
//
// A Traffic Mirror rule defines the Traffic Mirror source traffic to mirror.
//
// You need the Traffic Mirror filter ID when you create the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTrafficMirrorFilterRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnTrafficMirrorFilterRulePropsMixin(&CfnTrafficMirrorFilterRuleMixinProps{
//   	Description: jsii.String("description"),
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	DestinationPortRange: &TrafficMirrorPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	Protocol: jsii.Number(123),
//   	RuleAction: jsii.String("ruleAction"),
//   	RuleNumber: jsii.Number(123),
//   	SourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	SourcePortRange: &TrafficMirrorPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrafficDirection: jsii.String("trafficDirection"),
//   	TrafficMirrorFilterId: jsii.String("trafficMirrorFilterId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html
//
type CfnTrafficMirrorFilterRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTrafficMirrorFilterRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTrafficMirrorFilterRulePropsMixin
type jsiiProxy_CfnTrafficMirrorFilterRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTrafficMirrorFilterRulePropsMixin) Props() *CfnTrafficMirrorFilterRuleMixinProps {
	var returns *CfnTrafficMirrorFilterRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrafficMirrorFilterRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TrafficMirrorFilterRule`.
func NewCfnTrafficMirrorFilterRulePropsMixin(props *CfnTrafficMirrorFilterRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTrafficMirrorFilterRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTrafficMirrorFilterRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrafficMirrorFilterRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTrafficMirrorFilterRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TrafficMirrorFilterRule`.
func NewCfnTrafficMirrorFilterRulePropsMixin_Override(c CfnTrafficMirrorFilterRulePropsMixin, props *CfnTrafficMirrorFilterRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTrafficMirrorFilterRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTrafficMirrorFilterRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrafficMirrorFilterRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTrafficMirrorFilterRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrafficMirrorFilterRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTrafficMirrorFilterRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrafficMirrorFilterRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTrafficMirrorFilterRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

