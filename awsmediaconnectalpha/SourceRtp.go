package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for RTP.
//
// Example:
//   var stack Stack
//
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_Rtp(&SourceRtp{
//   		FlowSourceName: jsii.String("rtp-source"),
//   		Port: jsii.Number(5000),
//   		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
//   	}),
//   })
//
// Experimental.
type SourceRtp struct {
	// A description of the source.
	//
	// This description appears only on the MediaConnect console and will not be seen by the end user.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the source.
	// Default: - a name is generated automatically.
	//
	// Experimental.
	FlowSourceName *string `field:"optional" json:"flowSourceName" yaml:"flowSourceName"`
	// Defines networking configuration.
	// Experimental.
	Network NetworkConfiguration `field:"required" json:"network" yaml:"network"`
	// The port that the flow will be listening on for incoming content.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The maximum bitrate for RIST, RTP, and RTP-FEC streams.
	// Default: - no maximum bitrate.
	//
	// Experimental.
	MaxBitrate awscdk.Bitrate `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
}

