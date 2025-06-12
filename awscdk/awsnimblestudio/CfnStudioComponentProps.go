package awsnimblestudio


// Properties for defining a `CfnStudioComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioComponentProps := &CfnStudioComponentProps{
//   	Name: jsii.String("name"),
//   	StudioId: jsii.String("studioId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Configuration: &StudioComponentConfigurationProperty{
//   		ActiveDirectoryConfiguration: &ActiveDirectoryConfigurationProperty{
//   			ComputerAttributes: []interface{}{
//   				&ActiveDirectoryComputerAttributeProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			DirectoryId: jsii.String("directoryId"),
//   			OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		},
//   		ComputeFarmConfiguration: &ComputeFarmConfigurationProperty{
//   			ActiveDirectoryUser: jsii.String("activeDirectoryUser"),
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		LicenseServiceConfiguration: &LicenseServiceConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		SharedFileSystemConfiguration: &SharedFileSystemConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			FileSystemId: jsii.String("fileSystemId"),
//   			LinuxMountPoint: jsii.String("linuxMountPoint"),
//   			ShareName: jsii.String("shareName"),
//   			WindowsMountDrive: jsii.String("windowsMountDrive"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Ec2SecurityGroupIds: []*string{
//   		jsii.String("ec2SecurityGroupIds"),
//   	},
//   	InitializationScripts: []interface{}{
//   		&StudioComponentInitializationScriptProperty{
//   			LaunchProfileProtocolVersion: jsii.String("launchProfileProtocolVersion"),
//   			Platform: jsii.String("platform"),
//   			RunContext: jsii.String("runContext"),
//   			Script: jsii.String("script"),
//   		},
//   	},
//   	ScriptParameters: []interface{}{
//   		&ScriptParameterKeyValueProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Subtype: jsii.String("subtype"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html
//
type CfnStudioComponentProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-studioid
	//
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-ec2securitygroupids
	//
	Ec2SecurityGroupIds *[]*string `field:"optional" json:"ec2SecurityGroupIds" yaml:"ec2SecurityGroupIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-initializationscripts
	//
	InitializationScripts interface{} `field:"optional" json:"initializationScripts" yaml:"initializationScripts"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-scriptparameters
	//
	ScriptParameters interface{} `field:"optional" json:"scriptParameters" yaml:"scriptParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-subtype
	//
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-studiocomponent.html#cfn-nimblestudio-studiocomponent-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

