package awselasticloadbalancingv2


// Options for looking up an ApplicationLoadBalancer.
//
// Example:
//   loadBalancer := elbv2.applicationLoadBalancer.fromLookup(this, jsii.String("ALB"), &applicationLoadBalancerLookupOptions{
//   	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-east-2:123456789012:loadbalancer/app/my-load-balancer/1234567890123456"),
//   })
//
// Experimental.
type ApplicationLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

