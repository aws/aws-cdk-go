# AWS CodeDeploy Construct Library

## Table of Contents

* [Introduction](#introduction)
* Deploying to Amazon EC2 and on-premise instances

  * [EC2/on-premise Applications](#ec2on-premise-applications)
  * [EC2/on-premise Deployment Groups](#ec2on-premise-deployment-groups)
  * [EC2/on-premise Deployment Configurations](#ec2on-premise-deployment-configurations)
* Deploying to AWS Lambda functions

  * [Lambda Applications](#lambda-applications)
  * [Lambda Deployment Groups](#lambda-deployment-groups)
  * [Lambda Deployment Configurations](#lambda-deployment-configurations)
* Deploying to Amazon ECS services

  * [ECS Applications](#ecs-applications)
  * [ECS Deployment Groups](#ecs-deployment-groups)
  * [ECS Deployment Configurations](#ecs-deployment-configurations)
  * [ECS Deployments](#ecs-deployments)

## Introduction

AWS CodeDeploy is a deployment service that automates application deployments to
Amazon EC2 instances, on-premises instances, serverless Lambda functions, or
Amazon ECS services.

The CDK currently supports Amazon EC2, on-premise, AWS Lambda, and Amazon ECS applications.

## EC2/on-premise Applications

To create a new CodeDeploy Application that deploys to EC2/on-premise instances:

```go
application := codedeploy.NewServerApplication(this, jsii.String("CodeDeployApplication"), &ServerApplicationProps{
	ApplicationName: jsii.String("MyApplication"),
})
```

To import an already existing Application:

```go
application := codedeploy.ServerApplication_FromServerApplicationName(this, jsii.String("ExistingCodeDeployApplication"), jsii.String("MyExistingApplication"))
```

## EC2/on-premise Deployment Groups

To create a new CodeDeploy Deployment Group that deploys to EC2/on-premise instances:

```go
import autoscaling "github.com/aws/aws-cdk-go/awscdk"
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"

var application serverApplication
var asg autoScalingGroup
var alarm alarm

deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &ServerDeploymentGroupProps{
	Application: Application,
	DeploymentGroupName: jsii.String("MyDeploymentGroup"),
	AutoScalingGroups: []iAutoScalingGroup{
		asg,
	},
	// adds User Data that installs the CodeDeploy agent on your auto-scaling groups hosts
	// default: true
	InstallAgent: jsii.Boolean(true),
	// adds EC2 instances matching tags
	Ec2InstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
		// any instance with tags satisfying
		// key1=v1 or key1=v2 or key2 (any value) or value v3 (any key)
		// will match this group
		"key1": []*string{
			jsii.String("v1"),
			jsii.String("v2"),
		},
		"key2": []*string{
		},
		"": []*string{
			jsii.String("v3"),
		},
	}),
	// adds on-premise instances matching tags
	OnPremiseInstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
		"key1": []*string{
			jsii.String("v1"),
			jsii.String("v2"),
		},
	}, map[string][]*string{
		"key2": []*string{
			jsii.String("v3"),
		},
	}),
	// CloudWatch alarms
	Alarms: []iAlarm{
		alarm,
	},
	// whether to ignore failure to fetch the status of alarms from CloudWatch
	// default: false
	IgnorePollAlarmsFailure: jsii.Boolean(false),
	// auto-rollback configuration
	AutoRollback: &AutoRollbackConfig{
		FailedDeployment: jsii.Boolean(true),
		 // default: true
		StoppedDeployment: jsii.Boolean(true),
		 // default: false
		DeploymentInAlarm: jsii.Boolean(true),
	},
})
```

All properties are optional - if you don't provide an Application,
one will be automatically created.

To import an already existing Deployment Group:

```go
var application serverApplication

deploymentGroup := codedeploy.ServerDeploymentGroup_FromServerDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &ServerDeploymentGroupAttributes{
	Application: Application,
	DeploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
})
```

### Load balancers

You can [specify a load balancer](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html)
with the `loadBalancer` property when creating a Deployment Group.

`LoadBalancer` is an abstract class with static factory methods that allow you to create instances of it from various sources.

With Classic Elastic Load Balancer, you provide it directly:

```go
import elb "github.com/aws/aws-cdk-go/awscdk"

var lb loadBalancer

lb.AddListener(&LoadBalancerListener{
	ExternalPort: jsii.Number(80),
})

deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &ServerDeploymentGroupProps{
	LoadBalancer: codedeploy.LoadBalancer_Classic(lb),
})
```

With Application Load Balancer or Network Load Balancer,
you provide a Target Group as the load balancer:

```go
var alb applicationLoadBalancer

listener := alb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(80),
})
targetGroup := listener.AddTargets(jsii.String("Fleet"), &AddApplicationTargetsProps{
	Port: jsii.Number(80),
})

deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &ServerDeploymentGroupProps{
	LoadBalancer: codedeploy.LoadBalancer_Application(targetGroup),
})
```

## EC2/on-premise Deployment Configurations

You can also pass a Deployment Configuration when creating the Deployment Group:

```go
deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &ServerDeploymentGroupProps{
	DeploymentConfig: codedeploy.ServerDeploymentConfig_ALL_AT_ONCE(),
})
```

The default Deployment Configuration is `ServerDeploymentConfig.ONE_AT_A_TIME`.

You can also create a custom Deployment Configuration:

```go
deploymentConfig := codedeploy.NewServerDeploymentConfig(this, jsii.String("DeploymentConfiguration"), &ServerDeploymentConfigProps{
	DeploymentConfigName: jsii.String("MyDeploymentConfiguration"),
	 // optional property
	// one of these is required, but both cannot be specified at the same time
	MinimumHealthyHosts: codedeploy.MinimumHealthyHosts_Count(jsii.Number(2)),
})
```

Or import an existing one:

```go
deploymentConfig := codedeploy.ServerDeploymentConfig_FromServerDeploymentConfigName(this, jsii.String("ExistingDeploymentConfiguration"), jsii.String("MyExistingDeploymentConfiguration"))
```

## Lambda Applications

To create a new CodeDeploy Application that deploys to a Lambda function:

```go
application := codedeploy.NewLambdaApplication(this, jsii.String("CodeDeployApplication"), &LambdaApplicationProps{
	ApplicationName: jsii.String("MyApplication"),
})
```

To import an already existing Application:

```go
application := codedeploy.LambdaApplication_FromLambdaApplicationName(this, jsii.String("ExistingCodeDeployApplication"), jsii.String("MyExistingApplication"))
```

## Lambda Deployment Groups

To enable traffic shifting deployments for Lambda functions, CodeDeploy uses Lambda Aliases, which can balance incoming traffic between two different versions of your function.
Before deployment, the alias sends 100% of invokes to the version used in production.
When you publish a new version of the function to your stack, CodeDeploy will send a small percentage of traffic to the new version, monitor, and validate before shifting 100% of traffic to the new version.

To create a new CodeDeploy Deployment Group that deploys to a Lambda function:

```go
var myApplication lambdaApplication
var func function

version := func.currentVersion
version1Alias := lambda.NewAlias(this, jsii.String("alias"), &AliasProps{
	AliasName: jsii.String("prod"),
	Version: Version,
})

deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
	Application: myApplication,
	 // optional property: one will be created for you if not provided
	Alias: version1Alias,
	DeploymentConfig: codedeploy.LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
})
```

In order to deploy a new version of this function:

1. Reference the version with the latest changes `const version = func.currentVersion`.
2. Re-deploy the stack (this will trigger a deployment).
3. Monitor the CodeDeploy deployment as traffic shifts between the versions.

### Lambda Deployment Rollbacks and Alarms

CodeDeploy will roll back if the deployment fails. You can optionally trigger a rollback when one or more alarms are in a failed state:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var alias alias

// or add alarms to an existing group
var blueGreenAlias alias

alarm := cloudwatch.NewAlarm(this, jsii.String("Errors"), &AlarmProps{
	ComparisonOperator: cloudwatch.ComparisonOperator_GREATER_THAN_THRESHOLD,
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(1),
	Metric: alias.metricErrors(),
})
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
	Alias: Alias,
	DeploymentConfig: codedeploy.LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
	Alarms: []iAlarm{
		alarm,
	},
})
deploymentGroup.AddAlarm(cloudwatch.NewAlarm(this, jsii.String("BlueGreenErrors"), &AlarmProps{
	ComparisonOperator: cloudwatch.ComparisonOperator_GREATER_THAN_THRESHOLD,
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(1),
	Metric: blueGreenAlias.metricErrors(),
}))
```

### Pre and Post Hooks

CodeDeploy allows you to run an arbitrary Lambda function before traffic shifting actually starts (PreTraffic Hook) and after it completes (PostTraffic Hook).
With either hook, you have the opportunity to run logic that determines whether the deployment must succeed or fail.
For example, with PreTraffic hook you could run integration tests against the newly created Lambda version (but not serving traffic). With PostTraffic hook, you could run end-to-end validation checks.

```go
var warmUpUserCache function
var endToEndValidation function
var alias alias


// pass a hook whe creating the deployment group
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
	Alias: alias,
	DeploymentConfig: codedeploy.LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
	PreHook: warmUpUserCache,
})

// or configure one on an existing deployment group
deploymentGroup.AddPostHook(endToEndValidation)
```

### Import an existing Lambda Deployment Group

To import an already existing Deployment Group:

```go
var application lambdaApplication

deploymentGroup := codedeploy.LambdaDeploymentGroup_FromLambdaDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &LambdaDeploymentGroupAttributes{
	Application: Application,
	DeploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
})
```

## Lambda Deployment Configurations

CodeDeploy for Lambda comes with predefined configurations for traffic shifting.
The predefined configurations are available as LambdaDeploymentConfig constants.

```go
var application lambdaApplication
var alias alias
config := codedeploy.LambdaDeploymentConfig_CANARY_10PERCENT_30MINUTES()
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
	Application: Application,
	Alias: Alias,
	DeploymentConfig: config,
})
```

If you want to specify your own strategy,
you can do so with the LambdaDeploymentConfig construct,
letting you specify precisely how fast a new function version is deployed.

```go
var application lambdaApplication
var alias alias
config := codedeploy.NewLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &LambdaDeploymentConfigProps{
	TrafficRouting: codedeploy.NewTimeBasedCanaryTrafficRouting(&TimeBasedCanaryTrafficRoutingProps{
		Interval: awscdk.Duration_Minutes(jsii.Number(15)),
		Percentage: jsii.Number(5),
	}),
})
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
	Application: Application,
	Alias: Alias,
	DeploymentConfig: config,
})
```

You can specify a custom name for your deployment config, but if you do you will not be able to update the interval/percentage through CDK.

```go
config := codedeploy.NewLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &LambdaDeploymentConfigProps{
	TrafficRouting: codedeploy.NewTimeBasedCanaryTrafficRouting(&TimeBasedCanaryTrafficRoutingProps{
		Interval: awscdk.Duration_Minutes(jsii.Number(15)),
		Percentage: jsii.Number(5),
	}),
	DeploymentConfigName: jsii.String("MyDeploymentConfig"),
})
```

To import an already existing Deployment Config:

```go
deploymentConfig := codedeploy.LambdaDeploymentConfig_FromLambdaDeploymentConfigName(this, jsii.String("ExistingDeploymentConfiguration"), jsii.String("MyExistingDeploymentConfiguration"))
```

## ECS Applications

To create a new CodeDeploy Application that deploys an ECS service:

```go
application := codedeploy.NewEcsApplication(this, jsii.String("CodeDeployApplication"), &EcsApplicationProps{
	ApplicationName: jsii.String("MyApplication"),
})
```

To import an already existing Application:

```go
application := codedeploy.EcsApplication_FromEcsApplicationName(this, jsii.String("ExistingCodeDeployApplication"), jsii.String("MyExistingApplication"))
```

## ECS Deployment Groups

CodeDeploy can be used to deploy to load-balanced ECS services.
CodeDeploy performs ECS blue-green deployments by managing ECS task sets and load balancer
target groups.  During a blue-green deployment, one task set and target group runs the
original version of your ECS task definition ('blue') and another task set and target group
runs the new version of your ECS task definition ('green').

CodeDeploy orchestrates traffic shifting during ECS blue-green deployments by using
a load balancer listener to balance incoming traffic between the 'blue' and 'green' task sets/target groups
running two different versions of your ECS task definition.
Before deployment, the load balancer listener sends 100% of requests to the 'blue' target group.
When you publish a new version of the task definition and start a CodeDeploy deployment,
CodeDeploy can send a small percentage of traffic to the new 'green' task set behind the 'green' target group,
monitor, and validate before shifting 100% of traffic to the new version.

To create a new CodeDeploy Deployment Group that deploys to an ECS service:

```go
var myApplication ecsApplication
var cluster cluster
var taskDefinition fargateTaskDefinition
var blueTargetGroup iTargetGroup
var greenTargetGroup iTargetGroup
var listener iApplicationListener


