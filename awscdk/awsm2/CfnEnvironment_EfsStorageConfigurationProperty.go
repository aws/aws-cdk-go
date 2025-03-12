package awsm2


// Defines the storage configuration for an Amazon EFS file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   efsStorageConfigurationProperty := &EfsStorageConfigurationProperty{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	MountPoint: jsii.String("mountPoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-efsstorageconfiguration.html
//
type CfnEnvironment_EfsStorageConfigurationProperty struct {
	// The file system identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-efsstorageconfiguration.html#cfn-m2-environment-efsstorageconfiguration-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The mount point for the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-m2-environment-efsstorageconfiguration.html#cfn-m2-environment-efsstorageconfiguration-mountpoint
	//
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
}

