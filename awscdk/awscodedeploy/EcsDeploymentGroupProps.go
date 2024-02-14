package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for `EcsDeploymentGroup`.
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
type EcsDeploymentGroupProps struct {
	// The configuration options for blue-green ECS deployments.
	BlueGreenDeploymentConfig *EcsBlueGreenDeploymentConfig `field:"required" json:"blueGreenDeploymentConfig" yaml:"blueGreenDeploymentConfig"`
	// The ECS service to deploy with this Deployment Group.
	Service awsecs.IBaseService `field:"required" json:"service" yaml:"service"`
	// The CloudWatch alarms associated with this Deployment Group.
	//
	// CodeDeploy will stop (and optionally roll back)
	// a deployment if during it any of the alarms trigger.
	//
	// Alarms can also be added after the Deployment Group is created using the `#addAlarm` method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	// Default: [].
	//
	Alarms *[]awscloudwatch.IAlarm `field:"optional" json:"alarms" yaml:"alarms"`
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	// Default: One will be created for you.
	//
	Application IEcsApplication `field:"optional" json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Default: - default AutoRollbackConfig.
	//
	AutoRollback *AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The Deployment Configuration this Deployment Group uses.
	// Default: EcsDeploymentConfig.ALL_AT_ONCE
	//
	DeploymentConfig IEcsDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Default: An auto-generated name will be used.
	//
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Whether to skip the step of checking CloudWatch alarms during the deployment process.
	// Default: - false.
	//
	IgnoreAlarmConfiguration *bool `field:"optional" json:"ignoreAlarmConfiguration" yaml:"ignoreAlarmConfiguration"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Default: false.
	//
	IgnorePollAlarmsFailure *bool `field:"optional" json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// The service Role of this Deployment Group.
	// Default: - A new Role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

