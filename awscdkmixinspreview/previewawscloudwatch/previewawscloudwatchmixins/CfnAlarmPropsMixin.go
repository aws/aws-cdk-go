package previewawscloudwatchmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudwatch/previewawscloudwatchmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudWatch::Alarm` type specifies an alarm and associates it with the specified metric or metric math expression.
//
// When this operation creates an alarm, the alarm state is immediately set to `INSUFFICIENT_DATA` . The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed.
//
// When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAlarmPropsMixin := awscdkmixinspreview.Mixins.NewCfnAlarmPropsMixin(&CfnAlarmMixinProps{
//   	ActionsEnabled: jsii.Boolean(false),
//   	AlarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	AlarmDescription: jsii.String("alarmDescription"),
//   	AlarmName: jsii.String("alarmName"),
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	DatapointsToAlarm: jsii.Number(123),
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	EvaluateLowSampleCountPercentile: jsii.String("evaluateLowSampleCountPercentile"),
//   	EvaluationPeriods: jsii.Number(123),
//   	ExtendedStatistic: jsii.String("extendedStatistic"),
//   	InsufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Metrics: []interface{}{
//   		&MetricDataQueryProperty{
//   			AccountId: jsii.String("accountId"),
//   			Expression: jsii.String("expression"),
//   			Id: jsii.String("id"),
//   			Label: jsii.String("label"),
//   			MetricStat: &MetricStatProperty{
//   				Metric: &MetricProperty{
//   					Dimensions: []interface{}{
//   						&DimensionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   				Period: jsii.Number(123),
//   				Stat: jsii.String("stat"),
//   				Unit: jsii.String("unit"),
//   			},
//   			Period: jsii.Number(123),
//   			ReturnData: jsii.Boolean(false),
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	OkActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	Period: jsii.Number(123),
//   	Statistic: jsii.String("statistic"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Threshold: jsii.Number(123),
//   	ThresholdMetricId: jsii.String("thresholdMetricId"),
//   	TreatMissingData: jsii.String("treatMissingData"),
//   	Unit: jsii.String("unit"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html
//
type CfnAlarmPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAlarmMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAlarmPropsMixin
type jsiiProxy_CfnAlarmPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAlarmPropsMixin) Props() *CfnAlarmMixinProps {
	var returns *CfnAlarmMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlarmPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::Alarm`.
func NewCfnAlarmPropsMixin(props *CfnAlarmMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAlarmPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAlarmPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAlarmPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::Alarm`.
func NewCfnAlarmPropsMixin_Override(c CfnAlarmPropsMixin, props *CfnAlarmMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAlarmPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAlarmPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlarmPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAlarmPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAlarmPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAlarmPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

