# CDK Construct library for higher-level ECS Constructs

This library provides higher-level Amazon ECS constructs which follow common architectural patterns. It contains:

* Application Load Balanced Services
* Network Load Balanced Services
* Queue Processing Services
* Scheduled Tasks (cron jobs)
* Additional Examples

## Application Load Balanced Services

To define an Amazon ECS service that is behind an application load balancer, instantiate one of the following:

* `ApplicationLoadBalancedEc2Service`

```go
var cluster cluster

loadBalancedEcsService := ecsPatterns.NewApplicationLoadBalancedEc2Service(this, jsii.String("Service"), &applicationLoadBalancedEc2ServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("test")),
		environment: map[string]*string{
			"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
			"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
		},
		command: []*string{
			jsii.String("command"),
		},
		entryPoint: []*string{
			jsii.String("entry"),
			jsii.String("point"),
		},
	},
	desiredCount: jsii.Number(2),
})
```

* `ApplicationLoadBalancedFargateService`

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		command: []*string{
			jsii.String("command"),
		},
		entryPoint: []*string{
			jsii.String("entry"),
			jsii.String("point"),
		},
	},
})

loadBalancedFargateService.targetGroup.configureHealthCheck(&healthCheck{
	path: jsii.String("/custom-health-path"),
})
```

Instead of providing a cluster you can specify a VPC and CDK will create a new ECS cluster.
If you deploy multiple services CDK will only create one cluster per VPC.

You can omit `cluster` and `vpc` to let CDK create a new VPC with two AZs and create a cluster inside this VPC.

You can customize the health check for your target group; otherwise it defaults to `HTTP` over port `80` hitting path `/`.

Fargate services will use the `LATEST` platform version by default, but you can override by providing a value for the `platformVersion` property in the constructor.

Fargate services use the default VPC Security Group unless one or more are provided using the `securityGroups` property in the constructor.

By setting `redirectHTTP` to true, CDK will automatically create a listener on port 80 that redirects HTTP traffic to the HTTPS port.

If you specify the option `recordType` you can decide if you want the construct to use CNAME or Route53-Aliases as record sets.

If you need to encrypt the traffic between the load balancer and the ECS tasks, you can set the `targetProtocol` to `HTTPS`.

Additionally, if more than one application target group are needed, instantiate one of the following:

* `ApplicationMultipleTargetGroupsEc2Service`

```go
// One application load balancer with one listener and two target groups.
var cluster cluster

loadBalancedEc2Service := ecsPatterns.NewApplicationMultipleTargetGroupsEc2Service(this, jsii.String("Service"), &applicationMultipleTargetGroupsEc2ServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(256),
	taskImageOptions: &applicationLoadBalancedTaskImageProps{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	targetGroups: []applicationTargetProps{
		&applicationTargetProps{
			containerPort: jsii.Number(80),
		},
		&applicationTargetProps{
			containerPort: jsii.Number(90),
			pathPattern: jsii.String("a/b/c"),
			priority: jsii.Number(10),
		},
	},
})
```

* `ApplicationMultipleTargetGroupsFargateService`

```go
// One application load balancer with one listener and two target groups.
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("Service"), &applicationMultipleTargetGroupsFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageProps{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	targetGroups: []applicationTargetProps{
		&applicationTargetProps{
			containerPort: jsii.Number(80),
		},
		&applicationTargetProps{
			containerPort: jsii.Number(90),
			pathPattern: jsii.String("a/b/c"),
			priority: jsii.Number(10),
		},
	},
})
```

## Network Load Balanced Services

To define an Amazon ECS service that is behind a network load balancer, instantiate one of the following:

* `NetworkLoadBalancedEc2Service`

```go
var cluster cluster

loadBalancedEcsService := ecsPatterns.NewNetworkLoadBalancedEc2Service(this, jsii.String("Service"), &networkLoadBalancedEc2ServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	taskImageOptions: &networkLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("test")),
		environment: map[string]*string{
			"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
			"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
		},
	},
	desiredCount: jsii.Number(2),
})
```

* `NetworkLoadBalancedFargateService`

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewNetworkLoadBalancedFargateService(this, jsii.String("Service"), &networkLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	cpu: jsii.Number(512),
	taskImageOptions: &networkLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})
```

