package awsbatch


// An object that represents the node properties of a multi-node parallel job.
//
// > Node properties can't be specified for Amazon EKS based job definitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   nodePropertiesProperty := &NodePropertiesProperty{
//   	MainNode: jsii.Number(123),
//   	NodeRangeProperties: []interface{}{
//   		&NodeRangePropertyProperty{
//   			TargetNodes: jsii.String("targetNodes"),
//
//   			// the properties below are optional
//   			Container: &ContainerPropertiesProperty{
//   				Image: jsii.String("image"),
//
//   				// the properties below are optional
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Environment: []interface{}{
//   					&EnvironmentProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				EphemeralStorage: &EphemeralStorageProperty{
//   					SizeInGiB: jsii.Number(123),
//   				},
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				FargatePlatformConfiguration: &FargatePlatformConfigurationProperty{
//   					PlatformVersion: jsii.String("platformVersion"),
//   				},
//   				InstanceType: jsii.String("instanceType"),
//   				JobRoleArn: jsii.String("jobRoleArn"),
//   				LinuxParameters: &LinuxParametersProperty{
//   					Devices: []interface{}{
//   						&DeviceProperty{
//   							ContainerPath: jsii.String("containerPath"),
//   							HostPath: jsii.String("hostPath"),
//   							Permissions: []*string{
//   								jsii.String("permissions"),
//   							},
//   						},
//   					},
//   					InitProcessEnabled: jsii.Boolean(false),
//   					MaxSwap: jsii.Number(123),
//   					SharedMemorySize: jsii.Number(123),
//   					Swappiness: jsii.Number(123),
//   					Tmpfs: []interface{}{
//   						&TmpfsProperty{
//   							ContainerPath: jsii.String("containerPath"),
//   							Size: jsii.Number(123),
//
//   							// the properties below are optional
//   							MountOptions: []*string{
//   								jsii.String("mountOptions"),
//   							},
//   						},
//   					},
//   				},
//   				LogConfiguration: &LogConfigurationProperty{
//   					LogDriver: jsii.String("logDriver"),
//
//   					// the properties below are optional
//   					Options: options,
//   					SecretOptions: []interface{}{
//   						&SecretProperty{
//   							Name: jsii.String("name"),
//   							ValueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   				},
//   				Memory: jsii.Number(123),
//   				MountPoints: []interface{}{
//   					&MountPointsProperty{
//   						ContainerPath: jsii.String("containerPath"),
//   						ReadOnly: jsii.Boolean(false),
//   						SourceVolume: jsii.String("sourceVolume"),
//   					},
//   				},
//   				NetworkConfiguration: &NetworkConfigurationProperty{
//   					AssignPublicIp: jsii.String("assignPublicIp"),
//   				},
//   				Privileged: jsii.Boolean(false),
//   				ReadonlyRootFilesystem: jsii.Boolean(false),
//   				ResourceRequirements: []interface{}{
//   					&ResourceRequirementProperty{
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				RuntimePlatform: &RuntimePlatformProperty{
//   					CpuArchitecture: jsii.String("cpuArchitecture"),
//   					OperatingSystemFamily: jsii.String("operatingSystemFamily"),
//   				},
//   				Secrets: []interface{}{
//   					&SecretProperty{
//   						Name: jsii.String("name"),
//   						ValueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   				Ulimits: []interface{}{
//   					&UlimitProperty{
//   						HardLimit: jsii.Number(123),
//   						Name: jsii.String("name"),
//   						SoftLimit: jsii.Number(123),
//   					},
//   				},
//   				User: jsii.String("user"),
//   				Vcpus: jsii.Number(123),
//   				Volumes: []interface{}{
//   					&VolumesProperty{
//   						EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   							FileSystemId: jsii.String("fileSystemId"),
//
//   							// the properties below are optional
//   							AuthorizationConfig: &AuthorizationConfigProperty{
//   								AccessPointId: jsii.String("accessPointId"),
//   								Iam: jsii.String("iam"),
//   							},
//   							RootDirectory: jsii.String("rootDirectory"),
//   							TransitEncryption: jsii.String("transitEncryption"),
//   							TransitEncryptionPort: jsii.Number(123),
//   						},
//   						Host: &VolumesHostProperty{
//   							SourcePath: jsii.String("sourcePath"),
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	NumNodes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-nodeproperties.html
//
type CfnJobDefinition_NodePropertiesProperty struct {
	// Specifies the node index for the main node of a multi-node parallel job.
	//
	// This node index value must be fewer than the number of nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-nodeproperties.html#cfn-batch-jobdefinition-nodeproperties-mainnode
	//
	MainNode *float64 `field:"required" json:"mainNode" yaml:"mainNode"`
	// A list of node ranges and their properties that are associated with a multi-node parallel job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-nodeproperties.html#cfn-batch-jobdefinition-nodeproperties-noderangeproperties
	//
	NodeRangeProperties interface{} `field:"required" json:"nodeRangeProperties" yaml:"nodeRangeProperties"`
	// The number of nodes that are associated with a multi-node parallel job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-nodeproperties.html#cfn-batch-jobdefinition-nodeproperties-numnodes
	//
	NumNodes *float64 `field:"required" json:"numNodes" yaml:"numNodes"`
}

