package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a CMAF Ingest output group.
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
//   var nielsenId3Behavior NielsenId3Behavior
//   var outputDestination OutputDestination
//   var scte35Type Scte35Type
//   var segment Segment
//   var timedMetadataId3Frame TimedMetadataId3Frame
//   var timedMetadataPassthrough TimedMetadataPassthrough
//
//   cmafIngestOutputGroupProps := &CmafIngestOutputGroupProps{
//   	Destinations: []OutputDestination{
//   		outputDestination,
//   	},
//   	Name: jsii.String("name"),
//   	Outputs: []CmafIngestOutputDefinition{
//   		&CmafIngestOutputDefinition{
//   			Encode: encodeConfiguration,
//   			OutputName: jsii.String("outputName"),
//
//   			// the properties below are optional
//   			Captions: []EncodeConfiguration{
//   				encodeConfiguration,
//   			},
//   			NameModifier: jsii.String("nameModifier"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AdditionalDestinations: []OutputDestination{
//   		outputDestination,
//   	},
//   	CaptionLanguageMappings: []CmafCaptionLanguageMapping{
//   		&CmafCaptionLanguageMapping{
//   			CaptionChannel: jsii.Number(123),
//   			LanguageCode: jsii.String("languageCode"),
//   		},
//   	},
//   	Id3Behavior: id3Behavior,
//   	Id3NameModifier: jsii.String("id3NameModifier"),
//   	KlvBehavior: klvBehavior,
//   	KlvNameModifier: jsii.String("klvNameModifier"),
//   	NielsenId3Behavior: nielsenId3Behavior,
//   	NielsenId3NameModifier: jsii.String("nielsenId3NameModifier"),
//   	Scte35NameModifier: jsii.String("scte35NameModifier"),
//   	Scte35Type: scte35Type,
//   	Segment: segment,
//   	SendDelayMs: jsii.Number(123),
//   	TimedMetadataId3Frame: timedMetadataId3Frame,
//   	TimedMetadataId3Period: cdk.Duration_Minutes(jsii.Number(30)),
//   	TimedMetadataPassthrough: timedMetadataPassthrough,
//   }
//
// Experimental.
type CmafIngestOutputGroupProps struct {
	// The primary CMAF ingest destinations — one per pipeline.
	//
	// Array position determines the pipeline mapping:
	// - `destinations[0]` → Pipeline 0
	// - `destinations[1]` → Pipeline 1 (STANDARD channels only)
	//
	// For a SINGLE_PIPELINE channel, provide exactly 1 destination.
	// For a STANDARD channel, provide exactly 2 destinations to utilise both
	// pipelines for redundancy.
	// Experimental.
	Destinations *[]OutputDestination `field:"required" json:"destinations" yaml:"destinations"`
	// The name of this output group.
	//
	// Used as the destination reference ID. Underscores are normalised to hyphens internally.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The outputs for this CMAF Ingest output group.
	//
	// Each output should contain a single encode.
	// Experimental.
	Outputs *[]*CmafIngestOutputDefinition `field:"required" json:"outputs" yaml:"outputs"`
	// Configure additional destinations to fan out the CMAF ingest output to extra endpoints, for example for cross-region delivery or backup packaging.
	//
	// Standard channels support up to 2 additional destinations.
	// Single pipeline channels support 1 additional destination.
	// Default: - no additional destinations.
	//
	// Experimental.
	AdditionalDestinations *[]OutputDestination `field:"optional" json:"additionalDestinations" yaml:"additionalDestinations"`
	// Maps captions channels to languages for this CMAF Ingest output group.
	// Default: - no caption language mappings.
	//
	// Experimental.
	CaptionLanguageMappings *[]*CmafCaptionLanguageMapping `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// The ID3 behavior for the CMAF ingest output.
	// Default: Id3Behavior.DISABLED
	//
	// Experimental.
	Id3Behavior Id3Behavior `field:"optional" json:"id3Behavior" yaml:"id3Behavior"`
	// The name modifier for ID3 metadata.
	// Default: - service default.
	//
	// Experimental.
	Id3NameModifier *string `field:"optional" json:"id3NameModifier" yaml:"id3NameModifier"`
	// The KLV behavior for the CMAF ingest output.
	// Default: KlvBehavior.NO_PASSTHROUGH
	//
	// Experimental.
	KlvBehavior KlvBehavior `field:"optional" json:"klvBehavior" yaml:"klvBehavior"`
	// The name modifier for KLV metadata.
	// Default: - service default.
	//
	// Experimental.
	KlvNameModifier *string `field:"optional" json:"klvNameModifier" yaml:"klvNameModifier"`
	// The Nielsen ID3 behavior for the CMAF ingest output.
	// Default: NielsenId3Behavior.NO_PASSTHROUGH
	//
	// Experimental.
	NielsenId3Behavior NielsenId3Behavior `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The name modifier for Nielsen ID3 metadata.
	// Default: - service default.
	//
	// Experimental.
	NielsenId3NameModifier *string `field:"optional" json:"nielsenId3NameModifier" yaml:"nielsenId3NameModifier"`
	// The name modifier for SCTE-35 messages.
	// Default: - service default.
	//
	// Experimental.
	Scte35NameModifier *string `field:"optional" json:"scte35NameModifier" yaml:"scte35NameModifier"`
	// The SCTE-35 type for the CMAF ingest output.
	// Default: Scte35Type.SCTE_35_WITHOUT_SEGMENTATION
	//
	// Experimental.
	Scte35Type Scte35Type `field:"optional" json:"scte35Type" yaml:"scte35Type"`
	// The length of each media segment.
	// Default: - Segment.seconds(1)
	//
	// Experimental.
	Segment Segment `field:"optional" json:"segment" yaml:"segment"`
	// The number of milliseconds to delay the output from the second pipeline.
	// Default: - service default.
	//
	// Experimental.
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// Indicates the ID3 frame that has the timecode.
	// Default: TimedMetadataId3Frame.NONE
	//
	// Experimental.
	TimedMetadataId3Frame TimedMetadataId3Frame `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval.
	// Default: - Duration.seconds(10)
	//
	// Experimental.
	TimedMetadataId3Period awscdk.Duration `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// Whether timed metadata is passed through.
	// Default: TimedMetadataPassthrough.DISABLED
	//
	// Experimental.
	TimedMetadataPassthrough TimedMetadataPassthrough `field:"optional" json:"timedMetadataPassthrough" yaml:"timedMetadataPassthrough"`
}

