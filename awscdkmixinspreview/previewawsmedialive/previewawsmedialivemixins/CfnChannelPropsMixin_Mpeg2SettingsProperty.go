package previewawsmedialivemixins


// The settings for the MPEG-2 codec in the output.
//
// The parent of this entity is VideoCodecSetting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mpeg2SettingsProperty := &Mpeg2SettingsProperty{
//   	AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   	AfdSignaling: jsii.String("afdSignaling"),
//   	ColorMetadata: jsii.String("colorMetadata"),
//   	ColorSpace: jsii.String("colorSpace"),
//   	DisplayAspectRatio: jsii.String("displayAspectRatio"),
//   	FilterSettings: &Mpeg2FilterSettingsProperty{
//   		TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   			PostFilterSharpening: jsii.String("postFilterSharpening"),
//   			Strength: jsii.String("strength"),
//   		},
//   	},
//   	FixedAfd: jsii.String("fixedAfd"),
//   	FramerateDenominator: jsii.Number(123),
//   	FramerateNumerator: jsii.Number(123),
//   	GopClosedCadence: jsii.Number(123),
//   	GopNumBFrames: jsii.Number(123),
//   	GopSize: jsii.Number(123),
//   	GopSizeUnits: jsii.String("gopSizeUnits"),
//   	ScanType: jsii.String("scanType"),
//   	SubgopLength: jsii.String("subgopLength"),
//   	TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   		FontSize: jsii.String("fontSize"),
//   		Position: jsii.String("position"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	TimecodeInsertion: jsii.String("timecodeInsertion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html
//
type CfnChannelPropsMixin_Mpeg2SettingsProperty struct {
	// Choose Off to disable adaptive quantization.
	//
	// Or choose another value to enable the quantizer and set its strength. The strengths are: Auto, Off, Low, Medium, High. When you enable this field, MediaLive allows intra-frame quantizers to vary, which might improve visual quality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-adaptivequantization
	//
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Indicates the AFD values that MediaLive will write into the video encode.
	//
	// If you do not know what AFD signaling is, or if your downstream system has not given you guidance, choose AUTO.
	// AUTO: MediaLive will try to preserve the input AFD value (in cases where multiple AFD values are valid).
	// FIXED: MediaLive will use the value you specify in fixedAFD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-afdsignaling
	//
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Specifies whether to include the color space metadata.
	//
	// The metadata describes the color space that applies to the video (the colorSpace field). We recommend that you insert the metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-colormetadata
	//
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Choose the type of color space conversion to apply to the output.
	//
	// For detailed information on setting up both the input and the output to obtain the desired color space in the output, see the section on \"MediaLive Features - Video - color space\" in the MediaLive User Guide.
	// PASSTHROUGH: Keep the color space of the input content - do not convert it.
	// AUTO:Convert all content that is SD to rec 601, and convert all content that is HD to rec 709.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-colorspace
	//
	ColorSpace *string `field:"optional" json:"colorSpace" yaml:"colorSpace"`
	// Sets the pixel aspect ratio for the encode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-displayaspectratio
	//
	DisplayAspectRatio *string `field:"optional" json:"displayAspectRatio" yaml:"displayAspectRatio"`
	// Optionally specify a noise reduction filter, which can improve quality of compressed content.
	//
	// If you do not choose a filter, no filter will be applied.
	// TEMPORAL: This filter is useful for both source content that is noisy (when it has excessive digital artifacts) and source content that is clean.
	// When the content is noisy, the filter cleans up the source content before the encoding phase, with these two effects: First, it improves the output video quality because the content has been cleaned up. Secondly, it decreases the bandwidth because MediaLive does not waste bits on encoding noise.
	// When the content is reasonably clean, the filter tends to decrease the bitrate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-filtersettings
	//
	FilterSettings interface{} `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Complete this field only when afdSignaling is set to FIXED.
	//
	// Enter the AFD value (4 bits) to write on all frames of the video encode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-fixedafd
	//
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// description": "The framerate denominator.
	//
	// For example, 1001. The framerate is the numerator divided by the denominator. For example, 24000 / 1001 = 23.976 FPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-frameratedenominator
	//
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// The framerate numerator.
	//
	// For example, 24000. The framerate is the numerator divided by the denominator. For example, 24000 / 1001 = 23.976 FPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-frameratenumerator
	//
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// MPEG2: default is open GOP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-gopclosedcadence
	//
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// Relates to the GOP structure.
	//
	// The number of B-frames between reference frames. If you do not know what a B-frame is, use the default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-gopnumbframes
	//
	GopNumBFrames *float64 `field:"optional" json:"gopNumBFrames" yaml:"gopNumBFrames"`
	// Relates to the GOP structure.
	//
	// The GOP size (keyframe interval) in the units specified in gopSizeUnits. If you do not know what GOP is, use the default.
	// If gopSizeUnits is frames, then the gopSize must be an integer and must be greater than or equal to 1.
	// If gopSizeUnits is seconds, the gopSize must be greater than 0, but does not need to be an integer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-gopsize
	//
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Relates to the GOP structure.
	//
	// Specifies whether the gopSize is specified in frames or seconds. If you do not plan to change the default gopSize, leave the default. If you specify SECONDS, MediaLive will internally convert the gop size to a frame count.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-gopsizeunits
	//
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// Set the scan type of the output to PROGRESSIVE or INTERLACED (top field first).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-scantype
	//
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// Relates to the GOP structure.
	//
	// If you do not know what GOP is, use the default.
	// FIXED: Set the number of B-frames in each sub-GOP to the value in gopNumBFrames.
	// DYNAMIC: Let MediaLive optimize the number of B-frames in each sub-GOP, to improve visual quality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-subgoplength
	//
	SubgopLength *string `field:"optional" json:"subgopLength" yaml:"subgopLength"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-timecodeburninsettings
	//
	TimecodeBurninSettings interface{} `field:"optional" json:"timecodeBurninSettings" yaml:"timecodeBurninSettings"`
	// Determines how MediaLive inserts timecodes in the output video.
	//
	// For detailed information about setting up the input and the output for a timecode, see the section on \"MediaLive Features - Timecode configuration\" in the MediaLive User Guide.
	// DISABLED: do not include timecodes.
	// GOP_TIMECODE: Include timecode metadata in the GOP header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mpeg2settings.html#cfn-medialive-channel-mpeg2settings-timecodeinsertion
	//
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

