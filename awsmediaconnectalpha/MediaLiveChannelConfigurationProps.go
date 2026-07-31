package awsmediaconnectalpha


// Properties for MediaLive Channel Router Input configuration.
//
// Use this when the MediaLive channel already exists and you want to ingest
// from one of its outputs immediately.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//   var secret Secret
//
//   mediaLiveChannelConfigurationProps := &MediaLiveChannelConfigurationProps{
//   	MediaLiveChannelArn: jsii.String("mediaLiveChannelArn"),
//   	MediaLiveChannelOutputName: jsii.String("mediaLiveChannelOutputName"),
//   	MediaLivePipelineId: mediaconnect_alpha.MediaLivePipeline_PIPELINE_0,
//
//   	// the properties below are optional
//   	SourceTransitDecryption: &TransitEncryption{
//   		Secret: secret,
//
//   		// the properties below are optional
//   		Role: role,
//   	},
//   }
//
// Experimental.
type MediaLiveChannelConfigurationProps struct {
	// ARN of the MediaLive channel to use as input.
	//
	// Note: This will change to accept a typed MediaLive channel reference
	// when the.
	// Experimental.
	MediaLiveChannelArn *string `field:"required" json:"mediaLiveChannelArn" yaml:"mediaLiveChannelArn"`
	// The name of the MediaLive channel output to connect to this router input.
	// Experimental.
	MediaLiveChannelOutputName *string `field:"required" json:"mediaLiveChannelOutputName" yaml:"mediaLiveChannelOutputName"`
	// The MediaLive pipeline to connect to this router input.
	// Experimental.
	MediaLivePipelineId MediaLivePipeline `field:"required" json:"mediaLivePipelineId" yaml:"mediaLivePipelineId"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	SourceTransitDecryption *TransitEncryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

