package awsmedialivealpha


// Properties for EBU-TT-D caption output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var ebuTtDFillLineGap EbuTtDFillLineGap
//   var ebuTtDStyleControl EbuTtDStyleControl
//
//   ebuTtDDestinationProps := &EbuTtDDestinationProps{
//   	CopyrightHolder: jsii.String("copyrightHolder"),
//   	DefaultFontSize: jsii.Number(123),
//   	DefaultLineHeight: jsii.Number(123),
//   	FillLineGap: ebuTtDFillLineGap,
//   	FontFamily: jsii.String("fontFamily"),
//   	StyleControl: ebuTtDStyleControl,
//   }
//
// Experimental.
type EbuTtDDestinationProps struct {
	// The copyright holder included in the TTML copyright metadata tag.
	// Default: - no copyright holder.
	//
	// Experimental.
	CopyrightHolder *string `field:"optional" json:"copyrightHolder" yaml:"copyrightHolder"`
	// The default font size.
	// Default: - service default.
	//
	// Experimental.
	DefaultFontSize *float64 `field:"optional" json:"defaultFontSize" yaml:"defaultFontSize"`
	// The default line height.
	// Default: - service default.
	//
	// Experimental.
	DefaultLineHeight *float64 `field:"optional" json:"defaultLineHeight" yaml:"defaultLineHeight"`
	// How to handle the gap between multi-line captions.
	// Default: - service default.
	//
	// Experimental.
	FillLineGap EbuTtDFillLineGap `field:"optional" json:"fillLineGap" yaml:"fillLineGap"`
	// A comma-separated list of font families to include (valid only when `styleControl` is INCLUDE).
	// Default: - 'monospaced'.
	//
	// Experimental.
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// Whether source style information is included in the EBU-TT-D font data.
	// Default: - service default.
	//
	// Experimental.
	StyleControl EbuTtDStyleControl `field:"optional" json:"styleControl" yaml:"styleControl"`
}

