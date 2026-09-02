package awsmedialivealpha


// Settings for burning a timecode overlay into the video output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var timecodeBurninFontSize TimecodeBurninFontSize
//   var timecodeBurninPosition TimecodeBurninPosition
//
//   timecodeBurninSettings := &TimecodeBurninSettings{
//   	FontSize: timecodeBurninFontSize,
//   	Position: timecodeBurninPosition,
//   	Prefix: jsii.String("prefix"),
//   }
//
// Experimental.
type TimecodeBurninSettings struct {
	// The font size of the timecode overlay.
	// Default: - service default.
	//
	// Experimental.
	FontSize TimecodeBurninFontSize `field:"optional" json:"fontSize" yaml:"fontSize"`
	// The position of the timecode overlay on the video.
	// Default: - service default.
	//
	// Experimental.
	Position TimecodeBurninPosition `field:"optional" json:"position" yaml:"position"`
	// A string prepended to the timecode (e.g. a channel name).
	// Default: - no prefix.
	//
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

