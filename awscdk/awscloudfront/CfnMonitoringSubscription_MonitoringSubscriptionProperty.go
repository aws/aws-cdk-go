package awscloudfront


// A monitoring subscription.
//
// This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringSubscriptionProperty := &MonitoringSubscriptionProperty{
//   	RealtimeMetricsSubscriptionConfig: &RealtimeMetricsSubscriptionConfigProperty{
//   		RealtimeMetricsSubscriptionStatus: jsii.String("realtimeMetricsSubscriptionStatus"),
//   	},
//   }
//
type CfnMonitoringSubscription_MonitoringSubscriptionProperty struct {
	// A subscription configuration for additional CloudWatch metrics.
	RealtimeMetricsSubscriptionConfig interface{} `field:"optional" json:"realtimeMetricsSubscriptionConfig" yaml:"realtimeMetricsSubscriptionConfig"`
}

