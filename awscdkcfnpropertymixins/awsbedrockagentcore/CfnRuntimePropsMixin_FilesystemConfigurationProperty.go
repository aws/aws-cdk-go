package awsbedrockagentcore


// Filesystem configuration for the runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   filesystemConfigurationProperty := &FilesystemConfigurationProperty{
//   	CapacityProviderVolume: &CapacityProviderVolumeConfigurationProperty{
//   		MountPath: jsii.String("mountPath"),
//   		VolumeName: jsii.String("volumeName"),
//   	},
//   	EfsAccessPoint: &EfsAccessPointConfigurationProperty{
//   		AccessPointArn: jsii.String("accessPointArn"),
//   		MountPath: jsii.String("mountPath"),
//   	},
//   	S3FilesAccessPoint: &S3FilesAccessPointConfigurationProperty{
//   		AccessPointArn: jsii.String("accessPointArn"),
//   		MountPath: jsii.String("mountPath"),
//   	},
//   	SessionStorage: &SessionStorageConfigurationProperty{
//   		MountPath: jsii.String("mountPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-filesystemconfiguration.html
//
type CfnRuntimePropsMixin_FilesystemConfigurationProperty struct {
	// Configuration for a CapacityProvider-managed volume to mount into the agent runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-filesystemconfiguration.html#cfn-bedrockagentcore-runtime-filesystemconfiguration-capacityprovidervolume
	//
	CapacityProviderVolume interface{} `field:"optional" json:"capacityProviderVolume" yaml:"capacityProviderVolume"`
	// Configuration for EFS access point filesystem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-filesystemconfiguration.html#cfn-bedrockagentcore-runtime-filesystemconfiguration-efsaccesspoint
	//
	EfsAccessPoint interface{} `field:"optional" json:"efsAccessPoint" yaml:"efsAccessPoint"`
	// Configuration for S3 Files access point filesystem.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-filesystemconfiguration.html#cfn-bedrockagentcore-runtime-filesystemconfiguration-s3filesaccesspoint
	//
	S3FilesAccessPoint interface{} `field:"optional" json:"s3FilesAccessPoint" yaml:"s3FilesAccessPoint"`
	// Configuration for session storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-filesystemconfiguration.html#cfn-bedrockagentcore-runtime-filesystemconfiguration-sessionstorage
	//
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

