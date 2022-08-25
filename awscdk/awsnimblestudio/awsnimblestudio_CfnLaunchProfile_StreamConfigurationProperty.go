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
//   	maxSessionLengthInMinutes: jsii.Number(123),
//   	maxStoppedSessionLengthInMinutes: jsii.Number(123),
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
//   }
//
type CfnLaunchProfile_StreamConfigurationProperty struct {
	// Enable or disable the use of the system clipboard to copy and paste between the streaming session and streaming client.
	ClipboardMode *string `field:"required" json:"clipboardMode" yaml:"clipboardMode"`
	// The EC2 instance types that users can select from when launching a streaming session with this launch profile.
	Ec2InstanceTypes *[]*string `field:"required" json:"ec2InstanceTypes" yaml:"ec2InstanceTypes"`
	// The streaming images that users can select from when launching a streaming session with this launch profile.
	StreamingImageIds *[]*string `field:"required" json:"streamingImageIds" yaml:"streamingImageIds"`
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
	// (Optional) The upload storage for a streaming session.
	SessionStorage interface{} `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

