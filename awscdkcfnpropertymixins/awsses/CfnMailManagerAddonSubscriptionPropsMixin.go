package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a subscription for an Add On representing the acceptance of its terms of use and additional pricing.
//
// The subscription can then be used to create an instance for use in rule sets or traffic policies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMailManagerAddonSubscriptionPropsMixin := awscdkcfnpropertymixins.Aws_ses.NewCfnMailManagerAddonSubscriptionPropsMixin(&CfnMailManagerAddonSubscriptionMixinProps{
//   	AddonName: jsii.String("addonName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageraddonsubscription.html
//
type CfnMailManagerAddonSubscriptionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMailManagerAddonSubscriptionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMailManagerAddonSubscriptionPropsMixin
type jsiiProxy_CfnMailManagerAddonSubscriptionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMailManagerAddonSubscriptionPropsMixin) Props() *CfnMailManagerAddonSubscriptionMixinProps {
	var returns *CfnMailManagerAddonSubscriptionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMailManagerAddonSubscriptionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::MailManagerAddonSubscription`.
func NewCfnMailManagerAddonSubscriptionPropsMixin(props *CfnMailManagerAddonSubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMailManagerAddonSubscriptionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMailManagerAddonSubscriptionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMailManagerAddonSubscriptionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddonSubscriptionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::MailManagerAddonSubscription`.
func NewCfnMailManagerAddonSubscriptionPropsMixin_Override(c CfnMailManagerAddonSubscriptionPropsMixin, props *CfnMailManagerAddonSubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddonSubscriptionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMailManagerAddonSubscriptionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMailManagerAddonSubscriptionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddonSubscriptionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMailManagerAddonSubscriptionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerAddonSubscriptionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMailManagerAddonSubscriptionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMailManagerAddonSubscriptionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

