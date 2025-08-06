package awsconfig


// An object for you to specify your overrides for the recording mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordingModeOverrideProperty := &RecordingModeOverrideProperty{
//   	RecordingFrequency: jsii.String("recordingFrequency"),
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmodeoverride.html
//
type CfnConfigurationRecorder_RecordingModeOverrideProperty struct {
	// The recording frequency that will be applied to all the resource types specified in the override.
	//
	// - Continuous recording allows you to record configuration changes continuously whenever a change occurs.
	// - Daily recording allows you to receive a configuration item (CI) representing the most recent state of your resources over the last 24-hour period, only if it’s different from the previous CI recorded.
	//
	// > AWS Firewall Manager depends on continuous recording to monitor your resources. If you are using Firewall Manager, it is recommended that you set the recording frequency to Continuous.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmodeoverride.html#cfn-config-configurationrecorder-recordingmodeoverride-recordingfrequency
	//
	RecordingFrequency *string `field:"required" json:"recordingFrequency" yaml:"recordingFrequency"`
	// A comma-separated list that specifies which resource types AWS Config includes in the override.
	//
	// > Daily recording cannot be specified for the following resource types:
	// >
	// > - `AWS::Config::ResourceCompliance`
	// > - `AWS::Config::ConformancePackCompliance`
	// > - `AWS::Config::ConfigurationRecorder`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmodeoverride.html#cfn-config-configurationrecorder-recordingmodeoverride-resourcetypes
	//
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// A description that you provide for the override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingmodeoverride.html#cfn-config-configurationrecorder-recordingmodeoverride-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

