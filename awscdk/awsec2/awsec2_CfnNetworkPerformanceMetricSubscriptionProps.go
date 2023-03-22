package awsec2


// Properties for defining a `CfnNetworkPerformanceMetricSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkPerformanceMetricSubscriptionProps := &cfnNetworkPerformanceMetricSubscriptionProps{
//   	destination: jsii.String("destination"),
//   	metric: jsii.String("metric"),
//   	source: jsii.String("source"),
//   	statistic: jsii.String("statistic"),
//   }
//
type CfnNetworkPerformanceMetricSubscriptionProps struct {
	// The Region or Availability Zone that's the target for the subscription.
	//
	// For example, `eu-west-1` .
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The metric used for the subscription.
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The Region or Availability Zone that's the source for the subscription.
	//
	// For example, `us-east-1` .
	Source *string `field:"required" json:"source" yaml:"source"`
	// The statistic used for the subscription.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}

