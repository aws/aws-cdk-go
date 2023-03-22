package awsmedialive


// The settings for the MPEG-2 codec in the output.
//
// The parent of this entity is VideoCodecSetting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mpeg2SettingsProperty := &mpeg2SettingsProperty{
//   	adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   	afdSignaling: jsii.String("afdSignaling"),
//   	colorMetadata: jsii.String("colorMetadata"),
//   	colorSpace: jsii.String("colorSpace"),
//   	displayAspectRatio: jsii.String("displayAspectRatio"),
//   	filterSettings: &mpeg2FilterSettingsProperty{
//   		temporalFilterSettings: &temporalFilterSettingsProperty{
//   			postFilterSharpening: jsii.String("postFilterSharpening"),
//   			strength: jsii.String("strength"),
//   		},
//   	},
//   	fixedAfd: jsii.String("fixedAfd"),
//   	framerateDenominator: jsii.Number(123),
//   	framerateNumerator: jsii.Number(123),
//   	gopClosedCadence: jsii.Number(123),
//   	gopNumBFrames: jsii.Number(123),
//   	gopSize: jsii.Number(123),
//   	gopSizeUnits: jsii.String("gopSizeUnits"),
//   	scanType: jsii.String("scanType"),
//   	subgopLength: jsii.String("subgopLength"),
//   	timecodeInsertion: jsii.String("timecodeInsertion"),
//   }
//
type CfnChannel_Mpeg2SettingsProperty struct {
	// Choose Off to disable adaptive quantization.
	//
	// Or choose another value to enable the quantizer and set its strength. The strengths are: Auto, Off, Low, Medium, High. When you enable this field, MediaLive allows intra-frame quantizers to vary, which might improve visual quality.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Indicates the AFD values that MediaLive will write into the video encode.
	//
	// If you do not know what AFD signaling is, or if your downstream system has not given you guidance, choose AUTO.
	// AUTO: MediaLive will try to preserve the input AFD value (in cases where multiple AFD values are valid).
	// FIXED: MediaLive will use the value you specify in fixedAFD.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Specifies whether to include the color space metadata.
	//
	// The metadata describes the color space that applies to the video (the colorSpace field). We recommend that you insert the metadata.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Choose the type of color space conversion to apply to the output.
	//
	// For detailed information on setting up both the input and the output to obtain the desired color space in the output, see the section on \"MediaLive Features - Video - color space\" in the MediaLive User Guide.
	// PASSTHROUGH: Keep the color space of the input content - do not convert it.
	// AUTO:Convert all content that is SD to rec 601, and convert all content that is HD to rec 709.
	ColorSpace *string `field:"optional" json:"colorSpace" yaml:"colorSpace"`
	// Sets the pixel aspect ratio for the encode.
	DisplayAspectRatio *string `field:"optional" json:"displayAspectRatio" yaml:"displayAspectRatio"`
	// Optionally specify a noise reduction filter, which can improve quality of compressed content.
	//
	// If you do not choose a filter, no filter will be applied.
	// TEMPORAL: This filter is useful for both source content that is noisy (when it has excessive digital artifacts) and source content that is clean.
	// When the content is noisy, the filter cleans up the source content before the encoding phase, with these two effects: First, it improves the output video quality because the content has been cleaned up. Secondly, it decreases the bandwidth because MediaLive does not waste bits on encoding noise.
	// When the content is reasonably clean, the filter tends to decrease the bitrate.
	FilterSettings interface{} `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Complete this field only when afdSignaling is set to FIXED.
	//
	// Enter the AFD value (4 bits) to write on all frames of the video encode.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// description": "The framerate denominator.
	//
	// For example, 1001. The framerate is the numerator divided by the denominator. For example, 24000 / 1001 = 23.976 FPS.
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// The framerate numerator.
	//
	// For example, 24000. The framerate is the numerator divided by the denominator. For example, 24000 / 1001 = 23.976 FPS.
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// MPEG2: default is open GOP.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// Relates to the GOP structure.
	//
	// The number of B-frames between reference frames. If you do not know what a B-frame is, use the default.
	GopNumBFrames *float64 `field:"optional" json:"gopNumBFrames" yaml:"gopNumBFrames"`
	// Relates to the GOP structure.
	//
	// The GOP size (keyframe interval) in the units specified in gopSizeUnits. If you do not know what GOP is, use the default.
	// If gopSizeUnits is frames, then the gopSize must be an integer and must be greater than or equal to 1.
	// If gopSizeUnits is seconds, the gopSize must be greater than 0, but does not need to be an integer.
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Relates to the GOP structure.
	//
	// Specifies whether the gopSize is specified in frames or seconds. If you do not plan to change the default gopSize, leave the default. If you specify SECONDS, MediaLive will internally convert the gop size to a frame count.
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// Set the scan type of the output to PROGRESSIVE or INTERLACED (top field first).
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// Relates to the GOP structure.
	//
	// If you do not know what GOP is, use the default.
	// FIXED: Set the number of B-frames in each sub-GOP to the value in gopNumBFrames.
	// DYNAMIC: Let MediaLive optimize the number of B-frames in each sub-GOP, to improve visual quality.
	SubgopLength *string `field:"optional" json:"subgopLength" yaml:"subgopLength"`
	// Determines how MediaLive inserts timecodes in the output video.
	//
	// For detailed information about setting up the input and the output for a timecode, see the section on \"MediaLive Features - Timecode configuration\" in the MediaLive User Guide.
	// DISABLED: do not include timecodes.
	// GOP_TIMECODE: Include timecode metadata in the GOP header.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

