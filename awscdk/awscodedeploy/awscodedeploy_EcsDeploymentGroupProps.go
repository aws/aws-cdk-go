package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for {@link EcsDeploymentGroup}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myApplication ecsApplication
//   var cluster ecs.Cluster
//   var taskDefinition ecs.FargateTaskDefinition
//   var blueTargetGroup elbv2.ITargetGroup
//   var greenTargetGroup elbv2.ITargetGroup
//   var listener elbv2.IApplicationListener
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), map[string]interface{}{
//   	"cluster": cluster,
//   	"taskDefinition": taskDefinition,
//   	"deploymentController": map[string]interface{}{
//   		"type": ecs.DeploymentControllerType_CODE_DEPLOY,
//   	},
//   })
//
//   codedeploy.NewEcsDeploymentGroup(stack, jsii.String("BlueGreenDG"), &ecsDeploymentGroupProps{
//   	service: service,
//   	blueGreenDeploymentConfig: &ecsBlueGreenDeploymentConfig{
//   		blueTargetGroup: blueTargetGroup,
//   		greenTargetGroup: greenTargetGroup,
//   		listener: listener,
//   	},
//   	deploymentConfig: codedeploy.ecsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
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
	// Alarms can also be added after the Deployment Group is created using the {@link #addAlarm} method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	Alarms *[]awscloudwatch.IAlarm `field:"optional" json:"alarms" yaml:"alarms"`
	// The reference to the CodeDeploy ECS Application that this Deployment Group belongs to.
	Application IEcsApplication `field:"optional" json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	AutoRollback *AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The Deployment Configuration this Deployment Group uses.
	DeploymentConfig IEcsDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	IgnorePollAlarmsFailure *bool `field:"optional" json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// The service Role of this Deployment Group.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

