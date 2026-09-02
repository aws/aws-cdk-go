package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Transit encryption configuration for router integrations — securing the link between a router and a flow or a MediaLive channel/input.
//
// Uses AWS Secrets Manager for key
// management.
//
// The secret must live in the same AWS account and Region as the consuming resource.
// MediaConnect does not support cross-account or cross-Region secrets.
//
// **Trust-policy scope on routers.** Router I/O ids are service-generated (unknown at synth
// time), and pinning the live ARN would create a CloudFormation dependency cycle — so the
// auto-created role pins `aws:SourceArn` to a wildcarded ARN (`arn:...:routerInput:*` /
// `arn:...:routerOutput:*`) plus `aws:SourceAccount`. To pin a tighter trust policy, supply
// your own `role`.
//
// Example:
//   var stack Stack
//   var mediaLiveChannel IChannel
//   var transitSecret Secret
//   // must hold the same value as the channel's MediaConnectRouterSettings.shared() secret
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
//   		SourceTransitDecryption: &TransitEncryption{
//   			Secret: transitSecret,
//   		},
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/mediaconnect/latest/ug/cross-service-confused-deputy-prevention.html
//
// Experimental.
type TransitEncryption struct {
	// Secrets Manager secret containing the transit encryption key.
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
	// IAM role that MediaConnect assumes to access the Secrets Manager secret.
	//
	// If provided, the role is used as-is; you must grant it the necessary permissions
	// yourself.
	// Default: - a scoped role is auto-created with read access to the secret and a
	// confused-deputy trust condition. See the **Encryption** section of the module README
	// for the generated trust policy.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

