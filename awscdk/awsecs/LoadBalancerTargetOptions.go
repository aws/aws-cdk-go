package awsecs


// Properties for defining an ECS target.
//
// The port mapping for it must already have been created through addPortMapping().
//
// Example:
//   var cluster Cluster
//   var taskDefinition TaskDefinition
//   var blueTargetGroup ApplicationTargetGroup
//   var greenTargetGroup ApplicationTargetGroup
//   var prodListenerRule ApplicationListenerRule
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DeploymentStrategy: ecs.DeploymentStrategy_LINEAR,
//   	LinearConfiguration: &TrafficShiftConfig{
//   		StepPercent: jsii.Number(10),
//   		StepBakeTime: awscdk.Duration_Minutes(jsii.Number(5)),
//   	},
//   })
//
//   target := service.LoadBalancerTarget(&LoadBalancerTargetOptions{
//   	ContainerName: jsii.String("web"),
//   	ContainerPort: jsii.Number(80),
//   	AlternateTarget: ecs.NewAlternateTarget(jsii.String("AlternateTarget"), &AlternateTargetProps{
//   		AlternateTargetGroup: greenTargetGroup,
//   		ProductionListener: ecs.ListenerRuleConfiguration_ApplicationListenerRule(prodListenerRule),
//   	}),
//   })
//
//   target.AttachToApplicationTargetGroup(blueTargetGroup)
//
type LoadBalancerTargetOptions struct {
	// The name of the container.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Alternate target configuration for blue/green deployments.
	// Default: - No alternate target configuration.
	//
	AlternateTarget IAlternateTarget `field:"optional" json:"alternateTarget" yaml:"alternateTarget"`
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

