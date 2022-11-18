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

## Introduction

AWS CodeDeploy is a deployment service that automates application deployments to
Amazon EC2 instances, on-premises instances, serverless Lambda functions, or
Amazon ECS services.

The CDK currently supports Amazon EC2, on-premise, AWS Lambda, and Amazon ECS applications.

## EC2/on-premise Applications

To create a new CodeDeploy Application that deploys to EC2/on-premise instances:

```go
application := codedeploy.NewServerApplication(this, jsii.String("CodeDeployApplication"), &serverApplicationProps{
	applicationName: jsii.String("MyApplication"),
})
```

To import an already existing Application:

```go
application := codedeploy.serverApplication.fromServerApplicationName(this, jsii.String("ExistingCodeDeployApplication"), jsii.String("MyExistingApplication"))
```

## EC2/on-premise Deployment Groups

To create a new CodeDeploy Deployment Group that deploys to EC2/on-premise instances:

```go
import autoscaling "github.com/aws/aws-cdk-go/awscdk"
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"

var application serverApplication
var asg autoScalingGroup
var alarm alarm

deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &serverDeploymentGroupProps{
	application: application,
	deploymentGroupName: jsii.String("MyDeploymentGroup"),
	autoScalingGroups: []iAutoScalingGroup{
		asg,
	},
	// adds User Data that installs the CodeDeploy agent on your auto-scaling groups hosts
	// default: true
	installAgent: jsii.Boolean(true),
	// adds EC2 instances matching tags
	ec2InstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
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
	onPremiseInstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
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
	alarms: []iAlarm{
		alarm,
	},
	// whether to ignore failure to fetch the status of alarms from CloudWatch
	// default: false
	ignorePollAlarmsFailure: jsii.Boolean(false),
	// auto-rollback configuration
	autoRollback: &autoRollbackConfig{
		failedDeployment: jsii.Boolean(true),
		 // default: true
		stoppedDeployment: jsii.Boolean(true),
		 // default: false
		deploymentInAlarm: jsii.Boolean(true),
	},
})
```

All properties are optional - if you don't provide an Application,
one will be automatically created.

To import an already existing Deployment Group:

```go
var application serverApplication

deploymentGroup := codedeploy.serverDeploymentGroup.fromServerDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &serverDeploymentGroupAttributes{
	application: application,
	deploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
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

lb.addListener(&loadBalancerListener{
	externalPort: jsii.Number(80),
})

deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &serverDeploymentGroupProps{
	loadBalancer: codedeploy.loadBalancer.classic(lb),
})
```

With Application Load Balancer or Network Load Balancer,
you provide a Target Group as the load balancer:

```go
import elbv2 "github.com/aws/aws-cdk-go/awscdk"

var alb applicationLoadBalancer

listener := alb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
	port: jsii.Number(80),
})
targetGroup := listener.addTargets(jsii.String("Fleet"), &addApplicationTargetsProps{
	port: jsii.Number(80),
})

deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &serverDeploymentGroupProps{
	loadBalancer: codedeploy.loadBalancer.application(targetGroup),
})
```

## EC2/on-premise Deployment Configurations

You can also pass a Deployment Configuration when creating the Deployment Group:

```go
deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &serverDeploymentGroupProps{
	deploymentConfig: codedeploy.serverDeploymentConfig_ALL_AT_ONCE(),
})
```

The default Deployment Configuration is `ServerDeploymentConfig.ONE_AT_A_TIME`.

You can also create a custom Deployment Configuration:

```go
deploymentConfig := codedeploy.NewServerDeploymentConfig(this, jsii.String("DeploymentConfiguration"), &serverDeploymentConfigProps{
	deploymentConfigName: jsii.String("MyDeploymentConfiguration"),
	 // optional property
	// one of these is required, but both cannot be specified at the same time
	minimumHealthyHosts: codedeploy.minimumHealthyHosts.count(jsii.Number(2)),
})
```

Or import an existing one:

```go
deploymentConfig := codedeploy.serverDeploymentConfig.fromServerDeploymentConfigName(this, jsii.String("ExistingDeploymentConfiguration"), jsii.String("MyExistingDeploymentConfiguration"))
```

## Lambda Applications

To create a new CodeDeploy Application that deploys to a Lambda function:

```go
application := codedeploy.NewLambdaApplication(this, jsii.String("CodeDeployApplication"), &lambdaApplicationProps{
	applicationName: jsii.String("MyApplication"),
})
```

To import an already existing Application:

