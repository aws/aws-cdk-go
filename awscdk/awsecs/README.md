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
cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	Vpc: Vpc,
})

// Add capacity to it
cluster.AddCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &AddCapacityOptions{
	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
	DesiredCapacity: jsii.Number(3),
})

taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))

taskDefinition.AddContainer(jsii.String("DefaultContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	MemoryLimitMiB: jsii.Number(512),
})

// Instantiate an Amazon ECS Service
ecsService := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})
```

For a set of constructs defining common ECS architectural patterns, see the `aws-cdk-lib/aws-ecs-patterns` package.

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


cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	Vpc: Vpc,
})
```

The following code imports an existing cluster using the ARN which can be used to
import an Amazon ECS service either EC2 or Fargate.

```go
clusterArn := "arn:aws:ecs:us-east-1:012345678910:cluster/clusterName"

cluster := ecs.Cluster_FromClusterArn(this, jsii.String("Cluster"), clusterArn)
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


cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	Vpc: Vpc,
})

// Either add default capacity
cluster.AddCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &AddCapacityOptions{
	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
	DesiredCapacity: jsii.Number(3),
})

// Or add customized capacity. Be sure to start the Amazon ECS-optimized AMI.
autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
	MachineImage: ecs.EcsOptimizedImage_AmazonLinux(),
	// Or use Amazon ECS-Optimized Amazon Linux 2 AMI
	// machineImage: EcsOptimizedImage.amazonLinux2(),
	DesiredCapacity: jsii.Number(3),
})

capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &AsgCapacityProviderProps{
	AutoScalingGroup: AutoScalingGroup,
})
cluster.AddAsgCapacityProvider(capacityProvider)
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

autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	MachineImage: ecs.EcsOptimizedImage_AmazonLinux(&EcsOptimizedImageOptions{
		CachedInContext: jsii.Boolean(true),
	}),
	Vpc: Vpc,
	InstanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
})
```

To use `LaunchTemplate` with `AsgCapacityProvider`, make sure to specify the `userData` in the `LaunchTemplate`:

```go
var vpc vpc

launchTemplate := ec2.NewLaunchTemplate(this, jsii.String("ASG-LaunchTemplate"), &LaunchTemplateProps{
	InstanceType: ec2.NewInstanceType(jsii.String("t3.medium")),
	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(),
	UserData: ec2.UserData_ForLinux(),
})

autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	MixedInstancesPolicy: &MixedInstancesPolicy{
		InstancesDistribution: &InstancesDistribution{
			OnDemandPercentageAboveBaseCapacity: jsii.Number(50),
		},
		LaunchTemplate: launchTemplate,
	},
})

cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	Vpc: Vpc,
})

capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &AsgCapacityProviderProps{
	AutoScalingGroup: AutoScalingGroup,
	MachineImageType: ecs.MachineImageType_AMAZON_LINUX_2,
})

cluster.AddAsgCapacityProvider(capacityProvider)
```

The following code retrieve the Amazon Resource Names (ARNs) of tasks that are a part of a specified ECS cluster.
It's useful when you want to grant permissions to a task to access other AWS resources.

```go
var cluster cluster
var taskDefinition taskDefinition

taskARNs := cluster.ArnForTasks(jsii.String("*")) // arn:aws:ecs:<region>:<regionId>:task/<clusterName>/*

// Grant the task permission to access other AWS resources
taskDefinition.AddToTaskRolePolicy(
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("ecs:UpdateTaskProtection"),
	},
	Resources: []*string{
		taskARNs,
	},
}))
```

To manage task protection settings in an ECS cluster, you can use the `grantTaskProtection` method.
This method grants the `ecs:UpdateTaskProtection` permission to a specified IAM entity.

```go
// Assume 'cluster' is an instance of ecs.Cluster
var cluster cluster
var taskRole role


// Grant ECS Task Protection permissions to the role
// Now 'taskRole' has the 'ecs:UpdateTaskProtection' permission on all tasks in the cluster
cluster.GrantTaskProtection(taskRole)
```

### Bottlerocket

