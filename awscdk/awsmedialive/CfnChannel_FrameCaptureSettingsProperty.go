package awsmedialive


// The frame capture settings.
//
// The parent of this entity is VideoCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureSettingsProperty := &FrameCaptureSettingsProperty{
//   	CaptureInterval: jsii.Number(123),
//   	CaptureIntervalUnits: jsii.String("captureIntervalUnits"),
//   	TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   		FontSize: jsii.String("fontSize"),
//   		Position: jsii.String("position"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturesettings.html
//
type CfnChannel_FrameCaptureSettingsProperty struct {
	// The frequency, in seconds, for capturing frames for inclusion in the output.
	//
	// For example, "10" means capture a frame every 10 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturesettings.html#cfn-medialive-channel-framecapturesettings-captureinterval
	//
	CaptureInterval *float64 `field:"optional" json:"captureInterval" yaml:"captureInterval"`
	// Unit for the frame capture interval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturesettings.html#cfn-medialive-channel-framecapturesettings-captureintervalunits
	//
	CaptureIntervalUnits *string `field:"optional" json:"captureIntervalUnits" yaml:"captureIntervalUnits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturesettings.html#cfn-medialive-channel-framecapturesettings-timecodeburninsettings
	//
	TimecodeBurninSettings interface{} `field:"optional" json:"timecodeBurninSettings" yaml:"timecodeBurninSettings"`
}

