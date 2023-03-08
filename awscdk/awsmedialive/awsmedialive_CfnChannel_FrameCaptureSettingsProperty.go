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
//   frameCaptureSettingsProperty := &frameCaptureSettingsProperty{
//   	captureInterval: jsii.Number(123),
//   	captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   }
//
type CfnChannel_FrameCaptureSettingsProperty struct {
	// The frequency, in seconds, for capturing frames for inclusion in the output.
	//
	// For example, "10" means capture a frame every 10 seconds.
	CaptureInterval *float64 `field:"optional" json:"captureInterval" yaml:"captureInterval"`
	// Unit for the frame capture interval.
	CaptureIntervalUnits *string `field:"optional" json:"captureIntervalUnits" yaml:"captureIntervalUnits"`
}

