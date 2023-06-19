package awselasticloadbalancingv2


// Load balancing protocol for application load balancers.
//
// Example:
//   var listener applicationListener
//   var service baseService
//
//   service.RegisterLoadBalancerTargets(&EcsTarget{
//   	ContainerName: jsii.String("web"),
//   	ContainerPort: jsii.Number(80),
//   	NewTargetGroupId: jsii.String("ECS"),
//   	Listener: ecs.ListenerConfig_ApplicationListener(listener, &AddApplicationTargetsProps{
//   		Protocol: elbv2.ApplicationProtocol_HTTPS,
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

