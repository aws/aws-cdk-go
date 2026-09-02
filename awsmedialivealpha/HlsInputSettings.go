package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// HLS input settings for URL pull inputs.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   			NetworkInputSettings: &NetworkInputSettings{
//   				ServerValidation: medialive.ServerValidation_CHECK_CRYPTOGRAPHY_AND_VALIDATE_NAME(),
//   				HlsInputSettings: &HlsInputSettings{
//   					Bandwidth: awscdk.Bitrate_Mbps(jsii.Number(5)),
//   					Scte35Source: medialive.HlsScte35Source_MANIFEST(),
//   				},
//   			},
//   			LogicalInterfaceNames: []*string{
//   				jsii.String("eth0"),
//   				jsii.String("eth1"),
//   			},
//   		},
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type HlsInputSettings struct {
	// The bandwidth to select from the HLS manifest.
	//
	// MediaLive chooses the rendition
	// whose manifest bandwidth most closely matches this value.
	// Default: - highest bandwidth.
	//
	// Experimental.
	Bandwidth awscdk.Bitrate `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// When specified, reading of the HLS input begins this many buffer segments from the end (most recently written segment).
	// Default: - the HLS input begins with the first segment specified in the m3u8.
	//
	// Experimental.
	BufferSegments *float64 `field:"optional" json:"bufferSegments" yaml:"bufferSegments"`
	// Number of consecutive read failures before the input is considered unavailable.
	// Default: - service default.
	//
	// Experimental.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The interval between retry attempts.
	// Default: - service default.
	//
	// Experimental.
	RetryInterval awscdk.Duration `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// The source MediaLive ingests SCTE-35 messages from — the content segments or the manifest.
	// Default: - service default.
	//
	// Experimental.
	Scte35Source HlsScte35Source `field:"optional" json:"scte35Source" yaml:"scte35Source"`
}

