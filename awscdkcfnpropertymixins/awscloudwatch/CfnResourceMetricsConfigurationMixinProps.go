package awscloudwatch


// Properties for CfnResourceMetricsConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnResourceMetricsConfigurationMixinProps := &CfnResourceMetricsConfigurationMixinProps{
//   	MetricSelections: []interface{}{
//   		&ResourceMetricSelectionProperty{
//   			IncludeMetrics: []*string{
//   				jsii.String("includeMetrics"),
//   			},
//   		},
//   	},
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-resourcemetricsconfiguration.html
//
type CfnResourceMetricsConfigurationMixinProps struct {
	// The metric selections that define which metrics are enabled for detailed monitoring on the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-resourcemetricsconfiguration.html#cfn-cloudwatch-resourcemetricsconfiguration-metricselections
	//
	MetricSelections interface{} `field:"optional" json:"metricSelections" yaml:"metricSelections"`
	// The Amazon Resource Name (ARN) of the resource for which the detailed monitoring metrics configuration is managed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-resourcemetricsconfiguration.html#cfn-cloudwatch-resourcemetricsconfiguration-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

