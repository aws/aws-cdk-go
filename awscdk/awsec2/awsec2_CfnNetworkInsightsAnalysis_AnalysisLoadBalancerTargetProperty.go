package awsec2


// Describes a load balancer target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisLoadBalancerTargetProperty := &analysisLoadBalancerTargetProperty{
//   	address: jsii.String("address"),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	instance: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	port: jsii.Number(123),
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

