package awsmedialivealpha


// Output definition for a MediaConnect Router output group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var encodeConfiguration EncodeConfiguration
//   var m2tsSettings M2tsSettings
//
//   mediaConnectRouterOutputDefinition := &MediaConnectRouterOutputDefinition{
//   	Encodes: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	M2tsSettings: m2tsSettings,
//   }
//
// Experimental.
type MediaConnectRouterOutputDefinition struct {
	// The encode configurations to wire to this output.
	// Experimental.
	Encodes *[]EncodeConfiguration `field:"required" json:"encodes" yaml:"encodes"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// MPEG-TS (M2TS) container settings for this output.
	// Default: - service defaults.
	//
	// Experimental.
	M2tsSettings M2tsSettings `field:"optional" json:"m2tsSettings" yaml:"m2tsSettings"`
}

