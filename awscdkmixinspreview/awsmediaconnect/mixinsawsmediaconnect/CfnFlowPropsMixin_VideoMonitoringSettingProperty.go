package mixinsawsmediaconnect


// Specifies the configuration for video stream metrics monitoring.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   videoMonitoringSettingProperty := &VideoMonitoringSettingProperty{
//   	BlackFrames: &BlackFramesProperty{
//   		State: jsii.String("state"),
//   		ThresholdSeconds: jsii.Number(123),
//   	},
//   	FrozenFrames: &FrozenFramesProperty{
//   		State: jsii.String("state"),
//   		ThresholdSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-videomonitoringsetting.html
//
type CfnFlowPropsMixin_VideoMonitoringSettingProperty struct {
	// Detects video frames that are black.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-videomonitoringsetting.html#cfn-mediaconnect-flow-videomonitoringsetting-blackframes
	//
	BlackFrames interface{} `field:"optional" json:"blackFrames" yaml:"blackFrames"`
	// Detects video frames that have not changed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-videomonitoringsetting.html#cfn-mediaconnect-flow-videomonitoringsetting-frozenframes
	//
	FrozenFrames interface{} `field:"optional" json:"frozenFrames" yaml:"frozenFrames"`
}

