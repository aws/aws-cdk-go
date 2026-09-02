package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// Properties for creating a MediaLive Channel.
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
//   			AudioSelectors: []AudioSelector{
//   				medialive.AudioSelector_ByLanguage(jsii.String("eng"), jsii.String("eng"), medialive.AudioLanguageSelectionPolicy_STRICT()),
//   			},
//   			CaptionSelectors: []CaptionSelector{
//   				medialive.CaptionSelector_Embedded(jsii.String("embedded")),
//   			},
//   			VideoSelector: &VideoSelectorSettings{
//   				ColorSpace: medialive.VideoColorSpace_HDR10(),
//   				ColorSpaceUsage: medialive.VideoColorSpaceUsage_FORCE(),
//   				SelectBy: medialive.VideoSelection_ByProgramId(jsii.Number(1)),
//   			},
//   		},
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
type ChannelProps struct {
	// The input attachments for this channel.
	//
	// At least one is required.
	// Additional inputs can be added with `addInput()`.
	// Experimental.
	Inputs *[]*InputAttachment `field:"required" json:"inputs" yaml:"inputs"`
	// The initial output groups for this channel. At least one is required. Additional output groups can be added with `addOutputGroup()`.
	//
	// A single channel can contain multiple output groups with different codecs
	// (e.g. H.264 and H.265 ladders) sharing the same input.
	// Experimental.
	OutputGroups *[]OutputGroupConfiguration `field:"required" json:"outputGroups" yaml:"outputGroups"`
	// Anywhere settings for running the channel on AWS Elemental Anywhere.
	// Default: - not an Anywhere channel.
	//
	// Experimental.
	AnywhereSettings *AnywhereSettings `field:"optional" json:"anywhereSettings" yaml:"anywhereSettings"`
	// Settings for blanking video, audio, and captions during ad avails.
	// Default: - avail blanking disabled.
	//
	// Experimental.
	AvailBlanking *AvailBlanking `field:"optional" json:"availBlanking" yaml:"availBlanking"`
	// Ad avail handling configuration.
	//
	// Defines how SCTE-35 markers are processed.
	// Default: - no avail configuration.
	//
	// Experimental.
	AvailSettings AvailSettings `field:"optional" json:"availSettings" yaml:"availSettings"`
	// Blackout slate configuration.
	//
	// Controls what is displayed during blackout events.
	// Default: - blackout slate disabled.
	//
	// Experimental.
	BlackoutSlate *BlackoutSlate `field:"optional" json:"blackoutSlate" yaml:"blackoutSlate"`
	// The class of the channel (STANDARD for redundancy, SINGLE_PIPELINE for cost savings).
	// Default: ChannelClass.SINGLE_PIPELINE
	//
	// Experimental.
	ChannelClass ChannelClass `field:"optional" json:"channelClass" yaml:"channelClass"`
	// The engine version for the channel.
	// Default: - service default.
	//
	// Experimental.
	ChannelEngineVersion *string `field:"optional" json:"channelEngineVersion" yaml:"channelEngineVersion"`
	// The name of the channel.
	// Default: - auto-generated.
	//
	// Experimental.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Input security groups to associate with the channel.
	//
	// Controls which IP addresses can connect
	// to the channel's outputs (pull-style outputs where downstream systems initiate connections).
	// Default: - no channel security groups.
	//
	// Experimental.
	ChannelSecurityGroups *[]interfacesawsmedialive.IInputSecurityGroupRef `field:"optional" json:"channelSecurityGroups" yaml:"channelSecurityGroups"`
	// Global color correction rules applied to all outputs.
	// Default: - no color corrections.
	//
	// Experimental.
	ColorCorrections *[]*ColorCorrection `field:"optional" json:"colorCorrections" yaml:"colorCorrections"`
	// Feature activations for the channel (e.g. Input Prepare schedule actions).
	// Default: - all features disabled.
	//
	// Experimental.
	FeatureActivations *FeatureActivations `field:"optional" json:"featureActivations" yaml:"featureActivations"`
	// Global configuration settings for the channel.
	// Default: - default global configuration.
	//
	// Experimental.
	GlobalConfiguration *GlobalConfiguration `field:"optional" json:"globalConfiguration" yaml:"globalConfiguration"`
	// An AWS Elemental Inference feed to send this channel's media to for inference processing.
	//
	// Future breaking change: this will change when Elemental Inference is released as an L2 construct.
	// Default: - the channel is not associated to an inference feed.
	//
	// Experimental.
	InferenceFeedArn *string `field:"optional" json:"inferenceFeedArn" yaml:"inferenceFeedArn"`
	// The input specification for this channel.
	//
	// Defines the expected codec, bitrate, and
	// resolution of the inputs, and whether they are standard, CDI, or Elemental Link inputs.
	// Default: - InputSpecification.standard() (AVC codec, 20 Mbps max, HD resolution)
	//
	// Experimental.
	InputSpecification InputSpecification `field:"optional" json:"inputSpecification" yaml:"inputSpecification"`
	// Linked channel settings for primary/follower channel configurations.
	// Default: - not a linked channel.
	//
	// Experimental.
	LinkedChannelSettings LinkedChannelSettings `field:"optional" json:"linkedChannelSettings" yaml:"linkedChannelSettings"`
	// The log level for the channel.
	// Default: LogLevel.DISABLED
	//
	// Experimental.
	LogLevel LogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Maintenance window configuration for the channel.
	// Default: - default maintenance window.
	//
	// Experimental.
	Maintenance *MaintenanceSettings `field:"optional" json:"maintenance" yaml:"maintenance"`
	// Motion graphics overlay configuration.
	// Default: - motion graphics disabled.
	//
	// Experimental.
	MotionGraphicsConfiguration *MotionGraphicsConfiguration `field:"optional" json:"motionGraphicsConfiguration" yaml:"motionGraphicsConfiguration"`
	// Nielsen watermark configuration.
	// Default: - no Nielsen configuration.
	//
	// Experimental.
	NielsenConfiguration *NielsenConfiguration `field:"optional" json:"nielsenConfiguration" yaml:"nielsenConfiguration"`
	// The IAM role for MediaLive to assume when running this channel.
	//
	// [disable-awslint:prefer-ref-interface].
	// Default: - a role is auto-created with confused-deputy prevention.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Which output groups receive SCTE-35 segmentation cues.
	// Default: - service default.
	//
	// Experimental.
	Scte35SegmentationScope Scte35SegmentationScope `field:"optional" json:"scte35SegmentationScope" yaml:"scte35SegmentationScope"`
	// Tags to add to the channel.
	// Default: - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Thumbnail generation configuration.
	// Default: - thumbnails disabled.
	//
	// Experimental.
	ThumbnailConfiguration *ThumbnailConfiguration `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
	// Timecode configuration for the channel.
	// Default: - EMBEDDED source, no sync threshold.
	//
	// Experimental.
	TimecodeConfig *TimecodeConfig `field:"optional" json:"timecodeConfig" yaml:"timecodeConfig"`
	// VPC output settings.
	//
	// When set, all output endpoints are created in the specified VPC.
	// Default: - no VPC (outputs use public endpoints).
	//
	// Experimental.
	Vpc *VpcOutputSettings `field:"optional" json:"vpc" yaml:"vpc"`
}

