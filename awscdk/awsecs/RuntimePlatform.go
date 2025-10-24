package awsecs


// The interface for Runtime Platform.
//
// Example:
//   var cluster Cluster
//
//   applicationLoadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(512),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	MinHealthyPercent: jsii.Number(100),
//   	RuntimePlatform: &RuntimePlatform{
//   		CpuArchitecture: ecs.CpuArchitecture_ARM64(),
//   		OperatingSystemFamily: ecs.OperatingSystemFamily_LINUX(),
//   	},
//   })
//
type RuntimePlatform struct {
	// The CpuArchitecture for Fargate Runtime Platform.
	// Default: - Undefined.
	//
	CpuArchitecture CpuArchitecture `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// The operating system for Fargate Runtime Platform.
	// Default: - Undefined.
	//
	OperatingSystemFamily OperatingSystemFamily `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

