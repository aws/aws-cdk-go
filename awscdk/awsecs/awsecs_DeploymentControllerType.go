package awsecs


// The deployment controller type to use for the service.
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
type DeploymentControllerType string

const (
	// The rolling update (ECS) deployment type involves replacing the current running version of the container with the latest version.
	DeploymentControllerType_ECS DeploymentControllerType = "ECS"
	// The blue/green (CODE_DEPLOY) deployment type uses the blue/green deployment model powered by AWS CodeDeploy.
	DeploymentControllerType_CODE_DEPLOY DeploymentControllerType = "CODE_DEPLOY"
	// The external (EXTERNAL) deployment type enables you to use any third-party deployment controller.
	DeploymentControllerType_EXTERNAL DeploymentControllerType = "EXTERNAL"
)

