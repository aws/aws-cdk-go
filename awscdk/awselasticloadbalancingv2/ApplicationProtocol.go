package awselasticloadbalancingv2


// Load balancing protocol for application load balancers.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//   listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   service.RegisterLoadBalancerTargets(&EcsTarget{
//   	ContainerName: jsii.String("web"),
//   	ContainerPort: jsii.Number(80),
//   	NewTargetGroupId: jsii.String("ECS"),
//   	Listener: ecs.ListenerConfig_ApplicationListener(listener, &AddApplicationTargetsProps{
//   		Protocol: elbv2.ApplicationProtocol_HTTPS,
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

