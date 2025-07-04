package cloudassemblyschema


// Type of load balancer.
type LoadBalancerType string

const (
	// Network load balancer.
	LoadBalancerType_NETWORK LoadBalancerType = "NETWORK"
	// Application load balancer.
	LoadBalancerType_APPLICATION LoadBalancerType = "APPLICATION"
)

