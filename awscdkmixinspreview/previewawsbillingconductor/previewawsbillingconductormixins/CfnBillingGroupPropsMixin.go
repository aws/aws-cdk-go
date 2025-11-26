package previewawsbillingconductormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbillingconductor/previewawsbillingconductormixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a billing group that resembles a consolidated billing family that AWS charges, based off of the predefined pricing plan computation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBillingGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnBillingGroupPropsMixin(&CfnBillingGroupMixinProps{
//   	AccountGrouping: &AccountGroupingProperty{
//   		AutoAssociate: jsii.Boolean(false),
//   		LinkedAccountIds: []*string{
//   			jsii.String("linkedAccountIds"),
//   		},
//   		ResponsibilityTransferArn: jsii.String("responsibilityTransferArn"),
//   	},
//   	ComputationPreference: &ComputationPreferenceProperty{
//   		PricingPlanArn: jsii.String("pricingPlanArn"),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	PrimaryAccountId: jsii.String("primaryAccountId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-billingconductor-billinggroup.html
//
type CfnBillingGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBillingGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBillingGroupPropsMixin
type jsiiProxy_CfnBillingGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBillingGroupPropsMixin) Props() *CfnBillingGroupMixinProps {
	var returns *CfnBillingGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBillingGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BillingConductor::BillingGroup`.
func NewCfnBillingGroupPropsMixin(props *CfnBillingGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBillingGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBillingGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBillingGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnBillingGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BillingConductor::BillingGroup`.
func NewCfnBillingGroupPropsMixin_Override(c CfnBillingGroupPropsMixin, props *CfnBillingGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnBillingGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBillingGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBillingGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnBillingGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBillingGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_billingconductor.mixins.CfnBillingGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBillingGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnBillingGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

