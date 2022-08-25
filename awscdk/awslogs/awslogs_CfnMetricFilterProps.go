package awslogs


// Properties for defining a `CfnMetricFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMetricFilterProps := &cfnMetricFilterProps{
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//   	metricTransformations: []interface{}{
//   		&metricTransformationProperty{
//   			metricName: jsii.String("metricName"),
//   			metricNamespace: jsii.String("metricNamespace"),
//   			metricValue: jsii.String("metricValue"),
//
//   			// the properties below are optional
//   			defaultValue: jsii.Number(123),
//   			dimensions: []interface{}{
//   				&dimensionProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			unit: jsii.String("unit"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	filterName: jsii.String("filterName"),
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
	// `AWS::Logs::MetricFilter.FilterName`.
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
}

