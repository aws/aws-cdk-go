package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes Infrastructure Performance subscriptions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNetworkPerformanceMetricSubscriptionPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnNetworkPerformanceMetricSubscriptionPropsMixin(&CfnNetworkPerformanceMetricSubscriptionMixinProps{
//   	Destination: jsii.String("destination"),
//   	Metric: jsii.String("metric"),
//   	Source: jsii.String("source"),
//   	Statistic: jsii.String("statistic"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkperformancemetricsubscription.html
//
type CfnNetworkPerformanceMetricSubscriptionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNetworkPerformanceMetricSubscriptionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkPerformanceMetricSubscriptionPropsMixin
type jsiiProxy_CfnNetworkPerformanceMetricSubscriptionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNetworkPerformanceMetricSubscriptionPropsMixin) Props() *CfnNetworkPerformanceMetricSubscriptionMixinProps {
	var returns *CfnNetworkPerformanceMetricSubscriptionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkPerformanceMetricSubscriptionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::NetworkPerformanceMetricSubscription`.
func NewCfnNetworkPerformanceMetricSubscriptionPropsMixin(props *CfnNetworkPerformanceMetricSubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNetworkPerformanceMetricSubscriptionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkPerformanceMetricSubscriptionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkPerformanceMetricSubscriptionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnNetworkPerformanceMetricSubscriptionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::NetworkPerformanceMetricSubscription`.
func NewCfnNetworkPerformanceMetricSubscriptionPropsMixin_Override(c CfnNetworkPerformanceMetricSubscriptionPropsMixin, props *CfnNetworkPerformanceMetricSubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnNetworkPerformanceMetricSubscriptionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNetworkPerformanceMetricSubscriptionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkPerformanceMetricSubscriptionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnNetworkPerformanceMetricSubscriptionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkPerformanceMetricSubscriptionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnNetworkPerformanceMetricSubscriptionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkPerformanceMetricSubscriptionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNetworkPerformanceMetricSubscriptionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

