package awsec2


// Properties for defining a `CfnNetworkPerformanceMetricSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkPerformanceMetricSubscriptionProps := &CfnNetworkPerformanceMetricSubscriptionProps{
//   	Destination: jsii.String("destination"),
//   	Metric: jsii.String("metric"),
//   	Source: jsii.String("source"),
//   	Statistic: jsii.String("statistic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkperformancemetricsubscription.html
//
type CfnNetworkPerformanceMetricSubscriptionProps struct {
	// The Region or Availability Zone that's the target for the subscription.
	//
	// For example, `eu-west-1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkperformancemetricsubscription.html#cfn-ec2-networkperformancemetricsubscription-destination
	//
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The metric used for the subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkperformancemetricsubscription.html#cfn-ec2-networkperformancemetricsubscription-metric
	//
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The Region or Availability Zone that's the source for the subscription.
	//
	// For example, `us-east-1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkperformancemetricsubscription.html#cfn-ec2-networkperformancemetricsubscription-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The statistic used for the subscription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkperformancemetricsubscription.html#cfn-ec2-networkperformancemetricsubscription-statistic
	//
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}

