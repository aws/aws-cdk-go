package awsnimblestudio


// The configuration of the studio component, based on component type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioComponentConfigurationProperty := &StudioComponentConfigurationProperty{
//   	ActiveDirectoryConfiguration: &ActiveDirectoryConfigurationProperty{
//   		ComputerAttributes: []interface{}{
//   			&ActiveDirectoryComputerAttributeProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		DirectoryId: jsii.String("directoryId"),
//   		OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	},
//   	ComputeFarmConfiguration: &ComputeFarmConfigurationProperty{
//   		ActiveDirectoryUser: jsii.String("activeDirectoryUser"),
//   		Endpoint: jsii.String("endpoint"),
//   	},
//   	LicenseServiceConfiguration: &LicenseServiceConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   	},
//   	SharedFileSystemConfiguration: &SharedFileSystemConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		FileSystemId: jsii.String("fileSystemId"),
//   		LinuxMountPoint: jsii.String("linuxMountPoint"),
//   		ShareName: jsii.String("shareName"),
//   		WindowsMountDrive: jsii.String("windowsMountDrive"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html
//
type CfnStudioComponent_StudioComponentConfigurationProperty struct {
	// The configuration for a AWS Directory Service for Microsoft Active Directory studio resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html#cfn-nimblestudio-studiocomponent-studiocomponentconfiguration-activedirectoryconfiguration
	//
	ActiveDirectoryConfiguration interface{} `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// The configuration for a render farm that is associated with a studio resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html#cfn-nimblestudio-studiocomponent-studiocomponentconfiguration-computefarmconfiguration
	//
	ComputeFarmConfiguration interface{} `field:"optional" json:"computeFarmConfiguration" yaml:"computeFarmConfiguration"`
	// The configuration for a license service that is associated with a studio resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html#cfn-nimblestudio-studiocomponent-studiocomponentconfiguration-licenseserviceconfiguration
	//
	LicenseServiceConfiguration interface{} `field:"optional" json:"licenseServiceConfiguration" yaml:"licenseServiceConfiguration"`
	// The configuration for a shared file storage system that is associated with a studio resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html#cfn-nimblestudio-studiocomponent-studiocomponentconfiguration-sharedfilesystemconfiguration
	//
	SharedFileSystemConfiguration interface{} `field:"optional" json:"sharedFileSystemConfiguration" yaml:"sharedFileSystemConfiguration"`
}

