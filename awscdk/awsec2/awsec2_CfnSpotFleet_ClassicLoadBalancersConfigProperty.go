package awsec2


// Specifies the Classic Load Balancers to attach to a Spot Fleet.
//
// Spot Fleet registers the running Spot Instances with these Classic Load Balancers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   classicLoadBalancersConfigProperty := &classicLoadBalancersConfigProperty{
//   	classicLoadBalancers: []interface{}{
//   		&classicLoadBalancerProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnSpotFleet_ClassicLoadBalancersConfigProperty struct {
	// One or more Classic Load Balancers.
	ClassicLoadBalancers interface{} `field:"required" json:"classicLoadBalancers" yaml:"classicLoadBalancers"`
}

