package awselasticloadbalancingv2


// Options for looking up an NetworkLoadBalancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkLoadBalancerLookupOptions := &networkLoadBalancerLookupOptions{
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: map[string]*string{
//   		"loadBalancerTagsKey": jsii.String("loadBalancerTags"),
//   	},
//   }
//
// Experimental.
type NetworkLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

