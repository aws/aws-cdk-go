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

loadBalancedEcsService := ecsPatterns.NewApplicationLoadBalancedEc2Service(this, jsii.String("Service"), &ApplicationLoadBalancedEc2ServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
		Environment: map[string]*string{
			"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
			"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
		},
		Command: []*string{
			jsii.String("command"),
		},
		EntryPoint: []*string{
			jsii.String("entry"),
			jsii.String("point"),
		},
	},
	DesiredCount: jsii.Number(2),
})
```

* `ApplicationLoadBalancedFargateService`

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		Command: []*string{
			jsii.String("command"),
		},
		EntryPoint: []*string{
			jsii.String("entry"),
			jsii.String("point"),
		},
	},
})

loadBalancedFargateService.TargetGroup.ConfigureHealthCheck(&HealthCheck{
	Path: jsii.String("/custom-health-path"),
})
```

Instead of providing a cluster you can specify a VPC and CDK will create a new ECS cluster.
If you deploy multiple services CDK will only create one cluster per VPC.

You can omit `cluster` and `vpc` to let CDK create a new VPC with two AZs and create a cluster inside this VPC.

You can customize the health check for your target group; otherwise it defaults to `HTTP` over port `80` hitting path `/`.

You can customize the health check configuration of the container via the [`healthCheck`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ecs.HealthCheck.html) property; otherwise it defaults to the health check configuration from the container.

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

loadBalancedEc2Service := ecsPatterns.NewApplicationMultipleTargetGroupsEc2Service(this, jsii.String("Service"), &ApplicationMultipleTargetGroupsEc2ServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(256),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageProps{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	TargetGroups: []applicationTargetProps{
		&applicationTargetProps{
			ContainerPort: jsii.Number(80),
		},
		&applicationTargetProps{
			ContainerPort: jsii.Number(90),
			PathPattern: jsii.String("a/b/c"),
			Priority: jsii.Number(10),
		},
	},
})
```

* `ApplicationMultipleTargetGroupsFargateService`

```go
// One application load balancer with one listener and two target groups.
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("Service"), &ApplicationMultipleTargetGroupsFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageProps{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	TargetGroups: []applicationTargetProps{
		&applicationTargetProps{
			ContainerPort: jsii.Number(80),
		},
		&applicationTargetProps{
			ContainerPort: jsii.Number(90),
			PathPattern: jsii.String("a/b/c"),
			Priority: jsii.Number(10),
		},
	},
})
```

## Network Load Balanced Services

To define an Amazon ECS service that is behind a network load balancer, instantiate one of the following:

* `NetworkLoadBalancedEc2Service`

```go
var cluster cluster

loadBalancedEcsService := ecsPatterns.NewNetworkLoadBalancedEc2Service(this, jsii.String("Service"), &NetworkLoadBalancedEc2ServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
		Environment: map[string]*string{
			"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
			"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
		},
	},
	DesiredCount: jsii.Number(2),
})
```

* `NetworkLoadBalancedFargateService`

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewNetworkLoadBalancedFargateService(this, jsii.String("Service"), &NetworkLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	Cpu: jsii.Number(512),
	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
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

loadBalancedEc2Service := ecsPatterns.NewNetworkMultipleTargetGroupsEc2Service(this, jsii.String("Service"), &NetworkMultipleTargetGroupsEc2ServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(256),
	TaskImageOptions: &NetworkLoadBalancedTaskImageProps{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	LoadBalancers: []networkLoadBalancerProps{
		&networkLoadBalancerProps{
			Name: jsii.String("lb1"),
			Listeners: []networkListenerProps{
				&networkListenerProps{
					Name: jsii.String("listener1"),
				},
			},
		},
		&networkLoadBalancerProps{
			Name: jsii.String("lb2"),
			Listeners: []*networkListenerProps{
				&networkListenerProps{
					Name: jsii.String("listener2"),
				},
			},
		},
	},
	TargetGroups: []networkTargetProps{
		&networkTargetProps{
			ContainerPort: jsii.Number(80),
			Listener: jsii.String("listener1"),
		},
		&networkTargetProps{
			ContainerPort: jsii.Number(90),
			Listener: jsii.String("listener2"),
		},
	},
})
```

* NetworkMultipleTargetGroupsFargateService

