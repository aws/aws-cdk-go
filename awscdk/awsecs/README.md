# Amazon ECS Construct Library

This package contains constructs for working with **Amazon Elastic Container
Service** (Amazon ECS).

Amazon Elastic Container Service (Amazon ECS) is a fully managed container orchestration service.

For further information on Amazon ECS,
see the [Amazon ECS documentation](https://docs.aws.amazon.com/ecs)

The following example creates an Amazon ECS cluster, adds capacity to it, and
runs a service on it:

```go
var vpc vpc


// Create an ECS cluster
cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
	vpc: vpc,
})

// Add capacity to it
cluster.addCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &addCapacityOptions{
	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
	desiredCapacity: jsii.Number(3),
})

taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))

taskDefinition.addContainer(jsii.String("DefaultContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	memoryLimitMiB: jsii.Number(512),
})

// Instantiate an Amazon ECS Service
ecsService := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
})
```

For a set of constructs defining common ECS architectural patterns, see the `@aws-cdk/aws-ecs-patterns` package.

## Launch Types: AWS Fargate vs Amazon EC2 vs AWS ECS Anywhere

There are three sets of constructs in this library:

* Use the `Ec2TaskDefinition` and `Ec2Service` constructs to run tasks on Amazon EC2 instances running in your account.
* Use the `FargateTaskDefinition` and `FargateService` constructs to run tasks on
  instances that are managed for you by AWS.
* Use the `ExternalTaskDefinition` and `ExternalService` constructs to run AWS ECS Anywhere tasks on self-managed infrastructure.

Here are the main differences:

* **Amazon EC2**: instances are under your control. Complete control of task to host
  allocation. Required to specify at least a memory reservation or limit for
  every container. Can use Host, Bridge and AwsVpc networking modes. Can attach
  Classic Load Balancer. Can share volumes between container and host.
* **AWS Fargate**: tasks run on AWS-managed instances, AWS manages task to host
  allocation for you. Requires specification of memory and cpu sizes at the
  taskdefinition level. Only supports AwsVpc networking modes and
  Application/Network Load Balancers. Only the AWS log driver is supported.
  Many host features are not supported such as adding kernel capabilities
  and mounting host devices/volumes inside the container.
* **AWS ECS Anywhere**: tasks are run and managed by AWS ECS Anywhere on infrastructure
  owned by the customer. Bridge, Host and None networking modes are supported. Does not
  support autoscaling, load balancing, cloudmap or attachment of volumes.

For more information on Amazon EC2 vs AWS Fargate, networking and ECS Anywhere see the AWS Documentation:
[AWS Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS_Fargate.html),
[Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html),
[ECS Anywhere](https://aws.amazon.com/ecs/anywhere/)

## Clusters

A `Cluster` defines the infrastructure to run your
tasks on. You can run many tasks on a single cluster.

The following code creates a cluster that can run AWS Fargate tasks:

```go
var vpc vpc


cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
	vpc: vpc,
})
```

The following code imports an existing cluster using the ARN which can be used to
import an Amazon ECS service either EC2 or Fargate.

```go
clusterArn := "arn:aws:ecs:us-east-1:012345678910:cluster/clusterName"

cluster := ecs.cluster.fromClusterArn(this, jsii.String("Cluster"), clusterArn)
```

To use tasks with Amazon EC2 launch-type, you have to add capacity to
the cluster in order for tasks to be scheduled on your instances.  Typically,
you add an AutoScalingGroup with instances running the latest
Amazon ECS-optimized AMI to the cluster. There is a method to build and add such an
AutoScalingGroup automatically, or you can supply a customized AutoScalingGroup
that you construct yourself. It's possible to add multiple AutoScalingGroups
with various instance types.

The following example creates an Amazon ECS cluster and adds capacity to it:

```go
var vpc vpc


cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
	vpc: vpc,
})

// Either add default capacity
cluster.addCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &addCapacityOptions{
	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
	desiredCapacity: jsii.Number(3),
})

// Or add customized capacity. Be sure to start the Amazon ECS-optimized AMI.
autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
	machineImage: ecs.ecsOptimizedImage.amazonLinux(),
	// Or use Amazon ECS-Optimized Amazon Linux 2 AMI
	// machineImage: EcsOptimizedImage.amazonLinux2(),
	desiredCapacity: jsii.Number(3),
})

capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
	autoScalingGroup: autoScalingGroup,
})
cluster.addAsgCapacityProvider(capacityProvider)
```

If you omit the property `vpc`, the construct will create a new VPC with two AZs.

By default, all machine images will auto-update to the latest version
on each deployment, causing a replacement of the instances in your AutoScalingGroup
if the AMI has been updated since the last deployment.

If task draining is enabled, ECS will transparently reschedule tasks on to the new
instances before terminating your old instances. If you have disabled task draining,
the tasks will be terminated along with the instance. To prevent that, you
can pick a non-updating AMI by passing `cacheInContext: true`, but be sure
to periodically update to the latest AMI manually by using the [CDK CLI
context management commands](https://docs.aws.amazon.com/cdk/latest/guide/context.html):

```go
var vpc vpc

autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	machineImage: ecs.ecsOptimizedImage.amazonLinux(&ecsOptimizedImageOptions{
		cachedInContext: jsii.Boolean(true),
	}),
	vpc: vpc,
	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
})
```

To use `LaunchTemplate` with `AsgCapacityProvider`, make sure to specify the `userData` in the `LaunchTemplate`:

```go
var vpc vpc

launchTemplate := ec2.NewLaunchTemplate(this, jsii.String("ASG-LaunchTemplate"), &launchTemplateProps{
	instanceType: ec2.NewInstanceType(jsii.String("t3.medium")),
	machineImage: ecs.ecsOptimizedImage.amazonLinux2(),
	userData: ec2.userData.forLinux(),
})

autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	mixedInstancesPolicy: &mixedInstancesPolicy{
		instancesDistribution: &instancesDistribution{
			onDemandPercentageAboveBaseCapacity: jsii.Number(50),
		},
		launchTemplate: launchTemplate,
	},
})

cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
	vpc: vpc,
})

capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
	autoScalingGroup: autoScalingGroup,
	machineImageType: ecs.machineImageType_AMAZON_LINUX_2,
})

cluster.addAsgCapacityProvider(capacityProvider)
```

### Bottlerocket

[Bottlerocket](https://aws.amazon.com/bottlerocket/) is a Linux-based open source operating system that is
purpose-built by AWS for running containers. You can launch Amazon ECS container instances with the Bottlerocket AMI.

The following example will create a capacity with self-managed Amazon EC2 capacity of 2 `c5.large` Linux instances running with `Bottlerocket` AMI.

The following example adds Bottlerocket capacity to the cluster:

```go
var cluster cluster


cluster.addCapacity(jsii.String("bottlerocket-asg"), &addCapacityOptions{
	minCapacity: jsii.Number(2),
	instanceType: ec2.NewInstanceType(jsii.String("c5.large")),
	machineImage: ecs.NewBottleRocketImage(),
})
```

### ARM64 (Graviton) Instances

To launch instances with ARM64 hardware, you can use the Amazon ECS-optimized
Amazon Linux 2 (arm64) AMI. Based on Amazon Linux 2, this AMI is recommended
for use when launching your EC2 instances that are powered by Arm-based AWS
Graviton Processors.

```go
var cluster cluster


cluster.addCapacity(jsii.String("graviton-cluster"), &addCapacityOptions{
	minCapacity: jsii.Number(2),
	instanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
	machineImage: ecs.ecsOptimizedImage.amazonLinux2(ecs.amiHardwareType_ARM),
})
```

Bottlerocket is also supported:

```go
var cluster cluster


cluster.addCapacity(jsii.String("graviton-cluster"), &addCapacityOptions{
	minCapacity: jsii.Number(2),
	instanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
	machineImageType: ecs.machineImageType_BOTTLEROCKET,
})
```

### Spot Instances

To add spot instances into the cluster, you must specify the `spotPrice` in the `ecs.AddCapacityOptions` and optionally enable the `spotInstanceDraining` property.

```go
var cluster cluster


// Add an AutoScalingGroup with spot instances to the existing cluster
cluster.addCapacity(jsii.String("AsgSpot"), &addCapacityOptions{
	maxCapacity: jsii.Number(2),
	minCapacity: jsii.Number(2),
	desiredCapacity: jsii.Number(2),
	instanceType: ec2.NewInstanceType(jsii.String("c5.xlarge")),
	spotPrice: jsii.String("0.0735"),
	// Enable the Automated Spot Draining support for Amazon ECS
	spotInstanceDraining: jsii.Boolean(true),
})
```

### SNS Topic Encryption

When the `ecs.AddCapacityOptions` that you provide has a non-zero `taskDrainTime` (the default) then an SNS topic and Lambda are created to ensure that the
cluster's instances have been properly drained of tasks before terminating. The SNS Topic is sent the instance-terminating lifecycle event from the AutoScalingGroup,
and the Lambda acts on that event. If you wish to engage [server-side encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-data-encryption.html) for this SNS Topic
then you may do so by providing a KMS key for the `topicEncryptionKey` property of `ecs.AddCapacityOptions`.

```go
// Given
var cluster cluster
var key key

// Then, use that key to encrypt the lifecycle-event SNS Topic.
cluster.addCapacity(jsii.String("ASGEncryptedSNS"), &addCapacityOptions{
	instanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
	desiredCapacity: jsii.Number(3),
	topicEncryptionKey: key,
})
```

## Task definitions

A task definition describes what a single copy of a **task** should look like.
A task definition has one or more containers; typically, it has one
main container (the *default container* is the first one that's added
to the task definition, and it is marked *essential*) and optionally
some supporting containers which are used to support the main container,
doings things like upload logs or metrics to monitoring services.

To run a task or service with Amazon EC2 launch type, use the `Ec2TaskDefinition`. For AWS Fargate tasks/services, use the
`FargateTaskDefinition`. For AWS ECS Anywhere use the `ExternalTaskDefinition`. These classes
provide simplified APIs that only contain properties relevant for each specific launch type.

For a `FargateTaskDefinition`, specify the task size (`memoryLimitMiB` and `cpu`):

```go
fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
	memoryLimitMiB: jsii.Number(512),
	cpu: jsii.Number(256),
})
```

On Fargate Platform Version 1.4.0 or later, you may specify up to 200GiB of
[ephemeral storage](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-task-storage.html#fargate-task-storage-pv14):

```go
fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
	memoryLimitMiB: jsii.Number(512),
	cpu: jsii.Number(256),
	ephemeralStorageGiB: jsii.Number(100),
})
```

To add containers to a task definition, call `addContainer()`:

```go
fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
	memoryLimitMiB: jsii.Number(512),
	cpu: jsii.Number(256),
})
container := fargateTaskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
	// Use an image from DockerHub
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
})
```

For an `Ec2TaskDefinition`:

```go
ec2TaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"), &ec2TaskDefinitionProps{
	networkMode: ecs.networkMode_BRIDGE,
})

container := ec2TaskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
	// Use an image from DockerHub
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	memoryLimitMiB: jsii.Number(1024),
})
```

For an `ExternalTaskDefinition`:

```go
externalTaskDefinition := ecs.NewExternalTaskDefinition(this, jsii.String("TaskDef"))

container := externalTaskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
	// Use an image from DockerHub
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	memoryLimitMiB: jsii.Number(1024),
})
```

You can specify container properties when you add them to the task definition, or with various methods, e.g.:

To add a port mapping when adding a container to the task definition, specify the `portMappings` option:

```go
var taskDefinition taskDefinition


taskDefinition.addContainer(jsii.String("WebContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	memoryLimitMiB: jsii.Number(1024),
	portMappings: []portMapping{
		&portMapping{
			containerPort: jsii.Number(3000),
		},
	},
})
```

To add port mappings directly to a container definition, call `addPortMappings()`:

```go
var container containerDefinition


container.addPortMappings(&portMapping{
	containerPort: jsii.Number(3000),
})
```

To add data volumes to a task definition, call `addVolume()`:

```go
fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
	memoryLimitMiB: jsii.Number(512),
	cpu: jsii.Number(256),
})
volume := map[string]interface{}{
	// Use an Elastic FileSystem
	"name": jsii.String("mydatavolume"),
	"efsVolumeConfiguration": map[string]*string{
		"fileSystemId": jsii.String("EFS"),
	},
}

container := fargateTaskDefinition.addVolume(volume)
```

> Note: ECS Anywhere doesn't support volume attachments in the task definition.

To use a TaskDefinition that can be used with either Amazon EC2 or
AWS Fargate launch types, use the `TaskDefinition` construct.

When creating a task definition you have to specify what kind of
tasks you intend to run: Amazon EC2, AWS Fargate, or both.
The following example uses both:

```go
taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDef"), &taskDefinitionProps{
	memoryMiB: jsii.String("512"),
	cpu: jsii.String("256"),
	networkMode: ecs.networkMode_AWS_VPC,
	compatibility: ecs.compatibility_EC2_AND_FARGATE,
})
```

To grant a principal permission to run your `TaskDefinition`, you can use the `TaskDefinition.grantRun()` method:

```go
var role iGrantable

taskDef := ecs.NewTaskDefinition(this, jsii.String("TaskDef"), &taskDefinitionProps{
	cpu: jsii.String("512"),
	memoryMiB: jsii.String("512"),
	compatibility: ecs.compatibility_EC2_AND_FARGATE,
})

// Gives role required permissions to run taskDef
taskDef.grantRun(role)
```

### Images

Images supply the software that runs inside the container. Images can be
obtained from either DockerHub or from ECR repositories, built directly from a local Dockerfile, or use an existing tarball.

* `ecs.ContainerImage.fromRegistry(imageName)`: use a public image.
* `ecs.ContainerImage.fromRegistry(imageName, { credentials: mySecret })`: use a private image that requires credentials.
* `ecs.ContainerImage.fromEcrRepository(repo, tagOrDigest)`: use the given ECR repository as the image
  to start. If no tag or digest is provided, "latest" is assumed.
* `ecs.ContainerImage.fromAsset('./image')`: build and upload an
  image directly from a `Dockerfile` in your source directory.
* `ecs.ContainerImage.fromDockerImageAsset(asset)`: uses an existing
  `@aws-cdk/aws-ecr-assets.DockerImageAsset` as a container image.
* `ecs.ContainerImage.fromTarball(file)`: use an existing tarball.
* `new ecs.TagParameterContainerImage(repository)`: use the given ECR repository as the image
  but a CloudFormation parameter as the tag.

### Environment variables

To pass environment variables to the container, you can use the `environment`, `environmentFiles`, and `secrets` props.

```go
var secret secret
var dbSecret secret
var parameter stringParameter
var taskDefinition taskDefinition
var s3Bucket bucket


newContainer := taskDefinition.addContainer(jsii.String("container"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	memoryLimitMiB: jsii.Number(1024),
	environment: map[string]*string{
		 // clear text, not for sensitive data
		"STAGE": jsii.String("prod"),
	},
	environmentFiles: []environmentFile{
		ecs.*environmentFile.fromAsset(jsii.String("./demo-env-file.env")),
		ecs.*environmentFile.fromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
	},
	secrets: map[string]secret{
		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
		"SECRET": ecs.*secret.fromSecretsManager(secret),
		"DB_PASSWORD": ecs.*secret.fromSecretsManager(dbSecret, jsii.String("password")),
		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
		"API_KEY": ecs.*secret.fromSecretsManagerVersion(secret, &SecretVersionInfo{
			"versionId": jsii.String("12345"),
		}, jsii.String("apiKey")),
		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
		"PARAMETER": ecs.*secret.fromSsmParameter(parameter),
	},
})
newContainer.addEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
newContainer.addSecret(jsii.String("API_KEY"), ecs.secret.fromSecretsManager(secret))
newContainer.addSecret(jsii.String("DB_PASSWORD"), ecs.secret.fromSecretsManager(secret, jsii.String("password")))
```

The task execution role is automatically granted read permissions on the secrets/parameters. Support for environment
files is restricted to the EC2 launch type for files hosted on S3. Further details provided in the AWS documentation
about [specifying environment variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html).

### Linux parameters

To apply additional linux-specific options related to init process and memory management to the container, use the `linuxParameters` property:

```go
var taskDefinition taskDefinition


taskDefinition.addContainer(jsii.String("container"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	memoryLimitMiB: jsii.Number(1024),
	linuxParameters: ecs.NewLinuxParameters(this, jsii.String("LinuxParameters"), &linuxParametersProps{
		initProcessEnabled: jsii.Boolean(true),
		sharedMemorySize: jsii.Number(1024),
		maxSwap: awscdk.Size.mebibytes(jsii.Number(5000)),
		swappiness: jsii.Number(90),
	}),
})
```

### System controls

To set system controls (kernel parameters) on the container, use the `systemControls` prop:

```go
var taskDefinition taskDefinition


taskDefinition.addContainer(jsii.String("container"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	memoryLimitMiB: jsii.Number(1024),
	systemControls: []systemControl{
		&systemControl{
			namespace: jsii.String("net"),
			value: jsii.String("ipv4.tcp_tw_recycle"),
		},
	},
})
```

### Using Windows containers on Fargate

AWS Fargate supports Amazon ECS Windows containers. For more details, please see this [blog post](https://aws.amazon.com/tw/blogs/containers/running-windows-containers-with-amazon-ecs-on-aws-fargate/)

```go
// Create a Task Definition for the Windows container to start
taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
	runtimePlatform: &runtimePlatform{
		operatingSystemFamily: ecs.operatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
		cpuArchitecture: ecs.cpuArchitecture_X86_64(),
	},
	cpu: jsii.Number(1024),
	memoryLimitMiB: jsii.Number(2048),
})

taskDefinition.addContainer(jsii.String("windowsservercore"), &containerDefinitionOptions{
	logging: ecs.logDriver.awsLogs(&awsLogDriverProps{
		streamPrefix: jsii.String("win-iis-on-fargate"),
	}),
	portMappings: []portMapping{
		&portMapping{
			containerPort: jsii.Number(80),
		},
	},
	image: ecs.containerImage.fromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
})
```

### Using Graviton2 with Fargate

AWS Graviton2 supports AWS Fargate. For more details, please see this [blog post](https://aws.amazon.com/blogs/aws/announcing-aws-graviton2-support-for-aws-fargate-get-up-to-40-better-price-performance-for-your-serverless-containers/)

```go
// Create a Task Definition for running container on Graviton Runtime.
taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
	runtimePlatform: &runtimePlatform{
		operatingSystemFamily: ecs.operatingSystemFamily_LINUX(),
		cpuArchitecture: ecs.cpuArchitecture_ARM64(),
	},
	cpu: jsii.Number(1024),
	memoryLimitMiB: jsii.Number(2048),
})

taskDefinition.addContainer(jsii.String("webarm64"), &containerDefinitionOptions{
	logging: ecs.logDriver.awsLogs(&awsLogDriverProps{
		streamPrefix: jsii.String("graviton2-on-fargate"),
	}),
	portMappings: []portMapping{
		&portMapping{
			containerPort: jsii.Number(80),
		},
	},
	image: ecs.containerImage.fromRegistry(jsii.String("public.ecr.aws/nginx/nginx:latest-arm64v8")),
})
```

## Service

A `Service` instantiates a `TaskDefinition` on a `Cluster` a given number of
times, optionally associating them with a load balancer.
If a task fails,
Amazon ECS automatically restarts the task.

```go
var cluster cluster
var taskDefinition taskDefinition


service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	desiredCount: jsii.Number(5),
})
```

ECS Anywhere service definition looks like:

```go
var cluster cluster
var taskDefinition taskDefinition


service := ecs.NewExternalService(this, jsii.String("Service"), &externalServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	desiredCount: jsii.Number(5),
})
```

`Services` by default will create a security group if not provided.
If you'd like to specify which security groups to use you can override the `securityGroups` property.

### Deployment circuit breaker and rollback

Amazon ECS [deployment circuit breaker](https://aws.amazon.com/tw/blogs/containers/announcing-amazon-ecs-deployment-circuit-breaker/)
automatically rolls back unhealthy service deployments without the need for manual intervention. Use `circuitBreaker` to enable
deployment circuit breaker and optionally enable `rollback` for automatic rollback. See [Using the deployment circuit breaker](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html)
for more details.

```go
var cluster cluster
var taskDefinition taskDefinition

service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	circuitBreaker: &deploymentCircuitBreaker{
		rollback: jsii.Boolean(true),
	},
})
```

> Note: ECS Anywhere doesn't support deployment circuit breakers and rollback.

### Include an application/network load balancer

`Services` are load balancing targets and can be added to a target group, which will be attached to an application/network load balancers:

```go
var vpc vpc
var cluster cluster
var taskDefinition taskDefinition

service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
})

lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
	vpc: vpc,
	internetFacing: jsii.Boolean(true),
})
listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
	port: jsii.Number(80),
})
targetGroup1 := listener.addTargets(jsii.String("ECS1"), &addApplicationTargetsProps{
	port: jsii.Number(80),
	targets: []iApplicationLoadBalancerTarget{
		service,
	},
})
targetGroup2 := listener.addTargets(jsii.String("ECS2"), &addApplicationTargetsProps{
	port: jsii.Number(80),
	targets: []*iApplicationLoadBalancerTarget{
		service.loadBalancerTarget(&loadBalancerTargetOptions{
			containerName: jsii.String("MyContainer"),
			containerPort: jsii.Number(8080),
		}),
	},
})
```

> Note: ECS Anywhere doesn't support application/network load balancers.

Note that in the example above, the default `service` only allows you to register the first essential container or the first mapped port on the container as a target and add it to a new target group. To have more control over which container and port to register as targets, you can use `service.loadBalancerTarget()` to return a load balancing target for a specific container and port.

Alternatively, you can also create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.

```go
var cluster cluster
var taskDefinition taskDefinition
var vpc vpc

service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
})

lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
	vpc: vpc,
	internetFacing: jsii.Boolean(true),
})
listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
	port: jsii.Number(80),
})
service.registerLoadBalancerTargets(&ecsTarget{
	containerName: jsii.String("web"),
	containerPort: jsii.Number(80),
	newTargetGroupId: jsii.String("ECS"),
	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
		protocol: elbv2.applicationProtocol_HTTPS,
	}),
})
```

### Using a Load Balancer from a different Stack

If you want to put your Load Balancer and the Service it is load balancing to in
different stacks, you may not be able to use the convenience methods
`loadBalancer.addListener()` and `listener.addTargets()`.

The reason is that these methods will create resources in the same Stack as the
object they're called on, which may lead to cyclic references between stacks.
Instead, you will have to create an `ApplicationListener` in the service stack,
or an empty `TargetGroup` in the load balancer stack that you attach your
service to.

See the [ecs/cross-stack-load-balancer example](https://github.com/aws-samples/aws-cdk-examples/tree/master/typescript/ecs/cross-stack-load-balancer/)
for the alternatives.

### Include a classic load balancer

`Services` can also be directly attached to a classic load balancer as targets:

```go
var cluster cluster
var taskDefinition taskDefinition
var vpc vpc

service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
})

lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
	vpc: vpc,
})
lb.addListener(&loadBalancerListener{
	externalPort: jsii.Number(80),
})
lb.addTarget(service)
```

Similarly, if you want to have more control over load balancer targeting:

```go
var cluster cluster
var taskDefinition taskDefinition
var vpc vpc

service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
})

lb := elb.NewLoadBalancer(this, jsii.String("LB"), &loadBalancerProps{
	vpc: vpc,
})
lb.addListener(&loadBalancerListener{
	externalPort: jsii.Number(80),
})
lb.addTarget(service.loadBalancerTarget(&loadBalancerTargetOptions{
	containerName: jsii.String("MyContainer"),
	containerPort: jsii.Number(80),
}))
```

There are two higher-level constructs available which include a load balancer for you that can be found in the aws-ecs-patterns module:

* `LoadBalancedFargateService`
* `LoadBalancedEc2Service`

### Import existing services

`Ec2Service` and `FargateService` provide methods to import existing EC2/Fargate services.
The ARN of the existing service has to be specified to import the service.

Since AWS has changed the [ARN format for ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids),
feature flag `@aws-cdk/aws-ecs:arnFormatIncludesClusterName` must be enabled to use the new ARN format.
The feature flag changes behavior for the entire CDK project. Therefore it is not possible to mix the old and the new format in one CDK project.

```tss
declare const cluster: ecs.Cluster;

// Import service from EC2 service attributes
const service = ecs.Ec2Service.fromEc2ServiceAttributes(stack, 'EcsService', {
  serviceArn: 'arn:aws:ecs:us-west-2:123456789012:service/my-http-service',
  cluster,
});

// Import service from EC2 service ARN
const service = ecs.Ec2Service.fromEc2ServiceArn(stack, 'EcsService', 'arn:aws:ecs:us-west-2:123456789012:service/my-http-service');

// Import service from Fargate service attributes
const service = ecs.FargateService.fromFargateServiceAttributes(stack, 'EcsService', {
  serviceArn: 'arn:aws:ecs:us-west-2:123456789012:service/my-http-service',
  cluster,
});

// Import service from Fargate service ARN
const service = ecs.FargateService.fromFargateServiceArn(stack, 'EcsService', 'arn:aws:ecs:us-west-2:123456789012:service/my-http-service');
```

## Task Auto-Scaling

You can configure the task count of a service to match demand. Task auto-scaling is
configured by calling `autoScaleTaskCount()`:

```go
var target applicationTargetGroup
var service baseService

scaling := service.autoScaleTaskCount(&enableScalingProps{
	maxCapacity: jsii.Number(10),
})
scaling.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
	targetUtilizationPercent: jsii.Number(50),
})

scaling.scaleOnRequestCount(jsii.String("RequestScaling"), &requestCountScalingProps{
	requestsPerTarget: jsii.Number(10000),
	targetGroup: target,
})
```

Task auto-scaling is powered by *Application Auto-Scaling*.
See that section for details.

## Integration with CloudWatch Events

To start an Amazon ECS task on an Amazon EC2-backed Cluster, instantiate an
`@aws-cdk/aws-events-targets.EcsTask` instead of an `Ec2Service`:

```go
var cluster cluster

// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromAsset(path.resolve(__dirname, jsii.String(".."), jsii.String("eventhandler-image"))),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.NewAwsLogDriver(&awsLogDriverProps{
		streamPrefix: jsii.String("EventDemo"),
		mode: ecs.awsLogDriverMode_NON_BLOCKING,
	}),
})

// An Rule that describes the event trigger (in this case a scheduled run)
rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
	schedule: events.schedule.expression(jsii.String("rate(1 min)")),
})

// Pass an environment variable to the container 'TheContainer' in the task
rule.addTarget(targets.NewEcsTask(&ecsTaskProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	taskCount: jsii.Number(1),
	containerOverrides: []containerOverride{
		&containerOverride{
			containerName: jsii.String("TheContainer"),
			environment: []taskEnvironmentVariable{
				&taskEnvironmentVariable{
					name: jsii.String("I_WAS_TRIGGERED"),
					value: jsii.String("From CloudWatch Events"),
				},
			},
		},
	},
}))
```

## Log Drivers

Currently Supported Log Drivers:

* awslogs
* fluentd
* gelf
* journald
* json-file
* splunk
* syslog
* awsfirelens
* Generic

### awslogs Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.awsLogs(&awsLogDriverProps{
		streamPrefix: jsii.String("EventDemo"),
	}),
})
```

### fluentd Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.fluentd(),
})
```

### gelf Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.gelf(&gelfLogDriverProps{
		address: jsii.String("my-gelf-address"),
	}),
})
```

### journald Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.journald(),
})
```

### json-file Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.jsonFile(),
})
```

### splunk Log Driver

```go
var secret secret


// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.splunk(&splunkLogDriverProps{
		secretToken: secret,
		url: jsii.String("my-splunk-url"),
	}),
})
```

### syslog Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.syslog(),
})
```

### firelens Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.firelens(&fireLensLogDriverProps{
		options: map[string]*string{
			"Name": jsii.String("firehose"),
			"region": jsii.String("us-west-2"),
			"delivery_stream": jsii.String("my-stream"),
		},
	}),
})
```

To pass secrets to the log configuration, use the `secretOptions` property of the log configuration. The task execution role is automatically granted read permissions on the secrets/parameters.

```go
var secret secret
var parameter stringParameter


taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.logDrivers.firelens(&fireLensLogDriverProps{
		options: map[string]interface{}{
		},
		secretOptions: map[string]secret{
			 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store
			"apikey": ecs.*secret.fromSecretsManager(secret),
			"host": ecs.*secret.fromSsmParameter(parameter),
		},
	}),
})
```

### Generic Log Driver

A generic log driver object exists to provide a lower level abstraction of the log driver configuration.

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
	memoryLimitMiB: jsii.Number(256),
	logging: ecs.NewGenericLogDriver(&genericLogDriverProps{
		logDriver: jsii.String("fluentd"),
		options: map[string]*string{
			"tag": jsii.String("example-tag"),
		},
	}),
})
```

