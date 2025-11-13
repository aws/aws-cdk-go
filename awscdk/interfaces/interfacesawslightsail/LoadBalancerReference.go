package interfacesawslightsail


// A reference to a LoadBalancer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerReference := &LoadBalancerReference{
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   	LoadBalancerName: jsii.String("loadBalancerName"),
//   }
//
type LoadBalancerReference struct {
	// The ARN of the LoadBalancer resource.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The LoadBalancerName of the LoadBalancer resource.
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
}

