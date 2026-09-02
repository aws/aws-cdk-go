package awsmedialivealpha


// A display rectangle, expressed as percentages of the underlying video frame, for captions converted to EBU-TT-D or TTML.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   captionRectangle := &CaptionRectangle{
//   	Height: jsii.Number(123),
//   	LeftOffset: jsii.Number(123),
//   	TopOffset: jsii.Number(123),
//   	Width: jsii.Number(123),
//   }
//
// Experimental.
type CaptionRectangle struct {
	// Height of the rectangle, as a percentage of the frame height (0–100).
	// Experimental.
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// Left edge position, as a percentage of the frame width (0–100).
	// Experimental.
	LeftOffset *float64 `field:"required" json:"leftOffset" yaml:"leftOffset"`
	// Top edge position, as a percentage of the frame height (0–100).
	// Experimental.
	TopOffset *float64 `field:"required" json:"topOffset" yaml:"topOffset"`
	// Width of the rectangle, as a percentage of the frame width (0–100).
	// Experimental.
	Width *float64 `field:"required" json:"width" yaml:"width"`
}

