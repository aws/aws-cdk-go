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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-ec2subnetids
	//
	Ec2SubnetIds *[]*string `field:"required" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-launchprofileprotocolversions
	//
	LaunchProfileProtocolVersions *[]*string `field:"required" json:"launchProfileProtocolVersions" yaml:"launchProfileProtocolVersions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-streamconfiguration
	//
	StreamConfiguration interface{} `field:"required" json:"streamConfiguration" yaml:"streamConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-studiocomponentids
	//
	StudioComponentIds *[]*string `field:"required" json:"studioComponentIds" yaml:"studioComponentIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-studioid
	//
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

