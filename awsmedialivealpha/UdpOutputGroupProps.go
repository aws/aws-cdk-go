package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a UDP output group.
//
// Example:
//   var video EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Udp(&UdpOutputGroupProps{
//   	Name: jsii.String("udp"),
//   	Destinations: []UdpOutputDestination{
//   		medialive.UdpOutputDestination_Rtp(&TransportOutputDestinationProps{
//   			Address: jsii.String("203.0.113.5"),
//   			Port: jsii.Number(5000),
//   		}),
//   	},
//   	Outputs: []UdpOutputDefinition{
//   		&UdpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   			},
//   			OutputName: jsii.String("ts"),
//   			Fec: &FecOutputSettings{
//   				Mode: medialive.FecMode_COLUMN_AND_ROW(),
//   				ColumnDepth: jsii.Number(10),
//   				RowLength: jsii.Number(10),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type UdpOutputGroupProps struct {
	// The destinations for this output group — one per pipeline.
	//
	// Array position determines the pipeline mapping:
	// - `destinations[0]` → Pipeline 0
	// - `destinations[1]` → Pipeline 1 (STANDARD channels only)
	//
	// For a SINGLE_PIPELINE channel, provide exactly 1 destination.
	// For a STANDARD channel, provide exactly 2 destinations.
	// Experimental.
	Destinations *[]UdpOutputDestination `field:"required" json:"destinations" yaml:"destinations"`
	// The name of this output group.
	//
	// Used as the destination reference ID. Underscores are normalised to hyphens internally.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The output buffering.
	//
	// Applied at millisecond granularity.
	// Default: - service default.
	//
	// Experimental.
	Buffer awscdk.Duration `field:"optional" json:"buffer" yaml:"buffer"`
	// Action to take when the input is lost.
	// Default: UdpInputLossAction.EMIT_PROGRAM
	//
	// Experimental.
	InputLossAction UdpInputLossAction `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// The outputs for this UDP output group.
	// Default: - no initial outputs.
	//
	// Experimental.
	Outputs *[]*UdpOutputDefinition `field:"optional" json:"outputs" yaml:"outputs"`
	// Indicates the ID3 frame that has the timecode.
	// Default: UdpTimedMetadataId3Frame.PRIV
	//
	// Experimental.
	TimedMetadataId3Frame UdpTimedMetadataId3Frame `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval.
	// Default: - Duration.seconds(10)
	//
	// Experimental.
	TimedMetadataId3Period awscdk.Duration `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
}

