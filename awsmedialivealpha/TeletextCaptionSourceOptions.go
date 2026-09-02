package awsmedialivealpha


// Options for a Teletext caption source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   teletextCaptionSourceOptions := &TeletextCaptionSourceOptions{
//   	OutputRectangle: &CaptionRectangle{
//   		Height: jsii.Number(123),
//   		LeftOffset: jsii.Number(123),
//   		TopOffset: jsii.Number(123),
//   		Width: jsii.Number(123),
//   	},
//   	PageNumber: jsii.String("pageNumber"),
//   }
//
// Experimental.
type TeletextCaptionSourceOptions struct {
	// The caption rectangle to use when converting this source to EBU-TT-D or TTML.
	// Default: - no rectangle (service default).
	//
	// Experimental.
	OutputRectangle *CaptionRectangle `field:"optional" json:"outputRectangle" yaml:"outputRectangle"`
	// The Teletext page number to extract captions from, as a hexadecimal string with no `0x` prefix (range `100`–`8FF`).
	//
	// Unused for passthrough.
	// Default: - MediaLive service default.
	//
	// Experimental.
	PageNumber *string `field:"optional" json:"pageNumber" yaml:"pageNumber"`
}

