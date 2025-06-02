package awscloudfront


// A subscription configuration for additional CloudWatch metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   realtimeMetricsSubscriptionConfigProperty := &RealtimeMetricsSubscriptionConfigProperty{
//   	RealtimeMetricsSubscriptionStatus: jsii.String("realtimeMetricsSubscriptionStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-monitoringsubscription-realtimemetricssubscriptionconfig.html
//
type CfnMonitoringSubscription_RealtimeMetricsSubscriptionConfigProperty struct {
	// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-monitoringsubscription-realtimemetricssubscriptionconfig.html#cfn-cloudfront-monitoringsubscription-realtimemetricssubscriptionconfig-realtimemetricssubscriptionstatus
	//
	RealtimeMetricsSubscriptionStatus *string `field:"required" json:"realtimeMetricsSubscriptionStatus" yaml:"realtimeMetricsSubscriptionStatus"`
}

