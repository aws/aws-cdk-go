package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::ECS::TaskDefinition`.
//
// The `AWS::ECS::TaskDefinition` resource describes the container and volume definitions of a task. You can specify which Docker images to use, the required resources, and other configurations related to launching the task definition through an Amazon ECS service or task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskDefinition := awscdk.Aws_ecs.NewCfnTaskDefinition(this, jsii.String("MyCfnTaskDefinition"), &cfnTaskDefinitionProps{
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
//   					containerPortRange: jsii.String("containerPortRange"),
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
//   })
//
type CfnTaskDefinition interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrTaskDefinitionArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// A list of container definitions in JSON format that describe the different containers that make up your task.
	//
	// For more information about container definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide* .
	ContainerDefinitions() interface{}
	SetContainerDefinitions(val interface{})
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
	Cpu() *string
	SetCpu(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The ephemeral storage settings to use for tasks run with the task definition.
	EphemeralStorage() interface{}
	SetEphemeralStorage(val interface{})
	// The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf.
	//
	// The task execution IAM role is required depending on the requirements of your task. For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide* .
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// The name of a family that this task definition is registered to.
	//
	// Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	//
	// A family groups multiple versions of a task definition. Amazon ECS gives the first task definition that you registered to a family a revision number of 1. Amazon ECS gives sequential revision numbers to each task definition that you add.
	//
	// > To use revision numbers when you update a task definition, specify this property. If you don't specify a value, AWS CloudFormation generates a new task definition each time that you update it.
	Family() *string
	SetFamily(val *string)
	// The Elastic Inference accelerators to use for the containers in the task.
	InferenceAccelerators() interface{}
	SetInferenceAccelerators(val interface{})
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
	IpcMode() *string
	SetIpcMode(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
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
	Memory() *string
	SetMemory(val *string)
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
	NetworkMode() *string
	SetNetworkMode(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The process namespace to use for the containers in the task.
	//
	// The valid values are `host` or `task` . If `host` is specified, then all containers within the tasks that specified the `host` PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance. If `task` is specified, all containers within the specified task share the same process namespace. If no value is specified, the default is a private namespace. For more information, see [PID settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#pid-settings---pid) in the *Docker run reference* .
	//
	// If the `host` PID mode is used, be aware that there is a heightened risk of undesired process namespace expose. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/) .
	//
	// > This parameter is not supported for Windows containers or tasks run on AWS Fargate .
	PidMode() *string
	SetPidMode(val *string)
	// An array of placement constraint objects to use for tasks.
	//
	// > This parameter isn't supported for tasks run on AWS Fargate .
	PlacementConstraints() interface{}
	SetPlacementConstraints(val interface{})
	// The configuration details for the App Mesh proxy.
	//
	// Your Amazon ECS container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the `ecs-init` package to use a proxy configuration. If your container instances are launched from the Amazon ECS optimized AMI version `20190301` or later, they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
	ProxyConfiguration() interface{}
	SetProxyConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The task launch types the task definition was validated against.
	//
	// To determine which task launch types the task definition is validated for, see the `TaskDefinition$compatibilities` parameter.
	//
	// The valid values are:
	//
	// - `EC2` - The EC2 launch type
	// - `FARGATE` - The Fargate launch type
	// - `EXTERNAL` - The external instance (ECS Anywhere) launch type.
	RequiresCompatibilities() *[]*string
	SetRequiresCompatibilities(val *[]*string)
	// The operating system that your tasks definitions run on.
	//
	// A platform family is specified only for tasks using the Fargate launch type.
	//
	// When you specify a task definition in a service, this value must match the `runtimePlatform` value of the service.
	RuntimePlatform() interface{}
	SetRuntimePlatform(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
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
	Tags() awscdk.TagManager
	// The short name or full Amazon Resource Name (ARN) of the AWS Identity and Access Management role that grants containers in the task permission to call AWS APIs on your behalf.
	//
	// For more information, see [Amazon ECS Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// IAM roles for tasks on Windows require that the `-EnableTaskIAMRole` option is set when you launch the Amazon ECS-optimized Windows AMI. Your containers must also run some configuration code to use the feature. For more information, see [Windows IAM roles for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows_task_IAM_roles.html) in the *Amazon Elastic Container Service Developer Guide* .
	TaskRoleArn() *string
	SetTaskRoleArn(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The list of data volume definitions for the task.
	//
	// For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > The `host` and `sourcePath` parameters aren't supported for tasks run on AWS Fargate .
	Volumes() interface{}
	SetVolumes(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnTaskDefinition
type jsiiProxy_CfnTaskDefinition struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTaskDefinition) AttrTaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTaskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) ContainerDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containerDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) EphemeralStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) InferenceAccelerators() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) IpcMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipcMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) PidMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) PlacementConstraints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) ProxyConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) RequiresCompatibilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiresCompatibilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) RuntimePlatform() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskDefinition) Volumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Create a new `AWS::ECS::TaskDefinition`.
func NewCfnTaskDefinition(scope awscdk.Construct, id *string, props *CfnTaskDefinitionProps) CfnTaskDefinition {
	_init_.Initialize()

	if err := validateNewCfnTaskDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTaskDefinition{}

	_jsii_.Create(
		"monocdk.aws_ecs.CfnTaskDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ECS::TaskDefinition`.
func NewCfnTaskDefinition_Override(c CfnTaskDefinition, scope awscdk.Construct, id *string, props *CfnTaskDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.CfnTaskDefinition",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetContainerDefinitions(val interface{}) {
	if err := j.validateSetContainerDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetCpu(val *string) {
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetEphemeralStorage(val interface{}) {
	if err := j.validateSetEphemeralStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ephemeralStorage",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetFamily(val *string) {
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetInferenceAccelerators(val interface{}) {
	if err := j.validateSetInferenceAcceleratorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAccelerators",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetIpcMode(val *string) {
	_jsii_.Set(
		j,
		"ipcMode",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetMemory(val *string) {
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetNetworkMode(val *string) {
	_jsii_.Set(
		j,
		"networkMode",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetPidMode(val *string) {
	_jsii_.Set(
		j,
		"pidMode",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetPlacementConstraints(val interface{}) {
	if err := j.validateSetPlacementConstraintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementConstraints",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetProxyConfiguration(val interface{}) {
	if err := j.validateSetProxyConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetRequiresCompatibilities(val *[]*string) {
	_jsii_.Set(
		j,
		"requiresCompatibilities",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetRuntimePlatform(val interface{}) {
	if err := j.validateSetRuntimePlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimePlatform",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetTaskRoleArn(val *string) {
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTaskDefinition)SetVolumes(val interface{}) {
	if err := j.validateSetVolumesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumes",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnTaskDefinition_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTaskDefinition_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.CfnTaskDefinition",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTaskDefinition_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnTaskDefinition_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.CfnTaskDefinition",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTaskDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.CfnTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTaskDefinition_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ecs.CfnTaskDefinition",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTaskDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTaskDefinition) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTaskDefinition) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

