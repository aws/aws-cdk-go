package awselasticloadbalancingv2


// Load balancing algorithmm type for target groups.
// Experimental.
type TargetGroupLoadBalancingAlgorithmType string

const (
	// round_robin.
	// Experimental.
	TargetGroupLoadBalancingAlgorithmType_ROUND_ROBIN TargetGroupLoadBalancingAlgorithmType = "ROUND_ROBIN"
	// least_outstanding_requests.
	// Experimental.
	TargetGroupLoadBalancingAlgorithmType_LEAST_OUTSTANDING_REQUESTS TargetGroupLoadBalancingAlgorithmType = "LEAST_OUTSTANDING_REQUESTS"
)

