package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an AWS Lambda function as a deployment lifecycle hook target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var lambdaHook function
//   var blueTargetGroup applicationTargetGroup
//   var greenTargetGroup applicationTargetGroup
//   var prodListenerRule applicationListenerRule
//
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	DeploymentStrategy: ecs.DeploymentStrategy_BLUE_GREEN,
//   })
//
//   service.AddLifecycleHook(ecs.NewDeploymentLifecycleLambdaTarget(lambdaHook, jsii.String("PreScaleHook"), &DeploymentLifecycleLambdaTargetProps{
//   	LifecycleStages: []deploymentLifecycleStage{
//   		ecs.*deploymentLifecycleStage_PRE_SCALE_UP,
//   	},
//   }))
//
//   target := service.LoadBalancerTarget(&LoadBalancerTargetOptions{
//   	ContainerName: jsii.String("nginx"),
//   	ContainerPort: jsii.Number(80),
//   	Protocol: ecs.Protocol_TCP,
//   }, ecs.NewAlternateTarget(jsii.String("AlternateTarget"), &AlternateTargetProps{
//   	AlternateTargetGroup: greenTargetGroup,
//   	ProductionListener: ecs.ListenerRuleConfiguration_ApplicationListenerRule(prodListenerRule),
//   }))
//
//   target.AttachToApplicationTargetGroup(blueTargetGroup)
//
type DeploymentLifecycleLambdaTarget interface {
	IDeploymentLifecycleHookTarget
	// The IAM role for the deployment lifecycle hook target.
	Role() awsiam.IRole
	// Bind this target to a deployment lifecycle hook.
	Bind(scope constructs.IConstruct) *DeploymentLifecycleHookTargetConfig
}

// The jsii proxy struct for DeploymentLifecycleLambdaTarget
type jsiiProxy_DeploymentLifecycleLambdaTarget struct {
	jsiiProxy_IDeploymentLifecycleHookTarget
}

func (j *jsiiProxy_DeploymentLifecycleLambdaTarget) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewDeploymentLifecycleLambdaTarget(handler awslambda.IFunction, id *string, props *DeploymentLifecycleLambdaTargetProps) DeploymentLifecycleLambdaTarget {
	_init_.Initialize()

	if err := validateNewDeploymentLifecycleLambdaTargetParameters(handler, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentLifecycleLambdaTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.DeploymentLifecycleLambdaTarget",
		[]interface{}{handler, id, props},
		&j,
	)

	return &j
}

func NewDeploymentLifecycleLambdaTarget_Override(d DeploymentLifecycleLambdaTarget, handler awslambda.IFunction, id *string, props *DeploymentLifecycleLambdaTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.DeploymentLifecycleLambdaTarget",
		[]interface{}{handler, id, props},
		d,
	)
}

func (d *jsiiProxy_DeploymentLifecycleLambdaTarget) Bind(scope constructs.IConstruct) *DeploymentLifecycleHookTargetConfig {
	if err := d.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DeploymentLifecycleHookTargetConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

