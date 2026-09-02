package awsmedialivealpha


// Output definition for an Archive output group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var archiveContainer ArchiveContainer
//   var encodeConfiguration EncodeConfiguration
//
//   archiveOutputDefinition := &ArchiveOutputDefinition{
//   	Encodes: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	Container: archiveContainer,
//   	Extension: jsii.String("extension"),
//   	NameModifier: jsii.String("nameModifier"),
//   }
//
// Experimental.
type ArchiveOutputDefinition struct {
	// The encode configurations to wire to this output.
	// Experimental.
	Encodes *[]EncodeConfiguration `field:"required" json:"encodes" yaml:"encodes"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The container (transport stream) for this output — an MPEG-TS (M2TS) or raw container.
	//
	// Use the `ArchiveContainer` factory methods.
	// Default: - ArchiveContainer.m2ts() with service-default M2TS settings
	//
	// Experimental.
	Container ArchiveContainer `field:"optional" json:"container" yaml:"container"`
	// The output file extension.
	// Default: - auto-selected from container type.
	//
	// Experimental.
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// A string concatenated to the end of the destination file name.
	//
	// Required if the output group
	// contains more than one output of the same container type.
	// Default: - no name modifier.
	//
	// Experimental.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

