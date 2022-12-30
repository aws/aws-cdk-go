package awsnimblestudio


// A configuration for a streaming session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConfigurationProperty := &streamConfigurationProperty{
//   	clipboardMode: jsii.String("clipboardMode"),
//   	ec2InstanceTypes: []*string{
//   		jsii.String("ec2InstanceTypes"),
//   	},
//   	streamingImageIds: []*string{
//   		jsii.String("streamingImageIds"),
//   	},
//
//   	// the properties below are optional
//   	automaticTerminationMode: jsii.String("automaticTerminationMode"),
//   	maxSessionLengthInMinutes: jsii.Number(123),
//   	maxStoppedSessionLengthInMinutes: jsii.Number(123),
//   	sessionPersistenceMode: jsii.String("sessionPersistenceMode"),
//   	sessionStorage: &streamConfigurationSessionStorageProperty{
//   		mode: []*string{
//   			jsii.String("mode"),
//   		},
//
//   		// the properties below are optional
//   		root: &streamingSessionStorageRootProperty{
//   			linux: jsii.String("linux"),
//   			windows: jsii.String("windows"),
//   		},
//   	},
//   	volumeConfiguration: &volumeConfigurationProperty{
//   		iops: jsii.Number(123),
//   		size: jsii.Number(123),
//   		throughput: jsii.Number(123),
//   	},
//   }
//
type CfnLaunchProfile_StreamConfigurationProperty struct {
	// Enable or disable the use of the system clipboard to copy and paste between the streaming session and streaming client.
	ClipboardMode *string `field:"required" json:"clipboardMode" yaml:"clipboardMode"`
	// The EC2 instance types that users can select from when launching a streaming session with this launch profile.
	Ec2InstanceTypes *[]*string `field:"required" json:"ec2InstanceTypes" yaml:"ec2InstanceTypes"`
	// The streaming images that users can select from when launching a streaming session with this launch profile.
	StreamingImageIds *[]*string `field:"required" json:"streamingImageIds" yaml:"streamingImageIds"`
	// `CfnLaunchProfile.StreamConfigurationProperty.AutomaticTerminationMode`.
	AutomaticTerminationMode *string `field:"optional" json:"automaticTerminationMode" yaml:"automaticTerminationMode"`
	// The length of time, in minutes, that a streaming session can be active before it is stopped or terminated.
	//
	// After this point, Nimble Studio automatically terminates or stops the session. The default length of time is 690 minutes, and the maximum length of time is 30 days.
	MaxSessionLengthInMinutes *float64 `field:"optional" json:"maxSessionLengthInMinutes" yaml:"maxSessionLengthInMinutes"`
	// Integer that determines if you can start and stop your sessions and how long a session can stay in the STOPPED state.
	//
	// The default value is 0. The maximum value is 5760.
	//
	// If the value is missing or set to 0, your sessions can’t be stopped. If you then call `StopStreamingSession` , the session fails. If the time that a session stays in the READY state exceeds the `maxSessionLengthInMinutes` value, the session will automatically be terminated (instead of stopped).
	//
	// If the value is set to a positive number, the session can be stopped. You can call `StopStreamingSession` to stop sessions in the READY state. If the time that a session stays in the READY state exceeds the `maxSessionLengthInMinutes` value, the session will automatically be stopped (instead of terminated).
	MaxStoppedSessionLengthInMinutes *float64 `field:"optional" json:"maxStoppedSessionLengthInMinutes" yaml:"maxStoppedSessionLengthInMinutes"`
	// `CfnLaunchProfile.StreamConfigurationProperty.SessionPersistenceMode`.
	SessionPersistenceMode *string `field:"optional" json:"sessionPersistenceMode" yaml:"sessionPersistenceMode"`
	// (Optional) The upload storage for a streaming session.
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
	// `CfnLaunchProfile.StreamConfigurationProperty.VolumeConfiguration`.
	VolumeConfiguration interface{} `field:"optional" json:"volumeConfiguration" yaml:"volumeConfiguration"`
}