```go
// Two network load balancers, each with their own listener and target group.
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewNetworkMultipleTargetGroupsFargateService(this, jsii.String("Service"), &NetworkMultipleTargetGroupsFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(512),
	TaskImageOptions: &NetworkLoadBalancedTaskImageProps{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	LoadBalancers: []networkLoadBalancerProps{
		&networkLoadBalancerProps{
			Name: jsii.String("lb1"),
			Listeners: []networkListenerProps{
				&networkListenerProps{
					Name: jsii.String("listener1"),
				},
			},
		},
		&networkLoadBalancerProps{
			Name: jsii.String("lb2"),
			Listeners: []*networkListenerProps{
				&networkListenerProps{
					Name: jsii.String("listener2"),
				},
			},
		},
	},
	TargetGroups: []networkTargetProps{
		&networkTargetProps{
			ContainerPort: jsii.Number(80),
			Listener: jsii.String("listener1"),
		},
		&networkTargetProps{
			ContainerPort: jsii.Number(90),
			Listener: jsii.String("listener2"),
		},
	},
})
```

## Queue Processing Services

To define a service that creates a queue and reads from that queue, instantiate one of the following:

* `QueueProcessingEc2Service`

```go
var cluster cluster

queueProcessingEc2Service := ecsPatterns.NewQueueProcessingEc2Service(this, jsii.String("Service"), &QueueProcessingEc2ServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	Command: []*string{
		jsii.String("-c"),
		jsii.String("4"),
		jsii.String("amazon.com"),
	},
	EnableLogging: jsii.Boolean(false),
	DesiredTaskCount: jsii.Number(2),
	Environment: map[string]*string{
		"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
		"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
	},
	MaxScalingCapacity: jsii.Number(5),
	ContainerName: jsii.String("test"),
})
```

* `QueueProcessingFargateService`

```go
var cluster cluster

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	Command: []*string{
		jsii.String("-c"),
		jsii.String("4"),
		jsii.String("amazon.com"),
	},
	EnableLogging: jsii.Boolean(false),
	DesiredTaskCount: jsii.Number(2),
	Environment: map[string]*string{
		"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
		"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
	},
	MaxScalingCapacity: jsii.Number(5),
	ContainerName: jsii.String("test"),
})
```

when queue not provided by user, CDK will create a primary queue and a dead letter queue with default redrive policy and attach permission to the task to be able to access the primary queue.

NOTE: `QueueProcessingFargateService` adds a CPU Based scaling strategy by default. You can turn this off by setting `disableCpuBasedScaling: true`.

```go
var cluster cluster

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	Command: []*string{
		jsii.String("-c"),
		jsii.String("4"),
		jsii.String("amazon.com"),
	},
	EnableLogging: jsii.Boolean(false),
	DesiredTaskCount: jsii.Number(2),
	Environment: map[string]*string{
		"TEST_ENVIRONMENT_VARIABLE1": jsii.String("test environment variable 1 value"),
		"TEST_ENVIRONMENT_VARIABLE2": jsii.String("test environment variable 2 value"),
	},
	MaxScalingCapacity: jsii.Number(5),
	ContainerName: jsii.String("test"),
	DisableCpuBasedScaling: jsii.Boolean(true),
})
```

To specify a custom target CPU utilization percentage for the scaling strategy use the  `cpuTargetUtilizationPercent` property:

```go
var cluster cluster

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	Command: []*string{
		jsii.String("-c"),
		jsii.String("4"),
		jsii.String("amazon.com"),
	},
	EnableLogging: jsii.Boolean(false),
	DesiredTaskCount: jsii.Number(2),
	Environment: map[string]interface{}{
	},
	MaxScalingCapacity: jsii.Number(5),
	ContainerName: jsii.String("test"),
	CpuTargetUtilizationPercent: jsii.Number(90),
})
```

## Scheduled Tasks

To define a task that runs periodically, there are 2 options:

* `ScheduledEc2Task`

```go
// Instantiate an Amazon EC2 Task to run at a scheduled interval
var cluster cluster

ecsScheduledTask := ecsPatterns.NewScheduledEc2Task(this, jsii.String("ScheduledTask"), &ScheduledEc2TaskProps{
	Cluster: Cluster,
	ScheduledEc2TaskImageOptions: &ScheduledEc2TaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		MemoryLimitMiB: jsii.Number(256),
		Environment: map[string]*string{
			"name": jsii.String("TRIGGER"),
			"value": jsii.String("CloudWatch Events"),
		},
	},
	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
	Enabled: jsii.Boolean(true),
	RuleName: jsii.String("sample-scheduled-task-rule"),
})
```

