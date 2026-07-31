package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for RIST protocol configuration.
//
// Example:
//   var stack Stack
//   var networkInterface RouterNetworkInterface
//
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("FailoverInput"), &RouterInputProps{
//   	RouterInputName: jsii.String("failover-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_Failover(&FailoverConfigurationProps{
//   		NetworkInterface: networkInterface,
//   		Protocols: []RouterInputProtocol{
//   			awsmediaconnectalpha.RouterInputProtocol_Rist(&RistProtocolProps{
//   				Port: jsii.Number(5000),
//   				RecoveryLatency: awscdk.Duration_Millis(jsii.Number(1000)),
//   			}),
//   			awsmediaconnectalpha.RouterInputProtocol_*Rist(&RistProtocolProps{
//   				Port: jsii.Number(5002),
//   				 // Must not be consecutive with primary port
//   				RecoveryLatency: awscdk.Duration_*Millis(jsii.Number(1000)),
//   			}),
//   		},
//   		SourcePriority: awsmediaconnectalpha.SourcePriorityConfig_PrimarySecondary(awsmediaconnectalpha.PrimarySource_FIRST_SOURCE),
//   	}),
//   })
//
// Experimental.
type RistProtocolProps struct {
	// Port number for RIST traffic.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Recovery latency for RIST.
	// Experimental.
	RecoveryLatency awscdk.Duration `field:"required" json:"recoveryLatency" yaml:"recoveryLatency"`
}

