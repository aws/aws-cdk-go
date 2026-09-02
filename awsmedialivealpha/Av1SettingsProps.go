package awsmedialivealpha


// Properties for AV1 codec settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var afdSignaling AfdSignaling
//   var av1BitDepth Av1BitDepth
//   var av1ColorSpaceSettings Av1ColorSpaceSettings
//   var av1Level Av1Level
//   var av1RateControl Av1RateControl
//   var av1SceneChangeDetect Av1SceneChangeDetect
//   var av1SpatialAq Av1SpatialAq
//   var av1TemporalAq Av1TemporalAq
//   var av1TimecodeInsertion Av1TimecodeInsertion
//   var framerate Framerate
//   var gopSize GopSize
//   var lookAheadRateControl LookAheadRateControl
//   var pixelAspectRatio PixelAspectRatio
//   var timecodeBurninFontSize TimecodeBurninFontSize
//   var timecodeBurninPosition TimecodeBurninPosition
//
//   av1SettingsProps := &Av1SettingsProps{
//   	AfdSignaling: afdSignaling,
//   	BitDepth: av1BitDepth,
//   	BufSize: jsii.Number(123),
//   	ColorSpaceSettings: av1ColorSpaceSettings,
//   	FixedAfd: jsii.String("fixedAfd"),
//   	Framerate: framerate,
//   	GopSize: gopSize,
//   	Level: av1Level,
//   	LookAheadRateControl: lookAheadRateControl,
//   	MinBitrate: jsii.Number(123),
//   	MinIInterval: jsii.Number(123),
//   	PixelAspectRatio: pixelAspectRatio,
//   	RateControl: av1RateControl,
//   	SceneChangeDetect: av1SceneChangeDetect,
//   	SpatialAq: av1SpatialAq,
//   	TemporalAq: av1TemporalAq,
//   	TimecodeBurnin: &TimecodeBurninSettings{
//   		FontSize: timecodeBurninFontSize,
//   		Position: timecodeBurninPosition,
//   		Prefix: jsii.String("prefix"),
//   	},
//   	TimecodeInsertion: av1TimecodeInsertion,
//   }
//
// Experimental.
type Av1SettingsProps struct {
	// AFD signaling mode.
	// Default: AfdSignaling.NONE
	//
	// Experimental.
	AfdSignaling AfdSignaling `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Bit depth for the AV1 encode.
	// Default: - service default.
	//
	// Experimental.
	BitDepth Av1BitDepth `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// Size of buffer (HRD buffer model) in bits.
	// Default: - service default.
	//
	// Experimental.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Color space settings for the video.
	// Default: - service default.
	//
	// Experimental.
	ColorSpaceSettings Av1ColorSpaceSettings `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// Four-bit AFD value to write on all frames. Only valid when afdSignaling is FIXED.
	//
	// Valid values: FIXED_0000, FIXED_0010, FIXED_0011, FIXED_0100, FIXED_1000,
	// FIXED_1001, FIXED_1010, FIXED_1011, FIXED_1100, FIXED_1101, FIXED_1110, FIXED_1111.
	// Default: - service default.
	//
	// Experimental.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// The video frame rate.
	// Default: - follow source.
	//
	// Experimental.
	Framerate Framerate `field:"optional" json:"framerate" yaml:"framerate"`
	// The GOP size (keyframe interval).
	// Default: GopSize.seconds(1)
	//
	// Experimental.
	GopSize GopSize `field:"optional" json:"gopSize" yaml:"gopSize"`
	// The AV1 level.
	//
	// Valid values: AV1_LEVEL_2, AV1_LEVEL_2_1, AV1_LEVEL_3, AV1_LEVEL_3_1,
	// AV1_LEVEL_4, AV1_LEVEL_4_1, AV1_LEVEL_5, AV1_LEVEL_5_1, AV1_LEVEL_5_2,
	// AV1_LEVEL_5_3, AV1_LEVEL_6, AV1_LEVEL_6_1, AV1_LEVEL_6_2, AV1_LEVEL_6_3,
	// AV1_LEVEL_AUTO.
	// Default: Av1Level.AV1_LEVEL_AUTO
	//
	// Experimental.
	Level Av1Level `field:"optional" json:"level" yaml:"level"`
	// Amount of lookahead.
	//
	// Low decreases latency/memory; high can produce better quality.
	// Default: LookAheadRateControl.HIGH
	//
	// Experimental.
	LookAheadRateControl LookAheadRateControl `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// Minimum bitrate in bits/second.
	// Default: - service default.
	//
	// Experimental.
	MinBitrate *float64 `field:"optional" json:"minBitrate" yaml:"minBitrate"`
	// Only meaningful if sceneChangeDetect is enabled.
	//
	// Enforces separation between
	// repeated (cadence) I-frames and I-frames inserted by scene change detection.
	// Default: - service default.
	//
	// Experimental.
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// The pixel aspect ratio (PAR) of the video.
	// Default: - service default.
	//
	// Experimental.
	PixelAspectRatio PixelAspectRatio `field:"optional" json:"pixelAspectRatio" yaml:"pixelAspectRatio"`
	// The rate control configuration.
	// Default: - service default.
	//
	// Experimental.
	RateControl Av1RateControl `field:"optional" json:"rateControl" yaml:"rateControl"`
	// Scene change detection.
	// Default: Av1SceneChangeDetect.ENABLED
	//
	// Experimental.
	SceneChangeDetect Av1SceneChangeDetect `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// Spatial adaptive quantization.
	// Default: Av1SpatialAq.ENABLED
	//
	// Experimental.
	SpatialAq Av1SpatialAq `field:"optional" json:"spatialAq" yaml:"spatialAq"`
	// Temporal adaptive quantization.
	// Default: Av1TemporalAq.ENABLED
	//
	// Experimental.
	TemporalAq Av1TemporalAq `field:"optional" json:"temporalAq" yaml:"temporalAq"`
	// Timecode burn-in settings to overlay timecode on the video.
	// Default: - no timecode burn-in.
	//
	// Experimental.
	TimecodeBurnin *TimecodeBurninSettings `field:"optional" json:"timecodeBurnin" yaml:"timecodeBurnin"`
	// Timecode insertion mode.
	// Default: - service default.
	//
	// Experimental.
	TimecodeInsertion Av1TimecodeInsertion `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

