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
//   cfnTaskDefinitionProps := &cfnTaskDefinitionProps{
//   	containerDefinitions: []interface{}{
//   		&containerDefinitionProperty{
//   			image: jsii.String("image"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			cpu: jsii.Number(123),
//   			dependsOn: []interface{}{
//   				&containerDependencyProperty{
//   					condition: jsii.String("condition"),
//   					containerName: jsii.String("containerName"),
//   				},
//   			},
//   			disableNetworking: jsii.Boolean(false),
//   			dnsSearchDomains: []*string{
//   				jsii.String("dnsSearchDomains"),
//   			},
//   			dnsServers: []*string{
//   				jsii.String("dnsServers"),
//   			},
//   			dockerLabels: map[string]*string{
//   				"dockerLabelsKey": jsii.String("dockerLabels"),
//   			},
//   			dockerSecurityOptions: []*string{
//   				jsii.String("dockerSecurityOptions"),
//   			},
//   			entryPoint: []*string{
//   				jsii.String("entryPoint"),
//   			},
//   			environment: []interface{}{
//   				&keyValuePairProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			environmentFiles: []interface{}{
//   				&environmentFileProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			essential: jsii.Boolean(false),
//   			extraHosts: []interface{}{
//   				&hostEntryProperty{
//   					hostname: jsii.String("hostname"),
//   					ipAddress: jsii.String("ipAddress"),
//   				},
//   			},
//   			firelensConfiguration: &firelensConfigurationProperty{
//   				options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				type: jsii.String("type"),
//   			},
//   			healthCheck: &healthCheckProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				interval: jsii.Number(123),
//   				retries: jsii.Number(123),
//   				startPeriod: jsii.Number(123),
//   				timeout: jsii.Number(123),
//   			},
//   			hostname: jsii.String("hostname"),
//   			interactive: jsii.Boolean(false),
//   			links: []*string{
//   				jsii.String("links"),
//   			},
//   			linuxParameters: &linuxParametersProperty{
//   				capabilities: &kernelCapabilitiesProperty{
//   					add: []*string{
//   						jsii.String("add"),
//   					},
//   					drop: []*string{
//   						jsii.String("drop"),
//   					},
//   				},
//   				devices: []interface{}{
//   					&deviceProperty{
//   						containerPath: jsii.String("containerPath"),
//   						hostPath: jsii.String("hostPath"),
//   						permissions: []*string{
//   							jsii.String("permissions"),
//   						},
//   					},
//   				},
//   				initProcessEnabled: jsii.Boolean(false),
//   				maxSwap: jsii.Number(123),
//   				sharedMemorySize: jsii.Number(123),
//   				swappiness: jsii.Number(123),
//   				tmpfs: []interface{}{
//   					&tmpfsProperty{
//   						size: jsii.Number(123),
//
//   						// the properties below are optional
//   						containerPath: jsii.String("containerPath"),
//   						mountOptions: []*string{
//   							jsii.String("mountOptions"),
//   						},
//   					},
//   				},
//   			},
//   			logConfiguration: &logConfigurationProperty{
//   				logDriver: jsii.String("logDriver"),
//
//   				// the properties below are optional
//   				options: map[string]*string{
//   					"optionsKey": jsii.String("options"),
//   				},
//   				secretOptions: []interface{}{
//   					&secretProperty{
//   						name: jsii.String("name"),
//   						valueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   			memory: jsii.Number(123),
//   			memoryReservation: jsii.Number(123),
//   			mountPoints: []interface{}{
//   				&mountPointProperty{
//   					containerPath: jsii.String("containerPath"),
//   					readOnly: jsii.Boolean(false),
//   					sourceVolume: jsii.String("sourceVolume"),
//   				},
//   			},
//   			portMappings: []interface{}{
//   				&portMappingProperty{
//   					appProtocol: jsii.String("appProtocol"),
//   					containerPort: jsii.Number(123),
//   					hostPort: jsii.Number(123),
//   					name: jsii.String("name"),
//   					protocol: jsii.String("protocol"),
//   				},
//   			},
//   			privileged: jsii.Boolean(false),
//   			pseudoTerminal: jsii.Boolean(false),
//   			readonlyRootFilesystem: jsii.Boolean(false),
//   			repositoryCredentials: &repositoryCredentialsProperty{
//   				credentialsParameter: jsii.String("credentialsParameter"),
//   			},
//   			resourceRequirements: []interface{}{
//   				&resourceRequirementProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			secrets: []interface{}{
//   				&secretProperty{
//   					name: jsii.String("name"),
//   					valueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   			startTimeout: jsii.Number(123),
//   			stopTimeout: jsii.Number(123),
//   			systemControls: []interface{}{
//   				&systemControlProperty{
//   					namespace: jsii.String("namespace"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			ulimits: []interface{}{
//   				&ulimitProperty{
//   					hardLimit: jsii.Number(123),
//   					name: jsii.String("name"),
//   					softLimit: jsii.Number(123),
//   				},
//   			},
//   			user: jsii.String("user"),
//   			volumesFrom: []interface{}{
//   				&volumeFromProperty{
//   					readOnly: jsii.Boolean(false),
//   					sourceContainer: jsii.String("sourceContainer"),
//   				},
//   			},
//   			workingDirectory: jsii.String("workingDirectory"),
//   		},
//   	},
//   	cpu: jsii.String("cpu"),
//   	ephemeralStorage: &ephemeralStorageProperty{
//   		sizeInGiB: jsii.Number(123),
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	family: jsii.String("family"),
//   	inferenceAccelerators: []interface{}{
//   		&inferenceAcceleratorProperty{
//   			deviceName: jsii.String("deviceName"),
//   			deviceType: jsii.String("deviceType"),
//   		},
//   	},
//   	ipcMode: jsii.String("ipcMode"),
//   	memory: jsii.String("memory"),
//   	networkMode: jsii.String("networkMode"),
//   	pidMode: jsii.String("pidMode"),
//   	placementConstraints: []interface{}{
//   		&taskDefinitionPlacementConstraintProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			expression: jsii.String("expression"),
//   		},
//   	},
//   	proxyConfiguration: &proxyConfigurationProperty{
//   		containerName: jsii.String("containerName"),
//
//   		// the properties below are optional
//   		proxyConfigurationProperties: []interface{}{
//   			&keyValuePairProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		type: jsii.String("type"),
//   	},
//   	requiresCompatibilities: []*string{
//   		jsii.String("requiresCompatibilities"),
//   	},
//   	runtimePlatform: &runtimePlatformProperty{
//   		cpuArchitecture: jsii.String("cpuArchitecture"),
//   		operatingSystemFamily: jsii.String("operatingSystemFamily"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskRoleArn: jsii.String("taskRoleArn"),
//   	volumes: []interface{}{
//   		&volumeProperty{
//   			dockerVolumeConfiguration: &dockerVolumeConfigurationProperty{
//   				autoprovision: jsii.Boolean(false),
//   				driver: jsii.String("driver"),
//   				driverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   				scope: jsii.String("scope"),
//   			},
//   			efsVolumeConfiguration: &eFSVolumeConfigurationProperty{
//   				filesystemId: jsii.String("filesystemId"),
//
//   				// the properties below are optional
//   				authorizationConfig: &authorizationConfigProperty{
//   					accessPointId: jsii.String("accessPointId"),
//   					iam: jsii.String("iam"),
//   				},
//   				rootDirectory: jsii.String("rootDirectory"),
//   				transitEncryption: jsii.String("transitEncryption"),
//   				transitEncryptionPort: jsii.Number(123),
//   			},
//   			host: &hostVolumePropertiesProperty{
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnTaskDefinitionProps struct {
	// A list of container definitions in JSON format that describe the different containers that make up your task.
	//
	// For more information about container definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide* .
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
	// - 2048 (2 vCPU) - Available `memory` values: Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)
	// - 4096 (4 vCPU) - Available `memory` values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB).
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The ephemeral storage settings to use for tasks run with the task definition.
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf.
	//
	// The task execution IAM role is required depending on the requirements of your task. For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide* .
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of a family that this task definition is registered to.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	//
	// A family groups multiple versions of a task definition. Amazon ECS gives the first task definition that you registered to a family a revision number of 1. Amazon ECS gives sequential revision numbers to each task definition that you add.
	//
	// > To use revision numbers when you update a task definition, specify this property. If you don't specify a value, AWS CloudFormation generates a new task definition each time that you update it.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The Elastic Inference accelerators to use for the containers in the task.
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
	// - Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available `cpu` values: 4096 (4 vCPU).
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// The Docker networking mode to use for the containers in the task.
	//
	// The valid values are `none` , `bridge` , `awsvpc` , and `host` . The default Docker network mode is `bridge` . If you are using the Fargate launch type, the `awsvpc` network mode is required. If you are using the EC2 launch type, any network mode can be used. If the network mode is set to `none` , you cannot specify port mappings in your container definitions, and the tasks containers do not have external connectivity. The `host` and `awsvpc` network modes offer the highest networking performance for containers because they use the EC2 network stack instead of the virtualized network stack provided by the `bridge` mode.
	//
	// With the `host` and `awsvpc` network modes, exposed container ports are mapped directly to the corresponding host port (for the `host` network mode) or the attached elastic network interface port (for the `awsvpc` network mode), so you cannot take advantage of dynamic host port mappings.
	//
	// If the network mode is `awsvpc` , the task is allocated an elastic network interface, and you must specify a [NetworkConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_NetworkConfiguration.html) value when you create a service or run a task with the task definition. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > Currently, only Amazon ECS-optimized AMIs, other Amazon Linux variants with the `ecs-init` package, or AWS Fargate infrastructure support the `awsvpc` network mode.
	//
	// If the network mode is `host` , you cannot run multiple instantiations of the same task on a single container instance when port mappings are used.
	//
	// Docker for Windows uses different network modes than Docker for Linux. When you register a task definition with Windows containers, you must not specify a network mode. If you use the console to register a task definition with Windows containers, you must choose the `<default>` network mode object.
	//
	// For more information, see [Network settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#network-settings) in the *Docker run reference* .
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// The valid values are `host` or `task` . If `host` is specified, then all containers within the tasks that specified the `host` PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance. If `task` is specified, all containers within the specified task share the same process namespace. If no value is specified, the default is a private namespace. For more information, see [PID settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#pid-settings---pid) in the *Docker run reference* .
	//
	// If the `host` PID mode is used, be aware that there is a heightened risk of undesired process namespace expose. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/) .
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// An array of placement constraint objects to use for tasks.
	//
	// > This parameter isn't supported for tasks run on AWS Fargate .
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The `ProxyConfiguration` property specifies the configuration details for the App Mesh proxy.
	//
	// Your Amazon ECS container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to enable a proxy configuration. If your container instances are launched from the Amazon ECS-optimized AMI version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	ProxyConfiguration interface{} `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The task launch types the task definition was validated against.
	//
	// To determine which task launch types the task definition is validated for, see the `TaskDefinition$compatibilities` parameter.
	RequiresCompatibilities *[]*string `field:"optional" json:"requiresCompatibilities" yaml:"requiresCompatibilities"`
	// The operating system that your tasks definitions run on.
	//
	// A platform family is specified only for tasks using the Fargate launch type.
	//
	// When you specify a task definition in a service, this value must match the `runtimePlatform` value of the service.
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
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	// - Tag keys and values are case-sensitive.
	// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management role that grants containers in the task permission to call AWS APIs on your behalf.
	//
	// For more information, see [Amazon ECS Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// IAM roles for tasks on Windows require that the `-EnableTaskIAMRole` option is set when you launch the Amazon ECS-optimized Windows AMI. Your containers must also run some configuration code to use the feature. For more information, see [Windows IAM roles for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows_task_IAM_roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// The list of data volume definitions for the task.
	//
	// For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > The `host` and `sourcePath` parameters aren't supported for tasks run on AWS Fargate .
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