## CloudMap Service Discovery

To register your ECS service with a CloudMap Service Registry, you may add the
`cloudMapOptions` property to your service:

```go
var taskDefinition taskDefinition
var cluster cluster


service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	cloudMapOptions: &cloudMapOptions{
		// Create A records - useful for AWSVPC network mode.
		dnsRecordType: cloudmap.dnsRecordType_A,
	},
})
```

With `bridge` or `host` network modes, only `SRV` DNS record types are supported.
By default, `SRV` DNS record types will target the default container and default
port. However, you may target a different container and port on the same ECS task:

```go
var taskDefinition taskDefinition
var cluster cluster


// Add a container to the task definition
specificContainer := taskDefinition.addContainer(jsii.String("Container"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("/aws/aws-example-app")),
	memoryLimitMiB: jsii.Number(2048),
})

// Add a port mapping
specificContainer.addPortMappings(&portMapping{
	containerPort: jsii.Number(7600),
	protocol: ecs.protocol_TCP,
})

ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	cloudMapOptions: &cloudMapOptions{
		// Create SRV records - useful for bridge networking
		dnsRecordType: cloudmap.dnsRecordType_SRV,
		// Targets port TCP port 7600 `specificContainer`
		container: specificContainer,
		containerPort: jsii.Number(7600),
	},
})
```