The CDK will create a new Amazon ECS cluster if you specify a VPC and omit `cluster`. If you deploy multiple services the CDK will only create one cluster per VPC.

If `cluster` and `vpc` are omitted, the CDK creates a new VPC with subnets in two Availability Zones and a cluster within this VPC.

If you specify the option `recordType` you can decide if you want the construct to use CNAME or Route53-Aliases as record sets.

Additionally, if more than one network target group is needed, instantiate one of the following:

* NetworkMultipleTargetGroupsEc2Service

```go
// Two network load balancers, each with their own listener and target group.
var cluster cluster

loadBalancedEc2Service := ecsPatterns.NewNetworkMultipleTargetGroupsEc2Service(this, jsii.String("Service"), &networkMultipleTargetGroupsEc2ServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(256),
	taskImageOptions: &networkLoadBalancedTaskImageProps{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	loadBalancers: []networkLoadBalancerProps{
		&networkLoadBalancerProps{
			name: jsii.String("lb1"),
			listeners: []networkListenerProps{
				&networkListenerProps{
					name: jsii.String("listener1"),
				},
			},
		},
		&networkLoadBalancerProps{
			name: jsii.String("lb2"),
			listeners: []*networkListenerProps{
				&networkListenerProps{
					name: jsii.String("listener2"),
				},
			},
		},
	},
	targetGroups: []networkTargetProps{
		&networkTargetProps{
			containerPort: jsii.Number(80),
			listener: jsii.String("listener1"),
		},
		&networkTargetProps{
			containerPort: jsii.Number(90),
			listener: jsii.String("listener2"),
		},
	},
})
```

* NetworkMultipleTargetGroupsFargateService

```go
// Two network load balancers, each with their own listener and target group.
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewNetworkMultipleTargetGroupsFargateService(this, jsii.String("Service"), &networkMultipleTargetGroupsFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(512),
	taskImageOptions: &networkLoadBalancedTaskImageProps{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	loadBalancers: []networkLoadBalancerProps{
		&networkLoadBalancerProps{
			name: jsii.String("lb1"),
			listeners: []networkListenerProps{
				&networkListenerProps{
					name: jsii.String("listener1"),
				},
			},
		},
		&networkLoadBalancerProps{
			name: jsii.String("lb2"),
			listeners: []*networkListenerProps{
				&networkListenerProps{
					name: jsii.String("listener2"),
				},
			},
		},
	},
	targetGroups: []networkTargetProps{
		&networkTargetProps{
			containerPort: jsii.Number(80),
			listener: jsii.String("listener1"),
		},
		&networkTargetProps{
			containerPort: jsii.Number(90),
			listener: jsii.String("listener2"),
		},
	},
})
```

## Queue Processing Services

To define a service that creates a queue and reads from that queue, instantiate one of the following:

* `QueueProcessingEc2Service`

```go
var cluster cluster

queueProcessingEc2Service := ecsPatterns.NewQueueProcessingEc2Service(this, jsii.String("Service"), &queueProcessingEc2ServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	command: []*string{
		jsii.String("-c"),
		jsii.String("4"),
		jsii.String("amazon.com"),
	},
	enableLogging: jsii.Boolean(false),
	desiredTaskCount: jsii.Number(2),
	environment: map[string]*string{
		"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
		"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
	},
	maxScalingCapacity: jsii.Number(5),
	containerName: jsii.String("test"),
})
```

* `QueueProcessingFargateService`

```go
var cluster cluster

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	command: []*string{
		jsii.String("-c"),
		jsii.String("4"),
		jsii.String("amazon.com"),
	},
	enableLogging: jsii.Boolean(false),
	desiredTaskCount: jsii.Number(2),
	environment: map[string]*string{
		"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
		"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
	},
	maxScalingCapacity: jsii.Number(5),
	containerName: jsii.String("test"),
})
```

