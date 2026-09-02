package awsmedialivealpha


// Output definition for an MS Smooth output group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var encodeConfiguration EncodeConfiguration
//   var h265PackagingType H265PackagingType
//
//   msSmoothOutputDefinition := &MsSmoothOutputDefinition{
//   	Encodes: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	H265PackagingType: h265PackagingType,
//   	NameModifier: jsii.String("nameModifier"),
//   }
//
// Experimental.
type MsSmoothOutputDefinition struct {
	// The encode configurations to wire to this output.
	// Experimental.
	Encodes *[]EncodeConfiguration `field:"required" json:"encodes" yaml:"encodes"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// For H.265 video, whether to package as HEV1 or HVC1.
	// Default: - service default.
	//
	// Experimental.
	H265PackagingType H265PackagingType `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// A string concatenated to the end of the destination file name.
	//
	// Required if the output group
	// contains more than one output.
	// Default: - no name modifier.
	//
	// Experimental.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