### Associate With a Specific CloudMap Service

You may associate an ECS service with a specific CloudMap service. To do
this, use the service's `associateCloudMapService` method:

```go
var cloudMapService service
var ecsService fargateService


ecsService.associateCloudMapService(&associateCloudMapServiceOptions{
	service: cloudMapService,
})
```

## Capacity Providers

There are two major families of Capacity Providers: [AWS
Fargate](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-capacity-providers.html)
(including Fargate Spot) and EC2 [Auto Scaling
Group](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html)
Capacity Providers. Both are supported.

### Fargate Capacity Providers

To enable Fargate capacity providers, you can either set
`enableFargateCapacityProviders` to `true` when creating your cluster, or by
invoking the `enableFargateCapacityProviders()` method after creating your
cluster. This will add both `FARGATE` and `FARGATE_SPOT` as available capacity
providers on your cluster.

```go
var vpc vpc


cluster := ecs.NewCluster(this, jsii.String("FargateCPCluster"), &clusterProps{
	vpc: vpc,
	enableFargateCapacityProviders: jsii.Boolean(true),
})

taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"))

taskDefinition.addContainer(jsii.String("web"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
})

ecs.NewFargateService(this, jsii.String("FargateService"), &fargateServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
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

### Auto Scaling Group Capacity Providers

To add an Auto Scaling Group Capacity Provider, first create an EC2 Auto Scaling
Group. Then, create an `AsgCapacityProvider` and pass the Auto Scaling Group to
it in the constructor. Then add the Capacity Provider to the cluster. Finally,
you can refer to the Provider by its name in your service's or task's Capacity
Provider strategy.

By default, an Auto Scaling Group Capacity Provider will manage the Auto Scaling
Group's size for you. It will also enable managed termination protection, in
order to prevent EC2 Auto Scaling from terminating EC2 instances that have tasks
running on them. If you want to disable this behavior, set both
`enableManagedScaling` to and `enableManagedTerminationProtection` to `false`.

```go
var vpc vpc


cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
	vpc: vpc,
})

autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
	vpc: vpc,
	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
	machineImage: ecs.ecsOptimizedImage.amazonLinux2(),
	minCapacity: jsii.Number(0),
	maxCapacity: jsii.Number(100),
})

capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &asgCapacityProviderProps{
	autoScalingGroup: autoScalingGroup,
})
cluster.addAsgCapacityProvider(capacityProvider)

taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))

taskDefinition.addContainer(jsii.String("web"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	memoryReservationMiB: jsii.Number(256),
})

ecs.NewEc2Service(this, jsii.String("EC2Service"), &ec2ServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	capacityProviderStrategies: []capacityProviderStrategy{
		&capacityProviderStrategy{
			capacityProvider: capacityProvider.capacityProviderName,
			weight: jsii.Number(1),
		},
	},
})
```

## Elastic Inference Accelerators

Currently, this feature is only supported for services with EC2 launch types.

To add elastic inference accelerators to your EC2 instance, first add
`inferenceAccelerators` field to the Ec2TaskDefinition and set the `deviceName`
and `deviceType` properties.

```go
inferenceAccelerators := []map[string]*string{
	map[string]*string{
		"deviceName": jsii.String("device1"),
		"deviceType": jsii.String("eia2.medium"),
	},
}

taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("Ec2TaskDef"), &ec2TaskDefinitionProps{
	inferenceAccelerators: inferenceAccelerators,
})
```

To enable using the inference accelerators in the containers, add `inferenceAcceleratorResources`
field and set it to a list of device names used for the inference accelerators. Each value in the
list should match a `DeviceName` for an `InferenceAccelerator` specified in the task definition.

```go
var taskDefinition taskDefinition

inferenceAcceleratorResources := []*string{
	"device1",
}

taskDefinition.addContainer(jsii.String("cont"), &containerDefinitionOptions{
	image: ecs.containerImage.fromRegistry(jsii.String("test")),
	memoryLimitMiB: jsii.Number(1024),
	inferenceAcceleratorResources: inferenceAcceleratorResources,
})
```

## ECS Exec command

Please note, ECS Exec leverages AWS Systems Manager (SSM). So as a prerequisite for the exec command
to work, you need to have the SSM plugin for the AWS CLI installed locally. For more information, see
[Install Session Manager plugin for AWS CLI](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html).

To enable the ECS Exec feature for your containers, set the boolean flag `enableExecuteCommand` to `true` in
your `Ec2Service` or `FargateService`.

```go
var cluster cluster
var taskDefinition taskDefinition


service := ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	enableExecuteCommand: jsii.Boolean(true),
})
```

### Enabling logging

You can enable sending logs of your execute session commands to a CloudWatch log group or S3 bucket by configuring
the `executeCommandConfiguration` property for your cluster. The default configuration will send the
logs to the CloudWatch Logs using the `awslogs` log driver that is configured in your task definition. Please note,
when using your own `logConfiguration` the log group or S3 Bucket specified must already be created.

To encrypt data using your own KMS Customer Key (CMK), you must create a CMK and provide the key in the `kmsKey` field
of the `executeCommandConfiguration`. To use this key for encrypting CloudWatch log data or S3 bucket, make sure to associate the key
to these resources on creation.

```go
var vpc vpc

