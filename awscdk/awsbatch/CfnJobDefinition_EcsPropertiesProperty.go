package awsbatch


// An object that contains the properties for the Amazon ECS resources of a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   ecsPropertiesProperty := &EcsPropertiesProperty{
//   	TaskProperties: []interface{}{
//   		&EcsTaskPropertiesProperty{
//   			Containers: []interface{}{
//   				&TaskContainerPropertiesProperty{
//   					Image: jsii.String("image"),
//
//   					// the properties below are optional
//   					Command: []*string{
//   						jsii.String("command"),
//   					},
//   					DependsOn: []interface{}{
//   						&TaskContainerDependencyProperty{
//   							Condition: jsii.String("condition"),
//   							ContainerName: jsii.String("containerName"),
//   						},
//   					},
//   					Environment: []interface{}{
//   						&EnvironmentProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Essential: jsii.Boolean(false),
//   					LinuxParameters: &LinuxParametersProperty{
//   						Devices: []interface{}{
//   							&DeviceProperty{
//   								ContainerPath: jsii.String("containerPath"),
//   								HostPath: jsii.String("hostPath"),
//   								Permissions: []*string{
//   									jsii.String("permissions"),
//   								},
//   							},
//   						},
//   						InitProcessEnabled: jsii.Boolean(false),
//   						MaxSwap: jsii.Number(123),
//   						SharedMemorySize: jsii.Number(123),
//   						Swappiness: jsii.Number(123),
//   						Tmpfs: []interface{}{
//   							&TmpfsProperty{
//   								ContainerPath: jsii.String("containerPath"),
//   								Size: jsii.Number(123),
//
//   								// the properties below are optional
//   								MountOptions: []*string{
//   									jsii.String("mountOptions"),
//   								},
//   							},
//   						},
//   					},
//   					LogConfiguration: &LogConfigurationProperty{
//   						LogDriver: jsii.String("logDriver"),
//
//   						// the properties below are optional
//   						Options: options,
//   						SecretOptions: []interface{}{
//   							&SecretProperty{
//   								Name: jsii.String("name"),
//   								ValueFrom: jsii.String("valueFrom"),
//   							},
//   						},
//   					},
//   					MountPoints: []interface{}{
//   						&MountPointProperty{
//   							ContainerPath: jsii.String("containerPath"),
//   							ReadOnly: jsii.Boolean(false),
//   							SourceVolume: jsii.String("sourceVolume"),
//   						},
//   					},
//   					Name: jsii.String("name"),
//   					Privileged: jsii.Boolean(false),
//   					ReadonlyRootFilesystem: jsii.Boolean(false),
//   					RepositoryCredentials: &RepositoryCredentialsProperty{
//   						CredentialsParameter: jsii.String("credentialsParameter"),
//   					},
//   					ResourceRequirements: []interface{}{
//   						&ResourceRequirementProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					Secrets: []interface{}{
//   						&SecretProperty{
//   							Name: jsii.String("name"),
//   							ValueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   					Ulimits: []interface{}{
//   						&UlimitProperty{
//   							HardLimit: jsii.Number(123),
//   							Name: jsii.String("name"),
//   							SoftLimit: jsii.Number(123),
//   						},
//   					},
//   					User: jsii.String("user"),
//   				},
//   			},
//   			EphemeralStorage: &EphemeralStorageProperty{
//   				SizeInGiB: jsii.Number(123),
//   			},
//   			ExecutionRoleArn: jsii.String("executionRoleArn"),
//   			IpcMode: jsii.String("ipcMode"),
//   			NetworkConfiguration: &NetworkConfigurationProperty{
//   				AssignPublicIp: jsii.String("assignPublicIp"),
//   			},
//   			PidMode: jsii.String("pidMode"),
//   			PlatformVersion: jsii.String("platformVersion"),
//   			RuntimePlatform: &RuntimePlatformProperty{
//   				CpuArchitecture: jsii.String("cpuArchitecture"),
//   				OperatingSystemFamily: jsii.String("operatingSystemFamily"),
//   			},
//   			TaskRoleArn: jsii.String("taskRoleArn"),
//   			Volumes: []interface{}{
//   				&VolumesProperty{
//   					EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   						FileSystemId: jsii.String("fileSystemId"),
//
//   						// the properties below are optional
//   						AuthorizationConfig: &AuthorizationConfigProperty{
//   							AccessPointId: jsii.String("accessPointId"),
//   							Iam: jsii.String("iam"),
//   						},
//   						RootDirectory: jsii.String("rootDirectory"),
//   						TransitEncryption: jsii.String("transitEncryption"),
//   						TransitEncryptionPort: jsii.Number(123),
//   					},
//   					Host: &VolumesHostProperty{
//   						SourcePath: jsii.String("sourcePath"),
//   					},
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecsproperties.html
//
type CfnJobDefinition_EcsPropertiesProperty struct {
	// An object that contains the properties for the Amazon ECS task definition of a job.
	//
	// > This object is currently limited to one task element. However, the task element can run up to 10 containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ecsproperties.html#cfn-batch-jobdefinition-ecsproperties-taskproperties
	//
	TaskProperties interface{} `field:"required" json:"taskProperties" yaml:"taskProperties"`
}