[Bottlerocket](https://aws.amazon.com/bottlerocket/) is a Linux-based open source operating system that is
purpose-built by AWS for running containers. You can launch Amazon ECS container instances with the Bottlerocket AMI.

The following example will create a capacity with self-managed Amazon EC2 capacity of 2 `c5.large` Linux instances running with `Bottlerocket` AMI.

The following example adds Bottlerocket capacity to the cluster:

```go
var cluster cluster


cluster.AddCapacity(jsii.String("bottlerocket-asg"), &AddCapacityOptions{
	MinCapacity: jsii.Number(2),
	InstanceType: ec2.NewInstanceType(jsii.String("c5.large")),
	MachineImage: ecs.NewBottleRocketImage(),
})
```

You can also specify an NVIDIA-compatible AMI such as in this example:

```go
var cluster cluster


cluster.AddCapacity(jsii.String("bottlerocket-asg"), &AddCapacityOptions{
	InstanceType: ec2.NewInstanceType(jsii.String("p3.2xlarge")),
	MachineImage: ecs.NewBottleRocketImage(&BottleRocketImageProps{
		Variant: ecs.BottlerocketEcsVariant_AWS_ECS_2_NVIDIA,
	}),
})
```

### ARM64 (Graviton) Instances

To launch instances with ARM64 hardware, you can use the Amazon ECS-optimized
Amazon Linux 2 (arm64) AMI. Based on Amazon Linux 2, this AMI is recommended
for use when launching your EC2 instances that are powered by Arm-based AWS
Graviton Processors.

```go
var cluster cluster


cluster.AddCapacity(jsii.String("graviton-cluster"), &AddCapacityOptions{
	MinCapacity: jsii.Number(2),
	InstanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(ecs.AmiHardwareType_ARM),
})
```

Bottlerocket is also supported:

```go
var cluster cluster


cluster.AddCapacity(jsii.String("graviton-cluster"), &AddCapacityOptions{
	MinCapacity: jsii.Number(2),
	InstanceType: ec2.NewInstanceType(jsii.String("c6g.large")),
	MachineImageType: ecs.MachineImageType_BOTTLEROCKET,
})
```

### Amazon Linux 2 (Neuron) Instances

To launch Amazon EC2 Inf1, Trn1 or Inf2 instances, you can use the Amazon ECS optimized Amazon Linux 2 (Neuron) AMI. It comes pre-configured with AWS Inferentia and AWS Trainium drivers and the AWS Neuron runtime for Docker which makes running machine learning inference workloads easier on Amazon ECS.

```go
var cluster cluster


cluster.AddCapacity(jsii.String("neuron-cluster"), &AddCapacityOptions{
	MinCapacity: jsii.Number(2),
	InstanceType: ec2.NewInstanceType(jsii.String("inf1.xlarge")),
	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(ecs.AmiHardwareType_NEURON),
})
```

### Spot Instances

To add spot instances into the cluster, you must specify the `spotPrice` in the `ecs.AddCapacityOptions` and optionally enable the `spotInstanceDraining` property.

```go
var cluster cluster


// Add an AutoScalingGroup with spot instances to the existing cluster
cluster.AddCapacity(jsii.String("AsgSpot"), &AddCapacityOptions{
	MaxCapacity: jsii.Number(2),
	MinCapacity: jsii.Number(2),
	DesiredCapacity: jsii.Number(2),
	InstanceType: ec2.NewInstanceType(jsii.String("c5.xlarge")),
	SpotPrice: jsii.String("0.0735"),
	// Enable the Automated Spot Draining support for Amazon ECS
	SpotInstanceDraining: jsii.Boolean(true),
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
cluster.AddCapacity(jsii.String("ASGEncryptedSNS"), &AddCapacityOptions{
	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
	DesiredCapacity: jsii.Number(3),
	TopicEncryptionKey: key,
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
fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
	MemoryLimitMiB: jsii.Number(512),
	Cpu: jsii.Number(256),
})
```

On Fargate Platform Version 1.4.0 or later, you may specify up to 200GiB of
[ephemeral storage](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/fargate-task-storage.html#fargate-task-storage-pv14):

```go
fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
	MemoryLimitMiB: jsii.Number(512),
	Cpu: jsii.Number(256),
	EphemeralStorageGiB: jsii.Number(100),
})
```

To add containers to a task definition, call `addContainer()`:

```go
fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
	MemoryLimitMiB: jsii.Number(512),
	Cpu: jsii.Number(256),
})
container := fargateTaskDefinition.AddContainer(jsii.String("WebContainer"), &ContainerDefinitionOptions{
	// Use an image from DockerHub
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
})
```

For an `Ec2TaskDefinition`:

```go
ec2TaskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"), &Ec2TaskDefinitionProps{
	NetworkMode: ecs.NetworkMode_BRIDGE,
})

container := ec2TaskDefinition.AddContainer(jsii.String("WebContainer"), &ContainerDefinitionOptions{
	// Use an image from DockerHub
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	MemoryLimitMiB: jsii.Number(1024),
})
```

For an `ExternalTaskDefinition`:

```go
externalTaskDefinition := ecs.NewExternalTaskDefinition(this, jsii.String("TaskDef"))

container := externalTaskDefinition.AddContainer(jsii.String("WebContainer"), &ContainerDefinitionOptions{
	// Use an image from DockerHub
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	MemoryLimitMiB: jsii.Number(1024),
})
```

You can specify container properties when you add them to the task definition, or with various methods, e.g.:

To add a port mapping when adding a container to the task definition, specify the `portMappings` option:

```go
var taskDefinition taskDefinition


taskDefinition.AddContainer(jsii.String("WebContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	MemoryLimitMiB: jsii.Number(1024),
	PortMappings: []portMapping{
		&portMapping{
			ContainerPort: jsii.Number(3000),
		},
	},
})
```

To add port mappings directly to a container definition, call `addPortMappings()`:

```go
var container containerDefinition


container.AddPortMappings(&PortMapping{
	ContainerPort: jsii.Number(3000),
})
```

Sometimes it is useful to be able to configure port ranges for a container, e.g. to run applications such as game servers
and real-time streaming which typically require multiple ports to be opened simultaneously. This feature is supported on
both Linux and Windows operating systems for both the EC2 and AWS Fargate launch types. There is a maximum limit of 100
port ranges per container, and you cannot specify overlapping port ranges.

Docker recommends that you turn off the `docker-proxy` in the Docker daemon config file when you have a large number of ports.
For more information, see [Issue #11185](https://github.com/moby/moby/issues/11185) on the GitHub website.

```go
var container containerDefinition


container.AddPortMappings(&PortMapping{
	ContainerPort: ecs.*containerDefinition_CONTAINER_PORT_USE_RANGE(),
	ContainerPortRange: jsii.String("8080-8081"),
})
```

To add data volumes to a task definition, call `addVolume()`:

```go
fargateTaskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
	MemoryLimitMiB: jsii.Number(512),
	Cpu: jsii.Number(256),
})
volume := map[string]interface{}{
	// Use an Elastic FileSystem
	"name": jsii.String("mydatavolume"),
	"efsVolumeConfiguration": map[string]*string{
		"fileSystemId": jsii.String("EFS"),
	},
}

container := fargateTaskDefinition.AddVolume(volume)
```

> Note: ECS Anywhere doesn't support volume attachments in the task definition.

To use a TaskDefinition that can be used with either Amazon EC2 or
AWS Fargate launch types, use the `TaskDefinition` construct.

When creating a task definition you have to specify what kind of
tasks you intend to run: Amazon EC2, AWS Fargate, or both.
The following example uses both:

```go
taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDef"), &TaskDefinitionProps{
	MemoryMiB: jsii.String("512"),
	Cpu: jsii.String("256"),
	NetworkMode: ecs.NetworkMode_AWS_VPC,
	Compatibility: ecs.Compatibility_EC2_AND_FARGATE,
})
```

To grant a principal permission to run your `TaskDefinition`, you can use the `TaskDefinition.grantRun()` method:

```go
var role iGrantable

taskDef := ecs.NewTaskDefinition(this, jsii.String("TaskDef"), &TaskDefinitionProps{
	Cpu: jsii.String("512"),
	MemoryMiB: jsii.String("512"),
	Compatibility: ecs.Compatibility_EC2_AND_FARGATE,
})

// Gives role required permissions to run taskDef
taskDef.GrantRun(role)
```

To deploy containerized applications that require the allocation of standard input (stdin) or a terminal (tty), use the `interactive` property.

This parameter corresponds to `OpenStdin` in the [Create a container](https://docs.docker.com/engine/api/v1.35/#tag/Container/operation/ContainerCreate) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.35/)
and the `--interactive` option to [docker run](https://docs.docker.com/engine/reference/run/#security-configuration).

```go
var taskDefinition taskDefinition


taskDefinition.AddContainer(jsii.String("Container"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	Interactive: jsii.Boolean(true),
})
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
  `aws-cdk-lib/aws-ecr-assets.DockerImageAsset` as a container image.
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


newContainer := taskDefinition.AddContainer(jsii.String("container"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	MemoryLimitMiB: jsii.Number(1024),
	Environment: map[string]*string{
		 // clear text, not for sensitive data
		"STAGE": jsii.String("prod"),
	},
	EnvironmentFiles: []environmentFile{
		ecs.*environmentFile_FromAsset(jsii.String("./demo-env-file.env")),
		ecs.*environmentFile_FromBucket(s3Bucket, jsii.String("assets/demo-env-file.env")),
	},
	Secrets: map[string]secret{
		 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store at container start-up.
		"SECRET": ecs.*secret_fromSecretsManager(secret),
		"DB_PASSWORD": ecs.*secret_fromSecretsManager(dbSecret, jsii.String("password")),
		 // Reference a specific JSON field, (requires platform version 1.4.0 or later for Fargate tasks)
		"API_KEY": ecs.*secret_fromSecretsManagerVersion(secret, &SecretVersionInfo{
			"versionId": jsii.String("12345"),
		}, jsii.String("apiKey")),
		 // Reference a specific version of the secret by its version id or version stage (requires platform version 1.4.0 or later for Fargate tasks)
		"PARAMETER": ecs.*secret_fromSsmParameter(parameter),
	},
})
newContainer.AddEnvironment(jsii.String("QUEUE_NAME"), jsii.String("MyQueue"))
newContainer.AddSecret(jsii.String("API_KEY"), ecs.secret_FromSecretsManager(secret))
newContainer.AddSecret(jsii.String("DB_PASSWORD"), ecs.secret_FromSecretsManager(secret, jsii.String("password")))
```

The task execution role is automatically granted read permissions on the secrets/parameters. Further details provided in the AWS documentation
about [specifying environment variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html).

### Linux parameters

To apply additional linux-specific options related to init process and memory management to the container, use the `linuxParameters` property:

```go
var taskDefinition taskDefinition


taskDefinition.AddContainer(jsii.String("container"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	MemoryLimitMiB: jsii.Number(1024),
	LinuxParameters: ecs.NewLinuxParameters(this, jsii.String("LinuxParameters"), &LinuxParametersProps{
		InitProcessEnabled: jsii.Boolean(true),
		SharedMemorySize: jsii.Number(1024),
		MaxSwap: awscdk.Size_Mebibytes(jsii.Number(5000)),
		Swappiness: jsii.Number(90),
	}),
})
```

### System controls

To set system controls (kernel parameters) on the container, use the `systemControls` prop:

```go
var taskDefinition taskDefinition


taskDefinition.AddContainer(jsii.String("container"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	MemoryLimitMiB: jsii.Number(1024),
	SystemControls: []systemControl{
		&systemControl{
			Namespace: jsii.String("net.ipv6.conf.all.default.disable_ipv6"),
			Value: jsii.String("1"),
		},
	},
})
```

### Using Windows containers on Fargate

AWS Fargate supports Amazon ECS Windows containers. For more details, please see this [blog post](https://aws.amazon.com/tw/blogs/containers/running-windows-containers-with-amazon-ecs-on-aws-fargate/)

```go
// Create a Task Definition for the Windows container to start
taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
	RuntimePlatform: &RuntimePlatform{
		OperatingSystemFamily: ecs.OperatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
		CpuArchitecture: ecs.CpuArchitecture_X86_64(),
	},
	Cpu: jsii.Number(1024),
	MemoryLimitMiB: jsii.Number(2048),
})

taskDefinition.AddContainer(jsii.String("windowsservercore"), &ContainerDefinitionOptions{
	Logging: ecs.LogDriver_AwsLogs(&AwsLogDriverProps{
		StreamPrefix: jsii.String("win-iis-on-fargate"),
	}),
	PortMappings: []portMapping{
		&portMapping{
			ContainerPort: jsii.Number(80),
		},
	},
	Image: ecs.ContainerImage_FromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
})
```

### Using Windows authentication with gMSA

Amazon ECS supports Active Directory authentication for Linux containers through a special kind of service account called a group Managed Service Account (gMSA). For more details, please see the [product documentation on how to implement on Windows containers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows-gmsa.html), or this [blog post on how to implement on  Linux containers](https://aws.amazon.com/blogs/containers/using-windows-authentication-with-gmsa-on-linux-containers-on-amazon-ecs/).

There are two types of CredentialSpecs, domained-join or domainless. Both types support creation from a S3 bucket, a SSM parameter, or by directly specifying a location for the file in the constructor.

A domian-joined gMSA container looks like:

```go
// Make sure the task definition's execution role has permissions to read from the S3 bucket or SSM parameter where the CredSpec file is stored.
var parameter iParameter
var taskDefinition taskDefinition


// Domain-joined gMSA container from a SSM parameter
taskDefinition.AddContainer(jsii.String("gmsa-domain-joined-container"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	Cpu: jsii.Number(128),
	MemoryLimitMiB: jsii.Number(256),
	CredentialSpecs: []credentialSpec{
		ecs.DomainJoinedCredentialSpec_FromSsmParameter(parameter),
	},
})
```

A domianless gMSA container looks like:

```go
// Make sure the task definition's execution role has permissions to read from the S3 bucket or SSM parameter where the CredSpec file is stored.
var bucket bucket
var taskDefinition taskDefinition


// Domainless gMSA container from a S3 bucket object.
taskDefinition.AddContainer(jsii.String("gmsa-domainless-container"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	Cpu: jsii.Number(128),
	MemoryLimitMiB: jsii.Number(256),
	CredentialSpecs: []credentialSpec{
		ecs.DomainlessCredentialSpec_FromS3Bucket(bucket, jsii.String("credSpec")),
	},
})
```

### Using Graviton2 with Fargate

AWS Graviton2 supports AWS Fargate. For more details, please see this [blog post](https://aws.amazon.com/blogs/aws/announcing-aws-graviton2-support-for-aws-fargate-get-up-to-40-better-price-performance-for-your-serverless-containers/)

```go
// Create a Task Definition for running container on Graviton Runtime.
taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &FargateTaskDefinitionProps{
	RuntimePlatform: &RuntimePlatform{
		OperatingSystemFamily: ecs.OperatingSystemFamily_LINUX(),
		CpuArchitecture: ecs.CpuArchitecture_ARM64(),
	},
	Cpu: jsii.Number(1024),
	MemoryLimitMiB: jsii.Number(2048),
})

taskDefinition.AddContainer(jsii.String("webarm64"), &ContainerDefinitionOptions{
	Logging: ecs.LogDriver_AwsLogs(&AwsLogDriverProps{
		StreamPrefix: jsii.String("graviton2-on-fargate"),
	}),
	PortMappings: []portMapping{
		&portMapping{
			ContainerPort: jsii.Number(80),
		},
	},
	Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/nginx/nginx:latest-arm64v8")),
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


service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	DesiredCount: jsii.Number(5),
})
```

ECS Anywhere service definition looks like:

```go
var cluster cluster
var taskDefinition taskDefinition


service := ecs.NewExternalService(this, jsii.String("Service"), &ExternalServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	DesiredCount: jsii.Number(5),
})
```

`Services` by default will create a security group if not provided.
If you'd like to specify which security groups to use you can override the `securityGroups` property.

By default, the service will use the revision of the passed task definition generated when the `TaskDefinition`
is deployed by CloudFormation. However, this may not be desired if the revision is externally managed,
for example through CodeDeploy.

To set a specific revision number or the special `latest` revision, use the `taskDefinitionRevision` parameter:

```go
var cluster cluster
var taskDefinition taskDefinition


ecs.NewExternalService(this, jsii.String("Service"), &ExternalServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	DesiredCount: jsii.Number(5),
	TaskDefinitionRevision: ecs.TaskDefinitionRevision_Of(jsii.Number(1)),
})

ecs.NewExternalService(this, jsii.String("Service"), &ExternalServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	DesiredCount: jsii.Number(5),
	TaskDefinitionRevision: ecs.TaskDefinitionRevision_LATEST(),
})
```

### Deployment circuit breaker and rollback

Amazon ECS [deployment circuit breaker](https://aws.amazon.com/tw/blogs/containers/announcing-amazon-ecs-deployment-circuit-breaker/)
automatically rolls back unhealthy service deployments, eliminating the need for manual intervention.

Use `circuitBreaker` to enable the deployment circuit breaker which determines whether a service deployment
will fail if the service can't reach a steady state.
You can optionally enable `rollback` for automatic rollback.

See [Using the deployment circuit breaker](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html) for more details.

```go
var cluster cluster
var taskDefinition taskDefinition

service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	CircuitBreaker: &DeploymentCircuitBreaker{
		Enable: jsii.Boolean(true),
		Rollback: jsii.Boolean(true),
	},
})
```

> Note: ECS Anywhere doesn't support deployment circuit breakers and rollback.

### Deployment alarms

Amazon ECS [deployment alarms]
(https://aws.amazon.com/blogs/containers/automate-rollbacks-for-amazon-ecs-rolling-deployments-with-cloudwatch-alarms/)
allow monitoring and automatically reacting to changes during a rolling update
by using Amazon CloudWatch metric alarms.

Amazon ECS starts monitoring the configured deployment alarms as soon as one or
more tasks of the updated service are in a running state. The deployment process
continues until the primary deployment is healthy and has reached the desired
count and the active deployment has been scaled down to 0. Then, the deployment
remains in the IN_PROGRESS state for an additional "bake time." The length the
bake time is calculated based on the evaluation periods and period of the alarms.
After the bake time, if none of the alarms have been activated, then Amazon ECS
considers this to be a successful update and deletes the active deployment and
changes the status of the primary deployment to COMPLETED.

```go
import cw "github.com/aws/aws-cdk-go/awscdk"

var cluster cluster
var taskDefinition taskDefinition
var elbAlarm alarm


service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	DeploymentAlarms: &DeploymentAlarmConfig{
		AlarmNames: []*string{
			elbAlarm.AlarmName,
		},
		Behavior: ecs.AlarmBehavior_ROLLBACK_ON_ALARM,
	},
})

// Defining a deployment alarm after the service has been created
cpuAlarmName := "MyCpuMetricAlarm"
cw.NewAlarm(this, jsii.String("CPUAlarm"), &AlarmProps{
	AlarmName: cpuAlarmName,
	Metric: service.MetricCpuUtilization(),
	EvaluationPeriods: jsii.Number(2),
	Threshold: jsii.Number(80),
})
service.EnableDeploymentAlarms([]*string{
	cpuAlarmName,
}, &DeploymentAlarmOptions{
	Behavior: ecs.AlarmBehavior_FAIL_ON_ALARM,
})
```

> Note: Deployment alarms are only available when `deploymentController` is set
> to `DeploymentControllerType.ECS`, which is the default.

#### Troubleshooting circular dependencies

I saw this info message during synth time. What do I do?

```text
Deployment alarm ({"Ref":"MyAlarmABC1234"}) enabled on MyEcsService may cause a
circular dependency error when this stack deploys. The alarm name references the
alarm's logical id, or another resource. See the 'Deployment alarms' section in
the module README for more details.
```

If your app deploys successfully with this message, you can disregard it. But it
indicates that you could encounter a circular dependency error when you try to
deploy. If you want to alarm on metrics produced by the service, there will be a
circular dependency between the service and its deployment alarms. In this case,
there are two options to avoid the circular dependency.

1. Define the physical name for the alarm. Use a defined physical name that is
   unique within the deployment environment for the alarm name when creating the
   alarm, and re-use the defined name. This name could be a hardcoded string, a
   string generated based on the environment, or could reference another
   resource that does not depend on the service.
2. Define the physical name for the service. Then, don't use
   `metricCpuUtilization()` or similar methods. Create the metric object
   separately by referencing the service metrics using this name.

Option 1, defining a physical name for the alarm:

```go
import cw "github.com/aws/aws-cdk-go/awscdk"

var cluster cluster
var taskDefinition taskDefinition


service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})

cpuAlarmName := "MyCpuMetricAlarm"
myAlarm := cw.NewAlarm(this, jsii.String("CPUAlarm"), &AlarmProps{
	AlarmName: cpuAlarmName,
	Metric: service.MetricCpuUtilization(),
	EvaluationPeriods: jsii.Number(2),
	Threshold: jsii.Number(80),
})

// Using `myAlarm.alarmName` here will cause a circular dependency
service.EnableDeploymentAlarms([]*string{
	cpuAlarmName,
}, &DeploymentAlarmOptions{
	Behavior: ecs.AlarmBehavior_FAIL_ON_ALARM,
})
```

Option 2, defining a physical name for the service:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var cluster cluster
var taskDefinition taskDefinition

serviceName := "MyFargateService"
service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	ServiceName: jsii.String(ServiceName),
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})

cpuMetric := cw.NewMetric(&MetricProps{
	MetricName: jsii.String("CPUUtilization"),
	Namespace: jsii.String("AWS/ECS"),
	Period: awscdk.Duration_Minutes(jsii.Number(5)),
	Statistic: jsii.String("Average"),
	DimensionsMap: map[string]*string{
		"ClusterName": cluster.clusterName,
		// Using `service.serviceName` here will cause a circular dependency
		"ServiceName": serviceName,
	},
})
myAlarm := cw.NewAlarm(this, jsii.String("CPUAlarm"), &AlarmProps{
	AlarmName: jsii.String("cpuAlarmName"),
	Metric: cpuMetric,
	EvaluationPeriods: jsii.Number(2),
	Threshold: jsii.Number(80),
})

service.EnableDeploymentAlarms([]*string{
	myAlarm.AlarmName,
}, &DeploymentAlarmOptions{
	Behavior: ecs.AlarmBehavior_FAIL_ON_ALARM,
})
```

This issue only applies if the metrics to alarm on are emitted by the service
itself. If the metrics are emitted by a different resource, that does not depend
on the service, there will be no restrictions on the alarm name.

### Include an application/network load balancer

`Services` are load balancing targets and can be added to a target group, which will be attached to an application/network load balancers:

```go
var vpc vpc
var cluster cluster
var taskDefinition taskDefinition

service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})

lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
})
listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(80),
})
targetGroup1 := listener.AddTargets(jsii.String("ECS1"), &AddApplicationTargetsProps{
	Port: jsii.Number(80),
	Targets: []iApplicationLoadBalancerTarget{
		service,
	},
})
targetGroup2 := listener.AddTargets(jsii.String("ECS2"), &AddApplicationTargetsProps{
	Port: jsii.Number(80),
	Targets: []*iApplicationLoadBalancerTarget{
		service.LoadBalancerTarget(&LoadBalancerTargetOptions{
			ContainerName: jsii.String("MyContainer"),
			ContainerPort: jsii.Number(8080),
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

service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})

lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
})
listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(80),
})
service.RegisterLoadBalancerTargets(&EcsTarget{
	ContainerName: jsii.String("web"),
	ContainerPort: jsii.Number(80),
	NewTargetGroupId: jsii.String("ECS"),
	Listener: ecs.ListenerConfig_ApplicationListener(listener, &AddApplicationTargetsProps{
		Protocol: elbv2.ApplicationProtocol_HTTPS,
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

service := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})

lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
	Vpc: Vpc,
})
lb.AddListener(&LoadBalancerListener{
	ExternalPort: jsii.Number(80),
})
lb.AddTarget(service)
```

Similarly, if you want to have more control over load balancer targeting:

```go
var cluster cluster
var taskDefinition taskDefinition
var vpc vpc

service := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})

lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
	Vpc: Vpc,
})
lb.AddListener(&LoadBalancerListener{
	ExternalPort: jsii.Number(80),
})
lb.AddTarget(service.LoadBalancerTarget(&LoadBalancerTargetOptions{
	ContainerName: jsii.String("MyContainer"),
	ContainerPort: jsii.Number(80),
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
const service = ecs.Ec2Service.fromEc2ServiceAttributes(this, 'EcsService', {
  serviceArn: 'arn:aws:ecs:us-west-2:123456789012:service/my-http-service',
  cluster,
});

// Import service from EC2 service ARN
const service = ecs.Ec2Service.fromEc2ServiceArn(this, 'EcsService', 'arn:aws:ecs:us-west-2:123456789012:service/my-http-service');

// Import service from Fargate service attributes
const service = ecs.FargateService.fromFargateServiceAttributes(this, 'EcsService', {
  serviceArn: 'arn:aws:ecs:us-west-2:123456789012:service/my-http-service',
  cluster,
});

// Import service from Fargate service ARN
const service = ecs.FargateService.fromFargateServiceArn(this, 'EcsService', 'arn:aws:ecs:us-west-2:123456789012:service/my-http-service');
```

## Task Auto-Scaling

You can configure the task count of a service to match demand. Task auto-scaling is
configured by calling `autoScaleTaskCount()`:

```go
var target applicationTargetGroup
var service baseService

scaling := service.AutoScaleTaskCount(&EnableScalingProps{
	MaxCapacity: jsii.Number(10),
})
scaling.ScaleOnCpuUtilization(jsii.String("CpuScaling"), &CpuUtilizationScalingProps{
	TargetUtilizationPercent: jsii.Number(50),
})

scaling.ScaleOnRequestCount(jsii.String("RequestScaling"), &RequestCountScalingProps{
	RequestsPerTarget: jsii.Number(10000),
	TargetGroup: target,
})
```

Task auto-scaling is powered by *Application Auto-Scaling*.
See that section for details.

## Integration with CloudWatch Events

To start an Amazon ECS task on an Amazon EC2-backed Cluster, instantiate an
`aws-cdk-lib/aws-events-targets.EcsTask` instead of an `Ec2Service`:

```go
var cluster cluster

// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromAsset(path.resolve(__dirname, jsii.String(".."), jsii.String("eventhandler-image"))),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.NewAwsLogDriver(&AwsLogDriverProps{
		StreamPrefix: jsii.String("EventDemo"),
		Mode: ecs.AwsLogDriverMode_NON_BLOCKING,
	}),
})

