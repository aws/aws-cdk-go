package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Output definition for an RTMP output group.
//
// `outputName` is normalised (underscores → hyphens) for use as the destination reference ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var encodeConfiguration EncodeConfiguration
//   var rtmpCertificateMode RtmpCertificateMode
//   var rtmpDestination RtmpDestination
//
//   rtmpOutputDefinition := &RtmpOutputDefinition{
//   	Destinations: []RtmpDestination{
//   		rtmpDestination,
//   	},
//   	Encodes: []EncodeConfiguration{
//   		encodeConfiguration,
//   	},
//   	OutputName: jsii.String("outputName"),
//
//   	// the properties below are optional
//   	CertificateMode: rtmpCertificateMode,
//   	ConnectionRetryInterval: cdk.Duration_Minutes(jsii.Number(30)),
//   	NumRetries: jsii.Number(123),
//   }
//
// Experimental.
type RtmpOutputDefinition struct {
	// The encode configurations to wire to this output.
	// Experimental.
	Encodes *[]EncodeConfiguration `field:"required" json:"encodes" yaml:"encodes"`
	// The name of this output.
	//
	// Must be unique across all outputs in the channel.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The RTMP destination(s) for this output — one per channel pipeline.
	//
	// MediaLive publishes each RTMP output to one destination per pipeline (the console
	// calls these "Destination A" and "Destination B"). Provide a single destination for a
	// `SINGLE_PIPELINE` channel, or two (A then B) for a `STANDARD` channel.
	// Experimental.
	Destinations *[]RtmpDestination `field:"required" json:"destinations" yaml:"destinations"`
	// The TLS certificate verification mode.
	// Default: - service default.
	//
	// Experimental.
	CertificateMode RtmpCertificateMode `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The interval between connection retry attempts.
	// Default: - service default.
	//
	// Experimental.
	ConnectionRetryInterval awscdk.Duration `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The number of retry attempts.
	// Default: - service default.
	//
	// Experimental.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
}

