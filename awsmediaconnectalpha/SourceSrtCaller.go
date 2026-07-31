package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for SRT Caller.
//
// Example:
//   var stack Stack
//
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_SrtCaller(&SourceSrtCaller{
//   		FlowSourceName: jsii.String("remote-source"),
//   		SourceListenerAddress: jsii.String("203.0.113.50"),
//   		SourceListenerPort: jsii.Number(5000),
//   		MinLatency: awscdk.Duration_Millis(jsii.Number(200)),
//   	}),
//   })
//
// Experimental.
type SourceSrtCaller struct {
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
	// Source IP or domain name for SRT-caller protocol.
	// Experimental.
	SourceListenerAddress *string `field:"required" json:"sourceListenerAddress" yaml:"sourceListenerAddress"`
	// Source port for SRT-caller protocol.
	//
	// Valid range: 1024–65535. Ports 2077 and 2088 are reserved by MediaConnect for Zixi
	// traffic and cannot be used for SRT Caller.
	// Experimental.
	SourceListenerPort *float64 `field:"required" json:"sourceListenerPort" yaml:"sourceListenerPort"`
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
	// The maximum latency in milliseconds for SRT-based streams.
	// Default: - no maximum latency.
	//
	// Experimental.
	MaxLatency awscdk.Duration `field:"optional" json:"maxLatency" yaml:"maxLatency"`
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
	// The stream ID that you want to use for the transport.
	// Default: - no stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// Optional VPC interface for the outbound SRT Caller connection.
	//
	// SRT Caller initiates
	// the connection to the configured `sourceListenerAddress` and `sourceListenerPort`,
	// so no CIDR allow list is needed.
	// Default: - outbound connection via the public internet; no VPC interface.
	//
	// Experimental.
	VpcInterface *VpcInterfaceConfig `field:"optional" json:"vpcInterface" yaml:"vpcInterface"`
}

