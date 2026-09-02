package awsmedialivealpha


// Properties for WebVTT caption output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var webvttStyleControl WebvttStyleControl
//
//   webvttDestinationProps := &WebvttDestinationProps{
//   	StyleControl: webvttStyleControl,
//   }
//
// Experimental.
type WebvttDestinationProps struct {
	// Whether to pass through source color and position to the WebVTT output.
	//
	// PASSTHROUGH is only valid for EMBEDDED or TELETEXT sources.
	// Default: WebvttStyleControl.NO_STYLE_DATA
	//
	// Experimental.
	StyleControl WebvttStyleControl `field:"optional" json:"styleControl" yaml:"styleControl"`
}