when queue not provided by user, CDK will create a primary queue and a dead letter queue with default redrive policy and attach permission to the task to be able to access the primary queue.

## Scheduled Tasks

To define a task that runs periodically, there are 2 options:

* `ScheduledEc2Task`

```go
// Instantiate an Amazon EC2 Task to run at a scheduled interval
var cluster cluster

ecsScheduledTask := ecsPatterns.NewScheduledEc2Task(this, jsii.String("ScheduledTask"), &scheduledEc2TaskProps{
	cluster: cluster,
	scheduledEc2TaskImageOptions: &scheduledEc2TaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		memoryLimitMiB: jsii.Number(256),
		environment: map[string]*string{
			"name": jsii.String("TRIGGER"),
			"value": jsii.String("CloudWatch Events"),
		},
	},
	schedule: appscaling.schedule.expression(jsii.String("rate(1 minute)")),
	enabled: jsii.Boolean(true),
	ruleName: jsii.String("sample-scheduled-task-rule"),
})
```

* `ScheduledFargateTask`

```go
var cluster cluster

scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &scheduledFargateTaskProps{
	cluster: cluster,
	scheduledFargateTaskImageOptions: &scheduledFargateTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		memoryLimitMiB: jsii.Number(512),
	},
	schedule: appscaling.schedule.expression(jsii.String("rate(1 minute)")),
	platformVersion: ecs.fargatePlatformVersion_LATEST,
})
```

## Additional Examples

In addition to using the constructs, users can also add logic to customize these constructs:

### Configure HTTPS on an ApplicationLoadBalancedFargateService

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc
var cluster cluster


domainZone := awscdk.HostedZone.fromLookup(this, jsii.String("Zone"), &hostedZoneProviderProps{
	domainName: jsii.String("example.com"),
})
certificate := awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert"), jsii.String("arn:aws:acm:us-east-1:123456:certificate/abcdefg"))
loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	vpc: vpc,
	cluster: cluster,
	certificate: certificate,
	sslPolicy: awscdk.SslPolicy_RECOMMENDED,
	domainName: jsii.String("api.example.com"),
	domainZone: domainZone,
	redirectHTTP: jsii.Boolean(true),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})
```

### Set capacityProviderStrategies for ApplicationLoadBalancedFargateService

```go
var cluster cluster

cluster.enableFargateCapacityProviders()

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	capacityProviderStrategies: []capacityProviderStrategy{
		&capacityProviderStrategy{
			capacityProvider: jsii.String("FARGATE_SPOT"),
			weight: jsii.Number(2),
			base: jsii.Number(0),
		},
		&capacityProviderStrategy{
			capacityProvider: jsii.String("FARGATE"),
			weight: jsii.Number(1),
			base: jsii.Number(1),
		},
	},
})
```

### Add Schedule-Based Auto-Scaling to an ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	desiredCount: jsii.Number(1),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})

scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
	minCapacity: jsii.Number(5),
	maxCapacity: jsii.Number(20),
})

scalableTarget.scaleOnSchedule(jsii.String("DaytimeScaleDown"), &scalingSchedule{
	schedule: appscaling.schedule.cron(&cronOptions{
		hour: jsii.String("8"),
		minute: jsii.String("0"),
	}),
	minCapacity: jsii.Number(1),
})

scalableTarget.scaleOnSchedule(jsii.String("EveningRushScaleUp"), &scalingSchedule{
	schedule: appscaling.*schedule.cron(&cronOptions{
		hour: jsii.String("20"),
		minute: jsii.String("0"),
	}),
	minCapacity: jsii.Number(10),
})
```

### Add Metric-Based Auto-Scaling to an ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	desiredCount: jsii.Number(1),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})

scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
	minCapacity: jsii.Number(1),
	maxCapacity: jsii.Number(20),
})

scalableTarget.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
	targetUtilizationPercent: jsii.Number(50),
})

