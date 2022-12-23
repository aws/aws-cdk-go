package awscloudfront


// A subscription configuration for additional CloudWatch metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   realtimeMetricsSubscriptionConfigProperty := &realtimeMetricsSubscriptionConfigProperty{
//   	realtimeMetricsSubscriptionStatus: jsii.String("realtimeMetricsSubscriptionStatus"),
//   }
//
type CfnMonitoringSubscription_RealtimeMetricsSubscriptionConfigProperty struct {
	// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	RealtimeMetricsSubscriptionStatus *string `field:"required" json:"realtimeMetricsSubscriptionStatus" yaml:"realtimeMetricsSubscriptionStatus"`
}

