package awsecs


// The deployment controller type to use for the service.
//
// Example:
//   var myApplication ecsApplication
//   var cluster cluster
//   var taskDefinition fargateTaskDefinition
//   var blueTargetGroup iTargetGroup
//   var greenTargetGroup iTargetGroup
//   var listener iApplicationListener
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DeploymentController: &DeploymentController{
//   		Type: ecs.DeploymentControllerType_CODE_DEPLOY,
//   	},
//   })
//
//   codedeploy.NewEcsDeploymentGroup(this, jsii.String("BlueGreenDG"), &EcsDeploymentGroupProps{
//   	Service: Service,
//   	BlueGreenDeploymentConfig: &EcsBlueGreenDeploymentConfig{
//   		BlueTargetGroup: *BlueTargetGroup,
//   		GreenTargetGroup: *GreenTargetGroup,
//   		Listener: *Listener,
//   	},
//   	DeploymentConfig: codedeploy.EcsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
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

