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
	// `AWS::EC2::NetworkPerformanceMetricSubscription.Destination`.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// `AWS::EC2::NetworkPerformanceMetricSubscription.Metric`.
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// `AWS::EC2::NetworkPerformanceMetricSubscription.Source`.
	Source *string `field:"required" json:"source" yaml:"source"`
	// `AWS::EC2::NetworkPerformanceMetricSubscription.Statistic`.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}

