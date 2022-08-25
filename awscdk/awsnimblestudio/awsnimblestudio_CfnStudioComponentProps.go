package awsnimblestudio


// Properties for defining a `CfnStudioComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioComponentProps := &cfnStudioComponentProps{
//   	name: jsii.String("name"),
//   	studioId: jsii.String("studioId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	configuration: &studioComponentConfigurationProperty{
//   		activeDirectoryConfiguration: &activeDirectoryConfigurationProperty{
//   			computerAttributes: []interface{}{
//   				&activeDirectoryComputerAttributeProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			directoryId: jsii.String("directoryId"),
//   			organizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		},
//   		computeFarmConfiguration: &computeFarmConfigurationProperty{
//   			activeDirectoryUser: jsii.String("activeDirectoryUser"),
//   			endpoint: jsii.String("endpoint"),
//   		},
//   		licenseServiceConfiguration: &licenseServiceConfigurationProperty{
//   			endpoint: jsii.String("endpoint"),
//   		},
//   		sharedFileSystemConfiguration: &sharedFileSystemConfigurationProperty{
//   			endpoint: jsii.String("endpoint"),
//   			fileSystemId: jsii.String("fileSystemId"),
//   			linuxMountPoint: jsii.String("linuxMountPoint"),
//   			shareName: jsii.String("shareName"),
//   			windowsMountDrive: jsii.String("windowsMountDrive"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	ec2SecurityGroupIds: []*string{
//   		jsii.String("ec2SecurityGroupIds"),
//   	},
//   	initializationScripts: []interface{}{
//   		&studioComponentInitializationScriptProperty{
//   			launchProfileProtocolVersion: jsii.String("launchProfileProtocolVersion"),
//   			platform: jsii.String("platform"),
//   			runContext: jsii.String("runContext"),
//   			script: jsii.String("script"),
//   		},
//   	},
//   	scriptParameters: []interface{}{
//   		&scriptParameterKeyValueProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	subtype: jsii.String("subtype"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnStudioComponentProps struct {
	// A friendly name for the studio component resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// The type of the studio component.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration of the studio component, based on component type.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A human-readable description for the studio component resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds *[]*string `field:"optional" json:"ec2SecurityGroupIds" yaml:"ec2SecurityGroupIds"`
	// Initialization scripts for studio components.
	InitializationScripts interface{} `field:"optional" json:"initializationScripts" yaml:"initializationScripts"`
	// Parameters for the studio component scripts.
	ScriptParameters interface{} `field:"optional" json:"scriptParameters" yaml:"scriptParameters"`
	// The specific subtype of a studio component.
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

