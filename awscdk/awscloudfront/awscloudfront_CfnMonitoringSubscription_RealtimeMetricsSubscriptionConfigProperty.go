package awscloudfront


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
	// `CfnMonitoringSubscription.RealtimeMetricsSubscriptionConfigProperty.RealtimeMetricsSubscriptionStatus`.
	RealtimeMetricsSubscriptionStatus *string `field:"required" json:"realtimeMetricsSubscriptionStatus" yaml:"realtimeMetricsSubscriptionStatus"`
}

