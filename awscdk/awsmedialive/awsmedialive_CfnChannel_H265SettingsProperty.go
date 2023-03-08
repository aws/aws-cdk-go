package awsmedialive


// H265 Settings.
//
// The parent of this entity is VideoCodecSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   h265SettingsProperty := &h265SettingsProperty{
//   	adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   	afdSignaling: jsii.String("afdSignaling"),
//   	alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   	bitrate: jsii.Number(123),
//   	bufSize: jsii.Number(123),
//   	colorMetadata: jsii.String("colorMetadata"),
//   	colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   		colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   		},
//   		hdr10Settings: &hdr10SettingsProperty{
//   			maxCll: jsii.Number(123),
//   			maxFall: jsii.Number(123),
//   		},
//   		rec601Settings: &rec601SettingsProperty{
//   		},
//   		rec709Settings: &rec709SettingsProperty{
//   		},
//   	},
//   	filterSettings: &h265FilterSettingsProperty{
//   		temporalFilterSettings: &temporalFilterSettingsProperty{
//   			postFilterSharpening: jsii.String("postFilterSharpening"),
//   			strength: jsii.String("strength"),
//   		},
//   	},
//   	fixedAfd: jsii.String("fixedAfd"),
//   	flickerAq: jsii.String("flickerAq"),
//   	framerateDenominator: jsii.Number(123),
//   	framerateNumerator: jsii.Number(123),
//   	gopClosedCadence: jsii.Number(123),
//   	gopSize: jsii.Number(123),
//   	gopSizeUnits: jsii.String("gopSizeUnits"),
//   	level: jsii.String("level"),
//   	lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   	maxBitrate: jsii.Number(123),
//   	minIInterval: jsii.Number(123),
//   	parDenominator: jsii.Number(123),
//   	parNumerator: jsii.Number(123),
//   	profile: jsii.String("profile"),
//   	qvbrQualityLevel: jsii.Number(123),
//   	rateControlMode: jsii.String("rateControlMode"),
//   	scanType: jsii.String("scanType"),
//   	sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   	slices: jsii.Number(123),
//   	tier: jsii.String("tier"),
//   	timecodeInsertion: jsii.String("timecodeInsertion"),
//   }
//
type CfnChannel_H265SettingsProperty struct {
	// Adaptive quantization.
	//
	// Allows intra-frame quantizers to vary to improve visual quality.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Indicates that AFD values will be written into the output stream.
	//
	// If afdSignaling is "auto", the system will try to preserve the input AFD value (in cases where multiple AFD values are valid). If set to "fixed", the AFD value will be the value configured in the fixedAfd parameter.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Whether or not EML should insert an Alternative Transfer Function SEI message to support backwards compatibility with non-HDR decoders and displays.
	AlternativeTransferFunction *string `field:"optional" json:"alternativeTransferFunction" yaml:"alternativeTransferFunction"`
	// Average bitrate in bits/second.
	//
	// Required when the rate control mode is VBR or CBR. Not used for QVBR. In an MS Smooth output group, each output must have a unique value when its bitrate is rounded down to the nearest multiple of 1000.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Size of buffer (HRD buffer model) in bits.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Includes colorspace metadata in the output.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Color Space settings.
	ColorSpaceSettings interface{} `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// Optional filters that you can apply to an encode.
	FilterSettings interface{} `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Four bit AFD value to write on all frames of video in the output stream.
	//
	// Only valid when afdSignaling is set to 'Fixed'.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// If set to enabled, adjust quantization within each frame to reduce flicker or 'pop' on I-frames.
	FlickerAq *string `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// Framerate denominator.
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// Framerate numerator - framerate is a fraction, e.g. 24000 / 1001 = 23.976 fps.
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// Frequency of closed GOPs.
	//
	// In streaming applications, it is recommended that this be set to 1 so a decoder joining mid-stream will receive an IDR frame as quickly as possible. Setting this value to 0 will break output segmenting.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// GOP size (keyframe interval) in units of either frames or seconds per gopSizeUnits.
	//
	// If gopSizeUnits is frames, gopSize must be an integer and must be greater than or equal to 1.
	// If gopSizeUnits is seconds, gopSize must be greater than 0, but need not be an integer.
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Indicates if the gopSize is specified in frames or seconds.
	//
	// If seconds the system will convert the gopSize into a frame count at run time.
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// H.265 Level.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Amount of lookahead.
	//
	// A value of low can decrease latency and memory usage, while high can produce better quality for certain content.
	LookAheadRateControl *string `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// For QVBR: See the tooltip for Quality level.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// Only meaningful if sceneChangeDetect is set to enabled.
	//
	// Defaults to 5 if multiplex rate control is used. Enforces separation between repeated (cadence) I-frames and I-frames inserted by Scene Change Detection. If a scene change I-frame is within I-interval frames of a cadence I-frame, the GOP is shrunk and/or stretched to the scene change I-frame. GOP stretch requires enabling lookahead as well as setting I-interval. The normal cadence resumes for the next GOP. Note: Maximum GOP stretch = GOP size + Min-I-interval - 1
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// Pixel Aspect Ratio denominator.
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// Pixel Aspect Ratio numerator.
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// H.265 Profile.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Controls the target quality for the video encode.
	//
	// Applies only when the rate control mode is QVBR. Set values for the QVBR quality level field and Max bitrate field that suit your most important viewing devices. Recommended values are:
	// - Primary screen: Quality level: 8 to 10. Max bitrate: 4M
	// - PC or tablet: Quality level: 7. Max bitrate: 1.5M to 3M
	// - Smartphone: Quality level: 6. Max bitrate: 1M to 1.5M
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
	// Rate control mode.
	//
	// QVBR: Quality will match the specified quality level except when it is constrained by the
	// maximum bitrate. Recommended if you or your viewers pay for bandwidth. CBR: Quality varies, depending on the video complexity. Recommended only if you distribute
	// your assets to devices that cannot handle variable bitrates. Multiplex: This rate control mode is only supported (and is required) when the video is being
	// delivered to a MediaLive Multiplex in which case the rate control configuration is controlled
	// by the properties within the Multiplex Program.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Sets the scan type of the output to progressive or top-field-first interlaced.
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// Scene change detection.
	SceneChangeDetect *string `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// Number of slices per picture.
	//
	// Must be less than or equal to the number of macroblock rows for progressive pictures, and less than or equal to half the number of macroblock rows for interlaced pictures.
	// This field is optional; when no value is specified the encoder will choose the number of slices based on encode resolution.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// H.265 Tier.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// Determines how timecodes should be inserted into the video elementary stream.
	//
	// - 'disabled': Do not include timecodes
	// - 'picTimingSei': Pass through picture timing SEI messages from the source specified in Timecode Config.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