```go
application := codedeploy.lambdaApplication.fromLambdaApplicationName(this, jsii.String("ExistingCodeDeployApplication"), jsii.String("MyExistingApplication"))
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
version1Alias := lambda.NewAlias(this, jsii.String("alias"), &aliasProps{
	aliasName: jsii.String("prod"),
	version: version,
})

deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
	application: myApplication,
	 // optional property: one will be created for you if not provided
	alias: version1Alias,
	deploymentConfig: codedeploy.lambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
})
```

In order to deploy a new version of this function:

1. Reference the version with the latest changes `const version = func.currentVersion`.
2. Re-deploy the stack (this will trigger a deployment).
3. Monitor the CodeDeploy deployment as traffic shifts between the versions.

### Lambda Deployment Rollbacks and Alarms

CodeDeploy will roll back if the deployment fails. You can optionally trigger a rollback when one or more alarms are in a failed state:

```go
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"

var alias alias

// or add alarms to an existing group
var blueGreenAlias alias

alarm := cloudwatch.NewAlarm(this, jsii.String("Errors"), &alarmProps{
	comparisonOperator: cloudwatch.comparisonOperator_GREATER_THAN_THRESHOLD,
	threshold: jsii.Number(1),
	evaluationPeriods: jsii.Number(1),
	metric: alias.metricErrors(),
})
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
	alias: alias,
	deploymentConfig: codedeploy.lambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
	alarms: []iAlarm{
		alarm,
	},
})
deploymentGroup.addAlarm(cloudwatch.NewAlarm(this, jsii.String("BlueGreenErrors"), &alarmProps{
	comparisonOperator: cloudwatch.*comparisonOperator_GREATER_THAN_THRESHOLD,
	threshold: jsii.Number(1),
	evaluationPeriods: jsii.Number(1),
	metric: blueGreenAlias.metricErrors(),
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
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
	alias: alias,
	deploymentConfig: codedeploy.lambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
	preHook: warmUpUserCache,
})

// or configure one on an existing deployment group
deploymentGroup.addPostHook(endToEndValidation)
```

### Import an existing Lambda Deployment Group

To import an already existing Deployment Group:

```go
var application lambdaApplication

deploymentGroup := codedeploy.lambdaDeploymentGroup.fromLambdaDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &lambdaDeploymentGroupAttributes{
	application: application,
	deploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
})
```

## Lambda Deployment Configurations

CodeDeploy for Lambda comes with predefined configurations for traffic shifting.
The predefined configurations are available as LambdaDeploymentConfig constants.

```go
var application lambdaApplication
var alias alias
config := codedeploy.lambdaDeploymentConfig_CANARY_10PERCENT_30MINUTES()
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
	application: application,
	alias: alias,
	deploymentConfig: config,
})
```

If you want to specify your own strategy,
you can do so with the LambdaDeploymentConfig construct,
letting you specify precisely how fast a new function version is deployed.

```go
// Example automatically generated from non-compiling source. May contain errors.
var application lambdaApplication
var alias alias
config := codedeploy.NewLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &lambdaDeploymentConfigProps{
	trafficRoutingConfig: codedeploy.NewTimeBasedCanaryTrafficRoutingConfig(map[string]interface{}{
		"interval": cdk.Duration_minutes(jsii.Number(15)),
		"percentage": jsii.Number(5),
	}),
})
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
	application: application,
	alias: alias,
	deploymentConfig: config,
})
```

You can specify a custom name for your deployment config, but if you do you will not be able to update the interval/percentage through CDK.

```go
// Example automatically generated from non-compiling source. May contain errors.
config := codedeploy.NewLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &lambdaDeploymentConfigProps{
	trafficRoutingConfig: codedeploy.NewTimeBasedCanaryTrafficRoutingConfig(map[string]interface{}{
		"interval": cdk.Duration_minutes(jsii.Number(15)),
		"percentage": jsii.Number(5),
	}),
	deploymentConfigName: jsii.String("MyDeploymentConfig"),
})
```

To import an already existing Deployment Config:

```go
deploymentConfig := codedeploy.lambdaDeploymentConfig.fromLambdaDeploymentConfigName(this, jsii.String("ExistingDeploymentConfiguration"), jsii.String("MyExistingDeploymentConfiguration"))
```

## ECS Applications

To create a new CodeDeploy Application that deploys an ECS service:

```go
application := codedeploy.NewEcsApplication(this, jsii.String("CodeDeployApplication"), &ecsApplicationProps{
	applicationName: jsii.String("MyApplication"),
})
```

To import an already existing Application:

