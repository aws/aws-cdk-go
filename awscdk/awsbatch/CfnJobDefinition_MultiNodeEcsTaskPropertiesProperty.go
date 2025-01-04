package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   multiNodeEcsTaskPropertiesProperty := &MultiNodeEcsTaskPropertiesProperty{
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
//   				&MountPointProperty{
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
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	IpcMode: jsii.String("ipcMode"),
//   	PidMode: jsii.String("pidMode"),
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   	Volumes: []interface{}{
//   		&VolumeProperty{
//   			EfsVolumeConfiguration: &EFSVolumeConfigurationProperty{
//   				FileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				AuthorizationConfig: &EFSAuthorizationConfigProperty{
//   					AccessPointId: jsii.String("accessPointId"),
//   					Iam: jsii.String("iam"),
//   				},
//   				RootDirectory: jsii.String("rootDirectory"),
//   				TransitEncryption: jsii.String("transitEncryption"),
//   				TransitEncryptionPort: jsii.Number(123),
//   			},
//   			Host: &HostProperty{
//   				SourcePath: jsii.String("sourcePath"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-multinodeecstaskproperties.html
//
type CfnJobDefinition_MultiNodeEcsTaskPropertiesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-multinodeecstaskproperties.html#cfn-batch-jobdefinition-multinodeecstaskproperties-containers
	//
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-multinodeecstaskproperties.html#cfn-batch-jobdefinition-multinodeecstaskproperties-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-multinodeecstaskproperties.html#cfn-batch-jobdefinition-multinodeecstaskproperties-ipcmode
	//
	IpcMode *string `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-multinodeecstaskproperties.html#cfn-batch-jobdefinition-multinodeecstaskproperties-pidmode
	//
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-multinodeecstaskproperties.html#cfn-batch-jobdefinition-multinodeecstaskproperties-taskrolearn
	//
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-multinodeecstaskproperties.html#cfn-batch-jobdefinition-multinodeecstaskproperties-volumes
	//
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

