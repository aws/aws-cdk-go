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
type CfnEnvironment_EfsStorageConfigurationProperty struct {
	// The file system identifier.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The mount point for the file system.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
}