// An Rule that describes the event trigger (in this case a scheduled run)
rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Expression(jsii.String("rate(1 min)")),
})

// Pass an environment variable to the container 'TheContainer' in the task
rule.AddTarget(targets.NewEcsTask(&EcsTaskProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	TaskCount: jsii.Number(1),
	ContainerOverrides: []containerOverride{
		&containerOverride{
			ContainerName: jsii.String("TheContainer"),
			Environment: []taskEnvironmentVariable{
				&taskEnvironmentVariable{
					Name: jsii.String("I_WAS_TRIGGERED"),
					Value: jsii.String("From CloudWatch Events"),
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
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_AwsLogs(&AwsLogDriverProps{
		StreamPrefix: jsii.String("EventDemo"),
		Mode: ecs.AwsLogDriverMode_NON_BLOCKING,
		MaxBufferSize: awscdk.Size_Mebibytes(jsii.Number(25)),
	}),
})
```

### fluentd Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_Fluentd(),
})
```

### gelf Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_Gelf(&GelfLogDriverProps{
		Address: jsii.String("my-gelf-address"),
	}),
})
```

### journald Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_Journald(),
})
```

### json-file Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_JsonFile(),
})
```

### splunk Log Driver

```go
var secret secret


// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_Splunk(&SplunkLogDriverProps{
		SecretToken: secret,
		Url: jsii.String("my-splunk-url"),
	}),
})
```

### syslog Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_Syslog(),
})
```

