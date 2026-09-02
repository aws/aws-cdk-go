package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Output definition for an SRT output group.
//
// `outputName` is normalised (underscores → hyphens) for use as the destination reference ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encodeConfiguration EncodeConfiguration
//   var m2tsSettings M2tsSettings
//   var srtDestination SrtDestination
//   var srtEncryptionType SrtEncryptionType
//
//   srtOutputDefinition := &SrtOutputDefinition{
//   	Destinations: []SrtDestination{
//   		srtDestination,
//   	},
//   	Encodes: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	Buffer: cdk.Duration_Minutes(jsii.Number(30)),
//   	EncryptionType: srtEncryptionType,
//   	Latency: cdk.Duration_*Minutes(jsii.Number(30)),
//   	M2tsSettings: m2tsSettings,
//   }
//
// Experimental.
type SrtOutputDefinition struct {
	// The encode configurations to wire to this output.
	// Experimental.
	Encodes *[]EncodeConfiguration `field:"required" json:"encodes" yaml:"encodes"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The SRT destination(s) for this output — one per channel pipeline.
	//
	// MediaLive publishes each SRT output to one destination per pipeline (the console
	// calls these "Destination A" and "Destination B"). Provide a single destination for a
	// `SINGLE_PIPELINE` channel, or two (A then B) for a `STANDARD` channel. Each pipeline
	// can target a different listener and carry its own encryption passphrase.
	// Experimental.
	Destinations *[]SrtDestination `field:"required" json:"destinations" yaml:"destinations"`
	// The output buffering.
	//
	// Applied at millisecond granularity.
	// Default: - service default.
	//
	// Experimental.
	Buffer awscdk.Duration `field:"optional" json:"buffer" yaml:"buffer"`
	// The encryption type for the SRT output.
	// Default: - no encryption.
	//
	// Experimental.
	EncryptionType SrtEncryptionType `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// The SRT latency.
	// Default: - service default.
	//
	// Experimental.
	Latency awscdk.Duration `field:"optional" json:"latency" yaml:"latency"`
	// MPEG-TS (M2TS) container settings for this output.
	// Default: - service defaults.
	//
	// Experimental.
	M2tsSettings M2tsSettings `field:"optional" json:"m2tsSettings" yaml:"m2tsSettings"`
}

