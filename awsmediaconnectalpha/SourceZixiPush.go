package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for Zixi Push.
//
// No port option is exposed: MediaConnect assigns the Zixi Push ingest port itself —
// public sources are served on 2088, VPC sources are auto-assigned a port in 2090–2099.
// The service rejects any user-supplied port value, so the L2 surface doesn't accept one.
//
// Example:
//   var stack Stack
//
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_ZixiPush(&SourceZixiPush{
//   		FlowSourceName: jsii.String("zixi-source"),
//   		MaxLatency: awscdk.Duration_Millis(jsii.Number(2000)),
//   		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/mediaconnect/latest/ug/source-ports.html
//
// Experimental.
type SourceZixiPush struct {
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
	// Decrypt source with static keys.
	// Default: - no decryption.
	//
	// Experimental.
	Decryption *StaticKeyEncryption `field:"optional" json:"decryption" yaml:"decryption"`
	// The maximum latency in milliseconds for a Zixi-based source.
	// Default: - chosen by MediaConnect.
	//
	// Experimental.
	MaxLatency awscdk.Duration `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The stream ID that you want to use for the transport.
	//
	// This parameter applies only to Zixi-based streams.
	// Default: - no stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

