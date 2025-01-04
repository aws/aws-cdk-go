package awsbatch


// Container properties are used for Amazon ECS based job definitions.
//
// These properties to describe the container that's launched as part of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   containerPropertiesProperty := &ContainerPropertiesProperty{
//   	Image: jsii.String("image"),
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Environment: []interface{}{
//   		&EnvironmentProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	EphemeralStorage: &EphemeralStorageProperty{
//   		SizeInGiB: jsii.Number(123),
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	FargatePlatformConfiguration: &FargatePlatformConfigurationProperty{
//   		PlatformVersion: jsii.String("platformVersion"),
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	JobRoleArn: jsii.String("jobRoleArn"),
//   	LinuxParameters: &LinuxParametersProperty{
//   		Devices: []interface{}{
//   			&DeviceProperty{
//   				ContainerPath: jsii.String("containerPath"),
//   				HostPath: jsii.String("hostPath"),
//   				Permissions: []*string{
//   					jsii.String("permissions"),
//   				},
//   			},
//   		},
//   		InitProcessEnabled: jsii.Boolean(false),
//   		MaxSwap: jsii.Number(123),
//   		SharedMemorySize: jsii.Number(123),
//   		Swappiness: jsii.Number(123),
//   		Tmpfs: []interface{}{
//   			&TmpfsProperty{
//   				ContainerPath: jsii.String("containerPath"),
//   				Size: jsii.Number(123),
//
//   				// the properties below are optional
//   				MountOptions: []*string{
//   					jsii.String("mountOptions"),
//   				},
//   			},
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		LogDriver: jsii.String("logDriver"),
//
//   		// the properties below are optional
//   		Options: options,
//   		SecretOptions: []interface{}{
//   			&SecretProperty{
//   				Name: jsii.String("name"),
//   				ValueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	Memory: jsii.Number(123),
//   	MountPoints: []interface{}{
//   		&MountPointsProperty{
//   			ContainerPath: jsii.String("containerPath"),
//   			ReadOnly: jsii.Boolean(false),
//   			SourceVolume: jsii.String("sourceVolume"),
//   		},
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		AssignPublicIp: jsii.String("assignPublicIp"),
//   	},
//   	Privileged: jsii.Boolean(false),
//   	ReadonlyRootFilesystem: jsii.Boolean(false),
//   	RepositoryCredentials: &RepositoryCredentialsProperty{
//   		CredentialsParameter: jsii.String("credentialsParameter"),
//   	},
//   	ResourceRequirements: []interface{}{
//   		&ResourceRequirementProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	RuntimePlatform: &RuntimePlatformProperty{
//   		CpuArchitecture: jsii.String("cpuArchitecture"),
//   		OperatingSystemFamily: jsii.String("operatingSystemFamily"),
//   	},
//   	Secrets: []interface{}{
//   		&SecretProperty{
//   			Name: jsii.String("name"),
//   			ValueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   	Ulimits: []interface{}{
//   		&UlimitProperty{
//   			HardLimit: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			SoftLimit: jsii.Number(123),
//   		},
//   	},
//   	User: jsii.String("user"),
//   	Vcpus: jsii.Number(123),
//   	Volumes: []interface{}{
//   		&VolumesProperty{
//   			EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				AuthorizationConfig: &AuthorizationConfigProperty{
//   					AccessPointId: jsii.String("accessPointId"),
//   					Iam: jsii.String("iam"),
//   				},
//   				RootDirectory: jsii.String("rootDirectory"),
//   				TransitEncryption: jsii.String("transitEncryption"),
//   				TransitEncryptionPort: jsii.Number(123),
//   			},
//   			Host: &VolumesHostProperty{
//   				SourcePath: jsii.String("sourcePath"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html
//
type CfnJobDefinition_ContainerPropertiesProperty struct {
	// Required.
	//
	// The image used to start a container. This string is passed directly to the Docker daemon. Images in the Docker Hub registry are available by default. Other repositories are specified with `*repository-url* / *image* : *tag*` . It can be 255 characters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), underscores (_), colons (:), periods (.), forward slashes (/), and number signs (#). This parameter maps to `Image` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `IMAGE` parameter of [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	//
	// > Docker image architecture must match the processor architecture of the compute resources that they're scheduled on. For example, ARM-based Docker images can only run on ARM-based compute resources.
	//
	// - Images in Amazon ECR Public repositories use the full `registry/repository[:tag]` or `registry/repository[@digest]` naming conventions. For example, `public.ecr.aws/ *registry_alias* / *my-web-app* : *latest*` .
	// - Images in Amazon ECR repositories use the full registry and repository URI (for example, `123456789012.dkr.ecr.<region-name>.amazonaws.com/<repository-name>` ).
	// - Images in official repositories on Docker Hub use a single name (for example, `ubuntu` or `mongo` ).
	// - Images in other repositories on Docker Hub are qualified with an organization name (for example, `amazon/amazon-ecs-agent` ).
	// - Images in other online repositories are qualified further by a domain name (for example, `quay.io/assemblyline/ubuntu` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-image
	//
	Image *string `field:"required" json:"image" yaml:"image"`
	// The command that's passed to the container.
	//
	// This parameter maps to `Cmd` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `COMMAND` parameter to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . For more information, see [https://docs.docker.com/engine/reference/builder/#cmd](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#cmd) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to pass to a container.
	//
	// This parameter maps to `Env` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--env` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	//
	// > We don't recommend using plaintext environment variables for sensitive information, such as credential data. > Environment variables cannot start with " `AWS_BATCH` ". This naming convention is reserved for variables that AWS Batch sets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The amount of ephemeral storage to allocate for the task.
	//
	// This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-ephemeralstorage
	//
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// The Amazon Resource Name (ARN) of the execution role that AWS Batch can assume.
	//
	// For jobs that run on Fargate resources, you must provide an execution role. For more information, see [AWS Batch execution IAM role](https://docs.aws.amazon.com/batch/latest/userguide/execution-IAM-role.html) in the *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The platform configuration for jobs that are running on Fargate resources.
	//
	// Jobs that are running on Amazon EC2 resources must not specify this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-fargateplatformconfiguration
	//
	FargatePlatformConfiguration interface{} `field:"optional" json:"fargatePlatformConfiguration" yaml:"fargatePlatformConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The Amazon Resource Name (ARN) of the IAM role that the container can assume for AWS permissions.
	//
	// For more information, see [IAM roles for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-jobrolearn
	//
	JobRoleArn *string `field:"optional" json:"jobRoleArn" yaml:"jobRoleArn"`
	// Linux-specific modifications that are applied to the container, such as details for device mappings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-linuxparameters
	//
	LinuxParameters interface{} `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	//
	// This parameter maps to `LogConfig` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--log-driver` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . By default, containers use the same logging driver that the Docker daemon uses. However the container might use a different logging driver than the Docker daemon by specifying a log driver with this parameter in the container definition. To use a different logging driver for a container, the log system must be configured properly on the container instance (or on a different log server for remote logging options). For more information on the options for different supported log drivers, see [Configure logging drivers](https://docs.aws.amazon.com/https://docs.docker.com/engine/admin/logging/overview/) in the Docker documentation.
	//
	// > AWS Batch currently supports a subset of the logging drivers available to the Docker daemon (shown in the [LogConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-logconfiguration.html) data type).
	//
	// This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version | grep "Server API version"`
	//
	// > The Amazon ECS container agent running on a container instance must register the logging drivers available on that instance with the `ECS_AVAILABLE_LOGGING_DRIVERS` environment variable before containers placed on that instance can use these log configuration options. For more information, see [Amazon ECS container agent configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// This parameter is deprecated, use `resourceRequirements` to specify the memory requirements for the job definition.
	//
	// It's not supported for jobs running on Fargate resources. For jobs that run on Amazon EC2 resources, it specifies the memory hard limit (in MiB) for a container. If your container attempts to exceed the specified number, it's terminated. You must specify at least 4 MiB of memory for a job using this parameter. The memory hard limit can be specified in several places. It must be specified for each node at least once.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-memory
	//
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The mount points for data volumes in your container.
	//
	// This parameter maps to `Volumes` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--volume` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-mountpoints
	//
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// The network configuration for jobs that are running on Fargate resources.
	//
	// Jobs that are running on Amazon EC2 resources must not specify this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// When this parameter is true, the container is given elevated permissions on the host container instance (similar to the `root` user).
	//
	// This parameter maps to `Privileged` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--privileged` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . The default value is false.
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided, or specified as false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-privileged
	//
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	//
	// This parameter maps to `ReadonlyRootfs` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--read-only` option to `docker run` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-readonlyrootfilesystem
	//
	ReadonlyRootFilesystem interface{} `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The private repository authentication credentials to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-repositorycredentials
	//
	RepositoryCredentials interface{} `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// The type and amount of resources to assign to a container.
	//
	// The supported resources include `GPU` , `MEMORY` , and `VCPU` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-resourcerequirements
	//
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
	// An object that represents the compute environment architecture for AWS Batch jobs on Fargate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-runtimeplatform
	//
	RuntimePlatform interface{} `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The secrets for the container.
	//
	// For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data.html) in the *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-secrets
	//
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// A list of `ulimits` to set in the container.
	//
	// This parameter maps to `Ulimits` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--ulimit` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	//
	// > This parameter isn't applicable to jobs that are running on Fargate resources and shouldn't be provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-ulimits
	//
	Ulimits interface{} `field:"optional" json:"ulimits" yaml:"ulimits"`
	// The user name to use inside the container.
	//
	// This parameter maps to `User` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--user` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-user
	//
	User *string `field:"optional" json:"user" yaml:"user"`
	// This parameter is deprecated, use `resourceRequirements` to specify the vCPU requirements for the job definition.
	//
	// It's not supported for jobs running on Fargate resources. For jobs running on Amazon EC2 resources, it specifies the number of vCPUs reserved for the job.
	//
	// Each vCPU is equivalent to 1,024 CPU shares. This parameter maps to `CpuShares` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/#create-a-container) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.23/) and the `--cpu-shares` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/) . The number of vCPUs must be specified but can be specified in several places. You must specify it at least once for each node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-vcpus
	//
	Vcpus *float64 `field:"optional" json:"vcpus" yaml:"vcpus"`
	// A list of data volumes used in a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-volumes
	//
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

