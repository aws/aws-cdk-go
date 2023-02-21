package awscodedeploy


// The generations of AWS load balancing solutions.
type LoadBalancerGeneration string

const (
	// The first generation (ELB Classic).
	LoadBalancerGeneration_FIRST LoadBalancerGeneration = "FIRST"
	// The second generation (ALB and NLB).
	LoadBalancerGeneration_SECOND LoadBalancerGeneration = "SECOND"
)

