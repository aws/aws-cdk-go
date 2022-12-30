package cloudassemblyschema


// Type of load balancer.
// Experimental.
type LoadBalancerType string

const (
	// Network load balancer.
	// Experimental.
	LoadBalancerType_NETWORK LoadBalancerType = "NETWORK"
	// Application load balancer.
	// Experimental.
	LoadBalancerType_APPLICATION LoadBalancerType = "APPLICATION"
)

