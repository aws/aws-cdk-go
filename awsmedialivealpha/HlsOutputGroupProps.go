package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for an HLS output group.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   		},
//   	},
//   	AvailBlanking: &AvailBlanking{
//   		State: medialive.AvailBlankingState_ENABLED(),
//   		Image: medialive.FileLocation_FromBucket(bucket, jsii.String("slates/avail.png")),
//   	},
//   	BlackoutSlate: &BlackoutSlate{
//   		State: medialive.BlackoutSlateState_ENABLED(),
//   		Image: medialive.FileLocation_*FromBucket(bucket, jsii.String("slates/blackout.png")),
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type HlsOutputGroupProps struct {
	// The destinations for this output group — one per pipeline.
	//
	// Array position determines the pipeline mapping:
	// - `destinations[0]` → Pipeline 0
	// - `destinations[1]` → Pipeline 1 (STANDARD channels only)
	//
	// For a SINGLE_PIPELINE channel, provide exactly 1 destination.
	// For a STANDARD channel, provide exactly 2 destinations.
	// Experimental.
	Destinations *[]OutputDestination `field:"required" json:"destinations" yaml:"destinations"`
	// The name of this output group.
	//
	// Used as the destination reference ID. Underscores are normalised to hyphens internally.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Chooses one or more ad marker types to pass SCTE35 signals through to this group of Apple HLS outputs.
	// Default: - no ad markers.
	//
	// Experimental.
	AdMarkers *[]HlsAdMarkers `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// A partial URI prefix that will be prepended to each output in the media .m3u8 file.
	// Default: - no base URL content prefix.
	//
	// Experimental.
	BaseUrlContent *string `field:"optional" json:"baseUrlContent" yaml:"baseUrlContent"`
	// Optional base URL content for pipeline 1 if different from pipeline 0.
	// Default: - no base URL content 1.
	//
	// Experimental.
	BaseUrlContent1 *string `field:"optional" json:"baseUrlContent1" yaml:"baseUrlContent1"`
	// A partial URI prefix that will be prepended to each output in the media .m3u8 file for the manifest.
	// Default: - no base URL manifest prefix.
	//
	// Experimental.
	BaseUrlManifest *string `field:"optional" json:"baseUrlManifest" yaml:"baseUrlManifest"`
	// Optional base URL manifest for pipeline 1 if different from pipeline 0.
	// Default: - no base URL manifest 1.
	//
	// Experimental.
	BaseUrlManifest1 *string `field:"optional" json:"baseUrlManifest1" yaml:"baseUrlManifest1"`
	// A mapping of up to 4 captions channels to captions languages.
	//
	// Meaningful only if captionLanguageSetting is set to INSERT.
	// Default: - no caption language mappings.
	//
	// Experimental.
	CaptionLanguageMappings *[]*CaptionLanguageMapping `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// Applies only to 608 embedded output captions.
	// Default: - service default.
	//
	// Experimental.
	CaptionLanguageSetting HlsCaptionLanguageSetting `field:"optional" json:"captionLanguageSetting" yaml:"captionLanguageSetting"`
	// When set to DISABLED, sets the #EXT-X-ALLOW-CACHE:no tag in the manifest.
	// Default: HlsClientCache.ENABLED
	//
	// Experimental.
	ClientCache HlsClientCache `field:"optional" json:"clientCache" yaml:"clientCache"`
	// The specification to use (RFC-6381 or the default RFC-4281) during m3u8 playlist generation.
	// Default: HlsCodecSpecification.RFC_4281
	//
	// Experimental.
	CodecSpecification HlsCodecSpecification `field:"optional" json:"codecSpecification" yaml:"codecSpecification"`
	// A 128-bit, 16-byte hex value represented by a 32-character text string used as the IV for encryption.
	//
	// Used with encryptionType when ivSource is set to EXPLICIT.
	// Default: - no constant IV.
	//
	// Experimental.
	ConstantIv *string `field:"optional" json:"constantIv" yaml:"constantIv"`
	// Places segments in subdirectories.
	// Default: HlsDirectoryStructure.SINGLE_DIRECTORY
	//
	// Experimental.
	DirectoryStructure HlsDirectoryStructure `field:"optional" json:"directoryStructure" yaml:"directoryStructure"`
	// Specifies whether to insert EXT-X-DISCONTINUITY tags in the HLS child manifests.
	// Default: HlsDiscontinuityTags.INSERT
	//
	// Experimental.
	DiscontinuityTags HlsDiscontinuityTags `field:"optional" json:"discontinuityTags" yaml:"discontinuityTags"`
	// Encrypts the segments with the specified encryption scheme.
	// Default: - no encryption.
	//
	// Experimental.
	EncryptionType HlsEncryptionType `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Settings to configure the CDN for the HLS output.
	// Default: - service default.
	//
	// Experimental.
	HlsCdnSettings HlsCdnSettings `field:"optional" json:"hlsCdnSettings" yaml:"hlsCdnSettings"`
	// State of HLS ID3 Segment Tagging.
	// Default: HlsId3SegmentTaggingState.DISABLED
	//
	// Experimental.
	HlsId3SegmentTagging HlsId3SegmentTaggingState `field:"optional" json:"hlsId3SegmentTagging" yaml:"hlsId3SegmentTagging"`
	// Whether to create an I-frame-only manifest.
	// Default: HlsIFrameOnlyPlaylists.DISABLED
	//
	// Experimental.
	IFrameOnlyPlaylists HlsIFrameOnlyPlaylists `field:"optional" json:"iFrameOnlyPlaylists" yaml:"iFrameOnlyPlaylists"`
	// Specifies whether to include the final (incomplete) segment in the media output.
	// Default: HlsIncompleteSegmentBehavior.AUTO
	//
	// Experimental.
	IncompleteSegmentBehavior HlsIncompleteSegmentBehavior `field:"optional" json:"incompleteSegmentBehavior" yaml:"incompleteSegmentBehavior"`
	// The maximum number of segments in the media manifest (LIVE mode only).
	// Default: 10.
	//
	// Experimental.
	IndexNSegments *float64 `field:"optional" json:"indexNSegments" yaml:"indexNSegments"`
	// Action to take when the input is lost.
	// Default: HlsInputLossAction.EMIT_OUTPUT
	//
	// Experimental.
	InputLossAction HlsInputLossAction `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Whether the IV is listed in the manifest.
	// Default: - service default.
	//
	// Experimental.
	IvInManifest HlsIvInManifest `field:"optional" json:"ivInManifest" yaml:"ivInManifest"`
	// Whether the IV follows the segment number or is explicit.
	// Default: - service default.
	//
	// Experimental.
	IvSource HlsIvSource `field:"optional" json:"ivSource" yaml:"ivSource"`
	// The number of segments to retain in the destination directory (LIVE mode only).
	// Default: 21.
	//
	// Experimental.
	KeepSegments *float64 `field:"optional" json:"keepSegments" yaml:"keepSegments"`
	// Specifies how the key is represented in the resource identified by the URI.
	// Default: - service default.
	//
	// Experimental.
	KeyFormat *string `field:"optional" json:"keyFormat" yaml:"keyFormat"`
	// Either a single positive integer version value or a slash-delimited list of version values (1/2/3).
	// Default: - service default.
	//
	// Experimental.
	KeyFormatVersions *string `field:"optional" json:"keyFormatVersions" yaml:"keyFormatVersions"`
	// The key provider settings for HLS encryption.
	// Default: - no key provider settings.
	//
	// Experimental.
	KeyProviderSettings HlsKeyProviderSettings `field:"optional" json:"keyProviderSettings" yaml:"keyProviderSettings"`
	// When set to GZIP, compresses HLS playlist.
	// Default: HlsManifestCompression.NONE
	//
	// Experimental.
	ManifestCompression HlsManifestCompression `field:"optional" json:"manifestCompression" yaml:"manifestCompression"`
	// Indicates whether the output manifest should use floating point or integer values for segment duration.
	// Default: HlsManifestDurationFormat.FLOATING_POINT
	//
	// Experimental.
	ManifestDurationFormat HlsManifestDurationFormat `field:"optional" json:"manifestDurationFormat" yaml:"manifestDurationFormat"`
	// The minimum segment length.
	//
	// HLS supports whole-second segments only.
	// Default: - service default.
	//
	// Experimental.
	MinSegment Segment `field:"optional" json:"minSegment" yaml:"minSegment"`
	// The output mode — LIVE or VOD.
	// Default: HlsMode.LIVE
	//
	// Experimental.
	Mode HlsMode `field:"optional" json:"mode" yaml:"mode"`
	// The outputs for this HLS output group.
	// Default: - no initial outputs.
	//
	// Experimental.
	Outputs *[]*HlsOutputDefinition `field:"optional" json:"outputs" yaml:"outputs"`
	// Controls which manifests and segments are generated.
	// Default: HlsOutputSelection.MANIFESTS_AND_SEGMENTS
	//
	// Experimental.
	OutputSelection HlsOutputSelection `field:"optional" json:"outputSelection" yaml:"outputSelection"`
	// Includes or excludes the EXT-X-PROGRAM-DATE-TIME tag in .m3u8 manifest files.
	// Default: HlsProgramDateTime.INCLUDE
	//
	// Experimental.
	ProgramDateTime HlsProgramDateTime `field:"optional" json:"programDateTime" yaml:"programDateTime"`
	// Specifies the algorithm used to drive the HLS EXT-X-PROGRAM-DATE-TIME clock.
	// Default: - HlsProgramDateTimeClock.SYSTEM_CLOCK, or INITIALIZE_FROM_OUTPUT_TIMECODE with epoch locking
	//
	// Experimental.
	ProgramDateTimeClock HlsProgramDateTimeClock `field:"optional" json:"programDateTimeClock" yaml:"programDateTimeClock"`
	// The period of insertion of the EXT-X-PROGRAM-DATE-TIME entry.
	// Default: Duration.minutes(10)
	//
	// Experimental.
	ProgramDateTimePeriod awscdk.Duration `field:"optional" json:"programDateTimePeriod" yaml:"programDateTimePeriod"`
	// Whether the master manifest includes information about both pipelines.
	// Default: HlsRedundantManifest.DISABLED
	//
	// Experimental.
	RedundantManifest HlsRedundantManifest `field:"optional" json:"redundantManifest" yaml:"redundantManifest"`
	// The length of each media segment.
	//
	// HLS supports whole-second segments only.
	// Default: - Segment.seconds(2)
	//
	// Experimental.
	Segment Segment `field:"optional" json:"segment" yaml:"segment"`
	// The segmentation mode.
	// Default: HlsSegmentationMode.USE_SEGMENT_DURATION
	//
	// Experimental.
	SegmentationMode HlsSegmentationMode `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// The number of segments to write to a subdirectory before starting a new one.
	// Default: 10000.
	//
	// Experimental.
	SegmentsPerSubdirectory *float64 `field:"optional" json:"segmentsPerSubdirectory" yaml:"segmentsPerSubdirectory"`
	// Whether to include or exclude the RESOLUTION attribute for a video in the EXT-X-STREAM-INF tag.
	// Default: HlsStreamInfResolution.INCLUDE
	//
	// Experimental.
	StreamInfResolution HlsStreamInfResolution `field:"optional" json:"streamInfResolution" yaml:"streamInfResolution"`
	// Indicates the ID3 frame that has the timecode.
	// Default: HlsTimedMetadataId3Frame.PRIV
	//
	// Experimental.
	TimedMetadataId3Frame HlsTimedMetadataId3Frame `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval.
	// Default: Duration.seconds(10)
	//
	// Experimental.
	TimedMetadataId3Period awscdk.Duration `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// Provides an extra delta offset to fine tune the timestamps.
	// Default: - service default.
	//
	// Experimental.
	TimestampDelta awscdk.Duration `field:"optional" json:"timestampDelta" yaml:"timestampDelta"`
	// Whether to emit segmented files or a single file.
	// Default: HlsTsFileMode.SEGMENTED_FILES
	//
	// Experimental.
	TsFileMode HlsTsFileMode `field:"optional" json:"tsFileMode" yaml:"tsFileMode"`
}

