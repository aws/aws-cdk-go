package awsecs


// The deployment controller type to use for the service.
//
// Example:
//   var cluster cluster
//
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
//   	Cluster: Cluster,
//   	MemoryLimitMiB: jsii.Number(1024),
//   	DesiredCount: jsii.Number(1),
//   	Cpu: jsii.Number(512),
//   	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	DeploymentController: &DeploymentController{
//   		Type: ecs.DeploymentControllerType_CODE_DEPLOY,
//   	},
//   })
//
// Experimental.
type DeploymentControllerType string

const (
	// The rolling update (ECS) deployment type involves replacing the current running version of the container with the latest version.
	// Experimental.
	DeploymentControllerType_ECS DeploymentControllerType = "ECS"
	// The blue/green (CODE_DEPLOY) deployment type uses the blue/green deployment model powered by AWS CodeDeploy.
	// Experimental.
	DeploymentControllerType_CODE_DEPLOY DeploymentControllerType = "CODE_DEPLOY"
	// The external (EXTERNAL) deployment type enables you to use any third-party deployment controller.
	// Experimental.
	DeploymentControllerType_EXTERNAL DeploymentControllerType = "EXTERNAL"
)

