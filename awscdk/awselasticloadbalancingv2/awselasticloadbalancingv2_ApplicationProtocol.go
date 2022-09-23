package awselasticloadbalancingv2


// Load balancing protocol for application load balancers.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
type ApplicationProtocol string

const (
	// HTTP.
	ApplicationProtocol_HTTP ApplicationProtocol = "HTTP"
	// HTTPS.
	ApplicationProtocol_HTTPS ApplicationProtocol = "HTTPS"
)

