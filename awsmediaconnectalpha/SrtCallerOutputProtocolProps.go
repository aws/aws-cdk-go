package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for SRT Caller protocol configuration for outputs.
//
// Example:
//   var stack Stack
//   var networkInterface RouterNetworkInterface
//   var role IRole
//   var secret ISecret
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("EncryptedOutput"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("encrypted-output"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
//   		Protocol: awsmediaconnectalpha.RouterOutputProtocol_SrtCaller(&SrtCallerOutputProtocolProps{
//   			DestinationAddress: jsii.String("203.0.113.100"),
//   			DestinationPort: jsii.Number(9001),
//   			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
//   			EncryptionConfiguration: &RouterSrtEncryption{
//   				Role: *Role,
//   				Secret: *Secret,
//   			},
//   		}),
//   		NetworkInterface: networkInterface,
//   	}),
//   })
//
// Experimental.
type SrtCallerOutputProtocolProps struct {
	// Destination IP address to connect to.
	// Experimental.
	DestinationAddress *string `field:"required" json:"destinationAddress" yaml:"destinationAddress"`
	// Destination port to connect to.
	// Experimental.
	DestinationPort *float64 `field:"required" json:"destinationPort" yaml:"destinationPort"`
	// Minimum latency for SRT.
	// Experimental.
	MinimumLatency awscdk.Duration `field:"required" json:"minimumLatency" yaml:"minimumLatency"`
	// Optional encryption configuration for SRT streams.
	// Default: - No encryption.
	//
	// Experimental.
	EncryptionConfiguration *RouterSrtEncryption `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Optional stream ID for SRT connection.
	// Default: - No stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

