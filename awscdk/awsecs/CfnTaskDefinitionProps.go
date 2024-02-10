package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTaskDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskDefinitionProps := &CfnTaskDefinitionProps{
//   	ContainerDefinitions: []interface{}{
//   		&ContainerDefinitionProperty{
//   			Image: jsii.String("image"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			CredentialSpecs: []*string{
//   				jsii.String("credentialSpecs"),
//   			},
//   			DependsOn: []interface{}{
//   				&ContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			DisableNetworking: jsii.Boolean(false),
//   			DnsSearchDomains: []*string{
//   				jsii.String("dnsSearchDomains"),
//   			},
//   			DnsServers: []*string{
//   				jsii.String("dnsServers"),
//   			},
//   			DockerLabels: map[string]*string{
//   				"dockerLabelsKey": jsii.String("dockerLabels"),
//   			},
//   			DockerSecurityOptions: []*string{
//   				jsii.String("dockerSecurityOptions"),
//   			},
//   			EntryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			Environment: []interface{}{
//   				&KeyValuePairProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			EnvironmentFiles: []interface{}{
//   				&EnvironmentFileProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Essential: jsii.Boolean(false),
//   			ExtraHosts: []interface{}{
//   				&HostEntryProperty{
//   					Hostname: jsii.String("hostname"),
//   					IpAddress: jsii.String("ipAddress"),
//   				},
//   			},
//   			FirelensConfiguration: &FirelensConfigurationProperty{
//   				Options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			HealthCheck: &HealthCheckProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Interval: jsii.Number(123),
//   				Retries: jsii.Number(123),
//   				StartPeriod: jsii.Number(123),
//   				Timeout: jsii.Number(123),
//   			},
//   			Hostname: jsii.String("hostname"),
//   			Interactive: jsii.Boolean(false),
//   			Links: []*string{
//   				jsii.String("links"),
//   			},
//   			LinuxParameters: &LinuxParametersProperty{
//   				Capabilities: &KernelCapabilitiesProperty{
//   					Add: []*string{
//   						jsii.String("add"),
//   					},
//   					Drop: []*string{
//   						jsii.String("drop"),
//   					},
//   				},
//   				Devices: []interface{}{
//   					&DeviceProperty{
//   						ContainerPath: jsii.String("containerPath"),
//   						HostPath: jsii.String("hostPath"),
//   						Permissions: []*string{
//   							jsii.String("permissions"),
//   						},
//   					},
//   				},
//   				InitProcessEnabled: jsii.Boolean(false),
//   				MaxSwap: jsii.Number(123),
//   				SharedMemorySize: jsii.Number(123),
//   				Swappiness: jsii.Number(123),
//   				Tmpfs: []interface{}{
//   					&TmpfsProperty{
//   						Size: jsii.Number(123),
//
//   						// the properties below are optional
//   						ContainerPath: jsii.String("containerPath"),
//   						MountOptions: []*string{
//   							jsii.String("mountOptions"),
//   						},
//   					},
//   				},
//   			},
//   			LogConfiguration: &LogConfigurationProperty{
//   				LogDriver: jsii.String("logDriver"),
//
//   				// the properties below are optional
//   				Options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				SecretOptions: []interface{}{
//   					&SecretProperty{
//   						Name: jsii.String("name"),
//   						ValueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   			Memory: jsii.Number(123),
//   			MemoryReservation: jsii.Number(123),
//   			MountPoints: []interface{}{
//   				&MountPointProperty{
//   					ContainerPath: jsii.String("containerPath"),
//   					ReadOnly: jsii.Boolean(false),
//   					SourceVolume: jsii.String("sourceVolume"),
//   				},
//   			},
//   			PortMappings: []interface{}{
//   				&PortMappingProperty{
//   					AppProtocol: jsii.String("appProtocol"),
//   					ContainerPort: jsii.Number(123),
//   					ContainerPortRange: jsii.String("containerPortRange"),
//   					HostPort: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					Protocol: jsii.String("protocol"),
//   				},
//   			},
//   			Privileged: jsii.Boolean(false),
//   			PseudoTerminal: jsii.Boolean(false),
//   			ReadonlyRootFilesystem: jsii.Boolean(false),
//   			RepositoryCredentials: &RepositoryCredentialsProperty{
//   				CredentialsParameter: jsii.String("credentialsParameter"),
//   			},
//   			ResourceRequirements: []interface{}{
//   				&ResourceRequirementProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Secrets: []interface{}{
//   				&SecretProperty{
//   					Name: jsii.String("name"),
//   					ValueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   			StartTimeout: jsii.Number(123),
//   			StopTimeout: jsii.Number(123),
//   			SystemControls: []interface{}{
//   				&SystemControlProperty{
//   					Namespace: jsii.String("namespace"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Ulimits: []interface{}{
//   				&UlimitProperty{
//   					HardLimit: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					SoftLimit: jsii.Number(123),
//   				},
//   			},
//   			User: jsii.String("user"),
//   			VolumesFrom: []interface{}{
//   				&VolumeFromProperty{
//   					ReadOnly: jsii.Boolean(false),
//   					SourceContainer: jsii.String("sourceContainer"),
//   				},
//   			},
//   			WorkingDirectory: jsii.String("workingDirectory"),
//   		},
//   	},
//   	Cpu: jsii.String("cpu"),
//   	EphemeralStorage: &EphemeralStorageProperty{
//   		SizeInGiB: jsii.Number(123),
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Family: jsii.String("family"),
//   	InferenceAccelerators: []interface{}{
//   		&InferenceAcceleratorProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			DeviceType: jsii.String("deviceType"),
//   		},
//   	},
//   	IpcMode: jsii.String("ipcMode"),
//   	Memory: jsii.String("memory"),
//   	NetworkMode: jsii.String("networkMode"),
//   	PidMode: jsii.String("pidMode"),
//   	PlacementConstraints: []interface{}{
//   		&TaskDefinitionPlacementConstraintProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   	ProxyConfiguration: &ProxyConfigurationProperty{
//   		ContainerName: jsii.String("containerName"),
//
//   		// the properties below are optional
//   		ProxyConfigurationProperties: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	RequiresCompatibilities: []*string{
//   		jsii.String("requiresCompatibilities"),
//   	},
//   	RuntimePlatform: &RuntimePlatformProperty{
//   		CpuArchitecture: jsii.String("cpuArchitecture"),
//   		OperatingSystemFamily: jsii.String("operatingSystemFamily"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   	Volumes: []interface{}{
//   		&VolumeProperty{
//   			ConfiguredAtLaunch: jsii.Boolean(false),
//   			DockerVolumeConfiguration: &DockerVolumeConfigurationProperty{
//   				Autoprovision: jsii.Boolean(false),
//   				Driver: jsii.String("driver"),
//   				DriverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				Labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   				Scope: jsii.String("scope"),
//   			},
//   			EfsVolumeConfiguration: &EFSVolumeConfigurationProperty{
//   				FilesystemId: jsii.String("filesystemId"),
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
//   			Host: &HostVolumePropertiesProperty{
//   				SourcePath: jsii.String("sourcePath"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html
//
type CfnTaskDefinitionProps struct {
	// A list of container definitions in JSON format that describe the different containers that make up your task.
	//
	// For more information about container definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-containerdefinitions
	//
	ContainerDefinitions interface{} `field:"optional" json:"containerDefinitions" yaml:"containerDefinitions"`
	// The number of `cpu` units used by the task.
	//
	// If you use the EC2 launch type, this field is optional. Any value can be used. If you use the Fargate launch type, this field is required. You must use one of the following values. The value that you choose determines your range of valid values for the `memory` parameter.
	//
	// The CPU units cannot be less than 1 vCPU when you use Windows containers on Fargate.
	//
	// - 256 (.25 vCPU) - Available `memory` values: 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB)
	// - 512 (.5 vCPU) - Available `memory` values: 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB)
	// - 1024 (1 vCPU) - Available `memory` values: 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	// - 2048 (2 vCPU) - Available `memory` values: 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)
	// - 4096 (4 vCPU) - Available `memory` values: 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB)
	// - 8192 (8 vCPU) - Available `memory` values: 16 GB and 60 GB in 4 GB increments
	//
	// This option requires Linux platform `1.4.0` or later.
	// - 16384 (16vCPU) - Available `memory` values: 32GB and 120 GB in 8 GB increments
	//
	// This option requires Linux platform `1.4.0` or later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-cpu
	//
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The ephemeral storage settings to use for tasks run with the task definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-ephemeralstorage
	//
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf.
	//
	// The task execution IAM role is required depending on the requirements of your task. For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of a family that this task definition is registered to.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	//
	// A family groups multiple versions of a task definition. Amazon ECS gives the first task definition that you registered to a family a revision number of 1. Amazon ECS gives sequential revision numbers to each task definition that you add.
	//
	// > To use revision numbers when you update a task definition, specify this property. If you don't specify a value, AWS CloudFormation generates a new task definition each time that you update it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-family
	//
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The Elastic Inference accelerators to use for the containers in the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-inferenceaccelerators
	//
	InferenceAccelerators interface{} `field:"optional" json:"inferenceAccelerators" yaml:"inferenceAccelerators"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// The valid values are `host` , `task` , or `none` . If `host` is specified, then all containers within the tasks that specified the `host` IPC mode on the same container instance share the same IPC resources with the host Amazon EC2 instance. If `task` is specified, all containers within the specified task share the same IPC resources. If `none` is specified, then IPC resources within the containers of a task are private and not shared with other containers in a task or on the container instance. If no value is specified, then the IPC resource namespace sharing depends on the Docker daemon setting on the container instance. For more information, see [IPC settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#ipc-settings---ipc) in the *Docker run reference* .
	//
	// If the `host` IPC mode is used, be aware that there is a heightened risk of undesired IPC namespace expose. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/) .
	//
	// If you are setting namespaced kernel parameters using `systemControls` for the containers in the task, the following will apply to your IPC resource namespace. For more information, see [System Controls](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// - For tasks that use the `host` IPC mode, IPC namespace related `systemControls` are not supported.
	// - For tasks that use the `task` IPC mode, IPC namespace related `systemControls` will apply to all containers within a task.
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-ipcmode
	//
	IpcMode *string `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// The amount (in MiB) of memory used by the task.
	//
	// If your tasks runs on Amazon EC2 instances, you must specify either a task-level memory value or a container-level memory value. This field is optional and any value can be used. If a task-level memory value is specified, the container-level memory value is optional. For more information regarding container-level memory and memory reservation, see [ContainerDefinition](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) .
	//
	// If your tasks runs on AWS Fargate , this field is required. You must use one of the following values. The value you choose determines your range of valid values for the `cpu` parameter.
	//
	// - 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available `cpu` values: 256 (.25 vCPU)
	// - 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available `cpu` values: 512 (.5 vCPU)
	// - 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available `cpu` values: 1024 (1 vCPU)
	// - Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available `cpu` values: 2048 (2 vCPU)
	// - Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available `cpu` values: 4096 (4 vCPU)
	// - Between 16 GB and 60 GB in 4 GB increments - Available `cpu` values: 8192 (8 vCPU)
	//
	// This option requires Linux platform `1.4.0` or later.
	// - Between 32GB and 120 GB in 8 GB increments - Available `cpu` values: 16384 (16 vCPU)
	//
	// This option requires Linux platform `1.4.0` or later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-memory
	//
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// The Docker networking mode to use for the containers in the task.
	//
	// The valid values are `none` , `bridge` , `awsvpc` , and `host` . If no network mode is specified, the default is `bridge` .
	//
	// For Amazon ECS tasks on Fargate, the `awsvpc` network mode is required. For Amazon ECS tasks on Amazon EC2 Linux instances, any network mode can be used. For Amazon ECS tasks on Amazon EC2 Windows instances, `<default>` or `awsvpc` can be used. If the network mode is set to `none` , you cannot specify port mappings in your container definitions, and the tasks containers do not have external connectivity. The `host` and `awsvpc` network modes offer the highest networking performance for containers because they use the EC2 network stack instead of the virtualized network stack provided by the `bridge` mode.
	//
	// With the `host` and `awsvpc` network modes, exposed container ports are mapped directly to the corresponding host port (for the `host` network mode) or the attached elastic network interface port (for the `awsvpc` network mode), so you cannot take advantage of dynamic host port mappings.
	//
	// > When using the `host` network mode, you should not run containers using the root user (UID 0). It is considered best practice to use a non-root user.
	//
	// If the network mode is `awsvpc` , the task is allocated an elastic network interface, and you must specify a `NetworkConfiguration` value when you create a service or run a task with the task definition. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// If the network mode is `host` , you cannot run multiple instantiations of the same task on a single container instance when port mappings are used.
	//
	// For more information, see [Network settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#network-settings) in the *Docker run reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-networkmode
	//
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// The valid values are `host` or `task` . On Fargate for Linux containers, the only valid value is `task` . For example, monitoring sidecars might need `pidMode` to access information about other containers running in the same task.
	//
	// If `host` is specified, all containers within the tasks that specified the `host` PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance.
	//
	// If `task` is specified, all containers within the specified task share the same process namespace.
	//
	// If no value is specified, the default is a private namespace for each container. For more information, see [PID settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#pid-settings---pid) in the *Docker run reference* .
	//
	// If the `host` PID mode is used, there's a heightened risk of undesired process namespace exposure. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/) .
	//
	// > This parameter is not supported for Windows containers. > This parameter is only supported for tasks that are hosted on AWS Fargate if the tasks are using platform version `1.4.0` or later (Linux). This isn't supported for Windows containers on Fargate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-pidmode
	//
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// An array of placement constraint objects to use for tasks.
	//
	// > This parameter isn't supported for tasks run on AWS Fargate .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-placementconstraints
	//
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The configuration details for the App Mesh proxy.
	//
	// Your Amazon ECS container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to use a proxy configuration. If your container instances are launched from the Amazon ECS optimized AMI version `20190301` or later, they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-proxyconfiguration
	//
	ProxyConfiguration interface{} `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The task launch types the task definition was validated against.
	//
	// The valid values are `EC2` , `FARGATE` , and `EXTERNAL` . For more information, see [Amazon ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-requirescompatibilities
	//
	RequiresCompatibilities *[]*string `field:"optional" json:"requiresCompatibilities" yaml:"requiresCompatibilities"`
	// The operating system that your tasks definitions run on.
	//
	// A platform family is specified only for tasks using the Fargate launch type.
	//
	// When you specify a task definition in a service, this value must match the `runtimePlatform` value of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-runtimeplatform
	//
	RuntimePlatform interface{} `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The metadata that you apply to the task definition to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value. You define both of them.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8
	// - Maximum value length - 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management role that grants containers in the task permission to call AWS APIs on your behalf.
	//
	// For more information, see [Amazon ECS Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// IAM roles for tasks on Windows require that the `-EnableTaskIAMRole` option is set when you launch the Amazon ECS-optimized Windows AMI. Your containers must also run some configuration code to use the feature. For more information, see [Windows IAM roles for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows_task_IAM_roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-taskrolearn
	//
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// The list of data volume definitions for the task.
	//
	// For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > The `host` and `sourcePath` parameters aren't supported for tasks run on AWS Fargate .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskdefinition.html#cfn-ecs-taskdefinition-volumes
	//
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

