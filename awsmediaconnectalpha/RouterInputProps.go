package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for creating a Router Input.
//
// Example:
//   var stack Stack
//   var mediaLiveChannel IChannel
//
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("ChannelInput"), &RouterInputProps{
//   	RouterInputName: jsii.String("channel-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaLiveChannel(&MediaLiveChannelConfigurationProps{
//   		Channel: mediaLiveChannel,
//   		OutputName: jsii.String("router-ts"),
//   		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
//   	}),
//   })
//
// Experimental.
type RouterInputProps struct {
	// Configuration for the Router Input (standard, failover, merge, or MediaConnect flow).
	// Experimental.
	Configuration RouterInputConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The maximum bitrate for the router input.
	// Experimental.
	MaximumBitrate awscdk.Bitrate `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
	// Indicates whether the router input is configured for Regional or global routing.
	// Experimental.
	RoutingScope RoutingScope `field:"required" json:"routingScope" yaml:"routingScope"`
	// Maintenance window configuration.
	// Default: - Default maintenance window will be used.
	//
	// Experimental.
	MaintenanceConfiguration *MaintenanceConfiguration `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// The AWS Region where the router input is located.
	// Default: - Defaults to the same region as stack.
	//
	// Experimental.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Name of the Router Input.
	// Default: - Generated automatically.
	//
	// Experimental.
	RouterInputName *string `field:"optional" json:"routerInputName" yaml:"routerInputName"`
	// Tags to add to the Router Input.
	// Default: - No tagging.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Select a tier based on your maximum bitrate requirements.
	// Default: RouterInputTier.INPUT_20
	//
	// Experimental.
	Tier RouterInputTier `field:"optional" json:"tier" yaml:"tier"`
	// Transit encryption configuration using AWS Secrets Manager.
	//
	// When provided without a role, a scoped IAM role is automatically created with read
	// access to the secret.
	// Default: - Automatic encryption will be configured.
	//
	// Experimental.
	TransitEncryption *TransitEncryption `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
}