scalableTarget.scaleOnMemoryUtilization(jsii.String("MemoryScaling"), &memoryUtilizationScalingProps{
	targetUtilizationPercent: jsii.Number(50),
})
```

### Change the default Deployment Controller

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	desiredCount: jsii.Number(1),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	deploymentController: &deploymentController{
		type: ecs.deploymentControllerType_CODE_DEPLOY,
	},
})
```

### Deployment circuit breaker and rollback

Amazon ECS [deployment circuit breaker](https://aws.amazon.com/tw/blogs/containers/announcing-amazon-ecs-deployment-circuit-breaker/)
automatically rolls back unhealthy service deployments without the need for manual intervention. Use `circuitBreaker` to enable
deployment circuit breaker and optionally enable `rollback` for automatic rollback. See [Using the deployment circuit breaker](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html)
for more details.

```go
var cluster cluster

service := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	desiredCount: jsii.Number(1),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	circuitBreaker: &deploymentCircuitBreaker{
		rollback: jsii.Boolean(true),
	},
})
```

### Set deployment configuration on QueueProcessingService

```go
var cluster cluster

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	command: []*string{
		jsii.String("-c"),
		jsii.String("4"),
		jsii.String("amazon.com"),
	},
	enableLogging: jsii.Boolean(false),
	desiredTaskCount: jsii.Number(2),
	environment: map[string]interface{}{
	},
	maxScalingCapacity: jsii.Number(5),
	maxHealthyPercent: jsii.Number(200),
	minHealthyPercent: jsii.Number(66),
})
```

### Set taskSubnets and securityGroups for QueueProcessingFargateService

```go
var vpc vpc
var securityGroup securityGroup

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
	vpc: vpc,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	securityGroups: []iSecurityGroup{
		securityGroup,
	},
	taskSubnets: &subnetSelection{
		subnetType: ec2.subnetType_PRIVATE_ISOLATED,
	},
})
```

### Define tasks with public IPs for QueueProcessingFargateService

```go
var vpc vpc

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
	vpc: vpc,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	assignPublicIp: jsii.Boolean(true),
})
```

### Define tasks with custom queue parameters for QueueProcessingFargateService

```go
var vpc vpc

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
	vpc: vpc,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	maxReceiveCount: jsii.Number(42),
	retentionPeriod: awscdk.Duration.days(jsii.Number(7)),
	visibilityTimeout: awscdk.Duration.minutes(jsii.Number(5)),
})
```

### Set capacityProviderStrategies for QueueProcessingFargateService

```go
var cluster cluster

cluster.enableFargateCapacityProviders()

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	capacityProviderStrategies: []capacityProviderStrategy{
		&capacityProviderStrategy{
			capacityProvider: jsii.String("FARGATE_SPOT"),
			weight: jsii.Number(2),
		},
		&capacityProviderStrategy{
			capacityProvider: jsii.String("FARGATE"),
			weight: jsii.Number(1),
		},
	},
})
```

### Set a custom container-level Healthcheck for QueueProcessingFargateService

```go
var vpc vpc
var securityGroup securityGroup

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
	vpc: vpc,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	healthCheck: &healthCheck{
		command: []*string{
			jsii.String("CMD-SHELL"),
			jsii.String("curl -f http://localhost/ || exit 1"),
		},
		// the properties below are optional
		interval: awscdk.Duration.minutes(jsii.Number(30)),
		retries: jsii.Number(123),
		startPeriod: awscdk.Duration.minutes(jsii.Number(30)),
		timeout: awscdk.Duration.minutes(jsii.Number(30)),
	},
})
```

### Set capacityProviderStrategies for QueueProcessingEc2Service

```go
import autoscaling "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
	maxAzs: jsii.Number(1),
})
cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &clusterProps{
	vpc: vpc,
})
autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("asg"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_MICRO),
	machineImage: ecs.ecsOptimizedImage.amazonLinux2(),
})
capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("provider"), &asgCapacityProviderProps{
	autoScalingGroup: autoScalingGroup,
})
cluster.addAsgCapacityProvider(capacityProvider)

queueProcessingEc2Service := ecsPatterns.NewQueueProcessingEc2Service(this, jsii.String("Service"), &queueProcessingEc2ServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	capacityProviderStrategies: []capacityProviderStrategy{
		&capacityProviderStrategy{
			capacityProvider: capacityProvider.capacityProviderName,
		},
	},
})
```

### Select specific vpc subnets for ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	desiredCount: jsii.Number(1),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	taskSubnets: &subnetSelection{
		subnets: []iSubnet{
			ec2.subnet.fromSubnetId(this, jsii.String("subnet"), jsii.String("VpcISOLATEDSubnet1Subnet80F07FA0")),
		},
	},
})
```

### Select idleTimeout for ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	desiredCount: jsii.Number(1),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	idleTimeout: awscdk.Duration.seconds(jsii.Number(120)),
})
```