```go
application := codedeploy.ecsApplication.fromEcsApplicationName(this, jsii.String("ExistingCodeDeployApplication"), jsii.String("MyExistingApplication"))
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
// Example automatically generated from non-compiling source. May contain errors.
var myApplication ecsApplication
var cluster ecs.Cluster
var taskDefinition ecs.FargateTaskDefinition
var blueTargetGroup elbv2.ITargetGroup
var greenTargetGroup elbv2.ITargetGroup
var listener elbv2.IApplicationListener


service := ecs.NewFargateService(this, jsii.String("Service"), map[string]interface{}{
	"cluster": cluster,
	"taskDefinition": taskDefinition,
	"deploymentController": map[string]interface{}{
		"type": ecs.DeploymentControllerType_CODE_DEPLOY,
	},
})

codedeploy.NewEcsDeploymentGroup(stack, jsii.String("BlueGreenDG"), &ecsDeploymentGroupProps{
	service: service,
	blueGreenDeploymentConfig: &ecsBlueGreenDeploymentConfig{
		blueTargetGroup: blueTargetGroup,
		greenTargetGroup: greenTargetGroup,
		listener: listener,
	},
	deploymentConfig: codedeploy.ecsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
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
// Example automatically generated from non-compiling source. May contain errors.
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"


// Alarm on the number of unhealthy ECS tasks in each target group
blueUnhealthyHosts := cloudwatch.NewAlarm(stack, jsii.String("BlueUnhealthyHosts"), &alarmProps{
	alarmName: jsii.String(stack.stackName + "-Unhealthy-Hosts-Blue"),
	metric: blueTargetGroup.metricUnhealthyHostCount(),
	threshold: jsii.Number(1),
	evaluationPeriods: jsii.Number(2),
})

greenUnhealthyHosts := cloudwatch.NewAlarm(stack, jsii.String("GreenUnhealthyHosts"), &alarmProps{
	alarmName: jsii.String(stack.stackName + "-Unhealthy-Hosts-Green"),
	metric: greenTargetGroup.metricUnhealthyHostCount(),
	threshold: jsii.Number(1),
	evaluationPeriods: jsii.Number(2),
})

// Alarm on the number of HTTP 5xx responses returned by each target group
blueApiFailure := cloudwatch.NewAlarm(stack, jsii.String("Blue5xx"), &alarmProps{
	alarmName: jsii.String(stack.stackName + "-Http-5xx-Blue"),
	metric: blueTargetGroup.metricHttpCodeTarget(elbv2.httpCodeTarget_TARGET_5XX_COUNT, map[string]interface{}{
		"period": cdk.Duration_minutes(jsii.Number(1)),
	}),
	threshold: jsii.Number(1),
	evaluationPeriods: jsii.Number(1),
})

greenApiFailure := cloudwatch.NewAlarm(stack, jsii.String("Green5xx"), &alarmProps{
	alarmName: jsii.String(stack.stackName + "-Http-5xx-Green"),
	metric: greenTargetGroup.metricHttpCodeTarget(elbv2.*httpCodeTarget_TARGET_5XX_COUNT, map[string]interface{}{
		"period": cdk.Duration_minutes(jsii.Number(1)),
	}),
	threshold: jsii.Number(1),
	evaluationPeriods: jsii.Number(1),
})

codedeploy.NewEcsDeploymentGroup(stack, jsii.String("BlueGreenDG"), &ecsDeploymentGroupProps{
	// CodeDeploy will monitor these alarms during a deployment and automatically roll back
	alarms: []iAlarm{
		blueUnhealthyHosts,
		greenUnhealthyHosts,
		blueApiFailure,
		greenApiFailure,
	},
	autoRollback: &autoRollbackConfig{
		// CodeDeploy will automatically roll back if a deployment is stopped
		stoppedDeployment: jsii.Boolean(true),
	},
	service: service,
	blueGreenDeploymentConfig: &ecsBlueGreenDeploymentConfig{
		blueTargetGroup: blueTargetGroup,
		greenTargetGroup: greenTargetGroup,
		listener: listener,
	},
	deploymentConfig: codedeploy.ecsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
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
// Example automatically generated from non-compiling source. May contain errors.
var myApplication ecsApplication
var service ecs.FargateService
var blueTargetGroup elbv2.ITargetGroup
var greenTargetGroup elbv2.ITargetGroup
var listener elbv2.IApplicationListener
var testListener elbv2.IApplicationListener


codedeploy.NewEcsDeploymentGroup(stack, jsii.String("BlueGreenDG"), &ecsDeploymentGroupProps{
	service: service,
	blueGreenDeploymentConfig: &ecsBlueGreenDeploymentConfig{
		blueTargetGroup: blueTargetGroup,
		greenTargetGroup: greenTargetGroup,
		listener: listener,
		testListener: testListener,
	},
	deploymentConfig: codedeploy.ecsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
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
// Example automatically generated from non-compiling source. May contain errors.
codedeploy.NewEcsDeploymentGroup(stack, jsii.String("BlueGreenDG"), &ecsDeploymentGroupProps{
	// The deployment will wait for approval for up to 8 hours before stopping the deployment
	deploymentApprovalWaitTime: awscdk.Duration.hours(jsii.Number(8)),
	autoRollback: &autoRollbackConfig{
		// CodeDeploy will automatically roll back if the 8-hour approval period times out and the deployment stops
		stoppedDeployment: jsii.Boolean(true),
	},
	service: service,
	blueGreenDeploymentConfig: &ecsBlueGreenDeploymentConfig{
		blueTargetGroup: blueTargetGroup,
		greenTargetGroup: greenTargetGroup,
		listener: listener,
		testListener: testListener,
	},
	deploymentConfig: codedeploy.ecsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
})
```