### firelens Log Driver

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_Firelens(&FireLensLogDriverProps{
		Options: map[string]*string{
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
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_Firelens(&FireLensLogDriverProps{
		Options: map[string]interface{}{
		},
		SecretOptions: map[string]secret{
			 // Retrieved from AWS Secrets Manager or AWS Systems Manager Parameter Store
			"apikey": ecs.*secret_fromSecretsManager(secret),
			"host": ecs.*secret_fromSsmParameter(parameter),
		},
	}),
})
```

When forwarding logs to CloudWatch Logs using Fluent Bit, you can set the retention period for the newly created Log Group by specifying the `log_retention_days` parameter.
If a Fluent Bit container has not been added, CDK will automatically add it to the task definition, and the necessary IAM permissions will be added to the task role.
If you are adding the Fluent Bit container manually, ensure to add the `logs:PutRetentionPolicy` policy to the task role.

```go
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.LogDrivers_Firelens(&FireLensLogDriverProps{
		Options: map[string]*string{
			"Name": jsii.String("cloudwatch"),
			"region": jsii.String("us-west-2"),
			"log_group_name": jsii.String("firelens-fluent-bit"),
			"log_stream_prefix": jsii.String("from-fluent-bit"),
			"auto_create_group": jsii.String("true"),
			"log_retention_days": jsii.String("1"),
		},
	}),
})
```

> Visit [Fluent Bit CloudWatch Configuration Parameters](https://docs.fluentbit.io/manual/pipeline/outputs/cloudwatch#configuration-parameters)
> for more details.

### Generic Log Driver

A generic log driver object exists to provide a lower level abstraction of the log driver configuration.

```go
// Create a Task Definition for the container to start
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	MemoryLimitMiB: jsii.Number(256),
	Logging: ecs.NewGenericLogDriver(&GenericLogDriverProps{
		LogDriver: jsii.String("fluentd"),
		Options: map[string]*string{
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


service := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	CloudMapOptions: &CloudMapOptions{
		// Create A records - useful for AWSVPC network mode.
		DnsRecordType: cloudmap.DnsRecordType_A,
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
specificContainer := taskDefinition.AddContainer(jsii.String("Container"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("/aws/aws-example-app")),
	MemoryLimitMiB: jsii.Number(2048),
})

// Add a port mapping
specificContainer.AddPortMappings(&PortMapping{
	ContainerPort: jsii.Number(7600),
	Protocol: ecs.Protocol_TCP,
})

ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	CloudMapOptions: &CloudMapOptions{
		// Create SRV records - useful for bridge networking
		DnsRecordType: cloudmap.DnsRecordType_SRV,
		// Targets port TCP port 7600 `specificContainer`
		Container: specificContainer,
		ContainerPort: jsii.Number(7600),
	},
})
```

### Associate With a Specific CloudMap Service

You may associate an ECS service with a specific CloudMap service. To do
this, use the service's `associateCloudMapService` method:

```go
var cloudMapService service
var ecsService fargateService


ecsService.AssociateCloudMapService(&AssociateCloudMapServiceOptions{
	Service: cloudMapService,
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


cluster := ecs.NewCluster(this, jsii.String("FargateCPCluster"), &ClusterProps{
	Vpc: Vpc,
	EnableFargateCapacityProviders: jsii.Boolean(true),
})

taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"))

taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
})

ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
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

### Auto Scaling Group Capacity Providers

To add an Auto Scaling Group Capacity Provider, first create an EC2 Auto Scaling
Group. Then, create an `AsgCapacityProvider` and pass the Auto Scaling Group to
it in the constructor. Then add the Capacity Provider to the cluster. Finally,
you can refer to the Provider by its name in your service's or task's Capacity
Provider strategy.

By default, Auto Scaling Group Capacity Providers will manage the scale-in and
scale-out behavior of the auto scaling group based on the load your tasks put on
the cluster, this is called [Managed Scaling](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/asg-capacity-providers.html#asg-capacity-providers-managed-scaling). If you'd
rather manage scaling behavior yourself set `enableManagedScaling` to `false`.

Additionally [Managed Termination Protection](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-auto-scaling.html#managed-termination-protection) is enabled by default to
prevent scale-in behavior from terminating instances that have non-daemon tasks
running on them. This is ideal for tasks that can be run to completion. If your
tasks are safe to interrupt then this protection can be disabled by setting
`enableManagedTerminationProtection` to `false`. Managed Scaling must be enabled for
Managed Termination Protection to work.

> Currently there is a known [CloudFormation issue](https://github.com/aws/containers-roadmap/issues/631)
> that prevents CloudFormation from automatically deleting Auto Scaling Groups that
> have Managed Termination Protection enabled. To work around this issue you could set
> `enableManagedTerminationProtection` to `false` on the Auto Scaling Group Capacity
> Provider. If you'd rather not disable Managed Termination Protection, you can [manually
> delete the Auto Scaling Group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-process-shutdown.html).
> For other workarounds, see [this GitHub issue](https://github.com/aws/aws-cdk/issues/18179).

Managed instance draining facilitates graceful termination of Amazon ECS instances.
This allows your service workloads to stop safely and be rescheduled to non-terminating instances.
Infrastructure maintenance and updates are preformed without disruptions to workloads.
To use managed instance draining, set enableManagedDraining to true.

```go
var vpc vpc


cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	Vpc: Vpc,
})

autoScalingGroup := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
	Vpc: Vpc,
	InstanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
	MachineImage: ecs.EcsOptimizedImage_AmazonLinux2(),
	MinCapacity: jsii.Number(0),
	MaxCapacity: jsii.Number(100),
})

capacityProvider := ecs.NewAsgCapacityProvider(this, jsii.String("AsgCapacityProvider"), &AsgCapacityProviderProps{
	AutoScalingGroup: AutoScalingGroup,
	InstanceWarmupPeriod: jsii.Number(300),
})
cluster.AddAsgCapacityProvider(capacityProvider)

taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))

taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	MemoryReservationMiB: jsii.Number(256),
})

ecs.NewEc2Service(this, jsii.String("EC2Service"), &Ec2ServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	CapacityProviderStrategies: []capacityProviderStrategy{
		&capacityProviderStrategy{
			CapacityProvider: capacityProvider.CapacityProviderName,
			Weight: jsii.Number(1),
		},
	},
})
```

### Cluster Default Provider Strategy

A capacity provider strategy determines whether ECS tasks are launched on EC2 instances or Fargate/Fargate Spot. It can be specified at the cluster, service, or task level, and consists of one or more capacity providers. You can specify an optional base and weight value for finer control of how tasks are launched. The `base` specifies a minimum number of tasks on one capacity provider, and the `weight`s of each capacity provider determine how tasks are distributed after `base` is satisfied.

You can associate a default capacity provider strategy with an Amazon ECS cluster. After you do this, a default capacity provider strategy is used when creating a service or running a standalone task in the cluster and whenever a custom capacity provider strategy or a launch type isn't specified. We recommend that you define a default capacity provider strategy for each cluster.

For more information visit https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html

When the service does not have a capacity provider strategy, the cluster's default capacity provider strategy will be used. Default Capacity Provider Strategy can be added by using the method `addDefaultCapacityProviderStrategy`. A capacity provider strategy cannot contain a mix of EC2 Autoscaling Group capacity providers and Fargate providers.

```go
var capacityProvider asgCapacityProvider


cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
	EnableFargateCapacityProviders: jsii.Boolean(true),
})
cluster.AddAsgCapacityProvider(capacityProvider)

cluster.AddDefaultCapacityProviderStrategy([]capacityProviderStrategy{
	&capacityProviderStrategy{
		CapacityProvider: jsii.String("FARGATE"),
		Base: jsii.Number(10),
		Weight: jsii.Number(50),
	},
	&capacityProviderStrategy{
		CapacityProvider: jsii.String("FARGATE_SPOT"),
		Weight: jsii.Number(50),
	},
})
```

```go
var capacityProvider asgCapacityProvider


cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
	EnableFargateCapacityProviders: jsii.Boolean(true),
})
cluster.AddAsgCapacityProvider(capacityProvider)

cluster.AddDefaultCapacityProviderStrategy([]capacityProviderStrategy{
	&capacityProviderStrategy{
		CapacityProvider: capacityProvider.CapacityProviderName,
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

taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("Ec2TaskDef"), &Ec2TaskDefinitionProps{
	InferenceAccelerators: InferenceAccelerators,
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

taskDefinition.AddContainer(jsii.String("cont"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("test")),
	MemoryLimitMiB: jsii.Number(1024),
	InferenceAcceleratorResources: InferenceAcceleratorResources,
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


service := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	EnableExecuteCommand: jsii.Boolean(true),
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
logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &LogGroupProps{
	EncryptionKey: kmsKey,
})

// Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &BucketProps{
	EncryptionKey: kmsKey,
})

cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
	Vpc: Vpc,
	ExecuteCommandConfiguration: &ExecuteCommandConfiguration{
		KmsKey: *KmsKey,
		LogConfiguration: &ExecuteCommandLogConfiguration{
			CloudWatchLogGroup: logGroup,
			CloudWatchEncryptionEnabled: jsii.Boolean(true),
			S3Bucket: execBucket,
			S3EncryptionEnabled: jsii.Boolean(true),
			S3KeyPrefix: jsii.String("exec-command-output"),
		},
		Logging: ecs.ExecuteCommandLogging_OVERRIDE,
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
var cluster cluster
var taskDefinition taskDefinition
var containerOptions containerDefinitionOptions


container := taskDefinition.AddContainer(jsii.String("MyContainer"), containerOptions)

container.AddPortMappings(&PortMapping{
	Name: jsii.String("api"),
	ContainerPort: jsii.Number(8080),
})

cluster.AddDefaultCloudMapNamespace(&CloudMapNamespaceOptions{
	Name: jsii.String("local"),
})

service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	ServiceConnectConfiguration: &ServiceConnectProps{
		Services: []serviceConnectService{
			&serviceConnectService{
				PortMappingName: jsii.String("api"),
				DnsName: jsii.String("http-api"),
				Port: jsii.Number(80),
			},
		},
	},
})
```

Service Connect-enabled services may now reach this service at `http-api:80`. Traffic to this endpoint will
be routed to the container's port 8080.

To opt a service into using service connect without advertising a port, simply call the 'enableServiceConnect' method on an initialized service.

```go
var cluster cluster
var taskDefinition taskDefinition


service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})
service.EnableServiceConnect()
```

Service Connect also allows custom logging, Service Discovery name, and configuration of the port where service connect traffic is received.

```go
var cluster cluster
var taskDefinition taskDefinition


customService := ecs.NewFargateService(this, jsii.String("CustomizedService"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	ServiceConnectConfiguration: &ServiceConnectProps{
		LogDriver: ecs.LogDrivers_AwsLogs(&AwsLogDriverProps{
			StreamPrefix: jsii.String("sc-traffic"),
		}),
		Services: []serviceConnectService{
			&serviceConnectService{
				PortMappingName: jsii.String("api"),
				DnsName: jsii.String("customized-api"),
				Port: jsii.Number(80),
				IngressPortOverride: jsii.Number(20040),
				DiscoveryName: jsii.String("custom"),
			},
		},
	},
})
```

To set a timeout for service connect, use `idleTimeout` and `perRequestTimeout`.

**Note**: If `idleTimeout` is set to a time that is less than `perRequestTimeout`, the connection will close when
the `idleTimeout` is reached and not the `perRequestTimeout`.

```go
var cluster cluster
var taskDefinition taskDefinition


service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	ServiceConnectConfiguration: &ServiceConnectProps{
		Services: []serviceConnectService{
			&serviceConnectService{
				PortMappingName: jsii.String("api"),
				IdleTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
				PerRequestTimeout: awscdk.Duration_*Minutes(jsii.Number(5)),
			},
		},
	},
})
```

> Visit [Amazon ECS support for configurable timeout for services running with Service Connect](https://aws.amazon.com/about-aws/whats-new/2024/01/amazon-ecs-configurable-timeout-service-connect/) for more details.

## ServiceManagedVolume

Amazon ECS now supports the attachment of Amazon Elastic Block Store (EBS) volumes to ECS tasks,
allowing you to utilize persistent, high-performance block storage with your ECS services.
This feature supports various use cases, such as using EBS volumes as extended ephemeral storage or
loading data from EBS snapshots.
You can also specify `encrypted: true` so that ECS will manage the KMS key. If you want to use your own KMS key, you may do so by providing both `encrypted: true` and `kmsKeyId`.

You can only attach a single volume for each task in the ECS Service.

To add an empty EBS Volume to an ECS Service, call service.addVolume().

```go
var cluster cluster

taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"))

container := taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
	PortMappings: []portMapping{
		&portMapping{
			ContainerPort: jsii.Number(80),
			Protocol: ecs.Protocol_TCP,
		},
	},
})

volume := ecs.NewServiceManagedVolume(this, jsii.String("EBSVolume"), &ServiceManagedVolumeProps{
	Name: jsii.String("ebs1"),
	ManagedEBSVolume: &ServiceManagedEBSVolumeConfiguration{
		Size: awscdk.Size_Gibibytes(jsii.Number(15)),
		VolumeType: ec2.EbsDeviceVolumeType_GP3,
		FileSystemType: ecs.FileSystemType_XFS,
		TagSpecifications: []eBSTagSpecification{
			&eBSTagSpecification{
				Tags: map[string]*string{
					"purpose": jsii.String("production"),
				},
				PropagateTags: ecs.EbsPropagatedTagSource_SERVICE,
			},
		},
	},
})

volume.MountIn(container, &ContainerMountPoint{
	ContainerPath: jsii.String("/var/lib"),
	ReadOnly: jsii.Boolean(false),
})

taskDefinition.AddVolume(volume)

service := ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})

service.AddVolume(volume)
```

To create an EBS volume from an existing snapshot by specifying the `snapShotId` while adding a volume to the service.

```go
var container containerDefinition
var cluster cluster
var taskDefinition taskDefinition


volumeFromSnapshot := ecs.NewServiceManagedVolume(this, jsii.String("EBSVolume"), &ServiceManagedVolumeProps{
	Name: jsii.String("nginx-vol"),
	ManagedEBSVolume: &ServiceManagedEBSVolumeConfiguration{
		SnapShotId: jsii.String("snap-066877671789bd71b"),
		VolumeType: ec2.EbsDeviceVolumeType_GP3,
		FileSystemType: ecs.FileSystemType_XFS,
	},
})

volumeFromSnapshot.MountIn(container, &ContainerMountPoint{
	ContainerPath: jsii.String("/var/lib"),
	ReadOnly: jsii.Boolean(false),
})
taskDefinition.AddVolume(volumeFromSnapshot)
service := ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
})

service.AddVolume(volumeFromSnapshot)
```

## Enable pseudo-terminal (TTY) allocation

You can allocate a pseudo-terminal (TTY) for a container passing `pseudoTerminal` option while adding the container
to the task definition.
This maps to Tty option in the ["Create a container section"](https://docs.docker.com/engine/api/v1.38/#operation/ContainerCreate)
of the [Docker Remote API](https://docs.docker.com/engine/api/v1.38/) and the --tty option to [`docker run`](https://docs.docker.com/engine/reference/commandline/run/).

```go
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	PseudoTerminal: jsii.Boolean(true),
})
```

## Specify a container ulimit

You can specify a container `ulimits` by specifying them in the `ulimits` option while adding the container
to the task definition.

```go
taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
	Ulimits: []ulimit{
		&ulimit{
			HardLimit: jsii.Number(128),
			Name: ecs.UlimitName_RSS,
			SoftLimit: jsii.Number(128),
		},
	},
})
```
