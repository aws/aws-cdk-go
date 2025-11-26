package previewawslogsmixins


// Properties for CfnMetricFilterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMetricFilterMixinProps := &CfnMetricFilterMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
//
type CfnMetricFilterMixinProps struct {
	// This parameter is valid only for log groups that have an active log transformer.
	//
	// For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html) .
	//
	// If this value is `true` , the metric filter is applied on the transformed version of the log events instead of the original ingested log events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-applyontransformedlogs
	//
	ApplyOnTransformedLogs interface{} `field:"optional" json:"applyOnTransformedLogs" yaml:"applyOnTransformedLogs"`
	// The list of system fields that are emitted as additional dimensions in the generated metrics.
	//
	// Returns the `emitSystemFieldDimensions` value if it was specified when the metric filter was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-emitsystemfielddimensions
	//
	EmitSystemFieldDimensions *[]*string `field:"optional" json:"emitSystemFieldDimensions" yaml:"emitSystemFieldDimensions"`
	// The filter expression that specifies which log events are processed by this metric filter based on system fields.
	//
	// Returns the `fieldSelectionCriteria` value if it was specified when the metric filter was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-fieldselectioncriteria
	//
	FieldSelectionCriteria *string `field:"optional" json:"fieldSelectionCriteria" yaml:"fieldSelectionCriteria"`
	// The name of the metric filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-filtername
	//
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// A filter pattern for extracting metric data out of ingested log events.
	//
	// For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-filterpattern
	//
	FilterPattern *string `field:"optional" json:"filterPattern" yaml:"filterPattern"`
	// The name of an existing log group that you want to associate with this metric filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The metric transformations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-metrictransformations
	//
	MetricTransformations interface{} `field:"optional" json:"metricTransformations" yaml:"metricTransformations"`
}

