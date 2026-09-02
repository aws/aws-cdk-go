package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Common properties shared by the MediaPackage V2 output group variants.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var encodeConfiguration EncodeConfiguration
//   var id3Behavior Id3Behavior
//   var klvBehavior KlvBehavior
//   var mediaPackageV2Destination MediaPackageV2Destination
//   var mediaPackageV2HlsSetting MediaPackageV2HlsSetting
//   var nielsenId3Behavior NielsenId3Behavior
//   var scte35Type Scte35Type
//   var segment Segment
//   var timedMetadataId3Frame TimedMetadataId3Frame
//   var timedMetadataPassthrough TimedMetadataPassthrough
//
//   mediaPackageV2OutputGroupBaseProps := &MediaPackageV2OutputGroupBaseProps{
//   	Name: jsii.String("name"),
//   	Outputs: []MediaPackageV2OutputDefinition{
//   		&MediaPackageV2OutputDefinition{
//   			Encode: encodeConfiguration,
//   			OutputName: jsii.String("outputName"),
//
//   			// the properties below are optional
//   			AudioGroupId: jsii.String("audioGroupId"),
//   			AudioRenditionSets: jsii.String("audioRenditionSets"),
//   			Captions: []EncodeConfiguration{
//   				encodeConfiguration,
//   			},
//   			HlsAutoSelect: mediaPackageV2HlsSetting,
//   			HlsDefault: mediaPackageV2HlsSetting,
//   		},
//   	},
//
//   	// the properties below are optional
//   	AdditionalDestinations: []MediaPackageV2Destination{
//   		mediaPackageV2Destination,
//   	},
//   	CaptionLanguageMappings: []CaptionLanguageMapping{
//   		&CaptionLanguageMapping{
//   			CaptionChannel: jsii.Number(123),
//   			LanguageCode: jsii.String("languageCode"),
//   			LanguageDescription: jsii.String("languageDescription"),
//   		},
//   	},
//   	Id3Behavior: id3Behavior,
//   	KlvBehavior: klvBehavior,
//   	NielsenId3Behavior: nielsenId3Behavior,
//   	Scte35Type: scte35Type,
//   	Segment: segment,
//   	TimedMetadataId3Frame: timedMetadataId3Frame,
//   	TimedMetadataId3Period: cdk.Duration_Minutes(jsii.Number(30)),
//   	TimedMetadataPassthrough: timedMetadataPassthrough,
//   }
//
// See: https://docs.aws.amazon.com/medialive/latest/ug/creating-mediapackage-output-group.html
//
// Experimental.
type MediaPackageV2OutputGroupBaseProps struct {
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
}