* `ScheduledFargateTask`

```go
var cluster cluster

scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
	Cluster: Cluster,
	ScheduledFargateTaskImageOptions: &ScheduledFargateTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		MemoryLimitMiB: jsii.Number(512),
	},
	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
	PlatformVersion: ecs.FargatePlatformVersion_LATEST,
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


domainZone := awscdk.HostedZone_FromLookup(this, jsii.String("Zone"), &HostedZoneProviderProps{
	DomainName: jsii.String("example.com"),
})
certificate := awscdk.Certificate_FromCertificateArn(this, jsii.String("Cert"), jsii.String("arn:aws:acm:us-east-1:123456:certificate/abcdefg"))
loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Vpc: Vpc,
	Cluster: Cluster,
	Certificate: Certificate,
	SslPolicy: awscdk.SslPolicy_RECOMMENDED,
	DomainName: jsii.String("api.example.com"),
	DomainZone: DomainZone,
	RedirectHTTP: jsii.Boolean(true),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})
```

### Set capacityProviderStrategies for ApplicationLoadBalancedFargateService

```go
var cluster cluster

cluster.EnableFargateCapacityProviders()

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	CapacityProviderStrategies: []capacityProviderStrategy{
		&capacityProviderStrategy{
			CapacityProvider: jsii.String("FARGATE_SPOT"),
			Weight: jsii.Number(2),
			Base: jsii.Number(0),
		},
		&capacityProviderStrategy{
			CapacityProvider: jsii.String("FARGATE"),
			Weight: jsii.Number(1),
			Base: jsii.Number(1),
		},
	},
})
```

### Add Schedule-Based Auto-Scaling to an ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	DesiredCount: jsii.Number(1),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})

scalableTarget := loadBalancedFargateService.Service.AutoScaleTaskCount(&EnableScalingProps{
	MinCapacity: jsii.Number(5),
	MaxCapacity: jsii.Number(20),
})

scalableTarget.ScaleOnSchedule(jsii.String("DaytimeScaleDown"), &ScalingSchedule{
	Schedule: appscaling.Schedule_Cron(&CronOptions{
		Hour: jsii.String("8"),
		Minute: jsii.String("0"),
	}),
	MinCapacity: jsii.Number(1),
})

scalableTarget.ScaleOnSchedule(jsii.String("EveningRushScaleUp"), &ScalingSchedule{
	Schedule: appscaling.Schedule_*Cron(&CronOptions{
		Hour: jsii.String("20"),
		Minute: jsii.String("0"),
	}),
	MinCapacity: jsii.Number(10),
})
```

### Add Metric-Based Auto-Scaling to an ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	DesiredCount: jsii.Number(1),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})

scalableTarget := loadBalancedFargateService.Service.AutoScaleTaskCount(&EnableScalingProps{
	MinCapacity: jsii.Number(1),
	MaxCapacity: jsii.Number(20),
})

scalableTarget.ScaleOnCpuUtilization(jsii.String("CpuScaling"), &CpuUtilizationScalingProps{
	TargetUtilizationPercent: jsii.Number(50),
})

scalableTarget.ScaleOnMemoryUtilization(jsii.String("MemoryScaling"), &MemoryUtilizationScalingProps{
	TargetUtilizationPercent: jsii.Number(50),
})
```

