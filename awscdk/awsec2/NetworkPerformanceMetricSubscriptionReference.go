package awsec2


// A reference to a NetworkPerformanceMetricSubscription resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkPerformanceMetricSubscriptionReference := &NetworkPerformanceMetricSubscriptionReference{
//   	Destination: jsii.String("destination"),
//   	Metric: jsii.String("metric"),
//   	Source: jsii.String("source"),
//   	Statistic: jsii.String("statistic"),
//   }
//
type NetworkPerformanceMetricSubscriptionReference struct {
	// The Destination of the NetworkPerformanceMetricSubscription resource.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The Metric of the NetworkPerformanceMetricSubscription resource.
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The Source of the NetworkPerformanceMetricSubscription resource.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The Statistic of the NetworkPerformanceMetricSubscription resource.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}

