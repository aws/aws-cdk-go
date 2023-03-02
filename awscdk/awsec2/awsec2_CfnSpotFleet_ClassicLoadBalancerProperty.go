package awsec2


// Specifies a Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   classicLoadBalancerProperty := &classicLoadBalancerProperty{
//   	name: jsii.String("name"),
//   }
//
type CfnSpotFleet_ClassicLoadBalancerProperty struct {
	// The name of the load balancer.
	Name *string `field:"required" json:"name" yaml:"name"`
}

