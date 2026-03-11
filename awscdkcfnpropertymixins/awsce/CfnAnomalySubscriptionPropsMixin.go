package awsce

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CE::AnomalySubscription` resource (also referred to as an alert subscription) is a Cost Explorer resource type that sends notifications about specific anomalies that meet an alerting criteria defined by you.
//
// You can specify the frequency of the alerts and the subscribers to notify.
//
// Anomaly subscriptions can be associated with one or more [`AWS::CE::AnomalyMonitor`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalymonitor.html) resources, and they only send notifications about anomalies detected by those associated monitors. You can also configure a threshold to further control which anomalies are included in the notifications.
//
// Anomalies that don’t exceed the chosen threshold and therefore don’t trigger notifications from an anomaly subscription will still be available on the console and from the [`GetAnomalies`](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetAnomalies.html) API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAnomalySubscriptionPropsMixin := awscdkcfnpropertymixins.Aws_ce.NewCfnAnomalySubscriptionPropsMixin(&CfnAnomalySubscriptionMixinProps{
//   	Frequency: jsii.String("frequency"),
//   	MonitorArnList: []interface{}{
//   		jsii.String("monitorArnList"),
//   	},
//   	ResourceTags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Subscribers: []interface{}{
//   		&SubscriberProperty{
//   			Address: jsii.String("address"),
//   			Status: jsii.String("status"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SubscriptionName: jsii.String("subscriptionName"),
//   	Threshold: jsii.Number(123),
//   	ThresholdExpression: jsii.String("thresholdExpression"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ce-anomalysubscription.html
//
type CfnAnomalySubscriptionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAnomalySubscriptionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAnomalySubscriptionPropsMixin
type jsiiProxy_CfnAnomalySubscriptionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAnomalySubscriptionPropsMixin) Props() *CfnAnomalySubscriptionMixinProps {
	var returns *CfnAnomalySubscriptionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalySubscriptionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CE::AnomalySubscription`.
func NewCfnAnomalySubscriptionPropsMixin(props *CfnAnomalySubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAnomalySubscriptionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAnomalySubscriptionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAnomalySubscriptionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ce.CfnAnomalySubscriptionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CE::AnomalySubscription`.
func NewCfnAnomalySubscriptionPropsMixin_Override(c CfnAnomalySubscriptionPropsMixin, props *CfnAnomalySubscriptionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ce.CfnAnomalySubscriptionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAnomalySubscriptionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnomalySubscriptionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ce.CfnAnomalySubscriptionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnomalySubscriptionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ce.CfnAnomalySubscriptionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnomalySubscriptionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAnomalySubscriptionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

