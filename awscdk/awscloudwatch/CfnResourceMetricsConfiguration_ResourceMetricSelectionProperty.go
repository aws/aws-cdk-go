package awscloudwatch


// Defines which metrics are enabled for detailed monitoring on the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceMetricSelectionProperty := &ResourceMetricSelectionProperty{
//   	IncludeMetrics: []*string{
//   		jsii.String("includeMetrics"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-resourcemetricsconfiguration-resourcemetricselection.html
//
type CfnResourceMetricsConfiguration_ResourceMetricSelectionProperty struct {
	// The list of metric names to include in detailed monitoring for the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-resourcemetricsconfiguration-resourcemetricselection.html#cfn-cloudwatch-resourcemetricsconfiguration-resourcemetricselection-includemetrics
	//
	IncludeMetrics *[]*string `field:"required" json:"includeMetrics" yaml:"includeMetrics"`
}

