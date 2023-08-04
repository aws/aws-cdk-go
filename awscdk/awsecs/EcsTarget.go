package awsecs


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
type EcsTarget struct {
	// The name of the container.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Listener and properties for adding target group to the listener.
	Listener ListenerConfig `field:"required" json:"listener" yaml:"listener"`
	// ID for a target group to be created.
	NewTargetGroupId *string `field:"required" json:"newTargetGroupId" yaml:"newTargetGroupId"`
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	// Default: - Container port of the first added port mapping.
	//
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	// Default: Protocol.TCP
	//
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

