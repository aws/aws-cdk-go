package awsmedialivealpha


// Properties for creating an SDI Source.
//
// Example:
//   var stack Stack
//
//   sdi := medialive.NewSdiSource(stack, jsii.String("Sdi"), &SdiSourceProps{
//   	SdiSourceName: jsii.String("camera-1"),
//   	Type: medialive.SdiType_SINGLE(),
//   })
//
// Experimental.
type SdiSourceProps struct {
	// Type of SDI input.
	// Experimental.
	Type SdiType `field:"required" json:"type" yaml:"type"`
	// SDI Mode, only applicable for QUAD type.
	// Default: - no mode.
	//
	// Experimental.
	Mode SdiMode `field:"optional" json:"mode" yaml:"mode"`
	// The name of the SDI source.
	// Default: - auto-generated name.
	//
	// Experimental.
	SdiSourceName *string `field:"optional" json:"sdiSourceName" yaml:"sdiSourceName"`
	// Tags to add to the SDI source.
	// Default: - no tagging.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

