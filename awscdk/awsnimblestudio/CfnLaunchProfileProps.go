package awsnimblestudio


// Properties for defining a `CfnLaunchProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchProfileProps := &CfnLaunchProfileProps{
//   	Ec2SubnetIds: []*string{
//   		jsii.String("ec2SubnetIds"),
//   	},
//   	LaunchProfileProtocolVersions: []*string{
//   		jsii.String("launchProfileProtocolVersions"),
//   	},
//   	Name: jsii.String("name"),
//   	StreamConfiguration: &StreamConfigurationProperty{
//   		ClipboardMode: jsii.String("clipboardMode"),
//   		Ec2InstanceTypes: []*string{
//   			jsii.String("ec2InstanceTypes"),
//   		},
//   		StreamingImageIds: []*string{
//   			jsii.String("streamingImageIds"),
//   		},
//
//   		// the properties below are optional
//   		AutomaticTerminationMode: jsii.String("automaticTerminationMode"),
//   		MaxSessionLengthInMinutes: jsii.Number(123),
//   		MaxStoppedSessionLengthInMinutes: jsii.Number(123),
//   		SessionBackup: &StreamConfigurationSessionBackupProperty{
//   			MaxBackupsToRetain: jsii.Number(123),
//   			Mode: jsii.String("mode"),
//   		},
//   		SessionPersistenceMode: jsii.String("sessionPersistenceMode"),
//   		SessionStorage: &StreamConfigurationSessionStorageProperty{
//   			Mode: []*string{
//   				jsii.String("mode"),
//   			},
//
//   			// the properties below are optional
//   			Root: &StreamingSessionStorageRootProperty{
//   				Linux: jsii.String("linux"),
//   				Windows: jsii.String("windows"),
//   			},
//   		},
//   		VolumeConfiguration: &VolumeConfigurationProperty{
//   			Iops: jsii.Number(123),
//   			Size: jsii.Number(123),
//   			Throughput: jsii.Number(123),
//   		},
//   	},
//   	StudioComponentIds: []*string{
//   		jsii.String("studioComponentIds"),
//   	},
//   	StudioId: jsii.String("studioId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html
//
type CfnLaunchProfileProps struct {
	// Unique identifiers for a collection of EC2 subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-ec2subnetids
	//
	Ec2SubnetIds *[]*string `field:"required" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// The version number of the protocol that is used by the launch profile.
	//
	// The only valid version is "2021-03-31".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-launchprofileprotocolversions
	//
	LaunchProfileProtocolVersions *[]*string `field:"required" json:"launchProfileProtocolVersions" yaml:"launchProfileProtocolVersions"`
	// A friendly name for the launch profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A configuration for a streaming session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-streamconfiguration
	//
	StreamConfiguration interface{} `field:"required" json:"streamConfiguration" yaml:"streamConfiguration"`
	// Unique identifiers for a collection of studio components that can be used with this launch profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-studiocomponentids
	//
	StudioComponentIds *[]*string `field:"required" json:"studioComponentIds" yaml:"studioComponentIds"`
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio, all other resources are contained in a studio resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-studioid
	//
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// A human-readable description of the launch profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

