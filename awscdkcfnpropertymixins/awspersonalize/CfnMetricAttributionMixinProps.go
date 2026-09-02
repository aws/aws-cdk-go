package awspersonalize


// Properties for CfnMetricAttributionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMetricAttributionMixinProps := &CfnMetricAttributionMixinProps{
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
//   		S3DataDestination: &S3DataDestinationProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html
//
type CfnMetricAttributionMixinProps struct {
	// The ARN of the destination dataset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html#cfn-personalize-metricattribution-datasetgrouparn
	//
	DatasetGroupArn *string `field:"optional" json:"datasetGroupArn" yaml:"datasetGroupArn"`
	// A list of metric attributes for the metric attribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html#cfn-personalize-metricattribution-metrics
	//
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// The output configuration details for the metric attribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html#cfn-personalize-metricattribution-metricsoutputconfig
	//
	MetricsOutputConfig interface{} `field:"optional" json:"metricsOutputConfig" yaml:"metricsOutputConfig"`
	// The name of the metric attribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html#cfn-personalize-metricattribution-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