service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	DeploymentController: &DeploymentController{
		Type: ecs.DeploymentControllerType_CODE_DEPLOY,
	},
})

codedeploy.NewEcsDeploymentGroup(this, jsii.String("BlueGreenDG"), &EcsDeploymentGroupProps{
	Service: Service,
	BlueGreenDeploymentConfig: &EcsBlueGreenDeploymentConfig{
		BlueTargetGroup: *BlueTargetGroup,
		GreenTargetGroup: *GreenTargetGroup,
		Listener: *Listener,
	},
	DeploymentConfig: codedeploy.EcsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
})
```

In order to deploy a new task definition version to the ECS service,
deploy the changes directly through CodeDeploy using the CodeDeploy APIs or console.
When the `CODE_DEPLOY` deployment controller is used, the ECS service cannot be
deployed with a new task definition version through CloudFormation.

For more information on the behavior of CodeDeploy blue-green deployments for ECS, see
[What happens during an Amazon ECS deployment](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-steps-ecs.html#deployment-steps-what-happens)
in the CodeDeploy user guide.

Note: If you wish to deploy updates to your ECS service through CDK and CloudFormation instead of directly through CodeDeploy,
using the [`CfnCodeDeployBlueGreenHook`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.CfnCodeDeployBlueGreenHook.html)
construct is the recommended approach instead of using the `EcsDeploymentGroup` construct.  For a comparison
of ECS blue-green deployments through CodeDeploy (using `EcsDeploymentGroup`) and through CloudFormation (using `CfnCodeDeployBlueGreenHook`),
see [Create an Amazon ECS blue/green deployment through AWS CloudFormation](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployments-create-ecs-cfn.html#differences-ecs-bg-cfn)
in the CloudFormation user guide.

### ECS Deployment Rollbacks and Alarms

CodeDeploy will automatically roll back if a deployment fails.
You can optionally trigger an automatic rollback when one or more alarms are in a failed state during a deployment, or if the deployment stops.

In this example, CodeDeploy will monitor and roll back on alarms set for the
number of unhealthy ECS tasks in each of the blue and green target groups,
as well as alarms set for the number HTTP 5xx responses seen in each of the blue
and green target groups.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var service fargateService
var blueTargetGroup applicationTargetGroup
var greenTargetGroup applicationTargetGroup
var listener iApplicationListener


// Alarm on the number of unhealthy ECS tasks in each target group
blueUnhealthyHosts := cloudwatch.NewAlarm(this, jsii.String("BlueUnhealthyHosts"), &AlarmProps{
	AlarmName: jsii.String(awscdk.stack_Of(this).stackName + "-Unhealthy-Hosts-Blue"),
	Metric: blueTargetGroup.MetricUnhealthyHostCount(),
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(2),
})

greenUnhealthyHosts := cloudwatch.NewAlarm(this, jsii.String("GreenUnhealthyHosts"), &AlarmProps{
	AlarmName: jsii.String(awscdk.stack_Of(this).stackName + "-Unhealthy-Hosts-Green"),
	Metric: greenTargetGroup.*MetricUnhealthyHostCount(),
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(2),
})

// Alarm on the number of HTTP 5xx responses returned by each target group
blueApiFailure := cloudwatch.NewAlarm(this, jsii.String("Blue5xx"), &AlarmProps{
	AlarmName: jsii.String(awscdk.stack_Of(this).stackName + "-Http-5xx-Blue"),
	Metric: blueTargetGroup.MetricHttpCodeTarget(elbv2.HttpCodeTarget_TARGET_5XX_COUNT, &MetricOptions{
		Period: awscdk.Duration_Minutes(jsii.Number(1)),
	}),
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(1),
})

greenApiFailure := cloudwatch.NewAlarm(this, jsii.String("Green5xx"), &AlarmProps{
	AlarmName: jsii.String(awscdk.stack_Of(this).stackName + "-Http-5xx-Green"),
	Metric: greenTargetGroup.*MetricHttpCodeTarget(elbv2.HttpCodeTarget_TARGET_5XX_COUNT, &MetricOptions{
		Period: awscdk.Duration_*Minutes(jsii.Number(1)),
	}),
	Threshold: jsii.Number(1),
	EvaluationPeriods: jsii.Number(1),
})

codedeploy.NewEcsDeploymentGroup(this, jsii.String("BlueGreenDG"), &EcsDeploymentGroupProps{
	// CodeDeploy will monitor these alarms during a deployment and automatically roll back
	Alarms: []iAlarm{
		blueUnhealthyHosts,
		greenUnhealthyHosts,
		blueApiFailure,
		greenApiFailure,
	},
	AutoRollback: &AutoRollbackConfig{
		// CodeDeploy will automatically roll back if a deployment is stopped
		StoppedDeployment: jsii.Boolean(true),
	},
	Service: Service,
	BlueGreenDeploymentConfig: &EcsBlueGreenDeploymentConfig{
		BlueTargetGroup: *BlueTargetGroup,
		GreenTargetGroup: *GreenTargetGroup,
		Listener: *Listener,
	},
	DeploymentConfig: codedeploy.EcsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
})
```

