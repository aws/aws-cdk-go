package awsnimblestudio


// A configuration for a streaming session.
//
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
	// Allows or deactivates the use of the system clipboard to copy and paste between the streaming session and streaming client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-clipboardmode
	//
	ClipboardMode *string `field:"required" json:"clipboardMode" yaml:"clipboardMode"`
	// The EC2 instance types that users can select from when launching a streaming session with this launch profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-ec2instancetypes
	//
	Ec2InstanceTypes *[]*string `field:"required" json:"ec2InstanceTypes" yaml:"ec2InstanceTypes"`
	// The streaming images that users can select from when launching a streaming session with this launch profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-streamingimageids
	//
	StreamingImageIds *[]*string `field:"required" json:"streamingImageIds" yaml:"streamingImageIds"`
	// Indicates if a streaming session created from this launch profile should be terminated automatically or retained without termination after being in a `STOPPED` state.
	//
	// - When `ACTIVATED` , the streaming session is scheduled for termination after being in the `STOPPED` state for the time specified in `maxStoppedSessionLengthInMinutes` .
	// - When `DEACTIVATED` , the streaming session can remain in the `STOPPED` state indefinitely.
	//
	// This parameter is only allowed when `sessionPersistenceMode` is `ACTIVATED` . When allowed, the default value for this parameter is `DEACTIVATED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-automaticterminationmode
	//
	AutomaticTerminationMode *string `field:"optional" json:"automaticTerminationMode" yaml:"automaticTerminationMode"`
	// The length of time, in minutes, that a streaming session can be active before it is stopped or terminated.
	//
	// After this point, Nimble Studio automatically terminates or stops the session. The default length of time is 690 minutes, and the maximum length of time is 30 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-maxsessionlengthinminutes
	//
	MaxSessionLengthInMinutes *float64 `field:"optional" json:"maxSessionLengthInMinutes" yaml:"maxSessionLengthInMinutes"`
	// Integer that determines if you can start and stop your sessions and how long a session can stay in the `STOPPED` state.
	//
	// The default value is 0. The maximum value is 5760.
	//
	// This field is allowed only when `sessionPersistenceMode` is `ACTIVATED` and `automaticTerminationMode` is `ACTIVATED` .
	//
	// If the value is set to 0, your sessions can’t be `STOPPED` . If you then call `StopStreamingSession` , the session fails. If the time that a session stays in the `READY` state exceeds the `maxSessionLengthInMinutes` value, the session will automatically be terminated (instead of `STOPPED` ).
	//
	// If the value is set to a positive number, the session can be stopped. You can call `StopStreamingSession` to stop sessions in the `READY` state. If the time that a session stays in the `READY` state exceeds the `maxSessionLengthInMinutes` value, the session will automatically be stopped (instead of terminated).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-maxstoppedsessionlengthinminutes
	//
	MaxStoppedSessionLengthInMinutes *float64 `field:"optional" json:"maxStoppedSessionLengthInMinutes" yaml:"maxStoppedSessionLengthInMinutes"`
	// Information about the streaming session backup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionbackup
	//
	SessionBackup interface{} `field:"optional" json:"sessionBackup" yaml:"sessionBackup"`
	// Determine if a streaming session created from this launch profile can configure persistent storage.
	//
	// This means that `volumeConfiguration` and `automaticTerminationMode` are configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionpersistencemode
	//
	SessionPersistenceMode *string `field:"optional" json:"sessionPersistenceMode" yaml:"sessionPersistenceMode"`
	// The upload storage for a streaming session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-sessionstorage
	//
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
	// Custom volume configuration for the root volumes that are attached to streaming sessions.
	//
	// This parameter is only allowed when `sessionPersistenceMode` is `ACTIVATED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamconfiguration.html#cfn-nimblestudio-launchprofile-streamconfiguration-volumeconfiguration
	//
	VolumeConfiguration interface{} `field:"optional" json:"volumeConfiguration" yaml:"volumeConfiguration"`
}

