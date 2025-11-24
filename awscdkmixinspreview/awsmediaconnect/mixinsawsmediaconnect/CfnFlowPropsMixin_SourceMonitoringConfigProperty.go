package mixinsawsmediaconnect


// The `SourceMonitoringConfig` property type specifies the source monitoring settings for an `AWS::MediaConnect::Flow` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceMonitoringConfigProperty := &SourceMonitoringConfigProperty{
//   	AudioMonitoringSettings: []interface{}{
//   		&AudioMonitoringSettingProperty{
//   			SilentAudio: &SilentAudioProperty{
//   				State: jsii.String("state"),
//   				ThresholdSeconds: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ContentQualityAnalysisState: jsii.String("contentQualityAnalysisState"),
//   	ThumbnailState: jsii.String("thumbnailState"),
//   	VideoMonitoringSettings: []interface{}{
//   		&VideoMonitoringSettingProperty{
//   			BlackFrames: &BlackFramesProperty{
//   				State: jsii.String("state"),
//   				ThresholdSeconds: jsii.Number(123),
//   			},
//   			FrozenFrames: &FrozenFramesProperty{
//   				State: jsii.String("state"),
//   				ThresholdSeconds: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-sourcemonitoringconfig.html
//
type CfnFlowPropsMixin_SourceMonitoringConfigProperty struct {
	// Contains the settings for audio stream metrics monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-sourcemonitoringconfig.html#cfn-mediaconnect-flow-sourcemonitoringconfig-audiomonitoringsettings
	//
	AudioMonitoringSettings interface{} `field:"optional" json:"audioMonitoringSettings" yaml:"audioMonitoringSettings"`
	// Indicates whether content quality analysis is enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-sourcemonitoringconfig.html#cfn-mediaconnect-flow-sourcemonitoringconfig-contentqualityanalysisstate
	//
	ContentQualityAnalysisState *string `field:"optional" json:"contentQualityAnalysisState" yaml:"contentQualityAnalysisState"`
	// The current state of the thumbnail monitoring.
	//
	// - If you don't explicitly specify a value when creating a flow, no thumbnail state will be set.
	// - If you update an existing flow and remove a previously set thumbnail state, the value will change to `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-sourcemonitoringconfig.html#cfn-mediaconnect-flow-sourcemonitoringconfig-thumbnailstate
	//
	ThumbnailState *string `field:"optional" json:"thumbnailState" yaml:"thumbnailState"`
	// Contains the settings for video stream metrics monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-sourcemonitoringconfig.html#cfn-mediaconnect-flow-sourcemonitoringconfig-videomonitoringsettings
	//
	VideoMonitoringSettings interface{} `field:"optional" json:"videoMonitoringSettings" yaml:"videoMonitoringSettings"`
}

