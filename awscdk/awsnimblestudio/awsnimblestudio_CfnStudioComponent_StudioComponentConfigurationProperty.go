package awsnimblestudio


// The configuration of the studio component, based on component type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioComponentConfigurationProperty := &studioComponentConfigurationProperty{
//   	activeDirectoryConfiguration: &activeDirectoryConfigurationProperty{
//   		computerAttributes: []interface{}{
//   			&activeDirectoryComputerAttributeProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		directoryId: jsii.String("directoryId"),
//   		organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   	},
//   	computeFarmConfiguration: &computeFarmConfigurationProperty{
//   		activeDirectoryUser: jsii.String("activeDirectoryUser"),
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	licenseServiceConfiguration: &licenseServiceConfigurationProperty{
//   		endpoint: jsii.String("endpoint"),
//   	},
//   	sharedFileSystemConfiguration: &sharedFileSystemConfigurationProperty{
//   		endpoint: jsii.String("endpoint"),
//   		fileSystemId: jsii.String("fileSystemId"),
//   		linuxMountPoint: jsii.String("linuxMountPoint"),
//   		shareName: jsii.String("shareName"),
//   		windowsMountDrive: jsii.String("windowsMountDrive"),
//   	},
//   }
//
type CfnStudioComponent_StudioComponentConfigurationProperty struct {
	// The configuration for a Microsoft Active Directory (Microsoft AD) studio resource.
	ActiveDirectoryConfiguration interface{} `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// The configuration for a render farm that is associated with a studio resource.
	ComputeFarmConfiguration interface{} `field:"optional" json:"computeFarmConfiguration" yaml:"computeFarmConfiguration"`
	// The configuration for a license service that is associated with a studio resource.
	LicenseServiceConfiguration interface{} `field:"optional" json:"licenseServiceConfiguration" yaml:"licenseServiceConfiguration"`
	// The configuration for a shared file storage system that is associated with a studio resource.
	SharedFileSystemConfiguration interface{} `field:"optional" json:"sharedFileSystemConfiguration" yaml:"sharedFileSystemConfiguration"`
}