### Select idleTimeout for ApplicationMultipleTargetGroupsFargateService

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
	maxAzs: jsii.Number(1),
})
loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &applicationMultipleTargetGroupsFargateServiceProps{
	cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &clusterProps{
		vpc: vpc,
	}),
	memoryLimitMiB: jsii.Number(256),
	taskImageOptions: &applicationLoadBalancedTaskImageProps{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	enableExecuteCommand: jsii.Boolean(true),
	loadBalancers: []applicationLoadBalancerProps{
		&applicationLoadBalancerProps{
			name: jsii.String("lb"),
			idleTimeout: awscdk.Duration.seconds(jsii.Number(400)),
			domainName: jsii.String("api.example.com"),
			domainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
				zoneName: jsii.String("example.com"),
			}),
			listeners: []applicationListenerProps{
				&applicationListenerProps{
					name: jsii.String("listener"),
					protocol: awscdk.ApplicationProtocol_HTTPS,
					certificate: awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
					sslPolicy: awscdk.SslPolicy_TLS12_EXT,
				},
			},
		},
		&applicationLoadBalancerProps{
			name: jsii.String("lb2"),
			idleTimeout: awscdk.Duration.seconds(jsii.Number(120)),
			domainName: jsii.String("frontend.com"),
			domainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
				zoneName: jsii.String("frontend.com"),
			}),
			listeners: []*applicationListenerProps{
				&applicationListenerProps{
					name: jsii.String("listener2"),
					protocol: awscdk.ApplicationProtocol_HTTPS,
					certificate: awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
					sslPolicy: awscdk.SslPolicy_TLS12_EXT,
				},
			},
		},
	},
	targetGroups: []applicationTargetProps{
		&applicationTargetProps{
			containerPort: jsii.Number(80),
			listener: jsii.String("listener"),
		},
		&applicationTargetProps{
			containerPort: jsii.Number(90),
			pathPattern: jsii.String("a/b/c"),
			priority: jsii.Number(10),
			listener: jsii.String("listener"),
		},
		&applicationTargetProps{
			containerPort: jsii.Number(443),
			listener: jsii.String("listener2"),
		},
		&applicationTargetProps{
			containerPort: jsii.Number(80),
			pathPattern: jsii.String("a/b/c"),
			priority: jsii.Number(10),
			listener: jsii.String("listener2"),
		},
	},
})
```

### Set health checks for ApplicationMultipleTargetGroupsFargateService

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
	maxAzs: jsii.Number(1),
})

loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &applicationMultipleTargetGroupsFargateServiceProps{
	cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &clusterProps{
		vpc: vpc,
	}),
	memoryLimitMiB: jsii.Number(256),
	taskImageOptions: &applicationLoadBalancedTaskImageProps{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	enableExecuteCommand: jsii.Boolean(true),
	loadBalancers: []applicationLoadBalancerProps{
		&applicationLoadBalancerProps{
			name: jsii.String("lb"),
			idleTimeout: awscdk.Duration.seconds(jsii.Number(400)),
			domainName: jsii.String("api.example.com"),
			domainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
				zoneName: jsii.String("example.com"),
			}),
			listeners: []applicationListenerProps{
				&applicationListenerProps{
					name: jsii.String("listener"),
					protocol: awscdk.ApplicationProtocol_HTTPS,
					certificate: awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
					sslPolicy: awscdk.SslPolicy_TLS12_EXT,
				},
			},
		},
		&applicationLoadBalancerProps{
			name: jsii.String("lb2"),
			idleTimeout: awscdk.Duration.seconds(jsii.Number(120)),
			domainName: jsii.String("frontend.com"),
			domainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
				zoneName: jsii.String("frontend.com"),
			}),
			listeners: []*applicationListenerProps{
				&applicationListenerProps{
					name: jsii.String("listener2"),
					protocol: awscdk.ApplicationProtocol_HTTPS,
					certificate: awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
					sslPolicy: awscdk.SslPolicy_TLS12_EXT,
				},
			},
		},
	},
	targetGroups: []applicationTargetProps{
		&applicationTargetProps{
			containerPort: jsii.Number(80),
			listener: jsii.String("listener"),
		},
		&applicationTargetProps{
			containerPort: jsii.Number(90),
			pathPattern: jsii.String("a/b/c"),
			priority: jsii.Number(10),
			listener: jsii.String("listener"),
		},
		&applicationTargetProps{
			containerPort: jsii.Number(443),
			listener: jsii.String("listener2"),
		},
		&applicationTargetProps{
			containerPort: jsii.Number(80),
			pathPattern: jsii.String("a/b/c"),
			priority: jsii.Number(10),
			listener: jsii.String("listener2"),
		},
	},
})

loadBalancedFargateService.targetGroups[0].configureHealthCheck(&healthCheck{
	port: jsii.String("8050"),
	protocol: awscdk.Protocol_HTTP,
	healthyThresholdCount: jsii.Number(2),
	unhealthyThresholdCount: jsii.Number(2),
	timeout: awscdk.Duration.seconds(jsii.Number(10)),
	interval: awscdk.Duration.seconds(jsii.Number(30)),
	healthyHttpCodes: jsii.String("200"),
})

loadBalancedFargateService.targetGroups[1].configureHealthCheck(&healthCheck{
	port: jsii.String("8050"),
	protocol: awscdk.Protocol_HTTP,
	healthyThresholdCount: jsii.Number(2),
	unhealthyThresholdCount: jsii.Number(2),
	timeout: awscdk.Duration.seconds(jsii.Number(10)),
	interval: awscdk.Duration.seconds(jsii.Number(30)),
	healthyHttpCodes: jsii.String("200"),
})
```

