package mixinsawsconfig


// Specifies the default recording frequency that AWS Config uses to record configuration changes.
//
// AWS Config supports *Continuous recording* and *Daily recording* .
//
// - Continuous recording allows you to record configuration changes continuously whenever a change occurs.
// - Daily recording allows you to receive a configuration item (CI) representing the most recent state of your resources over the last 24-hour period, only if it’s different from the previous CI recorded.
//
// > AWS Firewall Manager depends on continuous recording to monitor your resources. If you are using Firewall Manager, it is recommended that you set the recording frequency to Continuous.
//
// You can also override the recording frequency for specific resource types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recordingModeProperty := &RecordingModeProperty{
//   	RecordingFrequency: jsii.String("recordingFrequency"),
//   	RecordingModeOverrides: []interface{}{
//   		&RecordingModeOverrideProperty{
//   			Description: jsii.String("description"),
//   			RecordingFrequency: jsii.String("recordingFrequency"),
//   			ResourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmode.html
//
type CfnConfigurationRecorderPropsMixin_RecordingModeProperty struct {
	// The default recording frequency that AWS Config uses to record configuration changes.
	//
	// > Daily recording cannot be specified for the following resource types:
	// >
	// > - `AWS::Config::ResourceCompliance`
	// > - `AWS::Config::ConformancePackCompliance`
	// > - `AWS::Config::ConfigurationRecorder`
	// >
	// > For the *allSupported* ( `ALL_SUPPORTED_RESOURCE_TYPES` ) recording strategy, these resource types will be set to Continuous recording.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmode.html#cfn-config-configurationrecorder-recordingmode-recordingfrequency
	//
	RecordingFrequency *string `field:"optional" json:"recordingFrequency" yaml:"recordingFrequency"`
	// An array of `recordingModeOverride` objects for you to specify your overrides for the recording mode.
	//
	// The `recordingModeOverride` object in the `recordingModeOverrides` array consists of three fields: a `description` , the new `recordingFrequency` , and an array of `resourceTypes` to override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmode.html#cfn-config-configurationrecorder-recordingmode-recordingmodeoverrides
	//
	RecordingModeOverrides interface{} `field:"optional" json:"recordingModeOverrides" yaml:"recordingModeOverrides"`
}

