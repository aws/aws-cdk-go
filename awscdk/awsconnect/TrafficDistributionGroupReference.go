package awsconnect


// A reference to a TrafficDistributionGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficDistributionGroupReference := &TrafficDistributionGroupReference{
//   	TrafficDistributionGroupArn: jsii.String("trafficDistributionGroupArn"),
//   }
//
type TrafficDistributionGroupReference struct {
	// The TrafficDistributionGroupArn of the TrafficDistributionGroup resource.
	TrafficDistributionGroupArn *string `field:"required" json:"trafficDistributionGroupArn" yaml:"trafficDistributionGroupArn"`
}