### Set runtimePlatform for ApplicationLoadBalancedFargateService

```go
var cluster cluster

applicationLoadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	runtimePlatform: &runtimePlatform{
		cpuArchitecture: ecs.cpuArchitecture_ARM64(),
		operatingSystemFamily: ecs.operatingSystemFamily_LINUX(),
	},
})
```

### Set PlatformVersion for ScheduledFargateTask

```go
var cluster cluster

scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &scheduledFargateTaskProps{
	cluster: cluster,
	scheduledFargateTaskImageOptions: &scheduledFargateTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		memoryLimitMiB: jsii.Number(512),
	},
	schedule: appscaling.schedule.expression(jsii.String("rate(1 minute)")),
	platformVersion: ecs.fargatePlatformVersion_VERSION1_4,
})
```

### Set SecurityGroups for ScheduledFargateTask

```go
vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
	maxAzs: jsii.Number(1),
})
cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &clusterProps{
	vpc: vpc,
})
securityGroup := ec2.NewSecurityGroup(this, jsii.String("SG"), &securityGroupProps{
	vpc: vpc,
})

scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &scheduledFargateTaskProps{
	cluster: cluster,
	scheduledFargateTaskImageOptions: &scheduledFargateTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		memoryLimitMiB: jsii.Number(512),
	},
	schedule: appscaling.schedule.expression(jsii.String("rate(1 minute)")),
	securityGroups: []iSecurityGroup{
		securityGroup,
	},
})
```

### Use the REMOVE_DEFAULT_DESIRED_COUNT feature flag

The REMOVE_DEFAULT_DESIRED_COUNT feature flag is used to override the default desiredCount that is autogenerated by the CDK. This will set the desiredCount of any service created by any of the following constructs to be undefined.