### Deployment bake time

You can specify how long CodeDeploy waits before it terminates the original 'blue' ECS task set when a blue-green deployment
is complete in order to let the deployment "bake" a while. During this bake time, CodeDeploy will continue to monitor any
CloudWatch alarms specified for the deployment group and will automatically roll back if those alarms go into a failed state.

```go
// Example automatically generated from non-compiling source. May contain errors.
codedeploy.NewEcsDeploymentGroup(stack, jsii.String("BlueGreenDG"), &ecsDeploymentGroupProps{
	service: service,
	blueGreenDeploymentConfig: &ecsBlueGreenDeploymentConfig{
		blueTargetGroup: blueTargetGroup,
		greenTargetGroup: greenTargetGroup,
		listener: listener,
		// CodeDeploy will wait for 30 minutes after completing the blue-green deployment before it terminates the blue tasks
		terminationWaitTime: awscdk.Duration.minutes(jsii.Number(30)),
	},
	// CodeDeploy will continue to monitor these alarms during the 30-minute bake time and will automatically
	// roll back if they go into a failed state at any point during the deployment.
	alarms: []iAlarm{
		blueUnhealthyHosts,
		greenUnhealthyHosts,
		blueApiFailure,
		greenApiFailure,
	},
	deploymentConfig: codedeploy.ecsDeploymentConfig_CANARY_10PERCENT_5MINUTES(),
})
```

### Import an existing ECS Deployment Group

To import an already existing Deployment Group:

```go
var application ecsApplication

deploymentGroup := codedeploy.ecsDeploymentGroup.fromEcsDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &ecsDeploymentGroupAttributes{
	application: application,
	deploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
})
```

## ECS Deployment Configurations

CodeDeploy for ECS comes with predefined configurations for traffic shifting.
The predefined configurations are available as LambdaDeploymentConfig constants.

```go
config := codedeploy.ecsDeploymentConfig_CANARY_10PERCENT_5MINUTES()
```

If you want to specify your own strategy,
you can do so with the EcsDeploymentConfig construct,
letting you specify precisely how fast an ECS service is deployed.

```go
// Example automatically generated from non-compiling source. May contain errors.
codedeploy.NewEcsDeploymentConfig(this, jsii.String("CustomConfig"), &ecsDeploymentConfigProps{
	trafficRoutingConfig: codedeploy.NewTimeBasedCanaryTrafficRoutingConfig(map[string]interface{}{
		"interval": cdk.Duration_minutes(jsii.Number(15)),
		"percentage": jsii.Number(5),
	}),
})
```

You can specify a custom name for your deployment config, but if you do you will not be able to update the interval/percentage through CDK.

```go
// Example automatically generated from non-compiling source. May contain errors.
config := codedeploy.NewEcsDeploymentConfig(this, jsii.String("CustomConfig"), &ecsDeploymentConfigProps{
	trafficRoutingConfig: codedeploy.NewTimeBasedCanaryTrafficRoutingConfig(map[string]interface{}{
		"interval": cdk.Duration_minutes(jsii.Number(15)),
		"percentage": jsii.Number(5),
	}),
	deploymentConfigName: jsii.String("MyDeploymentConfig"),
})
```

Or import an existing one:

```go
deploymentConfig := codedeploy.ecsDeploymentConfig.fromEcsDeploymentConfigName(this, jsii.String("ExistingDeploymentConfiguration"), jsii.String("MyExistingDeploymentConfiguration"))
```
