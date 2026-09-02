package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for frame capture codec settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var timecodeBurninFontSize TimecodeBurninFontSize
//   var timecodeBurninPosition TimecodeBurninPosition
//
//   frameCaptureSettingsProps := &FrameCaptureSettingsProps{
//   	CaptureInterval: cdk.Duration_Minutes(jsii.Number(30)),
//   	TimecodeBurnin: &TimecodeBurninSettings{
//   		FontSize: timecodeBurninFontSize,
//   		Position: timecodeBurninPosition,
//   		Prefix: jsii.String("prefix"),
//   	},
//   }
//
// Experimental.
type FrameCaptureSettingsProps struct {
	// The interval between frame captures.
	// Default: - service default.
	//
	// Experimental.
	CaptureInterval awscdk.Duration `field:"optional" json:"captureInterval" yaml:"captureInterval"`
	// Timecode burn-in settings to overlay timecode on the video.
	// Default: - no timecode burn-in.
	//
	// Experimental.
	TimecodeBurnin *TimecodeBurninSettings `field:"optional" json:"timecodeBurnin" yaml:"timecodeBurnin"`
}

