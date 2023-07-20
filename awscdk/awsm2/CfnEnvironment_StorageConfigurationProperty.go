package awsm2


// Defines the storage configuration for a runtime environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageConfigurationProperty := &StorageConfigurationProperty{
//   	Efs: &EfsStorageConfigurationProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//   		MountPoint: jsii.String("mountPoint"),
//   	},
//   	Fsx: &FsxStorageConfigurationProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//   		MountPoint: jsii.String("mountPoint"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-storageconfiguration.html
//
type CfnEnvironment_StorageConfigurationProperty struct {
	// Defines the storage configuration for an Amazon EFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-storageconfiguration.html#cfn-m2-environment-storageconfiguration-efs
	//
	Efs interface{} `field:"optional" json:"efs" yaml:"efs"`
	// Defines the storage configuration for an Amazon FSx file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-storageconfiguration.html#cfn-m2-environment-storageconfiguration-fsx
	//
	Fsx interface{} `field:"optional" json:"fsx" yaml:"fsx"`
}

