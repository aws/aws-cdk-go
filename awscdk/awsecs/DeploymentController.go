package awsecs


// The deployment controller to use for the service.
//
// Example:
//   var myApplication EcsApplication
//   var cluster Cluster
//   var taskDefinition FargateTaskDefinition
//   var blueTargetGroup ITargetGroup
//   var greenTargetGroup ITargetGroup
//   var listener IApplicationListener
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
type DeploymentController struct {
	// The deployment controller type to use.
	// Default: DeploymentControllerType.ECS
	//
	Type DeploymentControllerType `field:"optional" json:"type" yaml:"type"`
}

