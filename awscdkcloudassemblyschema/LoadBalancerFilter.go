package awscdkcloudassemblyschema


// Filters for selecting load balancers.
type LoadBalancerFilter struct {
	// Filter load balancers by their type.
	LoadBalancerType LoadBalancerType `field:"required" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	// Default: - does not search by load balancer arn.
	//
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Default: - does not match load balancers by tags.
	//
	LoadBalancerTags *[]*Tag `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
}

