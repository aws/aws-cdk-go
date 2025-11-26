package previewawscloudwatchmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudwatch/previewawscloudwatchmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CloudWatch::AnomalyDetector` type specifies an anomaly detection band for a certain metric and statistic.
//
// The band represents the expected "normal" range for the metric values. Anomaly detection bands can be used for visualization of a metric's expected values, and for alarms.
//
// For more information see [Using CloudWatch anomaly detection.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Anomaly_Detection.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAnomalyDetectorPropsMixin := awscdkmixinspreview.Mixins.NewCfnAnomalyDetectorPropsMixin(&CfnAnomalyDetectorMixinProps{
//   	Configuration: &ConfigurationProperty{
//   		ExcludedTimeRanges: []interface{}{
//   			&RangeProperty{
//   				EndTime: jsii.String("endTime"),
//   				StartTime: jsii.String("startTime"),
//   			},
//   		},
//   		MetricTimeZone: jsii.String("metricTimeZone"),
//   	},
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MetricCharacteristics: &MetricCharacteristicsProperty{
//   		PeriodicSpikes: jsii.Boolean(false),
//   	},
//   	MetricMathAnomalyDetector: &MetricMathAnomalyDetectorProperty{
//   		MetricDataQueries: []interface{}{
//   			&MetricDataQueryProperty{
//   				AccountId: jsii.String("accountId"),
//   				Expression: jsii.String("expression"),
//   				Id: jsii.String("id"),
//   				Label: jsii.String("label"),
//   				MetricStat: &MetricStatProperty{
//   					Metric: &MetricProperty{
//   						Dimensions: []interface{}{
//   							&DimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   					},
//   					Period: jsii.Number(123),
//   					Stat: jsii.String("stat"),
//   					Unit: jsii.String("unit"),
//   				},
//   				Period: jsii.Number(123),
//   				ReturnData: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Namespace: jsii.String("namespace"),
//   	SingleMetricAnomalyDetector: &SingleMetricAnomalyDetectorProperty{
//   		AccountId: jsii.String("accountId"),
//   		Dimensions: []interface{}{
//   			&DimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   		Stat: jsii.String("stat"),
//   	},
//   	Stat: jsii.String("stat"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-anomalydetector.html
//
type CfnAnomalyDetectorPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAnomalyDetectorMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAnomalyDetectorPropsMixin
type jsiiProxy_CfnAnomalyDetectorPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAnomalyDetectorPropsMixin) Props() *CfnAnomalyDetectorMixinProps {
	var returns *CfnAnomalyDetectorMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnomalyDetectorPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::AnomalyDetector`.
func NewCfnAnomalyDetectorPropsMixin(props *CfnAnomalyDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAnomalyDetectorPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAnomalyDetectorPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::AnomalyDetector`.
func NewCfnAnomalyDetectorPropsMixin_Override(c CfnAnomalyDetectorPropsMixin, props *CfnAnomalyDetectorMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAnomalyDetectorPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnomalyDetectorPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnomalyDetectorPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnAnomalyDetectorPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnomalyDetectorPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAnomalyDetectorPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

