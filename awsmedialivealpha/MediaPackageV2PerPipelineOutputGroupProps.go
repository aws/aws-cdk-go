package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a MediaPackage V2 output group with explicit per-pipeline destinations.
//
// Use this when you need to control the channel and endpoint each pipeline delivers to — for
// example cross-region delivery, or pinning pipeline 0 to a specific endpoint.
//
// Example:
//   var primary IChannel
//   var hdVideo EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_MediaPackageV2PerPipeline(&MediaPackageV2PerPipelineOutputGroupProps{
//   	Name: jsii.String("emp"),
//   	Destinations: []MediaPackageV2Destination{
//   		medialive.MediaPackageV2Destination_Channel(primary, medialive.MediaPackageV2EndpointId_ENDPOINT_2()),
//   		medialive.MediaPackageV2Destination_*Channel(primary, medialive.MediaPackageV2EndpointId_ENDPOINT_1()),
//   	},
//   	Outputs: []MediaPackageV2OutputDefinition{
//   		&MediaPackageV2OutputDefinition{
//   			Encode: hdVideo,
//   			OutputName: jsii.String("hd"),
//   		},
//   	},
//   })
//
// Experimental.
type MediaPackageV2PerPipelineOutputGroupProps struct {
	// The name of this output group.
	//
	// Used as the destination reference ID. Underscores are normalised to hyphens internally.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The outputs for this output group.
	//
	// MediaPackage V2 uses CMAF ingest which requires one track per output.
	// Create a separate output for each encode (e.g. one for HD video, one for SD video, one for audio).
	// Do NOT put multiple encodes in a single output.
	// Experimental.
	Outputs *[]*MediaPackageV2OutputDefinition `field:"required" json:"outputs" yaml:"outputs"`
	// Configure additional destinations to fan out the output to extra MediaPackage V2 channels, for example for cross-region delivery or backup packaging.
	//
	// These correspond to Destination 3/4 in the AWS console. Each additional destination is a
	// single, explicit entry (channel + endpoint), independent of the channel class.
	// See: https://docs.aws.amazon.com/medialive/latest/ug/creating-mediapackage-output-group.html
	//
	// Default: - no additional destinations.
	//
	// Experimental.
	AdditionalDestinations *[]MediaPackageV2Destination `field:"optional" json:"additionalDestinations" yaml:"additionalDestinations"`
	// Caption language mappings for the MediaPackage V2 output.
	// Default: - no caption language mappings.
	//
	// Experimental.
	CaptionLanguageMappings *[]*CaptionLanguageMapping `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// The ID3 behavior.
	// Default: Id3Behavior.DISABLED
	//
	// Experimental.
	Id3Behavior Id3Behavior `field:"optional" json:"id3Behavior" yaml:"id3Behavior"`
	// The KLV behavior.
	// Default: KlvBehavior.NO_PASSTHROUGH
	//
	// Experimental.
	KlvBehavior KlvBehavior `field:"optional" json:"klvBehavior" yaml:"klvBehavior"`
	// The Nielsen ID3 behavior.
	// Default: NielsenId3Behavior.NO_PASSTHROUGH
	//
	// Experimental.
	NielsenId3Behavior NielsenId3Behavior `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The SCTE-35 type.
	// Default: Scte35Type.SCTE_35_WITHOUT_SEGMENTATION
	//
	// Experimental.
	Scte35Type Scte35Type `field:"optional" json:"scte35Type" yaml:"scte35Type"`
	// The length of each media segment.
	// Default: - Segment.seconds(1)
	//
	// Experimental.
	Segment Segment `field:"optional" json:"segment" yaml:"segment"`
	// The timed metadata ID3 frame.
	// Default: TimedMetadataId3Frame.NONE
	//
	// Experimental.
	TimedMetadataId3Frame TimedMetadataId3Frame `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval.
	// Default: Duration.seconds(10)
	//
	// Experimental.
	TimedMetadataId3Period awscdk.Duration `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// Whether timed metadata is passed through.
	// Default: TimedMetadataPassthrough.DISABLED
	//
	// Experimental.
	TimedMetadataPassthrough TimedMetadataPassthrough `field:"optional" json:"timedMetadataPassthrough" yaml:"timedMetadataPassthrough"`
	// The primary MediaPackage V2 destinations — one per pipeline.
	//
	// Array position determines the pipeline mapping:
	// - `destinations[0]` → Pipeline 0
	// - `destinations[1]` → Pipeline 1 (STANDARD channels only)
	//
	// For a SINGLE_PIPELINE channel, provide exactly 1 destination.
	// For a STANDARD channel, provide exactly 2 destinations.
	// Experimental.
	Destinations *[]MediaPackageV2Destination `field:"required" json:"destinations" yaml:"destinations"`
}

