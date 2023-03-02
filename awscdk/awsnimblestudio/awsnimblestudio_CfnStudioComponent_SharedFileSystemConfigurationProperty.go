package awsnimblestudio


// The configuration for a shared file storage system that is associated with a studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharedFileSystemConfigurationProperty := &sharedFileSystemConfigurationProperty{
//   	endpoint: jsii.String("endpoint"),
//   	fileSystemId: jsii.String("fileSystemId"),
//   	linuxMountPoint: jsii.String("linuxMountPoint"),
//   	shareName: jsii.String("shareName"),
//   	windowsMountDrive: jsii.String("windowsMountDrive"),
//   }
//
type CfnStudioComponent_SharedFileSystemConfigurationProperty struct {
	// The endpoint of the shared file system that is accessed by the studio component resource.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The unique identifier for a file system.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The mount location for a shared file system on a Linux virtual workstation.
	LinuxMountPoint *string `field:"optional" json:"linuxMountPoint" yaml:"linuxMountPoint"`
	// The name of the file share.
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
	// The mount location for a shared file system on a Windows virtual workstation.
	WindowsMountDrive *string `field:"optional" json:"windowsMountDrive" yaml:"windowsMountDrive"`
}

