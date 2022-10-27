# AWS CodeDeploy Construct Library

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

### Rollbacks and Alarms

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

### Import an existing Deployment Group

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
