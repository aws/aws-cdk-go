package awscdkcloudassemblyschema


// Query input for looking up a load balancer listener.
type LoadBalancerListenerContextQuery struct {
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
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Find by listener's arn.
	// Default: - does not find by listener arn.
	//
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// Filter listeners by listener port.
	// Default: - does not filter by a listener port.
	//
	ListenerPort *float64 `field:"optional" json:"listenerPort" yaml:"listenerPort"`
	// Filter by listener protocol.
	// Default: - does not filter by listener protocol.
	//
	ListenerProtocol LoadBalancerListenerProtocol `field:"optional" json:"listenerProtocol" yaml:"listenerProtocol"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

