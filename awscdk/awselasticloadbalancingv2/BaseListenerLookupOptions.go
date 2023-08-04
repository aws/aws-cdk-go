package awselasticloadbalancingv2


// Options for listener lookup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseListenerLookupOptions := &BaseListenerLookupOptions{
//   	ListenerPort: jsii.Number(123),
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   	LoadBalancerTags: map[string]*string{
//   		"loadBalancerTagsKey": jsii.String("loadBalancerTags"),
//   	},
//   }
//
type BaseListenerLookupOptions struct {
	// Filter listeners by listener port.
	// Default: - does not filter by listener port.
	//
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter listeners by associated load balancer arn.
	// Default: - does not filter by load balancer arn.
	//
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Filter listeners by associated load balancer tags.
	// Default: - does not filter by load balancer tags.
	//
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

