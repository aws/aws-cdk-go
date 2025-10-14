package awsmedialive


// The settings for the H.264 codec in the output.
//
// The parent of this entity is VideoCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h264SettingsProperty := &H264SettingsProperty{
//   	AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   	AfdSignaling: jsii.String("afdSignaling"),
//   	Bitrate: jsii.Number(123),
//   	BufFillPct: jsii.Number(123),
//   	BufSize: jsii.Number(123),
//   	ColorMetadata: jsii.String("colorMetadata"),
//   	ColorSpaceSettings: &H264ColorSpaceSettingsProperty{
//   		ColorSpacePassthroughSettings: &ColorSpacePassthroughSettingsProperty{
//   		},
//   		Rec601Settings: &Rec601SettingsProperty{
//   		},
//   		Rec709Settings: &Rec709SettingsProperty{
//   		},
//   	},
//   	EntropyEncoding: jsii.String("entropyEncoding"),
//   	FilterSettings: &H264FilterSettingsProperty{
//   		BandwidthReductionFilterSettings: &BandwidthReductionFilterSettingsProperty{
//   			PostFilterSharpening: jsii.String("postFilterSharpening"),
//   			Strength: jsii.String("strength"),
//   		},
//   		TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   			PostFilterSharpening: jsii.String("postFilterSharpening"),
//   			Strength: jsii.String("strength"),
//   		},
//   	},
//   	FixedAfd: jsii.String("fixedAfd"),
//   	FlickerAq: jsii.String("flickerAq"),
//   	ForceFieldPictures: jsii.String("forceFieldPictures"),
//   	FramerateControl: jsii.String("framerateControl"),
//   	FramerateDenominator: jsii.Number(123),
//   	FramerateNumerator: jsii.Number(123),
//   	GopBReference: jsii.String("gopBReference"),
//   	GopClosedCadence: jsii.Number(123),
//   	GopNumBFrames: jsii.Number(123),
//   	GopSize: jsii.Number(123),
//   	GopSizeUnits: jsii.String("gopSizeUnits"),
//   	Level: jsii.String("level"),
//   	LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   	MaxBitrate: jsii.Number(123),
//   	MinBitrate: jsii.Number(123),
//   	MinIInterval: jsii.Number(123),
//   	MinQp: jsii.Number(123),
//   	NumRefFrames: jsii.Number(123),
//   	ParControl: jsii.String("parControl"),
//   	ParDenominator: jsii.Number(123),
//   	ParNumerator: jsii.Number(123),
//   	Profile: jsii.String("profile"),
//   	QualityLevel: jsii.String("qualityLevel"),
//   	QvbrQualityLevel: jsii.Number(123),
//   	RateControlMode: jsii.String("rateControlMode"),
//   	ScanType: jsii.String("scanType"),
//   	SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   	Slices: jsii.Number(123),
//   	Softness: jsii.Number(123),
//   	SpatialAq: jsii.String("spatialAq"),
//   	SubgopLength: jsii.String("subgopLength"),
//   	Syntax: jsii.String("syntax"),
//   	TemporalAq: jsii.String("temporalAq"),
//   	TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   		FontSize: jsii.String("fontSize"),
//   		Position: jsii.String("position"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	TimecodeInsertion: jsii.String("timecodeInsertion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html
//
type CfnChannel_H264SettingsProperty struct {
	// The adaptive quantization.
	//
	// This allows intra-frame quantizers to vary to improve visual quality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-adaptivequantization
	//
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Indicates that AFD values will be written into the output stream.
	//
	// If afdSignaling is auto, the system tries to preserve the input AFD value (in cases where multiple AFD values are valid). If set to fixed, the AFD value is the value configured in the fixedAfd parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-afdsignaling
	//
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// The average bitrate in bits/second.
	//
	// This is required when the rate control mode is VBR or CBR. It isn't used for QVBR. In a Microsoft Smooth output group, each output must have a unique value when its bitrate is rounded down to the nearest multiple of 1000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-bitrate
	//
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// The percentage of the buffer that should initially be filled (HRD buffer model).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-buffillpct
	//
	BufFillPct *float64 `field:"optional" json:"bufFillPct" yaml:"bufFillPct"`
	// The size of the buffer (HRD buffer model) in bits/second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-bufsize
	//
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Includes color space metadata in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-colormetadata
	//
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Settings to configure the color space handling for the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-colorspacesettings
	//
	ColorSpaceSettings interface{} `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// The entropy encoding mode.
	//
	// Use cabac (must be in Main or High profile) or cavlc.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-entropyencoding
	//
	EntropyEncoding *string `field:"optional" json:"entropyEncoding" yaml:"entropyEncoding"`
	// Optional filters that you can apply to an encode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-filtersettings
	//
	FilterSettings interface{} `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// A four-bit AFD value to write on all frames of video in the output stream.
	//
	// Valid only when afdSignaling is set to Fixed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-fixedafd
	//
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// If set to enabled, adjusts the quantization within each frame to reduce flicker or pop on I-frames.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-flickeraq
	//
	FlickerAq *string `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// This setting applies only when scan type is "interlaced." It controls whether coding is performed on a field basis or on a frame basis. (When the video is progressive, the coding is always performed on a frame basis.) enabled: Force MediaLive to code on a field basis, so that odd and even sets of fields are coded separately. disabled: Code the two sets of fields separately (on a field basis) or together (on a frame basis using PAFF), depending on what is most appropriate for the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-forcefieldpictures
	//
	ForceFieldPictures *string `field:"optional" json:"forceFieldPictures" yaml:"forceFieldPictures"`
	// Indicates how the output video frame rate is specified.
	//
	// If you select "specified," the output video frame rate is determined by framerateNumerator and framerateDenominator. If you select "initializeFromSource," the output video frame rate is set equal to the input video frame rate of the first input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-frameratecontrol
	//
	FramerateControl *string `field:"optional" json:"framerateControl" yaml:"framerateControl"`
	// The frame rate denominator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-frameratedenominator
	//
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// The frame rate numerator.
	//
	// The frame rate is a fraction, for example, 24000/1001 = 23.976 fps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-frameratenumerator
	//
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// If enabled, uses reference B frames for GOP structures that have B frames > 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopbreference
	//
	GopBReference *string `field:"optional" json:"gopBReference" yaml:"gopBReference"`
	// The frequency of closed GOPs.
	//
	// In streaming applications, we recommend that you set this to 1 so that a decoder joining mid-stream will receive an IDR frame as quickly as possible. Setting this value to 0 will break output segmenting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopclosedcadence
	//
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// The number of B-frames between reference frames.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopnumbframes
	//
	GopNumBFrames *float64 `field:"optional" json:"gopNumBFrames" yaml:"gopNumBFrames"`
	// The GOP size (keyframe interval) in units of either frames or seconds per gopSizeUnits.
	//
	// The value must be greater than zero.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopsize
	//
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Indicates if the gopSize is specified in frames or seconds.
	//
	// If seconds, the system converts the gopSize into a frame count at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopsizeunits
	//
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// The H.264 level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
	// The amount of lookahead.
	//
	// A value of low can decrease latency and memory usage, while high can produce better quality for certain content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-lookaheadratecontrol
	//
	LookAheadRateControl *string `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// For QVBR: See the tooltip for Quality level.
	//
	// For VBR: Set the maximum bitrate in order to accommodate expected spikes in the complexity of the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-maxbitrate
	//
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-minbitrate
	//
	MinBitrate *float64 `field:"optional" json:"minBitrate" yaml:"minBitrate"`
	// Meaningful only if sceneChangeDetect is set to enabled.
	//
	// This setting enforces separation between repeated (cadence) I-frames and I-frames inserted by Scene Change Detection. If a scene change I-frame is within I-interval frames of a cadence I-frame, the GOP is shrunk or stretched to the scene change I-frame. GOP stretch requires enabling lookahead as well as setting the I-interval. The normal cadence resumes for the next GOP. Note that the maximum GOP stretch = GOP size + Min-I-interval - 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-miniinterval
	//
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-minqp
	//
	MinQp *float64 `field:"optional" json:"minQp" yaml:"minQp"`
	// The number of reference frames to use.
	//
	// The encoder might use more than requested if you use B-frames or interlaced encoding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-numrefframes
	//
	NumRefFrames *float64 `field:"optional" json:"numRefFrames" yaml:"numRefFrames"`
	// Indicates how the output pixel aspect ratio is specified.
	//
	// If "specified" is selected, the output video pixel aspect ratio is determined by parNumerator and parDenominator. If "initializeFromSource" is selected, the output pixels aspect ratio will be set equal to the input video pixel aspect ratio of the first input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-parcontrol
	//
	ParControl *string `field:"optional" json:"parControl" yaml:"parControl"`
	// The Pixel Aspect Ratio denominator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-pardenominator
	//
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// The Pixel Aspect Ratio numerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-parnumerator
	//
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// An H.264 profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-profile
	//
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Leave as STANDARD_QUALITY or choose a different value (which might result in additional costs to run the channel).
	//
	// - ENHANCED_QUALITY: Produces a slightly better video quality without an increase in the bitrate. Has an effect only when the Rate control mode is QVBR or CBR. If this channel is in a MediaLive multiplex, the value must be ENHANCED_QUALITY.
	// - STANDARD_QUALITY: Valid for any Rate control mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-qualitylevel
	//
	QualityLevel *string `field:"optional" json:"qualityLevel" yaml:"qualityLevel"`
	// Controls the target quality for the video encode.
	//
	// This applies only when the rate control mode is QVBR. Set values for the QVBR quality level field and Max bitrate field that suit your most important viewing devices. Recommended values are: - Primary screen: Quality level: 8 to 10. Max bitrate: 4M - PC or tablet: Quality level: 7. Max bitrate: 1.5M to 3M - Smartphone: Quality level: 6. Max bitrate: 1M to 1.5M.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-qvbrqualitylevel
	//
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
	// The rate control mode.
	//
	// QVBR: The quality will match the specified quality level except when it is constrained by the maximum bitrate. We recommend this if you or your viewers pay for bandwidth. VBR: The quality and bitrate vary, depending on the video complexity. We recommend this instead of QVBR if you want to maintain a specific average bitrate over the duration of the channel. CBR: The quality varies, depending on the video complexity. We recommend this only if you distribute your assets to devices that can't handle variable bitrates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-ratecontrolmode
	//
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Sets the scan type of the output to progressive or top-field-first interlaced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-scantype
	//
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// The scene change detection.
	//
	// On: inserts I-frames when the scene change is detected. Off: does not force an I-frame when the scene change is detected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-scenechangedetect
	//
	SceneChangeDetect *string `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// The number of slices per picture.
	//
	// The number must be less than or equal to the number of macroblock rows for progressive pictures, and less than or equal to half the number of macroblock rows for interlaced pictures. This field is optional. If you don't specify a value, MediaLive chooses the number of slices based on the encode resolution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-slices
	//
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// Softness.
	//
	// Selects a quantizer matrix. Larger values reduce high-frequency content in the encoded image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-softness
	//
	Softness *float64 `field:"optional" json:"softness" yaml:"softness"`
	// If set to enabled, adjusts quantization within each frame based on the spatial variation of content complexity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-spatialaq
	//
	SpatialAq *string `field:"optional" json:"spatialAq" yaml:"spatialAq"`
	// If set to fixed, uses gopNumBFrames B-frames per sub-GOP.
	//
	// If set to dynamic, optimizes the number of B-frames used for each sub-GOP to improve visual quality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-subgoplength
	//
	SubgopLength *string `field:"optional" json:"subgopLength" yaml:"subgopLength"`
	// Produces a bitstream that is compliant with SMPTE RP-2027.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-syntax
	//
	Syntax *string `field:"optional" json:"syntax" yaml:"syntax"`
	// If set to enabled, adjusts quantization within each frame based on the temporal variation of content complexity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-temporalaq
	//
	TemporalAq *string `field:"optional" json:"temporalAq" yaml:"temporalAq"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-timecodeburninsettings
	//
	TimecodeBurninSettings interface{} `field:"optional" json:"timecodeBurninSettings" yaml:"timecodeBurninSettings"`
	// Determines how timecodes should be inserted into the video elementary stream.
	//
	// disabled: don't include timecodes. picTimingSei: pass through picture timing SEI messages from the source specified in Timecode Config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-timecodeinsertion
	//
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

