package awsmedialivealpha


// Properties for H.265 codec settings.
//
// Example:
//   // H.264
//   h264 := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("h264_720p"),
//   	Codec: medialive.VideoCodecSettings_H264(&H264SettingsProps{
//   		RateControl: medialive.H264RateControl_Cbr(&CbrRateControlProps{
//   			Bitrate: awscdk.Bitrate_Mbps(jsii.Number(3)),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   		Profile: medialive.H264Profile_HIGH(),
//   	}),
//   	Width: jsii.Number(1280),
//   	Height: jsii.Number(720),
//   })
//
//   // H.265
//   h265 := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("h265_1080p"),
//   	Codec: medialive.VideoCodecSettings_H265(&H265SettingsProps{
//   		RateControl: medialive.H265RateControl_Qvbr(&QvbrRateControlProps{
//   			MaxBitrate: awscdk.Bitrate_*Mbps(jsii.Number(5)),
//   			QvbrQualityLevel: jsii.Number(7),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   		Profile: medialive.H265Profile_MAIN(),
//   		Tier: medialive.H265Tier_HIGH(),
//   	}),
//   	Width: jsii.Number(1920),
//   	Height: jsii.Number(1080),
//   })
//
// Experimental.
type H265SettingsProps struct {
	// The video frame rate.
	//
	// Required for H.265.
	// Experimental.
	Framerate Framerate `field:"required" json:"framerate" yaml:"framerate"`
	// The adaptive quantization.
	//
	// Allows intra-frame quantizers to vary to improve visual quality.
	// Default: H265AdaptiveQuantization.AUTO
	//
	// Experimental.
	AdaptiveQuantization H265AdaptiveQuantization `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// AFD signaling mode.
	// Default: AfdSignaling.NONE
	//
	// Experimental.
	AfdSignaling AfdSignaling `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Whether to insert an Alternative Transfer Function SEI message for backwards compatibility with non-HDR decoders.
	// Default: - service default.
	//
	// Experimental.
	AlternativeTransferFunction H265AlternativeTransferFunction `field:"optional" json:"alternativeTransferFunction" yaml:"alternativeTransferFunction"`
	// Size of buffer (HRD buffer model) in bits.
	// Default: - service default.
	//
	// Experimental.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Whether to include color space metadata in the output.
	// Default: - service default.
	//
	// Experimental.
	ColorMetadata ColorMetadata `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Color space settings for the video.
	// Default: - service default.
	//
	// Experimental.
	ColorSpaceSettings H265ColorSpaceSettings `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// Deblocking filter control.
	// Default: - service default.
	//
	// Experimental.
	Deblocking H265Deblocking `field:"optional" json:"deblocking" yaml:"deblocking"`
	// Optional video filter settings.
	// Default: - service default.
	//
	// Experimental.
	FilterSettings H265FilterSettings `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Four-bit AFD value to write on all frames. Only valid when afdSignaling is FIXED.
	//
	// Valid values: FIXED_0000, FIXED_0010, FIXED_0011, FIXED_0100, FIXED_1000,
	// FIXED_1001, FIXED_1010, FIXED_1011, FIXED_1100, FIXED_1101, FIXED_1110, FIXED_1111.
	// Default: - service default.
	//
	// Experimental.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// If enabled, adjusts quantization within each frame to reduce flicker on I-frames.
	// Default: - service default.
	//
	// Experimental.
	FlickerAq FlickerAq `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// If enabled, uses reference B frames for GOP structures that have B frames > 1.
	// Default: - service default.
	//
	// Experimental.
	GopBReference GopBReference `field:"optional" json:"gopBReference" yaml:"gopBReference"`
	// Frequency of closed GOPs.
	//
	// Set to 1 for streaming so decoders joining mid-stream get an IDR frame quickly.
	// Default: - service default.
	//
	// Experimental.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// Number of B-frames between reference frames.
	// Default: - service default.
	//
	// Experimental.
	GopNumBFrames *float64 `field:"optional" json:"gopNumBFrames" yaml:"gopNumBFrames"`
	// The GOP size (keyframe interval).
	// Default: GopSize.seconds(1)
	//
	// Experimental.
	GopSize GopSize `field:"optional" json:"gopSize" yaml:"gopSize"`
	// The H.265 level.
	//
	// Valid values: H265_LEVEL_1, H265_LEVEL_2, H265_LEVEL_2_1, H265_LEVEL_3,
	// H265_LEVEL_3_1, H265_LEVEL_4, H265_LEVEL_4_1, H265_LEVEL_5, H265_LEVEL_5_1,
	// H265_LEVEL_5_2, H265_LEVEL_6, H265_LEVEL_6_1, H265_LEVEL_6_2, H265_LEVEL_AUTO.
	// Default: - service default (auto).
	//
	// Experimental.
	Level H265Level `field:"optional" json:"level" yaml:"level"`
	// Amount of lookahead.
	//
	// Low decreases latency/memory; high can produce better quality.
	// Default: - service default.
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
	// Minimum QP value.
	// Default: - service default.
	//
	// Experimental.
	MinQp *float64 `field:"optional" json:"minQp" yaml:"minQp"`
	// Whether motion vectors can cross picture boundaries.
	// Default: - service default.
	//
	// Experimental.
	MvOverPictureBoundaries H265MvOverPictureBoundaries `field:"optional" json:"mvOverPictureBoundaries" yaml:"mvOverPictureBoundaries"`
	// Whether to use temporal motion vector prediction.
	// Default: - service default.
	//
	// Experimental.
	MvTemporalPredictor H265MvTemporalPredictor `field:"optional" json:"mvTemporalPredictor" yaml:"mvTemporalPredictor"`
	// The pixel aspect ratio (PAR) of the video.
	// Default: - square pixels.
	//
	// Experimental.
	PixelAspectRatio PixelAspectRatio `field:"optional" json:"pixelAspectRatio" yaml:"pixelAspectRatio"`
	// The H.265 profile.
	// Default: H265Profile.MAIN
	//
	// Experimental.
	Profile H265Profile `field:"optional" json:"profile" yaml:"profile"`
	// The rate control configuration.
	// Default: - CBR with no bitrate (service default).
	//
	// Experimental.
	RateControl H265RateControl `field:"optional" json:"rateControl" yaml:"rateControl"`
	// Sets the scan type of the output.
	// Default: ScanType.PROGRESSIVE
	//
	// Experimental.
	ScanType ScanType `field:"optional" json:"scanType" yaml:"scanType"`
	// Whether scene change detection inserts I-frames on scene changes.
	// Default: H265SceneChangeDetect.ENABLED
	//
	// Experimental.
	SceneChangeDetect H265SceneChangeDetect `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// Number of slices per picture.
	// Default: - encoder chooses based on resolution.
	//
	// Experimental.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// Sub-GOP length mode.
	// Default: - service default.
	//
	// Experimental.
	SubgopLength SubgopLength `field:"optional" json:"subgopLength" yaml:"subgopLength"`
	// The H.265 tier.
	// Default: H265Tier.MAIN
	//
	// Experimental.
	Tier H265Tier `field:"optional" json:"tier" yaml:"tier"`
	// Tile height in pixels.
	//
	// Must be a multiple of the CTU size.
	// Default: - service default.
	//
	// Experimental.
	TileHeight *float64 `field:"optional" json:"tileHeight" yaml:"tileHeight"`
	// Tile padding mode.
	// Default: - service default.
	//
	// Experimental.
	TilePadding H265TilePadding `field:"optional" json:"tilePadding" yaml:"tilePadding"`
	// Tile width in pixels.
	//
	// Must be a multiple of the CTU size.
	// Default: - service default.
	//
	// Experimental.
	TileWidth *float64 `field:"optional" json:"tileWidth" yaml:"tileWidth"`
	// Timecode burn-in settings to overlay timecode on the video.
	// Default: - no timecode burn-in.
	//
	// Experimental.
	TimecodeBurnin *TimecodeBurninSettings `field:"optional" json:"timecodeBurnin" yaml:"timecodeBurnin"`
	// Determines how timecodes are inserted into the video elementary stream.
	//
	// This controls insertion into the output elementary stream. The channel's `timecodeConfig` controls the
	// source of the timecode used for output.
	// Default: - service default.
	//
	// Experimental.
	TimecodeInsertion TimecodeInsertion `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
	// Treeblock size for the encoder.
	// Default: - service default.
	//
	// Experimental.
	TreeblockSize H265TreeblockSize `field:"optional" json:"treeblockSize" yaml:"treeblockSize"`
}

