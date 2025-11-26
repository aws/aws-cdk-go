package previewawsnimblestudiomixins


// Properties for CfnLaunchProfilePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLaunchProfileMixinProps := &CfnLaunchProfileMixinProps{
//   	Description: jsii.String("description"),
//   	Ec2SubnetIds: []*string{
//   		jsii.String("ec2SubnetIds"),
//   	},
//   	LaunchProfileProtocolVersions: []*string{
//   		jsii.String("launchProfileProtocolVersions"),
//   	},
//   	Name: jsii.String("name"),
//   	StreamConfiguration: &StreamConfigurationProperty{
//   		AutomaticTerminationMode: jsii.String("automaticTerminationMode"),
//   		ClipboardMode: jsii.String("clipboardMode"),
//   		Ec2InstanceTypes: []*string{
//   			jsii.String("ec2InstanceTypes"),
//   		},
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
//   			Root: &StreamingSessionStorageRootProperty{
//   				Linux: jsii.String("linux"),
//   				Windows: jsii.String("windows"),
//   			},
//   		},
//   		StreamingImageIds: []*string{
//   			jsii.String("streamingImageIds"),
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
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html
//
type CfnLaunchProfileMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-ec2subnetids
	//
	Ec2SubnetIds *[]*string `field:"optional" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-launchprofileprotocolversions
	//
	LaunchProfileProtocolVersions *[]*string `field:"optional" json:"launchProfileProtocolVersions" yaml:"launchProfileProtocolVersions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-streamconfiguration
	//
	StreamConfiguration interface{} `field:"optional" json:"streamConfiguration" yaml:"streamConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-studiocomponentids
	//
	StudioComponentIds *[]*string `field:"optional" json:"studioComponentIds" yaml:"studioComponentIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-studioid
	//
	StudioId *string `field:"optional" json:"studioId" yaml:"studioId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-nimblestudio-launchprofile.html#cfn-nimblestudio-launchprofile-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

