package awsecs


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
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

