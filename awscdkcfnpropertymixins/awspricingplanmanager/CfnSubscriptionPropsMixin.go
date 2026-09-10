package awspricingplanmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspricingplanmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::PricingPlanManager::Subscription.
//
// Deleting an activated subscription does not terminate it immediately; it schedules a cancellation that takes effect at the end of the current billing period. Until that date the subscription remains active and billing continues. Deleting a subscription that has not yet been activated (PENDING_APPROVAL status) removes it immediately with no further charges.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSubscriptionPropsMixin := awscdkcfnpropertymixins.Aws_pricingplanmanager.NewCfnSubscriptionPropsMixin(&CfnSubscriptionMixinProps{
//   	PlanFamily: jsii.String("planFamily"),
//   	PlanTier: jsii.String("planTier"),
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	UsageLevel: jsii.String("usageLevel"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pricingplanmanager-subscription.html
//
type CfnSubscriptionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSubscriptionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSubscriptionPropsMixin
type jsiiProxy_CfnSubscriptionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSubscriptionPropsMixin) Props() *CfnSubscriptionMixinProps {
	var returns *CfnSubscriptionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriptionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::PricingPlanManager::Subscription`.
func NewCfnSubscriptionPropsMixin(props *CfnSubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSubscriptionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSubscriptionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSubscriptionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_pricingplanmanager.CfnSubscriptionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::PricingPlanManager::Subscription`.
func NewCfnSubscriptionPropsMixin_Override(c CfnSubscriptionPropsMixin, props *CfnSubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_pricingplanmanager.CfnSubscriptionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSubscriptionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubscriptionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_pricingplanmanager.CfnSubscriptionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriptionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_pricingplanmanager.CfnSubscriptionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriptionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSubscriptionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

