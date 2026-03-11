package awssecurityhub

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SecurityHub::ProductSubscription` resource creates a subscription to a third-party product that generates findings that you want to receive in AWS Security Hub CSPM .
//
// For a list of integrations to third-party products, see [Available third-party partner product integrations](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-partner-providers.html) in the *AWS Security Hub CSPM User Guide* .
//
// To change a product subscription, remove the current product subscription resource, and then create a new one.
//
// Tags aren't supported for this resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnProductSubscriptionPropsMixin := awscdkcfnpropertymixins.Aws_securityhub.NewCfnProductSubscriptionPropsMixin(&CfnProductSubscriptionMixinProps{
//   	ProductArn: jsii.String("productArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-productsubscription.html
//
type CfnProductSubscriptionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnProductSubscriptionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProductSubscriptionPropsMixin
type jsiiProxy_CfnProductSubscriptionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnProductSubscriptionPropsMixin) Props() *CfnProductSubscriptionMixinProps {
	var returns *CfnProductSubscriptionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProductSubscriptionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityHub::ProductSubscription`.
func NewCfnProductSubscriptionPropsMixin(props *CfnProductSubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnProductSubscriptionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProductSubscriptionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProductSubscriptionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnProductSubscriptionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityHub::ProductSubscription`.
func NewCfnProductSubscriptionPropsMixin_Override(c CfnProductSubscriptionPropsMixin, props *CfnProductSubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnProductSubscriptionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnProductSubscriptionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProductSubscriptionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnProductSubscriptionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProductSubscriptionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_securityhub.CfnProductSubscriptionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProductSubscriptionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProductSubscriptionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

