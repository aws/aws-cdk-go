package awscloudfront


// Properties for defining a `CfnMonitoringSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitoringSubscriptionProps := &cfnMonitoringSubscriptionProps{
//   	distributionId: jsii.String("distributionId"),
//   	monitoringSubscription: &monitoringSubscriptionProperty{
//   		realtimeMetricsSubscriptionConfig: &realtimeMetricsSubscriptionConfigProperty{
//   			realtimeMetricsSubscriptionStatus: jsii.String("realtimeMetricsSubscriptionStatus"),
//   		},
//   	},
//   }
//
type CfnMonitoringSubscriptionProps struct {
	// The ID of the distribution that you are enabling metrics for.
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
	// A subscription configuration for additional CloudWatch metrics.
	MonitoringSubscription interface{} `field:"required" json:"monitoringSubscription" yaml:"monitoringSubscription"`
}

