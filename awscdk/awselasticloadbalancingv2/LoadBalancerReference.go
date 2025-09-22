package awselasticloadbalancingv2


// A reference to a LoadBalancer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerReference := &LoadBalancerReference{
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   }
//
type LoadBalancerReference struct {
	// The LoadBalancerArn of the LoadBalancer resource.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
}

