package awsmedialivealpha


// Properties for a temporal filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var temporalFilterPostFilterSharpening TemporalFilterPostFilterSharpening
//   var temporalFilterStrength TemporalFilterStrength
//
//   temporalFilterProps := &TemporalFilterProps{
//   	PostFilterSharpening: temporalFilterPostFilterSharpening,
//   	Strength: temporalFilterStrength,
//   }
//
// Experimental.
type TemporalFilterProps struct {
	// Post-filter sharpening control.
	// Default: - service default.
	//
	// Experimental.
	PostFilterSharpening TemporalFilterPostFilterSharpening `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// Filter strength.
	//
	// We recommend 1 or 2. Higher values may remove useful detail.
	// Default: - service default.
	//
	// Experimental.
	Strength TemporalFilterStrength `field:"optional" json:"strength" yaml:"strength"`
}

