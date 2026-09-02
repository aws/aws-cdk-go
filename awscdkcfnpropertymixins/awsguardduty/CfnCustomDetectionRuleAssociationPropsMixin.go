package awsguardduty

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::GuardDuty::CustomDetectionRuleAssociation.
//
// Associates a GuardDuty custom detection rule with the caller's account, enabling the rule in either LIVE or DRY_RUN mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCustomDetectionRuleAssociationPropsMixin := awscdkcfnpropertymixins.Aws_guardduty.NewCfnCustomDetectionRuleAssociationPropsMixin(&CfnCustomDetectionRuleAssociationMixinProps{
//   	Mode: jsii.String("mode"),
//   	RuleId: jsii.String("ruleId"),
//   	Tags: []TagItemProperty{
//   		&TagItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-customdetectionruleassociation.html
//
type CfnCustomDetectionRuleAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCustomDetectionRuleAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCustomDetectionRuleAssociationPropsMixin
type jsiiProxy_CfnCustomDetectionRuleAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCustomDetectionRuleAssociationPropsMixin) Props() *CfnCustomDetectionRuleAssociationMixinProps {
	var returns *CfnCustomDetectionRuleAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomDetectionRuleAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GuardDuty::CustomDetectionRuleAssociation`.
func NewCfnCustomDetectionRuleAssociationPropsMixin(props *CfnCustomDetectionRuleAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCustomDetectionRuleAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCustomDetectionRuleAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomDetectionRuleAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_guardduty.CfnCustomDetectionRuleAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GuardDuty::CustomDetectionRuleAssociation`.
func NewCfnCustomDetectionRuleAssociationPropsMixin_Override(c CfnCustomDetectionRuleAssociationPropsMixin, props *CfnCustomDetectionRuleAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_guardduty.CfnCustomDetectionRuleAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCustomDetectionRuleAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomDetectionRuleAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_guardduty.CfnCustomDetectionRuleAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomDetectionRuleAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_guardduty.CfnCustomDetectionRuleAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomDetectionRuleAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCustomDetectionRuleAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