kmsKey := kms.NewKey(this, jsii.String("KmsKey"))

// Pass the KMS key in the `encryptionKey` field to associate the key to the log group
logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &logGroupProps{
	encryptionKey: kmsKey,
})

// Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &bucketProps{
	encryptionKey: kmsKey,
})

cluster := ecs.NewCluster(this, jsii.String("Cluster"), &clusterProps{
	vpc: vpc,
	executeCommandConfiguration: &executeCommandConfiguration{
		kmsKey: kmsKey,
		logConfiguration: &executeCommandLogConfiguration{
			cloudWatchLogGroup: logGroup,
			cloudWatchEncryptionEnabled: jsii.Boolean(true),
			s3Bucket: execBucket,
			s3EncryptionEnabled: jsii.Boolean(true),
			s3KeyPrefix: jsii.String("exec-command-output"),
		},
		logging: ecs.executeCommandLogging_OVERRIDE,
	},
})
```

## Amazon ECS Service Connect

Service Connect is a managed AWS mesh network offering. It simplifies DNS queries and inter-service communication for
ECS Services by allowing customers to set up simple DNS aliases for their services, which are accessible to all
services that have enabled Service Connect.

To enable Service Connect, you must have created a CloudMap namespace. The CDK can infer your cluster's default CloudMap namespace,
or you can specify a custom namespace. You must also have created a named port mapping on at least one container in your Task Definition.

```go
// Example automatically generated from non-compiling source. May contain errors.
var cluster cluster
var taskDefinition taskDefinition
var container containerDefinition


