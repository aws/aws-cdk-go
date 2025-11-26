package previewawscloudwatchmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudwatch/previewawscloudwatchmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a metric stream.
//
// Metrics streams can automatically stream CloudWatch metrics to AWS destinations including Amazon S3 and to many third-party solutions. For more information, see [Metric streams](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Metric-Streams.html) .
//
// To create a metric stream, you must be logged on to an account that has the `iam:PassRole` permission and either the *CloudWatchFullAccess* policy or the `cloudwatch:PutMetricStream` permission.
//
// When you create or update a metric stream, you choose one of the following:
//
// - Stream metrics from all metric namespaces in the account.
// - Stream metrics from all metric namespaces in the account, except for the namespaces that you list in `ExcludeFilters` .
// - Stream metrics from only the metric namespaces that you list in `IncludeFilters` .
//
// When you create a metric stream, the stream is created in the `running` state. If you update an existing metric stream, the state does not change.
//
// If you create a metric stream in an account that has been set up as a monitoring account in CloudWatch cross-account observability, you can choose whether to include metrics from linked source accounts in the metric stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMetricStreamPropsMixin := awscdkmixinspreview.Mixins.NewCfnMetricStreamPropsMixin(&CfnMetricStreamMixinProps{
//   	ExcludeFilters: []interface{}{
//   		&MetricStreamFilterProperty{
//   			MetricNames: []*string{
//   				jsii.String("metricNames"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   	FirehoseArn: jsii.String("firehoseArn"),
//   	IncludeFilters: []interface{}{
//   		&MetricStreamFilterProperty{
//   			MetricNames: []*string{
//   				jsii.String("metricNames"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   	IncludeLinkedAccountsMetrics: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	OutputFormat: jsii.String("outputFormat"),
//   	RoleArn: jsii.String("roleArn"),
//   	StatisticsConfigurations: []interface{}{
//   		&MetricStreamStatisticsConfigurationProperty{
//   			AdditionalStatistics: []*string{
//   				jsii.String("additionalStatistics"),
//   			},
//   			IncludeMetrics: []interface{}{
//   				&MetricStreamStatisticsMetricProperty{
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html
//
type CfnMetricStreamPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMetricStreamMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMetricStreamPropsMixin
type jsiiProxy_CfnMetricStreamPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMetricStreamPropsMixin) Props() *CfnMetricStreamMixinProps {
	var returns *CfnMetricStreamMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricStreamPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::MetricStream`.
func NewCfnMetricStreamPropsMixin(props *CfnMetricStreamMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMetricStreamPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMetricStreamPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMetricStreamPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::MetricStream`.
func NewCfnMetricStreamPropsMixin_Override(c CfnMetricStreamPropsMixin, props *CfnMetricStreamMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMetricStreamPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMetricStreamPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMetricStreamPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudwatch.mixins.CfnMetricStreamPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMetricStreamPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMetricStreamPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

