package awsm2


// > AWS Mainframe Modernization Service (Managed Runtime Environment experience) will no longer be open to new customers starting on November 7, 2025.
//
// If you would like to use the service, please sign up prior to November 7, 2025. For capabilities similar to AWS Mainframe Modernization Service (Managed Runtime Environment experience) explore AWS Mainframe Modernization Service (Self-Managed Experience). Existing customers can continue to use the service as normal. For more information, see [AWS Mainframe Modernization availability change](https://docs.aws.amazon.com/m2/latest/userguide/mainframe-modernization-availability-change.html) .
//
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

