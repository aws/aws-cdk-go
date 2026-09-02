package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for creating a Router Output.
//
// Example:
//   var stack Stack
//   var networkInterface RouterNetworkInterface
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("SrtOutput"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("srt-output"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	// tier defaults to RouterOutputTier.OUTPUT_20 (lowest cost)
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
//   		Protocol: awsmediaconnectalpha.RouterOutputProtocol_SrtListener(&SrtListenerOutputProtocolProps{
//   			Port: jsii.Number(9001),
//   			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
//   		}),
//   		NetworkInterface: networkInterface,
//   	}),
//   })
//
// Experimental.
type RouterOutputProps struct {
	// Configuration for the Router Output (standard, MediaConnect Flow, or MediaLive).
	// Experimental.
	Configuration RouterOutputConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The maximum bitrate for the router output.
	// Experimental.
	MaximumBitrate awscdk.Bitrate `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
	// Indicates whether the router output is configured for Regional or global routing.
	// Experimental.
	RoutingScope RoutingScope `field:"required" json:"routingScope" yaml:"routingScope"`
	// Maintenance window configuration.
	// Default: - Default maintenance window will be used.
	//
	// Experimental.
	MaintenanceConfiguration *MaintenanceConfiguration `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// The AWS Region where the router output is located.
	// Default: - Defaults to the same region as stack.
	//
	// Experimental.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Name of the Router Output.
	// Default: - Generated automatically.
	//
	// Experimental.
	RouterOutputName *string `field:"optional" json:"routerOutputName" yaml:"routerOutputName"`
	// Tags to add to the Router Output.
	// Default: - No tagging.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Routing tier that determines the maximum bitrate (in Mbps) for this output.
	// Default: RouterOutputTier.OUTPUT_20
	//
	// Experimental.
	Tier RouterOutputTier `field:"optional" json:"tier" yaml:"tier"`
}

