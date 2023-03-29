package awselasticloadbalancingv2


// Load balancing algorithmm type for target groups.
type TargetGroupLoadBalancingAlgorithmType string

const (
	// round_robin.
	TargetGroupLoadBalancingAlgorithmType_ROUND_ROBIN TargetGroupLoadBalancingAlgorithmType = "ROUND_ROBIN"
	// least_outstanding_requests.
	TargetGroupLoadBalancingAlgorithmType_LEAST_OUTSTANDING_REQUESTS TargetGroupLoadBalancingAlgorithmType = "LEAST_OUTSTANDING_REQUESTS"
)

