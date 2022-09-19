package awsapplicationautoscaling


// Properties for enabling Application Auto Scaling.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
//   	minCapacity: jsii.Number(1),
//   	maxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
//   scalableTarget.scaleOnMemoryUtilization(jsii.String("MemoryScaling"), &memoryUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
type EnableScalingProps struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

