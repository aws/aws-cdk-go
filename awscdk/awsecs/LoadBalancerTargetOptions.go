package awsecs


// Properties for defining an ECS target.
//
// The port mapping for it must already have been created through addPortMapping().
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
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
// Experimental.
type LoadBalancerTargetOptions struct {
	// The name of the container.
	// Experimental.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The port number of the container.
	//
	// Only applicable when using application/network load balancers.
	// Experimental.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// The protocol used for the port mapping.
	//
	// Only applicable when using application load balancers.
	// Experimental.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