### Deployment validation and manual deployment approval

CodeDeploy blue-green deployments provide an opportunity to validate the new task definition version running on
the 'green' ECS task set prior to shifting any production traffic to the new version.  A second 'test' listener
serving traffic on a different port be added to the load balancer. For example, the test listener can serve
test traffic on port 9001 while the main listener serves production traffic on port 443.
During a blue-green deployment, CodeDeploy can then shift 100% of test traffic over to the 'green'
task set/target group prior to shifting any production traffic during the deployment.

```go
var myApplication ecsApplication
var service fargateService
var blueTargetGroup iTargetGroup
var greenTargetGroup iTargetGroup
var listener iApplicationListener
var testListener iApplicationListener


codedeploy.NewEcsDeploymentGroup(this, jsii.String("BlueGreenDG"), &EcsDeploymentGroupProps{
	Service: Service,
	BlueGreenDeploymentConfig: &EcsBlueGreenDeploymentConfig{
		BlueTargetGroup: *BlueTargetGroup,
		GreenTargetGroup: *GreenTargetGroup,
		Listener: *Listener,
		TestListener: *TestListener,
	},
	DeploymentConfig: codedeploy.EcsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
})
```

Automated validation steps can run during the CodeDeploy deployment after shifting test traffic and before
shifting production traffic.  CodeDeploy supports registering Lambda functions as lifecycle hooks for
an ECS deployment.  These Lambda functions can run automated validation steps against the test traffic
port, for example in response to the `AfterAllowTestTraffic` lifecycle hook.  For more information about
how to specify the Lambda functions to run for each CodeDeploy lifecycle hook in an ECS deployment, see the
[AppSpec 'hooks' for an Amazon ECS deployment](https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-ecs)
section in the CodeDeploy user guide.

