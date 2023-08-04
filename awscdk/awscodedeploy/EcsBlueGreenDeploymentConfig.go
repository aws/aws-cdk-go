package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Specify how the deployment behaves and how traffic is routed to the ECS service during a blue-green ECS deployment.
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
// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-ecs
//
type EcsBlueGreenDeploymentConfig struct {
	// The target group that will be associated with the 'blue' ECS task set during a blue-green deployment.
	BlueTargetGroup awselasticloadbalancingv2.ITargetGroup `field:"required" json:"blueTargetGroup" yaml:"blueTargetGroup"`
	// The target group that will be associated with the 'green' ECS task set during a blue-green deployment.
	GreenTargetGroup awselasticloadbalancingv2.ITargetGroup `field:"required" json:"greenTargetGroup" yaml:"greenTargetGroup"`
	// The load balancer listener used to serve production traffic and to shift production traffic from the 'blue' ECS task set to the 'green' ECS task set during a blue-green deployment.
	Listener awselasticloadbalancingv2.IListener `field:"required" json:"listener" yaml:"listener"`
	// Specify how long CodeDeploy waits for approval to continue a blue-green deployment before it stops the deployment.
	//
	// After provisioning the 'green' ECS task set and re-routing test traffic, CodeDeploy can wait for approval before
	// continuing the deployment and re-routing production traffic.  During this wait time, validation such as manual
	// testing or running integration tests can occur using the test traffic port, prior to exposing the new 'green' task
	// set to production traffic.  To approve the deployment, validation steps use the CodeDeploy
	// [ContinueDeployment API(https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_ContinueDeployment.html).
	// If the ContinueDeployment API is not called within the wait time period, CodeDeploy will stop the deployment.
	//
	// By default, CodeDeploy will not wait for deployment approval.  After re-routing test traffic to the 'green' ECS task set
	// and running any 'AfterAllowTestTraffic' and 'BeforeAllowTraffic' lifecycle hooks, the deployment will immediately
	// re-route production traffic to the 'green' ECS task set.
	// Default: 0.
	//
	DeploymentApprovalWaitTime awscdk.Duration `field:"optional" json:"deploymentApprovalWaitTime" yaml:"deploymentApprovalWaitTime"`
	// Specify how long CodeDeploy waits before it terminates the original 'blue' ECS task set when a blue-green deployment is complete.
	//
	// During this wait time, CodeDeploy will continue to monitor any CloudWatch alarms specified for the deployment group,
	// and the deployment group can be configured to automatically roll back if those alarms fire.  Once CodeDeploy begins to
	// terminate the 'blue' ECS task set, the deployment can no longer be rolled back, manually or automatically.
	//
	// By default, the deployment will immediately terminate the 'blue' ECS task set after production traffic is successfully
	// routed to the 'green' ECS task set.
	// Default: 0.
	//
	TerminationWaitTime awscdk.Duration `field:"optional" json:"terminationWaitTime" yaml:"terminationWaitTime"`
	// The load balancer listener used to route test traffic to the 'green' ECS task set during a blue-green deployment.
	//
	// During a blue-green deployment, validation can occur after test traffic has been re-routed and before production
	// traffic has been re-routed to the 'green' ECS task set.  You can specify one or more Lambda funtions in the
	// deployment's AppSpec file that run during the AfterAllowTestTraffic hook. The functions can run validation tests.
	// If a validation test fails, a deployment rollback is triggered. If the validation tests succeed, the next hook in
	// the deployment lifecycle, BeforeAllowTraffic, is triggered.
	//
	// If a test listener is not specified, the deployment will proceed to routing the production listener to the 'green' ECS task set
	// and will skip the AfterAllowTestTraffic hook.
	// Default: No test listener will be added.
	//
	TestListener awselasticloadbalancingv2.IListener `field:"optional" json:"testListener" yaml:"testListener"`
}

