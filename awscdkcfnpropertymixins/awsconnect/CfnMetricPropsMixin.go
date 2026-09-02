package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Connect::Metric, a custom metric configured for an Amazon Connect instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMetricPropsMixin := awscdkcfnpropertymixins.Aws_connect.NewCfnMetricPropsMixin(&CfnMetricMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	MetricCalculation: &MetricCalculationProperty{
//   		Calculation: jsii.String("calculation"),
//   		CalculationComponents: []interface{}{
//   			&CalculationComponentProperty{
//   				Alias: jsii.String("alias"),
//   				MetricFilters: []interface{}{
//   					&MetricFilterProperty{
//   						BooleanCondition: &MetricFilterBooleanConditionProperty{
//   							Comparison: jsii.String("comparison"),
//   						},
//   						MetricFilterKey: jsii.String("metricFilterKey"),
//   						Negate: jsii.Boolean(false),
//   						NumberCondition: &MetricFilterNumberConditionProperty{
//   							Comparison: jsii.String("comparison"),
//   							Values: []interface{}{
//   								jsii.Number(123),
//   							},
//   						},
//   						StringCondition: &MetricFilterStringConditionProperty{
//   							Comparison: jsii.String("comparison"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   				MetricId: jsii.String("metricId"),
//   				MetricName: jsii.String("metricName"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PositiveTrendIndicator: jsii.String("positiveTrendIndicator"),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Unit: jsii.String("unit"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-metric.html
//
type CfnMetricPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMetricMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMetricPropsMixin
type jsiiProxy_CfnMetricPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMetricPropsMixin) Props() *CfnMetricMixinProps {
	var returns *CfnMetricMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::Metric`.
func NewCfnMetricPropsMixin(props *CfnMetricMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMetricPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMetricPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMetricPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnMetricPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::Metric`.
func NewCfnMetricPropsMixin_Override(c CfnMetricPropsMixin, props *CfnMetricMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnMetricPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMetricPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMetricPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnMetricPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMetricPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnMetricPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMetricPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMetricPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

