package awsecs


// Properties for defining an ECS target.
//
// The port mapping for it must already have been created through addPortMapping().
//
// Example:
//   var cluster Cluster
//   var taskDefinition TaskDefinition
//   var vpc Vpc
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   })
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
//   	Vpc: Vpc,
//   })
//   lb.AddListener(&LoadBalancerListener{
//   	ExternalPort: jsii.Number(80),
//   })
//   lb.AddTarget(service.LoadBalancerTarget(&LoadBalancerTargetOptions{
//   	ContainerName: jsii.String("MyContainer"),
//   	ContainerPort: jsii.Number(80),
//   }))
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

