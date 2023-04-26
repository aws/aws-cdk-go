package awsec2


// Describes a load balancer target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisLoadBalancerTargetProperty := &AnalysisLoadBalancerTargetProperty{
//   	Address: jsii.String("address"),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Instance: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Port: jsii.Number(123),
//   }
//
type CfnNetworkInsightsAnalysis_AnalysisLoadBalancerTargetProperty struct {
	// The IP address.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The Availability Zone.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Information about the instance.
	Instance interface{} `field:"optional" json:"instance" yaml:"instance"`
	// The port on which the target is listening.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

