package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// Properties for MediaLive Channel Router Input configuration.
//
// Use this when the MediaLive channel already exists and you want to ingest
// from one of its outputs immediately.
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
type MediaLiveChannelConfigurationProps struct {
	// The MediaLive channel to use as input.
	// Experimental.
	Channel interfacesawsmedialive.IChannelRef `field:"required" json:"channel" yaml:"channel"`
	// The name of the individual output (within the channel's MediaConnect Router output group) to connect to this router input — not the name of the output group itself.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The MediaLive pipeline to connect to this router input.
	// Experimental.
	Pipeline MediaLivePipeline `field:"required" json:"pipeline" yaml:"pipeline"`
	// Optional transit encryption configuration.
	//
	// Must match the encryption type configured on the
	// MediaLive channel's MediaConnect Router output group.
	// Default: - automatic encryption.
	//
	// Experimental.
	SourceTransitDecryption *TransitEncryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

