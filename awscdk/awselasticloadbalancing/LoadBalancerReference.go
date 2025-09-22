package awselasticloadbalancing


// A reference to a LoadBalancer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerReference := &LoadBalancerReference{
//   	LoadBalancerId: jsii.String("loadBalancerId"),
//   }
//
type LoadBalancerReference struct {
	// The Id of the LoadBalancer resource.
	LoadBalancerId *string `field:"required" json:"loadBalancerId" yaml:"loadBalancerId"`
}

