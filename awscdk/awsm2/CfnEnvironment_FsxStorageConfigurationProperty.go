package awsm2


// Defines the storage configuration for an Amazon FSx file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fsxStorageConfigurationProperty := &FsxStorageConfigurationProperty{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	MountPoint: jsii.String("mountPoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-fsxstorageconfiguration.html
//
type CfnEnvironment_FsxStorageConfigurationProperty struct {
	// The file system identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-fsxstorageconfiguration.html#cfn-m2-environment-fsxstorageconfiguration-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The mount point for the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-fsxstorageconfiguration.html#cfn-m2-environment-fsxstorageconfiguration-mountpoint
	//
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
}

