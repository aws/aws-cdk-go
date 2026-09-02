package awspersonalize


// Properties for defining a `CfnMetricAttribution`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMetricAttributionProps := &CfnMetricAttributionProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	Metrics: []interface{}{
//   		&MetricAttributeProperty{
//   			EventType: jsii.String("eventType"),
//   			Expression: jsii.String("expression"),
//   			MetricName: jsii.String("metricName"),
//   		},
//   	},
//   	MetricsOutputConfig: &MetricsOutputConfigProperty{
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		S3DataDestination: &S3DataDestinationProperty{
//   			Path: jsii.String("path"),
//
//   			// the properties below are optional
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html
//
type CfnMetricAttributionProps struct {
	// The ARN of the destination dataset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html#cfn-personalize-metricattribution-datasetgrouparn
	//
	DatasetGroupArn *string `field:"required" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// A list of metric attributes for the metric attribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html#cfn-personalize-metricattribution-metrics
	//
	Metrics interface{} `field:"required" json:"metrics" yaml:"metrics"`
	// The output configuration details for the metric attribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html#cfn-personalize-metricattribution-metricsoutputconfig
	//
	MetricsOutputConfig interface{} `field:"required" json:"metricsOutputConfig" yaml:"metricsOutputConfig"`
	// The name of the metric attribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html#cfn-personalize-metricattribution-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

