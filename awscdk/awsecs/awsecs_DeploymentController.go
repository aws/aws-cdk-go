package awsecs


// The deployment controller to use for the service.
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
//   	deploymentController: &deploymentController{
//   		type: ecs.deploymentControllerType_CODE_DEPLOY,
//   	},
//   })
//
type DeploymentController struct {
	// The deployment controller type to use.
	Type DeploymentControllerType `field:"optional" json:"type" yaml:"type"`
}

