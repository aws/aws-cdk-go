package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Source priority configuration for failover Router Input configurations.
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
type SourcePriorityConfig interface {
}

// The jsii proxy struct for SourcePriorityConfig
type jsiiProxy_SourcePriorityConfig struct {
	_ byte // padding
}

// Treat both sources with equal priority — MediaConnect picks one when needed.
// Experimental.
func SourcePriorityConfig_None() SourcePriorityConfig {
	_init_.Initialize()

	var returns SourcePriorityConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourcePriorityConfig",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Designate one of the two sources as primary;
//
// the other is treated as backup.
// Experimental.
func SourcePriorityConfig_PrimarySecondary(primary PrimarySource) SourcePriorityConfig {
	_init_.Initialize()

	if err := validateSourcePriorityConfig_PrimarySecondaryParameters(primary); err != nil {
		panic(err)
	}
	var returns SourcePriorityConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.SourcePriorityConfig",
		"primarySecondary",
		[]interface{}{primary},
		&returns,
	)

	return returns
}

