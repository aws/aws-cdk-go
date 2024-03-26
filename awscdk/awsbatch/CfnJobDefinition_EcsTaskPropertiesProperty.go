package awsbatch


// The properties for a task definition that describes the container and volume definitions of an Amazon ECS task.
//
// You can specify which Docker images to use, the required resources, and other configurations related to launching the task definition through an Amazon ECS service or task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   ecsTaskPropertiesProperty := &EcsTaskPropertiesProperty{
//   	Containers: []interface{}{
//   		&TaskContainerPropertiesProperty{
//   			Image: jsii.String("image"),
//
//   			// the properties below are optional
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			DependsOn: []interface{}{
//   				&TaskContainerDependencyProperty{
//   					Condition: jsii.String("condition"),
//   					ContainerName: jsii.String("containerName"),
//   				},
//   			},
//   			Environment: []interface{}{
//   				&EnvironmentProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Essential: jsii.Boolean(false),
//   			LinuxParameters: &LinuxParametersProperty{
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
//   						ContainerPath: jsii.String("containerPath"),
//   						Size: jsii.Number(123),
//
//   						// the properties below are optional
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
//   				Options: options,
//   				SecretOptions: []interface{}{
//   					&SecretProperty{
//   						Name: jsii.String("name"),
//   						ValueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   			},
//   			MountPoints: []interface{}{
//   				&MountPointsProperty{
//   					ContainerPath: jsii.String("containerPath"),
//   					ReadOnly: jsii.Boolean(false),
//   					SourceVolume: jsii.String("sourceVolume"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Privileged: jsii.Boolean(false),
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
//   			Ulimits: []interface{}{
//   				&UlimitProperty{
//   					HardLimit: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					SoftLimit: jsii.Number(123),
//   				},
//   			},
//   			User: jsii.String("user"),
//   		},
//   	},
//   	EphemeralStorage: &EphemeralStorageProperty{
//   		SizeInGiB: jsii.Number(123),
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	IpcMode: jsii.String("ipcMode"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		AssignPublicIp: jsii.String("assignPublicIp"),
//   	},
//   	PidMode: jsii.String("pidMode"),
//   	PlatformVersion: jsii.String("platformVersion"),
//   	RuntimePlatform: &RuntimePlatformProperty{
//   		CpuArchitecture: jsii.String("cpuArchitecture"),
//   		OperatingSystemFamily: jsii.String("operatingSystemFamily"),
//   	},
//   	TaskRoleArn: jsii.String("taskRoleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html
//
type CfnJobDefinition_EcsTaskPropertiesProperty struct {
	// This object is a list of containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-containers
	//
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// The amount of ephemeral storage to allocate for the task.
	//
	// This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-ephemeralstorage
	//
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// The Amazon Resource Name (ARN) of the execution role that AWS Batch can assume.
	//
	// For jobs that run on Fargate resources, you must provide an execution role. For more information, see [AWS Batch execution IAM role](https://docs.aws.amazon.com/batch/latest/userguide/execution-IAM-role.html) in the *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// The valid values are `host` , `task` , or `none` .
	//
	// If `host` is specified, all containers within the tasks that specified the `host` IPC mode on the same container instance share the same IPC resources with the host Amazon EC2 instance.
	//
	// If `task` is specified, all containers within the specified `task` share the same IPC resources.
	//
	// If `none` is specified, the IPC resources within the containers of a task are private, and are not shared with other containers in a task or on the container instance.
	//
	// If no value is specified, then the IPC resource namespace sharing depends on the Docker daemon setting on the container instance. For more information, see [IPC settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#ipc-settings---ipc) in the Docker run reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-ipcmode
	//
	IpcMode *string `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// The network configuration for jobs that are running on Fargate resources.
	//
	// Jobs that are running on Amazon EC2 resources must not specify this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The process namespace to use for the containers in the task.
	//
	// The valid values are `host` or `task` . For example, monitoring sidecars might need `pidMode` to access information about other containers running in the same task.
	//
	// If `host` is specified, all containers within the tasks that specified the `host` PID mode on the same container instance share the process namespace with the host Amazon EC2 instance.
	//
	// If `task` is specified, all containers within the specified task share the same process namespace.
	//
	// If no value is specified, the default is a private namespace for each container. For more information, see [PID settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#pid-settings---pid) in the Docker run reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-pidmode
	//
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// The Fargate platform version where the jobs are running.
	//
	// A platform version is specified only for jobs that are running on Fargate resources. If one isn't specified, the `LATEST` platform version is used by default. This uses a recent, approved version of the Fargate platform for compute resources. For more information, see [AWS Fargate platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-platformversion
	//
	PlatformVersion *string `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// An object that represents the compute environment architecture for AWS Batch jobs on Fargate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-runtimeplatform
	//
	RuntimePlatform interface{} `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The Amazon Resource Name (ARN) that's associated with the Amazon ECS task.
	//
	// > This is object is comparable to [ContainerProperties:jobRoleArn](https://docs.aws.amazon.com/batch/latest/APIReference/API_ContainerProperties.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-taskrolearn
	//
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// A list of volumes that are associated with the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecstaskproperties.html#cfn-batch-jobdefinition-ecstaskproperties-volumes
	//
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