After provisioning the 'green' ECS task set and re-routing test traffic during a blue-green deployment,
CodeDeploy can wait for approval before continuing the deployment and re-routing production traffic.
During this approval wait time, you can complete additional validation steps prior to exposing the new
'green' task set to production traffic, such as manual testing through the test listener port or
running automated integration test suites.

To approve the deployment, validation steps use the CodeDeploy
[ContinueDeployment API(https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_ContinueDeployment.html).
If the ContinueDeployment API is not called within the approval wait time period, CodeDeploy will stop the
deployment and can automatically roll back the deployment.

```go
var service fargateService
var blueTargetGroup iTargetGroup
var greenTargetGroup iTargetGroup
var listener iApplicationListener
var testListener iApplicationListener


codedeploy.NewEcsDeploymentGroup(this, jsii.String("BlueGreenDG"), &EcsDeploymentGroupProps{
	AutoRollback: &AutoRollbackConfig{
		// CodeDeploy will automatically roll back if the 8-hour approval period times out and the deployment stops
		StoppedDeployment: jsii.Boolean(true),
	},
	Service: Service,
	BlueGreenDeploymentConfig: &EcsBlueGreenDeploymentConfig{
		// The deployment will wait for approval for up to 8 hours before stopping the deployment
		DeploymentApprovalWaitTime: awscdk.Duration_Hours(jsii.Number(8)),
		BlueTargetGroup: *BlueTargetGroup,
		GreenTargetGroup: *GreenTargetGroup,
		Listener: *Listener,
		TestListener: *TestListener,
	},
	DeploymentConfig: codedeploy.EcsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
})
```

### Deployment bake time

You can specify how long CodeDeploy waits before it terminates the original 'blue' ECS task set when a blue-green deployment
is complete in order to let the deployment "bake" a while. During this bake time, CodeDeploy will continue to monitor any
CloudWatch alarms specified for the deployment group and will automatically roll back if those alarms go into a failed state.

```go
import "github.com/aws/aws-cdk-go/awscdk"

var service fargateService
var blueTargetGroup iTargetGroup
var greenTargetGroup iTargetGroup
var listener iApplicationListener
var blueUnhealthyHosts alarm
var greenUnhealthyHosts alarm
var blueApiFailure alarm
var greenApiFailure alarm


codedeploy.NewEcsDeploymentGroup(this, jsii.String("BlueGreenDG"), &EcsDeploymentGroupProps{
	Service: Service,
	BlueGreenDeploymentConfig: &EcsBlueGreenDeploymentConfig{
		BlueTargetGroup: *BlueTargetGroup,
		GreenTargetGroup: *GreenTargetGroup,
		Listener: *Listener,
		// CodeDeploy will wait for 30 minutes after completing the blue-green deployment before it terminates the blue tasks
		TerminationWaitTime: awscdk.Duration_Minutes(jsii.Number(30)),
	},
	// CodeDeploy will continue to monitor these alarms during the 30-minute bake time and will automatically
	// roll back if they go into a failed state at any point during the deployment.
	Alarms: []iAlarm{
		blueUnhealthyHosts,
		greenUnhealthyHosts,
		blueApiFailure,
		greenApiFailure,
	},
	DeploymentConfig: codedeploy.EcsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
})
```

### Import an existing ECS Deployment Group

To import an already existing Deployment Group:

```go
var application ecsApplication

deploymentGroup := codedeploy.EcsDeploymentGroup_FromEcsDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &EcsDeploymentGroupAttributes{
	Application: Application,
	DeploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
})
```

## ECS Deployment Configurations

CodeDeploy for ECS comes with predefined configurations for traffic shifting.
The predefined configurations are available as LambdaDeploymentConfig constants.

```go
config := codedeploy.EcsDeploymentConfig_CANARY_10PERCENT_5MINUTES()
```

If you want to specify your own strategy,
you can do so with the EcsDeploymentConfig construct,
letting you specify precisely how fast an ECS service is deployed.

```go
codedeploy.NewEcsDeploymentConfig(this, jsii.String("CustomConfig"), &EcsDeploymentConfigProps{
	TrafficRouting: codedeploy.NewTimeBasedCanaryTrafficRouting(&TimeBasedCanaryTrafficRoutingProps{
		Interval: awscdk.Duration_Minutes(jsii.Number(15)),
		Percentage: jsii.Number(5),
	}),
})
```

You can specify a custom name for your deployment config, but if you do you will not be able to update the interval/percentage through CDK.

```go
config := codedeploy.NewEcsDeploymentConfig(this, jsii.String("CustomConfig"), &EcsDeploymentConfigProps{
	TrafficRouting: codedeploy.NewTimeBasedCanaryTrafficRouting(&TimeBasedCanaryTrafficRoutingProps{
		Interval: awscdk.Duration_Minutes(jsii.Number(15)),
		Percentage: jsii.Number(5),
	}),
	DeploymentConfigName: jsii.String("MyDeploymentConfig"),
})
```

Or import an existing one:

```go
deploymentConfig := codedeploy.EcsDeploymentConfig_FromEcsDeploymentConfigName(this, jsii.String("ExistingDeploymentConfiguration"), jsii.String("MyExistingDeploymentConfiguration"))
```

## ECS Deployments

[![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg)](https://constructs.dev/packages/@cdklabs/cdk-ecs-codedeploy)

An experimental construct is available on the Construct Hub called [@cdklabs/cdk-ecs-codedeploy](https://constructs.dev/packages/@cdklabs/cdk-ecs-codedeploy) that manages ECS CodeDeploy deployments.

```go
var deploymentGroup iEcsDeploymentGroup
var taskDefinition iTaskDefinition


NewEcsDeployment(map[string]interface{}{
	"deploymentGroup": deploymentGroup,
	"targetService": map[string]interface{}{
		"taskDefinition": taskDefinition,
		"containerName": jsii.String("mycontainer"),
		"containerPort": jsii.Number(80),
	},
})
```

The deployment will use the AutoRollbackConfig for the EcsDeploymentGroup unless it is overridden in the deployment:

```go
var deploymentGroup iEcsDeploymentGroup
var taskDefinition iTaskDefinition


NewEcsDeployment(map[string]interface{}{
	"deploymentGroup": deploymentGroup,
	"targetService": map[string]interface{}{
		"taskDefinition": taskDefinition,
		"containerName": jsii.String("mycontainer"),
		"containerPort": jsii.Number(80),
	},
	"autoRollback": map[string]*bool{
		"failedDeployment": jsii.Boolean(true),
		"deploymentInAlarm": jsii.Boolean(true),
		"stoppedDeployment": jsii.Boolean(false),
	},
})
```

By default, the CodeDeploy Deployment will timeout after 30 minutes. The timeout value can be overridden:

```go
var deploymentGroup iEcsDeploymentGroup
var taskDefinition iTaskDefinition


NewEcsDeployment(map[string]interface{}{
	"deploymentGroup": deploymentGroup,
	"targetService": map[string]interface{}{
		"taskDefinition": taskDefinition,
		"containerName": jsii.String("mycontainer"),
		"containerPort": jsii.Number(80),
	},
	"timeout": awscdk.Duration_minutes(jsii.Number(60)),
})
```
