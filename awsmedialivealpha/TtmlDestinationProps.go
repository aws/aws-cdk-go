package awsmedialivealpha


// Properties for TTML caption output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var ttmlStyleControl TtmlStyleControl
//
//   ttmlDestinationProps := &TtmlDestinationProps{
//   	StyleControl: ttmlStyleControl,
//   }
//
// Experimental.
type TtmlDestinationProps struct {
	// Whether to pass through style and position from a TTML-like source.
	// Default: TtmlStyleControl.PASSTHROUGH
	//
	// Experimental.
	StyleControl TtmlStyleControl `field:"optional" json:"styleControl" yaml:"styleControl"`
}