### Change the default Deployment Controller

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	DesiredCount: jsii.Number(1),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	DeploymentController: &DeploymentController{
		Type: ecs.DeploymentControllerType_CODE_DEPLOY,
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

service := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	DesiredCount: jsii.Number(1),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	CircuitBreaker: &DeploymentCircuitBreaker{
		Rollback: jsii.Boolean(true),
	},
})
```

### Set deployment configuration on QueueProcessingService

```go
var cluster cluster

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	Command: []*string{
		jsii.String("-c"),
		jsii.String("4"),
		jsii.String("amazon.com"),
	},
	EnableLogging: jsii.Boolean(false),
	DesiredTaskCount: jsii.Number(2),
	Environment: map[string]interface{}{
	},
	MaxScalingCapacity: jsii.Number(5),
	MaxHealthyPercent: jsii.Number(200),
	MinHealthyPercent: jsii.Number(66),
})
```

### Set taskSubnets and securityGroups for QueueProcessingFargateService

```go
var vpc vpc
var securityGroup securityGroup

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Vpc: Vpc,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	SecurityGroups: []iSecurityGroup{
		securityGroup,
	},
	TaskSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
	},
})
```

### Define tasks with public IPs for QueueProcessingFargateService

```go
var vpc vpc

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Vpc: Vpc,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	AssignPublicIp: jsii.Boolean(true),
})
```

### Define tasks with custom queue parameters for QueueProcessingFargateService

```go
var vpc vpc

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Vpc: Vpc,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	MaxReceiveCount: jsii.Number(42),
	RetentionPeriod: awscdk.Duration_Days(jsii.Number(7)),
	VisibilityTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
})
```

### Set cooldown for QueueProcessingFargateService

The cooldown period is the amount of time to wait for a previous scaling activity to take effect.
To specify something other than the default cooldown period of 300 seconds, use the `cooldown` parameter:

```go
var vpc vpc

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Vpc: Vpc,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	AssignPublicIp: jsii.Boolean(true),
	Cooldown: awscdk.Duration_Seconds(jsii.Number(500)),
})
```

### Set capacityProviderStrategies for QueueProcessingFargateService

```go
var cluster cluster

cluster.EnableFargateCapacityProviders()

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	CapacityProviderStrategies: []capacityProviderStrategy{
		&capacityProviderStrategy{
			CapacityProvider: jsii.String("FARGATE_SPOT"),
			Weight: jsii.Number(2),
		},
		&capacityProviderStrategy{
			CapacityProvider: jsii.String("FARGATE"),
			Weight: jsii.Number(1),
		},
	},
})
```

### Set a custom container-level Healthcheck for QueueProcessingFargateService

```go
var vpc vpc
var securityGroup securityGroup

queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &QueueProcessingFargateServiceProps{
	Vpc: Vpc,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	HealthCheck: &HealthCheck{
		Command: []*string{
			jsii.String("CMD-SHELL"),
			jsii.String("curl -f http://localhost/ || exit 1"),
		},
		// the properties below are optional
		Interval: awscdk.Duration_Minutes(jsii.Number(30)),
		Retries: jsii.Number(123),
		StartPeriod: awscdk.Duration_*Minutes(jsii.Number(30)),
		Timeout: awscdk.Duration_*Minutes(jsii.Number(30)),
	},
})
```

### Set capacityProviderStrategies for QueueProcessingEc2Service

```go
import autoscaling "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	MaxAzs: jsii.Number(1),
})
cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
	Vpc: Vpc,
})
autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("asg"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(),
})
capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("provider"), &AsgCapacityProviderProps{
	AutoScalingGroup: AutoScalingGroup,
})
cluster.AddAsgCapacityProvider(capacityProvider)

queueProcessingEc2Service := ecsPatterns.NewQueueProcessingEc2Service(this, jsii.String("Service"), &QueueProcessingEc2ServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(512),
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	CapacityProviderStrategies: []capacityProviderStrategy{
		&capacityProviderStrategy{
			CapacityProvider: capacityProvider.CapacityProviderName,
		},
	},
})
```

### Select specific vpc subnets for ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	DesiredCount: jsii.Number(1),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	TaskSubnets: &SubnetSelection{
		Subnets: []iSubnet{
			ec2.Subnet_FromSubnetId(this, jsii.String("subnet"), jsii.String("VpcISOLATEDSubnet1Subnet80F07FA0")),
		},
	},
})
```

