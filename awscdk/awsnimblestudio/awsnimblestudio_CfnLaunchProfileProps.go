package awsnimblestudio


// Properties for defining a `CfnLaunchProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchProfileProps := &cfnLaunchProfileProps{
//   	ec2SubnetIds: []*string{
//   		jsii.String("ec2SubnetIds"),
//   	},
//   	launchProfileProtocolVersions: []*string{
//   		jsii.String("launchProfileProtocolVersions"),
//   	},
//   	name: jsii.String("name"),
//   	streamConfiguration: &streamConfigurationProperty{
//   		clipboardMode: jsii.String("clipboardMode"),
//   		ec2InstanceTypes: []*string{
//   			jsii.String("ec2InstanceTypes"),
//   		},
//   		streamingImageIds: []*string{
//   			jsii.String("streamingImageIds"),
//   		},
//
//   		// the properties below are optional
//   		maxSessionLengthInMinutes: jsii.Number(123),
//   		maxStoppedSessionLengthInMinutes: jsii.Number(123),
//   		sessionStorage: &streamConfigurationSessionStorageProperty{
//   			mode: []*string{
//   				jsii.String("mode"),
//   			},
//
//   			// the properties below are optional
//   			root: &streamingSessionStorageRootProperty{
//   				linux: jsii.String("linux"),
//   				windows: jsii.String("windows"),
//   			},
//   		},
//   	},
//   	studioComponentIds: []*string{
//   		jsii.String("studioComponentIds"),
//   	},
//   	studioId: jsii.String("studioId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnLaunchProfileProps struct {
	// Unique identifiers for a collection of EC2 subnets.
	Ec2SubnetIds *[]*string `field:"required" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// The version number of the protocol that is used by the launch profile.
	//
	// The only valid version is "2021-03-31".
	LaunchProfileProtocolVersions *[]*string `field:"required" json:"launchProfileProtocolVersions" yaml:"launchProfileProtocolVersions"`
	// A friendly name for the launch profile.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A configuration for a streaming session.
	StreamConfiguration interface{} `field:"required" json:"streamConfiguration" yaml:"streamConfiguration"`
	// Unique identifiers for a collection of studio components that can be used with this launch profile.
	StudioComponentIds *[]*string `field:"required" json:"studioComponentIds" yaml:"studioComponentIds"`
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// A human-readable description of the launch profile.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

