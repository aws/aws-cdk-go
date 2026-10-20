package awscloudwatch


// Properties for defining a `CfnResourceMetricsConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceMetricsConfigurationProps := &CfnResourceMetricsConfigurationProps{
//   	ResourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	MetricSelections: []interface{}{
//   		&ResourceMetricSelectionProperty{
//   			IncludeMetrics: []*string{
//   				jsii.String("includeMetrics"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-resourcemetricsconfiguration.html
//
type CfnResourceMetricsConfigurationProps struct {
	// The Amazon Resource Name (ARN) of the resource for which the detailed monitoring metrics configuration is managed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-resourcemetricsconfiguration.html#cfn-cloudwatch-resourcemetricsconfiguration-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The metric selections that define which metrics are enabled for detailed monitoring on the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-resourcemetricsconfiguration.html#cfn-cloudwatch-resourcemetricsconfiguration-metricselections
	//
	MetricSelections interface{} `field:"optional" json:"metricSelections" yaml:"metricSelections"`
}