### Select idleTimeout for ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	DesiredCount: jsii.Number(1),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	IdleTimeout: awscdk.Duration_Seconds(jsii.Number(120)),
})
```

### Select idleTimeout for ApplicationMultipleTargetGroupsFargateService

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	MaxAzs: jsii.Number(1),
})
loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &ApplicationMultipleTargetGroupsFargateServiceProps{
	Cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
		Vpc: *Vpc,
	}),
	MemoryLimitMiB: jsii.Number(256),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageProps{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	EnableExecuteCommand: jsii.Boolean(true),
	LoadBalancers: []applicationLoadBalancerProps{
		&applicationLoadBalancerProps{
			Name: jsii.String("lb"),
			IdleTimeout: awscdk.Duration_Seconds(jsii.Number(400)),
			DomainName: jsii.String("api.example.com"),
			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
				ZoneName: jsii.String("example.com"),
			}),
			Listeners: []applicationListenerProps{
				&applicationListenerProps{
					Name: jsii.String("listener"),
					Protocol: awscdk.ApplicationProtocol_HTTPS,
					Certificate: awscdk.Certificate_FromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
				},
			},
		},
		&applicationLoadBalancerProps{
			Name: jsii.String("lb2"),
			IdleTimeout: awscdk.Duration_*Seconds(jsii.Number(120)),
			DomainName: jsii.String("frontend.com"),
			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
				ZoneName: jsii.String("frontend.com"),
			}),
			Listeners: []*applicationListenerProps{
				&applicationListenerProps{
					Name: jsii.String("listener2"),
					Protocol: awscdk.ApplicationProtocol_HTTPS,
					Certificate: awscdk.Certificate_*FromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
				},
			},
		},
	},
	TargetGroups: []applicationTargetProps{
		&applicationTargetProps{
			ContainerPort: jsii.Number(80),
			Listener: jsii.String("listener"),
		},
		&applicationTargetProps{
			ContainerPort: jsii.Number(90),
			PathPattern: jsii.String("a/b/c"),
			Priority: jsii.Number(10),
			Listener: jsii.String("listener"),
		},
		&applicationTargetProps{
			ContainerPort: jsii.Number(443),
			Listener: jsii.String("listener2"),
		},
		&applicationTargetProps{
			ContainerPort: jsii.Number(80),
			PathPattern: jsii.String("a/b/c"),
			Priority: jsii.Number(10),
			Listener: jsii.String("listener2"),
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

vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	MaxAzs: jsii.Number(1),
})

loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &ApplicationMultipleTargetGroupsFargateServiceProps{
	Cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
		Vpc: *Vpc,
	}),
	MemoryLimitMiB: jsii.Number(256),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageProps{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	EnableExecuteCommand: jsii.Boolean(true),
	LoadBalancers: []applicationLoadBalancerProps{
		&applicationLoadBalancerProps{
			Name: jsii.String("lb"),
			IdleTimeout: awscdk.Duration_Seconds(jsii.Number(400)),
			DomainName: jsii.String("api.example.com"),
			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
				ZoneName: jsii.String("example.com"),
			}),
			Listeners: []applicationListenerProps{
				&applicationListenerProps{
					Name: jsii.String("listener"),
					Protocol: awscdk.ApplicationProtocol_HTTPS,
					Certificate: awscdk.Certificate_FromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
				},
			},
		},
		&applicationLoadBalancerProps{
			Name: jsii.String("lb2"),
			IdleTimeout: awscdk.Duration_*Seconds(jsii.Number(120)),
			DomainName: jsii.String("frontend.com"),
			DomainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
				ZoneName: jsii.String("frontend.com"),
			}),
			Listeners: []*applicationListenerProps{
				&applicationListenerProps{
					Name: jsii.String("listener2"),
					Protocol: awscdk.ApplicationProtocol_HTTPS,
					Certificate: awscdk.Certificate_*FromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
					SslPolicy: awscdk.SslPolicy_TLS12_EXT,
				},
			},
		},
	},
	TargetGroups: []applicationTargetProps{
		&applicationTargetProps{
			ContainerPort: jsii.Number(80),
			Listener: jsii.String("listener"),
		},
		&applicationTargetProps{
			ContainerPort: jsii.Number(90),
			PathPattern: jsii.String("a/b/c"),
			Priority: jsii.Number(10),
			Listener: jsii.String("listener"),
		},
		&applicationTargetProps{
			ContainerPort: jsii.Number(443),
			Listener: jsii.String("listener2"),
		},
		&applicationTargetProps{
			ContainerPort: jsii.Number(80),
			PathPattern: jsii.String("a/b/c"),
			Priority: jsii.Number(10),
			Listener: jsii.String("listener2"),
		},
	},
})

loadBalancedFargateService.TargetGroups[0].ConfigureHealthCheck(&HealthCheck{
	Port: jsii.String("8050"),
	Protocol: awscdk.Protocol_HTTP,
	HealthyThresholdCount: jsii.Number(2),
	UnhealthyThresholdCount: jsii.Number(2),
	Timeout: awscdk.Duration_*Seconds(jsii.Number(10)),
	Interval: awscdk.Duration_*Seconds(jsii.Number(30)),
	HealthyHttpCodes: jsii.String("200"),
})

loadBalancedFargateService.TargetGroups[1].ConfigureHealthCheck(&HealthCheck{
	Port: jsii.String("8050"),
	Protocol: awscdk.Protocol_HTTP,
	HealthyThresholdCount: jsii.Number(2),
	UnhealthyThresholdCount: jsii.Number(2),
	Timeout: awscdk.Duration_*Seconds(jsii.Number(10)),
	Interval: awscdk.Duration_*Seconds(jsii.Number(30)),
	HealthyHttpCodes: jsii.String("200"),
})
```

