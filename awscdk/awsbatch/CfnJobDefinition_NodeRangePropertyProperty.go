package awsbatch


// This is an object that represents the properties of the node range for a multi-node parallel job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   nodeRangePropertyProperty := &NodeRangePropertyProperty{
//   	TargetNodes: jsii.String("targetNodes"),
//
//   	// the properties below are optional
//   	Container: &ContainerPropertiesProperty{
//   		Image: jsii.String("image"),
//
//   		// the properties below are optional
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Environment: []interface{}{
//   			&EnvironmentProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		EphemeralStorage: &EphemeralStorageProperty{
//   			SizeInGiB: jsii.Number(123),
//   		},
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		FargatePlatformConfiguration: &FargatePlatformConfigurationProperty{
//   			PlatformVersion: jsii.String("platformVersion"),
//   		},
//   		InstanceType: jsii.String("instanceType"),
//   		JobRoleArn: jsii.String("jobRoleArn"),
//   		LinuxParameters: &LinuxParametersProperty{
//   			Devices: []interface{}{
//   				&DeviceProperty{
//   					ContainerPath: jsii.String("containerPath"),
//   					HostPath: jsii.String("hostPath"),
//   					Permissions: []*string{
//   						jsii.String("permissions"),
//   					},
//   				},
//   			},
//   			InitProcessEnabled: jsii.Boolean(false),
//   			MaxSwap: jsii.Number(123),
//   			SharedMemorySize: jsii.Number(123),
//   			Swappiness: jsii.Number(123),
//   			Tmpfs: []interface{}{
//   				&TmpfsProperty{
//   					ContainerPath: jsii.String("containerPath"),
//   					Size: jsii.Number(123),
//
//   					// the properties below are optional
//   					MountOptions: []*string{
//   						jsii.String("mountOptions"),
//   					},
//   				},
//   			},
//   		},
//   		LogConfiguration: &LogConfigurationProperty{
//   			LogDriver: jsii.String("logDriver"),
//
//   			// the properties below are optional
//   			Options: options,
//   			SecretOptions: []interface{}{
//   				&SecretProperty{
//   					Name: jsii.String("name"),
//   					ValueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		Memory: jsii.Number(123),
//   		MountPoints: []interface{}{
//   			&MountPointsProperty{
//   				ContainerPath: jsii.String("containerPath"),
//   				ReadOnly: jsii.Boolean(false),
//   				SourceVolume: jsii.String("sourceVolume"),
//   			},
//   		},
//   		NetworkConfiguration: &NetworkConfigurationProperty{
//   			AssignPublicIp: jsii.String("assignPublicIp"),
//   		},
//   		Privileged: jsii.Boolean(false),
//   		ReadonlyRootFilesystem: jsii.Boolean(false),
//   		RepositoryCredentials: &RepositoryCredentialsProperty{
//   			CredentialsParameter: jsii.String("credentialsParameter"),
//   		},
//   		ResourceRequirements: []interface{}{
//   			&ResourceRequirementProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RuntimePlatform: &RuntimePlatformProperty{
//   			CpuArchitecture: jsii.String("cpuArchitecture"),
//   			OperatingSystemFamily: jsii.String("operatingSystemFamily"),
//   		},
//   		Secrets: []interface{}{
//   			&SecretProperty{
//   				Name: jsii.String("name"),
//   				ValueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   		Ulimits: []interface{}{
//   			&UlimitProperty{
//   				HardLimit: jsii.Number(123),
//   				Name: jsii.String("name"),
//   				SoftLimit: jsii.Number(123),
//   			},
//   		},
//   		User: jsii.String("user"),
//   		Vcpus: jsii.Number(123),
//   		Volumes: []interface{}{
//   			&VolumesProperty{
//   				EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   					FileSystemId: jsii.String("fileSystemId"),
//
//   					// the properties below are optional
//   					AuthorizationConfig: &AuthorizationConfigProperty{
//   						AccessPointId: jsii.String("accessPointId"),
//   						Iam: jsii.String("iam"),
//   					},
//   					RootDirectory: jsii.String("rootDirectory"),
//   					TransitEncryption: jsii.String("transitEncryption"),
//   					TransitEncryptionPort: jsii.Number(123),
//   				},
//   				Host: &VolumesHostProperty{
//   					SourcePath: jsii.String("sourcePath"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html
//
type CfnJobDefinition_NodeRangePropertyProperty struct {
	// The range of nodes, using node index values.
	//
	// A range of `0:3` indicates nodes with index values of `0` through `3` . If the starting range value is omitted ( `:n` ), then `0` is used to start the range. If the ending range value is omitted ( `n:` ), then the highest possible node index is used to end the range. Your accumulative node ranges must account for all nodes ( `0:n` ). You can nest node ranges (for example, `0:10` and `4:5` ). In this case, the `4:5` range properties override the `0:10` properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-targetnodes
	//
	TargetNodes *string `field:"required" json:"targetNodes" yaml:"targetNodes"`
	// The container details for the node range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-container
	//
	Container interface{} `field:"optional" json:"container" yaml:"container"`
}

