package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for SRT Listener.
//
// Example:
//   var stack Stack
//
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_SrtListener(&SourceSrtListener{
//   		FlowSourceName: jsii.String("live-encoder-source"),
//   		Description: jsii.String("Live encoder feed"),
//   		Port: jsii.Number(5000),
//   		MinLatency: awscdk.Duration_Millis(jsii.Number(2000)),
//   		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
//   	}),
//   })
//
// Experimental.
type SourceSrtListener struct {
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
	// The port that the flow listens on for incoming content.
	//
	// Valid range: 1024–65535. Ports 2077 and 2088 are reserved by MediaConnect for Zixi
	// traffic and cannot be used for SRT Listener.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// SRT Decryption options.
	// Default: - no decryption.
	//
	// Experimental.
	Decryption *SrtPasswordEncryption `field:"optional" json:"decryption" yaml:"decryption"`
	// The maximum bitrate for streams.
	// Default: - no maximum bitrate.
	//
	// Experimental.
	MaxBitrate awscdk.Bitrate `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// The value you set on your
	// MediaConnect source represents the minimal potential latency of that connection. The
	// latency of the stream is set to the higher of the sender's minimum latency and the
	// receiver's minimum latency.
	// Default: - no minimum latency.
	//
	// Experimental.
	MinLatency awscdk.Duration `field:"optional" json:"minLatency" yaml:"minLatency"`
}