### Set runtimePlatform for ApplicationLoadBalancedFargateService

```go
var cluster cluster

applicationLoadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	RuntimePlatform: &RuntimePlatform{
		CpuArchitecture: ecs.CpuArchitecture_ARM64(),
		OperatingSystemFamily: ecs.OperatingSystemFamily_LINUX(),
	},
})
```

### Customize Container Name for ScheduledFargateTask

```go
var cluster cluster

scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
	Cluster: Cluster,
	ScheduledFargateTaskImageOptions: &ScheduledFargateTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		ContainerName: jsii.String("customContainerName"),
		MemoryLimitMiB: jsii.Number(512),
	},
	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
	PlatformVersion: ecs.FargatePlatformVersion_LATEST,
})
```

### Customize Container Name for ScheduledEc2Task

```go
var cluster cluster

ecsScheduledTask := ecsPatterns.NewScheduledEc2Task(this, jsii.String("ScheduledTask"), &ScheduledEc2TaskProps{
	Cluster: Cluster,
	ScheduledEc2TaskImageOptions: &ScheduledEc2TaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		ContainerName: jsii.String("customContainerName"),
		MemoryLimitMiB: jsii.Number(256),
		Environment: map[string]*string{
			"name": jsii.String("TRIGGER"),
			"value": jsii.String("CloudWatch Events"),
		},
	},
	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
	Enabled: jsii.Boolean(true),
	RuleName: jsii.String("sample-scheduled-task-rule"),
})
```

### Set PlatformVersion for ScheduledFargateTask

```go
var cluster cluster

scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
	Cluster: Cluster,
	ScheduledFargateTaskImageOptions: &ScheduledFargateTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		MemoryLimitMiB: jsii.Number(512),
	},
	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
	PlatformVersion: ecs.FargatePlatformVersion_VERSION1_4,
})
```

### Set SecurityGroups for ScheduledFargateTask

```go
vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	MaxAzs: jsii.Number(1),
})
cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
	Vpc: Vpc,
})
securityGroup := ec2.NewSecurityGroup(this, jsii.String("SG"), &SecurityGroupProps{
	Vpc: Vpc,
})

scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
	Cluster: Cluster,
	ScheduledFargateTaskImageOptions: &ScheduledFargateTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		MemoryLimitMiB: jsii.Number(512),
	},
	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
	SecurityGroups: []iSecurityGroup{
		securityGroup,
	},
})
```

### Deploy application and metrics sidecar

The following is an example of deploying an application along with a metrics sidecar container that utilizes `dockerLabels` for discovery:

```go
var cluster cluster
var vpc vpc

service := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	Vpc: Vpc,
	DesiredCount: jsii.Number(1),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		DockerLabels: map[string]*string{
			"application.label.one": jsii.String("first_label"),
			"application.label.two": jsii.String("second_label"),
		},
	},
})

service.TaskDefinition.AddContainer(jsii.String("Sidecar"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_*FromRegistry(jsii.String("example/metrics-sidecar")),
})
```

### Select specific load balancer name ApplicationLoadBalancedFargateService

```go
var cluster cluster

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	DesiredCount: jsii.Number(1),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	TaskSubnets: &SubnetSelection{
		Subnets: []iSubnet{
			ec2.Subnet_FromSubnetId(this, jsii.String("subnet"), jsii.String("VpcISOLATEDSubnet1Subnet80F07FA0")),
		},
	},
	LoadBalancerName: jsii.String("application-lb-name"),
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

loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	DesiredCount: jsii.Number(1),
	Cpu: jsii.Number(512),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	EnableExecuteCommand: jsii.Boolean(true),
})
```

