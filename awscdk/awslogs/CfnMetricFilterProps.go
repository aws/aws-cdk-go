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
//   	FilterName: jsii.String("filterName"),
//   }
//
type CfnMetricFilterProps struct {
	// A filter pattern for extracting metric data out of ingested log events.
	//
	// For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	FilterPattern *string `field:"required" json:"filterPattern" yaml:"filterPattern"`
	// The name of an existing log group that you want to associate with this metric filter.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// The metric transformations.
	MetricTransformations interface{} `field:"required" json:"metricTransformations" yaml:"metricTransformations"`
	// The name of the metric filter.
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
}