* ApplicationLoadBalancedEc2Service
* ApplicationLoadBalancedFargateService
* NetworkLoadBalancedEc2Service
* NetworkLoadBalancedFargateService
* QueueProcessingEc2Service
* QueueProcessingFargateService

If a desiredCount is not passed in as input to the above constructs, CloudFormation will either create a new service to start up with a desiredCount of 1, or update an existing service to start up with the same desiredCount as prior to the update.

To enable the feature flag, ensure that the REMOVE_DEFAULT_DESIRED_COUNT flag within an application stack context is set to true, like so:

```go
var stack stack

stack.node.setContext(cxapi.eCS_REMOVE_DEFAULT_DESIRED_COUNT, jsii.Boolean(true))
```

The following is an example of an application with the REMOVE_DEFAULT_DESIRED_COUNT feature flag enabled:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import ecs "github.com/aws/aws-cdk-go/awscdk"
import ecsPatterns "github.com/aws/aws-cdk-go/awscdk"
import cxapi "github.com/aws/aws-cdk-go/awscdk"
import path "github.com/aws-samples/dummy/path"

app := awscdk.NewApp()

stack := awscdk.NewStack(app, jsii.String("aws-ecs-patterns-queue"))
stack.node.setContext(cxapi.eCS_REMOVE_DEFAULT_DESIRED_COUNT, jsii.Boolean(true))

vpc := ec2.NewVpc(stack, jsii.String("VPC"), &vpcProps{
	maxAzs: jsii.Number(2),
})

ecsPatterns.NewQueueProcessingFargateService(stack, jsii.String("QueueProcessingService"), &queueProcessingFargateServiceProps{
	vpc: vpc,
	memoryLimitMiB: jsii.Number(512),
	image: ecs.NewAssetImage(path.join(__dirname, jsii.String(".."), jsii.String("sqs-reader"))),
})
```

### Deploy application and metrics sidecar

The following is an example of deploying an application along with a metrics sidecar container that utilizes `dockerLabels` for discovery:

```go
var cluster cluster
var vpc vpc

service := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	vpc: vpc,
	desiredCount: jsii.Number(1),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		dockerLabels: map[string]*string{
			"application.label.one": jsii.String("first_label"),
			"application.label.two": jsii.String("second_label"),
		},
	},
})

service.taskDefinition.addContainer(jsii.String("Sidecar"), &containerDefinitionOptions{
	image: ecs.*containerImage.fromRegistry(jsii.String("example/metrics-sidecar")),
})
```

### Select specific load balancer name ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	desiredCount: jsii.Number(1),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	taskSubnets: &subnetSelection{
		subnets: []iSubnet{
			ec2.subnet.fromSubnetId(this, jsii.String("subnet"), jsii.String("VpcISOLATEDSubnet1Subnet80F07FA0")),
		},
	},
	loadBalancerName: jsii.String("application-lb-name"),
})
```

### ECS Exec

You can use ECS Exec to run commands in or get a shell to a container running on an Amazon EC2 instance or on
AWS Fargate. Enable ECS Exec, by setting `enableExecuteCommand` to `true`.

ECS Exec is supported by all Services i.e. `ApplicationLoadBalanced(Fargate|Ec2)Service`, `ApplicationMultipleTargetGroups(Fargate|Ec2)Service`, `NetworkLoadBalanced(Fargate|Ec2)Service`, `NetworkMultipleTargetGroups(Fargate|Ec2)Service`, `QueueProcessing(Fargate|Ec2)Service`. It is not supported for `ScheduledTask`s.

Read more about ECS Exec in the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-exec.html).

Example:

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
	cluster: cluster,
	memoryLimitMiB: jsii.Number(1024),
	desiredCount: jsii.Number(1),
	cpu: jsii.Number(512),
	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	enableExecuteCommand: jsii.Boolean(true),
})
```

Please note, ECS Exec leverages AWS Systems Manager (SSM). So as a prerequisite for the exec command
to work, you need to have the SSM plugin for the AWS CLI installed locally. For more information, see
[Install Session Manager plugin for AWS CLI](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html).
