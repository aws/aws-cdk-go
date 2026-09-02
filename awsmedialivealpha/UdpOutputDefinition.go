package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Output definition for a UDP output group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var encodeConfiguration EncodeConfiguration
//   var fecMode FecMode
//   var m2tsSettings M2tsSettings
//
//   udpOutputDefinition := &UdpOutputDefinition{
//   	Encodes: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	Buffer: cdk.Duration_Minutes(jsii.Number(30)),
//   	Fec: &FecOutputSettings{
//   		ColumnDepth: jsii.Number(123),
//   		Mode: fecMode,
//   		RowLength: jsii.Number(123),
//   	},
//   	M2tsSettings: m2tsSettings,
//   }
//
// Experimental.
type UdpOutputDefinition struct {
	// The encode configurations to wire to this output.
	// Experimental.
	Encodes *[]EncodeConfiguration `field:"required" json:"encodes" yaml:"encodes"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The output buffering (overrides group-level setting).
	//
	// Applied at millisecond granularity.
	// Default: - uses group-level buffer.
	//
	// Experimental.
	Buffer awscdk.Duration `field:"optional" json:"buffer" yaml:"buffer"`
	// Forward Error Correction (FEC) settings for this output.
	// Default: - no FEC.
	//
	// Experimental.
	Fec *FecOutputSettings `field:"optional" json:"fec" yaml:"fec"`
	// MPEG-TS (M2TS) container settings for this output.
	// Default: - service defaults.
	//
	// Experimental.
	M2tsSettings M2tsSettings `field:"optional" json:"m2tsSettings" yaml:"m2tsSettings"`
}

