package awsmedialivealpha


// Properties for H.264 codec settings.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//
//
//   video := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("video_720p"),
//   	Codec: medialive.VideoCodecSettings_H264(&H264SettingsProps{
//   		RateControl: medialive.H264RateControl_Cbr(&CbrRateControlProps{
//   			Bitrate: awscdk.Bitrate_Mbps(jsii.Number(3)),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   	}),
//   	Width: jsii.Number(1280),
//   	Height: jsii.Number(720),
//   })
//
//   audio := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
//   	Name: jsii.String("audio_aac"),
//   	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
//   		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
//   	}),
//   })
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
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
//   						audio,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type H264SettingsProps struct {
	// The adaptive quantization.
	//
	// This allows intra-frame quantizers to vary to improve visual quality.
	// Default: H264AdaptiveQuantization.AUTO
	//
	// Experimental.
	AdaptiveQuantization H264AdaptiveQuantization `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Indicates that AFD values will be written into the output stream.
	//
	// If afdSignaling is auto, the
	// system tries to preserve the input AFD value (in cases where multiple AFD values are valid). If
	// set to fixed, the AFD value is the value configured in the fixedAfd parameter.
	// Default: AfdSignaling.NONE
	//
	// Experimental.
	AfdSignaling AfdSignaling `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Percentage of the buffer that should initially be filled (HRD buffer model).
	// Default: - service default.
	//
	// Experimental.
	BufFillPct *float64 `field:"optional" json:"bufFillPct" yaml:"bufFillPct"`
	// Size of the buffer (HRD buffer model) in bits/second.
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
	ColorSpaceSettings H264ColorSpaceSettings `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// The entropy encoding mode.
	//
	// CABAC requires Main or High profile.
	// Default: - service default.
	//
	// Experimental.
	EntropyEncoding H264EntropyEncoding `field:"optional" json:"entropyEncoding" yaml:"entropyEncoding"`
	// Optional video filter settings.
	// Default: - service default.
	//
	// Experimental.
	FilterSettings H264FilterSettings `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Four-bit AFD value to write on all frames. Only valid when afdSignaling is FIXED.
	//
	// Valid values: FIXED_0000, FIXED_0010, FIXED_0011, FIXED_0100, FIXED_1000,
	// FIXED_1001, FIXED_1010, FIXED_1011, FIXED_1100, FIXED_1101, FIXED_1110, FIXED_1111.
	// Default: - service default.
	//
	// Experimental.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// If enabled, adjusts quantization within each frame to reduce flicker on I-frames.
	// Default: FlickerAq.ENABLED
	//
	// Experimental.
	FlickerAq FlickerAq `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// Controls whether coding is on a field basis or frame basis when scan type is interlaced.
	// Default: - service default.
	//
	// Experimental.
	ForceFieldPictures H264ForceFieldPictures `field:"optional" json:"forceFieldPictures" yaml:"forceFieldPictures"`
	// The video frame rate.
	// Default: - follow source.
	//
	// Experimental.
	Framerate Framerate `field:"optional" json:"framerate" yaml:"framerate"`
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
	// The number of B-frames between reference frames.
	// Default: - service default.
	//
	// Experimental.
	GopNumBFrames *float64 `field:"optional" json:"gopNumBFrames" yaml:"gopNumBFrames"`
	// The GOP size (keyframe interval).
	// Default: GopSize.seconds(1)
	//
	// Experimental.
	GopSize GopSize `field:"optional" json:"gopSize" yaml:"gopSize"`
	// The H.264 level.
	//
	// Valid values: H264_LEVEL_1, H264_LEVEL_1_1, H264_LEVEL_1_2, H264_LEVEL_1_3,
	// H264_LEVEL_2, H264_LEVEL_2_1, H264_LEVEL_2_2, H264_LEVEL_3, H264_LEVEL_3_1,
	// H264_LEVEL_3_2, H264_LEVEL_4, H264_LEVEL_4_1, H264_LEVEL_4_2, H264_LEVEL_5,
	// H264_LEVEL_5_1, H264_LEVEL_5_2, H264_LEVEL_AUTO.
	// Default: H264Level.H264_LEVEL_AUTO
	//
	// Experimental.
	Level H264Level `field:"optional" json:"level" yaml:"level"`
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
	// Minimum QP value.
	//
	// Sets a floor on the quantization parameter.
	// Default: - service default.
	//
	// Experimental.
	MinQp *float64 `field:"optional" json:"minQp" yaml:"minQp"`
	// The number of reference frames to use.
	//
	// The encoder might use more if B-frames or interlaced encoding is used.
	// Default: - service default.
	//
	// Experimental.
	NumRefFrames *float64 `field:"optional" json:"numRefFrames" yaml:"numRefFrames"`
	// The pixel aspect ratio (PAR) of the video.
	// Default: - follow source (or square pixels when framerate is specified).
	//
	// Experimental.
	PixelAspectRatio PixelAspectRatio `field:"optional" json:"pixelAspectRatio" yaml:"pixelAspectRatio"`
	// The H.264 profile.
	// Default: H264Profile.MAIN
	//
	// Experimental.
	Profile H264Profile `field:"optional" json:"profile" yaml:"profile"`
	// Quality level.
	//
	// ENHANCED_QUALITY produces slightly better video without increasing bitrate.
	// Default: - service default.
	//
	// Experimental.
	QualityLevel H264QualityLevel `field:"optional" json:"qualityLevel" yaml:"qualityLevel"`
	// The rate control configuration.
	// Default: - CBR with no bitrate (service default).
	//
	// Experimental.
	RateControl H264RateControl `field:"optional" json:"rateControl" yaml:"rateControl"`
	// Sets the scan type of the output.
	// Default: ScanType.PROGRESSIVE
	//
	// Experimental.
	ScanType ScanType `field:"optional" json:"scanType" yaml:"scanType"`
	// Whether scene change detection inserts I-frames on scene changes.
	// Default: H264SceneChangeDetect.ENABLED
	//
	// Experimental.
	SceneChangeDetect H264SceneChangeDetect `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// Number of slices per picture.
	//
	// Must be <= macroblock rows (progressive) or half (interlaced).
	// Default: - encoder chooses based on resolution.
	//
	// Experimental.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// Softness.
	//
	// Selects a quantizer matrix; larger values reduce high-frequency content.
	// Default: - service default.
	//
	// Experimental.
	Softness *float64 `field:"optional" json:"softness" yaml:"softness"`
	// Whether spatial adaptive quantization adjusts quantization within each frame based on spatial variation.
	// Default: H264SpatialAq.ENABLED
	//
	// Experimental.
	SpatialAq H264SpatialAq `field:"optional" json:"spatialAq" yaml:"spatialAq"`
	// Sub-GOP length mode.
	// Default: - service default.
	//
	// Experimental.
	SubgopLength SubgopLength `field:"optional" json:"subgopLength" yaml:"subgopLength"`
	// Produces a bitstream compliant with SMPTE RP-2027.
	// Default: H264Syntax.DEFAULT
	//
	// Experimental.
	Syntax H264Syntax `field:"optional" json:"syntax" yaml:"syntax"`
	// Whether temporal adaptive quantization adjusts quantization based on temporal variation between frames.
	// Default: H264TemporalAq.ENABLED
	//
	// Experimental.
	TemporalAq H264TemporalAq `field:"optional" json:"temporalAq" yaml:"temporalAq"`
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
}

