package awselasticloadbalancingv2


// Load balancing algorithmm type for target groups.
//
// Example:
//   var vpc Vpc
//
//
//   tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TargetGroup"), &ApplicationTargetGroupProps{
//   	Vpc: Vpc,
//   	LoadBalancingAlgorithmType: elbv2.TargetGroupLoadBalancingAlgorithmType_WEIGHTED_RANDOM,
//   	EnableAnomalyMitigation: jsii.Boolean(true),
//   })
//
type TargetGroupLoadBalancingAlgorithmType string

const (
	// round_robin.
	TargetGroupLoadBalancingAlgorithmType_ROUND_ROBIN TargetGroupLoadBalancingAlgorithmType = "ROUND_ROBIN"
	// least_outstanding_requests.
	TargetGroupLoadBalancingAlgorithmType_LEAST_OUTSTANDING_REQUESTS TargetGroupLoadBalancingAlgorithmType = "LEAST_OUTSTANDING_REQUESTS"
	// weighted_random.
	TargetGroupLoadBalancingAlgorithmType_WEIGHTED_RANDOM TargetGroupLoadBalancingAlgorithmType = "WEIGHTED_RANDOM"
)

