package awsm2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   efsStorageConfigurationProperty := &efsStorageConfigurationProperty{
//   	fileSystemId: jsii.String("fileSystemId"),
//   	mountPoint: jsii.String("mountPoint"),
//   }
//
type CfnEnvironment_EfsStorageConfigurationProperty struct {
	// `CfnEnvironment.EfsStorageConfigurationProperty.FileSystemId`.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// `CfnEnvironment.EfsStorageConfigurationProperty.MountPoint`.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
}

