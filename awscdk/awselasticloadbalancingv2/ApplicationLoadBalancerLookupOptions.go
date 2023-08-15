package awselasticloadbalancingv2


// Options for looking up an ApplicationLoadBalancer.
//
// Example:
//   loadBalancer := elbv2.ApplicationLoadBalancer_FromLookup(this, jsii.String("ALB"), &ApplicationLoadBalancerLookupOptions{
//   	LoadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-east-2:123456789012:loadbalancer/app/my-load-balancer/1234567890123456"),
//   })
//
type ApplicationLoadBalancerLookupOptions struct {
	// Find by load balancer's ARN.
	// Default: - does not search by load balancer arn.
	//
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Default: - does not match load balancers by tags.
	//
	LoadBalancerTags *map[string]*string `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

