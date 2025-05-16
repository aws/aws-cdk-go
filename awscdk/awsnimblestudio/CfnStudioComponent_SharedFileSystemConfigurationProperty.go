package awsnimblestudio


// The configuration for a shared file storage system that is associated with a studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharedFileSystemConfigurationProperty := &SharedFileSystemConfigurationProperty{
//   	Endpoint: jsii.String("endpoint"),
//   	FileSystemId: jsii.String("fileSystemId"),
//   	LinuxMountPoint: jsii.String("linuxMountPoint"),
//   	ShareName: jsii.String("shareName"),
//   	WindowsMountDrive: jsii.String("windowsMountDrive"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-sharedfilesystemconfiguration.html
//
type CfnStudioComponent_SharedFileSystemConfigurationProperty struct {
	// The endpoint of the shared file system that is accessed by the studio component resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-sharedfilesystemconfiguration.html#cfn-nimblestudio-studiocomponent-sharedfilesystemconfiguration-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The unique identifier for a file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-sharedfilesystemconfiguration.html#cfn-nimblestudio-studiocomponent-sharedfilesystemconfiguration-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The mount location for a shared file system on a Linux virtual workstation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-sharedfilesystemconfiguration.html#cfn-nimblestudio-studiocomponent-sharedfilesystemconfiguration-linuxmountpoint
	//
	LinuxMountPoint *string `field:"optional" json:"linuxMountPoint" yaml:"linuxMountPoint"`
	// The name of the file share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-sharedfilesystemconfiguration.html#cfn-nimblestudio-studiocomponent-sharedfilesystemconfiguration-sharename
	//
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
	// The mount location for a shared file system on a Windows virtual workstation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-sharedfilesystemconfiguration.html#cfn-nimblestudio-studiocomponent-sharedfilesystemconfiguration-windowsmountdrive
	//
	WindowsMountDrive *string `field:"optional" json:"windowsMountDrive" yaml:"windowsMountDrive"`
}

