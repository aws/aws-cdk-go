package previewawslogsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs/previewawslogsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Logs::MetricFilter` resource specifies a metric filter that describes how CloudWatch Logs extracts information from logs and transforms it into Amazon CloudWatch metrics.
//
// If you have multiple metric filters that are associated with a log group, all the filters are applied to the log streams in that group.
//
// The maximum number of metric filters that can be associated with a log group is 100.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMetricFilterPropsMixin := awscdkmixinspreview.Mixins.NewCfnMetricFilterPropsMixin(&CfnMetricFilterMixinProps{
//   	ApplyOnTransformedLogs: jsii.Boolean(false),
//   	EmitSystemFieldDimensions: []*string{
//   		jsii.String("emitSystemFieldDimensions"),
//   	},
//   	FieldSelectionCriteria: jsii.String("fieldSelectionCriteria"),
//   	FilterName: jsii.String("filterName"),
//   	FilterPattern: jsii.String("filterPattern"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	MetricTransformations: []interface{}{
//   		&MetricTransformationProperty{
//   			DefaultValue: jsii.Number(123),
//   			Dimensions: []interface{}{
//   				&DimensionProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MetricName: jsii.String("metricName"),
//   			MetricNamespace: jsii.String("metricNamespace"),
//   			MetricValue: jsii.String("metricValue"),
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
//
type CfnMetricFilterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMetricFilterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMetricFilterPropsMixin
type jsiiProxy_CfnMetricFilterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMetricFilterPropsMixin) Props() *CfnMetricFilterMixinProps {
	var returns *CfnMetricFilterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricFilterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Logs::MetricFilter`.
func NewCfnMetricFilterPropsMixin(props *CfnMetricFilterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMetricFilterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMetricFilterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMetricFilterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnMetricFilterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Logs::MetricFilter`.
func NewCfnMetricFilterPropsMixin_Override(c CfnMetricFilterPropsMixin, props *CfnMetricFilterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnMetricFilterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMetricFilterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMetricFilterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnMetricFilterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMetricFilterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnMetricFilterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMetricFilterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMetricFilterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

