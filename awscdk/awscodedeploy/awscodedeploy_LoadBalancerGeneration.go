package awscodedeploy


// The generations of AWS load balancing solutions.
// Experimental.
type LoadBalancerGeneration string

const (
	// The first generation (ELB Classic).
	// Experimental.
	LoadBalancerGeneration_FIRST LoadBalancerGeneration = "FIRST"
	// The second generation (ALB and NLB).
	// Experimental.
	LoadBalancerGeneration_SECOND LoadBalancerGeneration = "SECOND"
)

