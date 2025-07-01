package awsnimblestudio


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConfigurationProperty := &StreamConfigurationProperty{
//   	ClipboardMode: jsii.String("clipboardMode"),
//   	Ec2InstanceTypes: []*string{
//   		jsii.String("ec2InstanceTypes"),
//   	},
//   	StreamingImageIds: []*string{
//   		jsii.String("streamingImageIds"),
//   	},
//
//   	// the properties below are optional
//   	AutomaticTerminationMode: jsii.String("automaticTerminationMode"),
//   	MaxSessionLengthInMinutes: jsii.Number(123),
//   	MaxStoppedSessionLengthInMinutes: jsii.Number(123),
//   	SessionBackup: &StreamConfigurationSessionBackupProperty{
//   		MaxBackupsToRetain: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   	},
//   	SessionPersistenceMode: jsii.String("sessionPersistenceMode"),
//   	SessionStorage: &StreamConfigurationSessionStorageProperty{
//   		Mode: []*string{
//   			jsii.String("mode"),
//   		},
//
//   		// the properties below are optional
//   		Root: &StreamingSessionStorageRootProperty{
//   			Linux: jsii.String("linux"),
//   			Windows: jsii.String("windows"),
//   		},
//   	},
//   	VolumeConfiguration: &VolumeConfigurationProperty{
//   		Iops: jsii.Number(123),
//   		Size: jsii.Number(123),
//   		Throughput: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html
//
type CfnLaunchProfile_StreamConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-clipboardmode
	//
	ClipboardMode *string `field:"required" json:"clipboardMode" yaml:"clipboardMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-ec2instancetypes
	//
	Ec2InstanceTypes *[]*string `field:"required" json:"ec2InstanceTypes" yaml:"ec2InstanceTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-streamingimageids
	//
	StreamingImageIds *[]*string `field:"required" json:"streamingImageIds" yaml:"streamingImageIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-automaticterminationmode
	//
	AutomaticTerminationMode *string `field:"optional" json:"automaticTerminationMode" yaml:"automaticTerminationMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-maxsessionlengthinminutes
	//
	MaxSessionLengthInMinutes *float64 `field:"optional" json:"maxSessionLengthInMinutes" yaml:"maxSessionLengthInMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-maxstoppedsessionlengthinminutes
	//
	MaxStoppedSessionLengthInMinutes *float64 `field:"optional" json:"maxStoppedSessionLengthInMinutes" yaml:"maxStoppedSessionLengthInMinutes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionbackup
	//
	SessionBackup interface{} `field:"optional" json:"sessionBackup" yaml:"sessionBackup"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionpersistencemode
	//
	SessionPersistenceMode *string `field:"optional" json:"sessionPersistenceMode" yaml:"sessionPersistenceMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionstorage
	//
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-volumeconfiguration
	//
	VolumeConfiguration interface{} `field:"optional" json:"volumeConfiguration" yaml:"volumeConfiguration"`
}

