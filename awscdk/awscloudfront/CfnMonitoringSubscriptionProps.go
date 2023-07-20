package awscloudfront


// Properties for defining a `CfnMonitoringSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitoringSubscriptionProps := &CfnMonitoringSubscriptionProps{
//   	DistributionId: jsii.String("distributionId"),
//   	MonitoringSubscription: &MonitoringSubscriptionProperty{
//   		RealtimeMetricsSubscriptionConfig: &RealtimeMetricsSubscriptionConfigProperty{
//   			RealtimeMetricsSubscriptionStatus: jsii.String("realtimeMetricsSubscriptionStatus"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-monitoringsubscription.html
//
type CfnMonitoringSubscriptionProps struct {
	// The ID of the distribution that you are enabling metrics for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-monitoringsubscription.html#cfn-cloudfront-monitoringsubscription-distributionid
	//
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// A subscription configuration for additional CloudWatch metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-monitoringsubscription.html#cfn-cloudfront-monitoringsubscription-monitoringsubscription
	//
	MonitoringSubscription interface{} `field:"required" json:"monitoringSubscription" yaml:"monitoringSubscription"`
}

