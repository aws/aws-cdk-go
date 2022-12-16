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
type NetworkLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

