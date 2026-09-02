package awsmedialivealpha


// Properties for a bandwidth reduction filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var bandwidthReductionPostFilterSharpening BandwidthReductionPostFilterSharpening
//   var bandwidthReductionStrength BandwidthReductionStrength
//
//   bandwidthReductionFilterProps := &BandwidthReductionFilterProps{
//   	PostFilterSharpening: bandwidthReductionPostFilterSharpening,
//   	Strength: bandwidthReductionStrength,
//   }
//
// Experimental.
type BandwidthReductionFilterProps struct {
	// Post-filter sharpening control.
	// Default: - service default.
	//
	// Experimental.
	PostFilterSharpening BandwidthReductionPostFilterSharpening `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// Bandwidth reduction strength.
	// Default: - service default.
	//
	// Experimental.
	Strength BandwidthReductionStrength `field:"optional" json:"strength" yaml:"strength"`
}