Please note, ECS Exec leverages AWS Systems Manager (SSM). So as a prerequisite for the exec command
to work, you need to have the SSM plugin for the AWS CLI installed locally. For more information, see
[Install Session Manager plugin for AWS CLI](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html).

### Propagate Tags from task definition for ScheduledFargateTask

For tasks that are defined by a Task Definition, tags applied to the definition will not be applied
to the running task by default. To get this behavior, set `propagateTags` to `ecs.PropagatedTagSource.TASK_DEFINITION` as
shown below:

```go
import "github.com/aws/aws-cdk-go/awscdk"


vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	MaxAzs: jsii.Number(1),
})
cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
	Vpc: Vpc,
})
taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
	MemoryLimitMiB: jsii.Number(512),
	Cpu: jsii.Number(256),
})
taskDefinition.AddContainer(jsii.String("WebContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
})
awscdk.Tags_Of(taskDefinition).Add(jsii.String("my-tag"), jsii.String("my-tag-value"))
scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
	Cluster: Cluster,
	TaskDefinition: taskDefinition,
	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
	PropagateTags: ecs.PropagatedTagSource_TASK_DEFINITION,
})
```

### Pass a list of tags for ScheduledFargateTask

You can pass a list of tags to be applied to a Fargate task directly. These tags are in addition to any tags
that could be applied to the task definition and propagated using the `propagateTags` attribute.

```go
vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	MaxAzs: jsii.Number(1),
})
cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
	Vpc: Vpc,
})
scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
	Cluster: Cluster,
	ScheduledFargateTaskImageOptions: &ScheduledFargateTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
		MemoryLimitMiB: jsii.Number(512),
	},
	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
	Tags: []tag{
		&tag{
			Key: jsii.String("my-tag"),
			Value: jsii.String("my-tag-value"),
		},
	},
})
```

### Use custom ephemeral storage for ECS Fargate tasks

You can pass a custom ephemeral storage (21GiB - 200GiB) to ECS Fargate tasks on Fargate Platform Version 1.4.0 or later.

```go
vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	MaxAzs: jsii.Number(2),
	RestrictDefaultSecurityGroup: jsii.Boolean(false),
})
cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &ClusterProps{
	Vpc: Vpc,
})

applicationLoadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("ALBFargateServiceWithCustomEphemeralStorage"), &ApplicationLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	Cpu: jsii.Number(512),
	EphemeralStorageGiB: jsii.Number(21),
	TaskImageOptions: &ApplicationLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})

networkLoadBalancedFargateService := ecsPatterns.NewNetworkLoadBalancedFargateService(this, jsii.String("NLBFargateServiceWithCustomEphemeralStorage"), &NetworkLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	MemoryLimitMiB: jsii.Number(1024),
	Cpu: jsii.Number(512),
	EphemeralStorageGiB: jsii.Number(200),
	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_*FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
})
```

### Set securityGroups for NetworkLoadBalancedFargateService

```go
var vpc vpc
var securityGroup securityGroup

queueProcessingFargateService := ecsPatterns.NewNetworkLoadBalancedFargateService(this, jsii.String("Service"), &NetworkLoadBalancedFargateServiceProps{
	Vpc: Vpc,
	MemoryLimitMiB: jsii.Number(512),
	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	SecurityGroups: []iSecurityGroup{
		securityGroup,
	},
})
```

### Use dualstack NLB

```go
import "github.com/aws/aws-cdk-go/awscdk"


// The VPC and subnet must have associated IPv6 CIDR blocks.
vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
	IpProtocol: ec2.IpProtocol_DUAL_STACK,
})
cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
	Vpc: Vpc,
})

networkLoadbalancedFargateService := ecsPatterns.NewNetworkLoadBalancedFargateService(this, jsii.String("NlbFargateService"), &NetworkLoadBalancedFargateServiceProps{
	Cluster: Cluster,
	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
})

networkLoadbalancedEc2Service := ecsPatterns.NewNetworkLoadBalancedEc2Service(this, jsii.String("NlbEc2Service"), &NetworkLoadBalancedEc2ServiceProps{
	Cluster: Cluster,
	TaskImageOptions: &NetworkLoadBalancedTaskImageOptions{
		Image: ecs.ContainerImage_*FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	},
	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
})
```
