package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringSubscriptionProperty := &monitoringSubscriptionProperty{
//   	realtimeMetricsSubscriptionConfig: &realtimeMetricsSubscriptionConfigProperty{
//   		realtimeMetricsSubscriptionStatus: jsii.String("realtimeMetricsSubscriptionStatus"),
//   	},
//   }
//
type CfnMonitoringSubscription_MonitoringSubscriptionProperty struct {
	// `CfnMonitoringSubscription.MonitoringSubscriptionProperty.RealtimeMetricsSubscriptionConfig`.
	RealtimeMetricsSubscriptionConfig interface{} `field:"optional" json:"realtimeMetricsSubscriptionConfig" yaml:"realtimeMetricsSubscriptionConfig"`
}

