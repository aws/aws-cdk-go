package awsmediaconnectalpha


// Properties for failover Router Input configuration.
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
type FailoverConfigurationProps struct {
	// Network interface for the Router Input.
	// Experimental.
	NetworkInterface IRouterNetworkInterface `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// Array of exactly 2 protocol configurations for failover (must be same protocol type).
	// Experimental.
	Protocols *[]RouterInputProtocol `field:"required" json:"protocols" yaml:"protocols"`
	// Source priority configuration for failover.
	// Default: SourcePriorityConfig.none()
	//
	// Experimental.
	SourcePriority SourcePriorityConfig `field:"optional" json:"sourcePriority" yaml:"sourcePriority"`
}