container.addPortMappings(&portMapping{
	name: jsii.String("api"),
	containerPort: jsii.Number(8080),
})

taskDefinition.addContainer(container)

cluster.addDefaultCloudMapNamespace(&cloudMapNamespaceOptions{
	name: jsii.String("local"),
})

service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	serviceConnectConfiguration: &serviceConnectProps{
		services: []serviceConnectService{
			&serviceConnectService{
				portMappingName: jsii.String("api"),
				dnsName: jsii.String("http-api"),
				port: jsii.Number(80),
			},
		},
	},
})
```

Service Connect-enabled services may now reach this service at `http-api:80`. Traffic to this endpoint will
be routed to the container's port 8080.

To opt a service into using service connect without advertising a port, simply call the 'enableServiceConnect' method on an initialized service.

```go
// Example automatically generated from non-compiling source. May contain errors.
service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
})
service.enableServiceConnect()
```

Service Connect also allows custom logging, Service Discovery name, and configuration of the port where service connect traffic is received.

```go
// Example automatically generated from non-compiling source. May contain errors.
customService := ecs.NewFargateService(this, jsii.String("CustomizedService"), &fargateServiceProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	serviceConnectConfiguration: &serviceConnectProps{
		logDriver: ecs.logDrivers.awslogs(map[string]*string{
			"streamPrefix": jsii.String("sc-traffic"),
		}),
		services: []serviceConnectService{
			&serviceConnectService{
				portMappingName: jsii.String("api"),
				dnsName: jsii.String("customized-api"),
				port: jsii.Number(80),
				ingressPortOverride: jsii.Number(20040),
				discoveryName: jsii.String("custom"),
			},
		},
	},
})
```
