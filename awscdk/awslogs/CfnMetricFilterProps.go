package awslogs


// Properties for defining a `CfnMetricFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMetricFilterProps := &CfnMetricFilterProps{
//   	FilterPattern: jsii.String("filterPattern"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	MetricTransformations: []interface{}{
//   		&MetricTransformationProperty{
//   			MetricName: jsii.String("metricName"),
//   			MetricNamespace: jsii.String("metricNamespace"),
//   			MetricValue: jsii.String("metricValue"),
//
//   			// the properties below are optional
//   			DefaultValue: jsii.Number(123),
//   			Dimensions: []interface{}{
//   				&DimensionProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ApplyOnTransformedLogs: jsii.Boolean(false),
//   	FilterName: jsii.String("filterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
//
type CfnMetricFilterProps struct {
	// A filter pattern for extracting metric data out of ingested log events.
	//
	// For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-filterpattern
	//
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The name of an existing log group that you want to associate with this metric filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-loggroupname
	//
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The metric transformations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-metrictransformations
	//
	MetricTransformations interface{} `field:"required" json:"metricTransformations" yaml:"metricTransformations"`
	// This parameter is valid only for log groups that have an active log transformer.
	//
	// For more information about log transformers, see [PutTransformer](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html) .
	//
	// If this value is `true` , the metric filter is applied on the transformed version of the log events instead of the original ingested log events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-applyontransformedlogs
	//
	ApplyOnTransformedLogs interface{} `field:"optional" json:"applyOnTransformedLogs" yaml:"applyOnTransformedLogs"`
	// The name of the metric filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-logs-metricfilter-filtername
	//
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
}

