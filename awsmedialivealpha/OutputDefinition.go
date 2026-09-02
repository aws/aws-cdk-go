package awsmedialivealpha


// Base output definition — shared by all output group types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var encodeConfiguration EncodeConfiguration
//
//   outputDefinition := &OutputDefinition{
//   	Encodes: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	OutputName: jsii.String("outputName"),
//   }
//
// Experimental.
type OutputDefinition struct {
	// The encode configurations to wire to this output.
	// Experimental.
	Encodes *[]EncodeConfiguration `field:"required" json:"encodes" yaml:"encodes"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
}

