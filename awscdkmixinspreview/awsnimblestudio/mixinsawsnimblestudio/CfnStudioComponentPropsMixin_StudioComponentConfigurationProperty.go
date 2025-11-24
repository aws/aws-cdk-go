package mixinsawsnimblestudio


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnStudioComponentPropsMixin_StudioComponentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html#cfn-nimblestudio-studiocomponent-studiocomponentconfiguration-activedirectoryconfiguration
	//
	ActiveDirectoryConfiguration interface{} `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html#cfn-nimblestudio-studiocomponent-studiocomponentconfiguration-computefarmconfiguration
	//
	ComputeFarmConfiguration interface{} `field:"optional" json:"computeFarmConfiguration" yaml:"computeFarmConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html#cfn-nimblestudio-studiocomponent-studiocomponentconfiguration-licenseserviceconfiguration
	//
	LicenseServiceConfiguration interface{} `field:"optional" json:"licenseServiceConfiguration" yaml:"licenseServiceConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studiocomponent-studiocomponentconfiguration.html#cfn-nimblestudio-studiocomponent-studiocomponentconfiguration-sharedfilesystemconfiguration
	//
	SharedFileSystemConfiguration interface{} `field:"optional" json:"sharedFileSystemConfiguration" yaml:"sharedFileSystemConfiguration"`
}

