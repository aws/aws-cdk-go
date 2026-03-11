package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoT::CustomMetric` resource to define a custom metric published by your devices to Device Defender.
//
// For API reference, see [CreateCustomMetric](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCustomMetric.html) and for general information, see [Custom metrics](https://docs.aws.amazon.com/iot/latest/developerguide/dd-detect-custom-metrics.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCustomMetricPropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnCustomMetricPropsMixin(&CfnCustomMetricMixinProps{
//   	DisplayName: jsii.String("displayName"),
//   	MetricName: jsii.String("metricName"),
//   	MetricType: jsii.String("metricType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-custommetric.html
//
type CfnCustomMetricPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCustomMetricMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCustomMetricPropsMixin
type jsiiProxy_CfnCustomMetricPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCustomMetricPropsMixin) Props() *CfnCustomMetricMixinProps {
	var returns *CfnCustomMetricMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomMetricPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::CustomMetric`.
func NewCfnCustomMetricPropsMixin(props *CfnCustomMetricMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCustomMetricPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCustomMetricPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomMetricPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnCustomMetricPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::CustomMetric`.
func NewCfnCustomMetricPropsMixin_Override(c CfnCustomMetricPropsMixin, props *CfnCustomMetricMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnCustomMetricPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCustomMetricPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomMetricPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnCustomMetricPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomMetricPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnCustomMetricPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomMetricPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCustomMetricPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

