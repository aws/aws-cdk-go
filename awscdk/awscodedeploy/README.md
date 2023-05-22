# AWS CodeDeploy Construct Library

AWS CodeDeploy is a deployment service that automates application deployments to
Amazon EC2 instances, on-premises instances, serverless Lambda functions, or
Amazon ECS services.

The CDK currently supports Amazon EC2, on-premise and AWS Lambda applications.

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
import elbv2 "github.com/aws/aws-cdk-go/awscdk"

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

## Deployment Configurations

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

### Create a custom Deployment Config

CodeDeploy for Lambda comes with built-in configurations for traffic shifting.
If you want to specify your own strategy,
you can do so with the CustomLambdaDeploymentConfig construct,
letting you specify precisely how fast a new function version is deployed.

```go
var application lambdaApplication
var alias alias
config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &CustomLambdaDeploymentConfigProps{
	Type: codedeploy.CustomLambdaDeploymentConfigType_CANARY,
	Interval: awscdk.Duration_Minutes(jsii.Number(1)),
	Percentage: jsii.Number(5),
})
deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
	Application: Application,
	Alias: Alias,
	DeploymentConfig: config,
})
```

You can specify a custom name for your deployment config, but if you do you will not be able to update the interval/percentage through CDK.

```go
config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &CustomLambdaDeploymentConfigProps{
	Type: codedeploy.CustomLambdaDeploymentConfigType_CANARY,
	Interval: awscdk.Duration_Minutes(jsii.Number(1)),
	Percentage: jsii.Number(5),
	DeploymentConfigName: jsii.String("MyDeploymentConfig"),
})
```

### Rollbacks and Alarms

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

### Import an existing Deployment Group

To import an already existing Deployment Group:

```go
var application lambdaApplication

deploymentGroup := codedeploy.LambdaDeploymentGroup_FromLambdaDeploymentGroupAttributes(this, jsii.String("ExistingCodeDeployDeploymentGroup"), &LambdaDeploymentGroupAttributes{
	Application: Application,
	DeploymentGroupName: jsii.String("MyExistingDeploymentGroup"),
})
```
