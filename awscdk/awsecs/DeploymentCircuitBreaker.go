package awsecs


// The deployment circuit breaker to use for the service.
//
// Example:
//   var cluster Cluster
//
//   service := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	DesiredCount: jsii.Number(1),
//   	Cpu: jsii.Number(512),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	MinHealthyPercent: jsii.Number(100),
//   	CircuitBreaker: &DeploymentCircuitBreaker{
//   		Rollback: jsii.Boolean(true),
//   	},
//   })
//
type DeploymentCircuitBreaker struct {
	// Whether to enable the deployment circuit breaker logic.
	// Default: true.
	//
	Enable *bool `field:"optional" json:"enable" yaml:"enable"`
	// Whether to enable rollback on deployment failure.
	// Default: false.
	//
	Rollback *bool `field:"optional" json:"rollback" yaml:"rollback"`
}

