package awselasticloadbalancingv2


// Load balancing protocol for application load balancers.
//
// Example:
//   var listener applicationListener
//   var service baseService
//
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
// Experimental.
type ApplicationProtocol string

const (
	// HTTP.
	// Experimental.
	ApplicationProtocol_HTTP ApplicationProtocol = "HTTP"
	// HTTPS.
	// Experimental.
	ApplicationProtocol_HTTPS ApplicationProtocol = "HTTPS"
)

