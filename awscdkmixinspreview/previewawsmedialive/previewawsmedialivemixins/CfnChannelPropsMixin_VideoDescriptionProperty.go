package previewawsmedialivemixins


// Encoding information for one output video.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var colorSpacePassthroughSettings interface{}
//   var hlg2020Settings interface{}
//   var rec601Settings interface{}
//   var rec709Settings interface{}
//
//   videoDescriptionProperty := &VideoDescriptionProperty{
//   	CodecSettings: &VideoCodecSettingsProperty{
//   		Av1Settings: &Av1SettingsProperty{
//   			AfdSignaling: jsii.String("afdSignaling"),
//   			Bitrate: jsii.Number(123),
//   			BufSize: jsii.Number(123),
//   			ColorSpaceSettings: &Av1ColorSpaceSettingsProperty{
//   				ColorSpacePassthroughSettings: colorSpacePassthroughSettings,
//   				Hdr10Settings: &Hdr10SettingsProperty{
//   					MaxCll: jsii.Number(123),
//   					MaxFall: jsii.Number(123),
//   				},
//   				Rec601Settings: rec601Settings,
//   				Rec709Settings: rec709Settings,
//   			},
//   			FixedAfd: jsii.String("fixedAfd"),
//   			FramerateDenominator: jsii.Number(123),
//   			FramerateNumerator: jsii.Number(123),
//   			GopSize: jsii.Number(123),
//   			GopSizeUnits: jsii.String("gopSizeUnits"),
//   			Level: jsii.String("level"),
//   			LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   			MaxBitrate: jsii.Number(123),
//   			MinBitrate: jsii.Number(123),
//   			MinIInterval: jsii.Number(123),
//   			ParDenominator: jsii.Number(123),
//   			ParNumerator: jsii.Number(123),
//   			QvbrQualityLevel: jsii.Number(123),
//   			RateControlMode: jsii.String("rateControlMode"),
//   			SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   			SpatialAq: jsii.String("spatialAq"),
//   			TemporalAq: jsii.String("temporalAq"),
//   			TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   				FontSize: jsii.String("fontSize"),
//   				Position: jsii.String("position"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			TimecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   		FrameCaptureSettings: &FrameCaptureSettingsProperty{
//   			CaptureInterval: jsii.Number(123),
//   			CaptureIntervalUnits: jsii.String("captureIntervalUnits"),
//   			TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   				FontSize: jsii.String("fontSize"),
//   				Position: jsii.String("position"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   		H264Settings: &H264SettingsProperty{
//   			AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			AfdSignaling: jsii.String("afdSignaling"),
//   			Bitrate: jsii.Number(123),
//   			BufFillPct: jsii.Number(123),
//   			BufSize: jsii.Number(123),
//   			ColorMetadata: jsii.String("colorMetadata"),
//   			ColorSpaceSettings: &H264ColorSpaceSettingsProperty{
//   				ColorSpacePassthroughSettings: &ColorSpacePassthroughSettingsProperty{
//   				},
//   				Rec601Settings: &Rec601SettingsProperty{
//   				},
//   				Rec709Settings: &Rec709SettingsProperty{
//   				},
//   			},
//   			EntropyEncoding: jsii.String("entropyEncoding"),
//   			FilterSettings: &H264FilterSettingsProperty{
//   				BandwidthReductionFilterSettings: &BandwidthReductionFilterSettingsProperty{
//   					PostFilterSharpening: jsii.String("postFilterSharpening"),
//   					Strength: jsii.String("strength"),
//   				},
//   				TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   					PostFilterSharpening: jsii.String("postFilterSharpening"),
//   					Strength: jsii.String("strength"),
//   				},
//   			},
//   			FixedAfd: jsii.String("fixedAfd"),
//   			FlickerAq: jsii.String("flickerAq"),
//   			ForceFieldPictures: jsii.String("forceFieldPictures"),
//   			FramerateControl: jsii.String("framerateControl"),
//   			FramerateDenominator: jsii.Number(123),
//   			FramerateNumerator: jsii.Number(123),
//   			GopBReference: jsii.String("gopBReference"),
//   			GopClosedCadence: jsii.Number(123),
//   			GopNumBFrames: jsii.Number(123),
//   			GopSize: jsii.Number(123),
//   			GopSizeUnits: jsii.String("gopSizeUnits"),
//   			Level: jsii.String("level"),
//   			LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   			MaxBitrate: jsii.Number(123),
//   			MinBitrate: jsii.Number(123),
//   			MinIInterval: jsii.Number(123),
//   			MinQp: jsii.Number(123),
//   			NumRefFrames: jsii.Number(123),
//   			ParControl: jsii.String("parControl"),
//   			ParDenominator: jsii.Number(123),
//   			ParNumerator: jsii.Number(123),
//   			Profile: jsii.String("profile"),
//   			QualityLevel: jsii.String("qualityLevel"),
//   			QvbrQualityLevel: jsii.Number(123),
//   			RateControlMode: jsii.String("rateControlMode"),
//   			ScanType: jsii.String("scanType"),
//   			SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   			Slices: jsii.Number(123),
//   			Softness: jsii.Number(123),
//   			SpatialAq: jsii.String("spatialAq"),
//   			SubgopLength: jsii.String("subgopLength"),
//   			Syntax: jsii.String("syntax"),
//   			TemporalAq: jsii.String("temporalAq"),
//   			TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   				FontSize: jsii.String("fontSize"),
//   				Position: jsii.String("position"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			TimecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   		H265Settings: &H265SettingsProperty{
//   			AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			AfdSignaling: jsii.String("afdSignaling"),
//   			AlternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   			Bitrate: jsii.Number(123),
//   			BufSize: jsii.Number(123),
//   			ColorMetadata: jsii.String("colorMetadata"),
//   			ColorSpaceSettings: &H265ColorSpaceSettingsProperty{
//   				ColorSpacePassthroughSettings: &ColorSpacePassthroughSettingsProperty{
//   				},
//   				DolbyVision81Settings: &DolbyVision81SettingsProperty{
//   				},
//   				Hdr10Settings: &Hdr10SettingsProperty{
//   					MaxCll: jsii.Number(123),
//   					MaxFall: jsii.Number(123),
//   				},
//   				Hlg2020Settings: hlg2020Settings,
//   				Rec601Settings: &Rec601SettingsProperty{
//   				},
//   				Rec709Settings: &Rec709SettingsProperty{
//   				},
//   			},
//   			Deblocking: jsii.String("deblocking"),
//   			FilterSettings: &H265FilterSettingsProperty{
//   				BandwidthReductionFilterSettings: &BandwidthReductionFilterSettingsProperty{
//   					PostFilterSharpening: jsii.String("postFilterSharpening"),
//   					Strength: jsii.String("strength"),
//   				},
//   				TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   					PostFilterSharpening: jsii.String("postFilterSharpening"),
//   					Strength: jsii.String("strength"),
//   				},
//   			},
//   			FixedAfd: jsii.String("fixedAfd"),
//   			FlickerAq: jsii.String("flickerAq"),
//   			FramerateDenominator: jsii.Number(123),
//   			FramerateNumerator: jsii.Number(123),
//   			GopBReference: jsii.String("gopBReference"),
//   			GopClosedCadence: jsii.Number(123),
//   			GopNumBFrames: jsii.Number(123),
//   			GopSize: jsii.Number(123),
//   			GopSizeUnits: jsii.String("gopSizeUnits"),
//   			Level: jsii.String("level"),
//   			LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   			MaxBitrate: jsii.Number(123),
//   			MinBitrate: jsii.Number(123),
//   			MinIInterval: jsii.Number(123),
//   			MinQp: jsii.Number(123),
//   			MvOverPictureBoundaries: jsii.String("mvOverPictureBoundaries"),
//   			MvTemporalPredictor: jsii.String("mvTemporalPredictor"),
//   			ParDenominator: jsii.Number(123),
//   			ParNumerator: jsii.Number(123),
//   			Profile: jsii.String("profile"),
//   			QvbrQualityLevel: jsii.Number(123),
//   			RateControlMode: jsii.String("rateControlMode"),
//   			ScanType: jsii.String("scanType"),
//   			SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   			Slices: jsii.Number(123),
//   			SubgopLength: jsii.String("subgopLength"),
//   			Tier: jsii.String("tier"),
//   			TileHeight: jsii.Number(123),
//   			TilePadding: jsii.String("tilePadding"),
//   			TileWidth: jsii.Number(123),
//   			TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   				FontSize: jsii.String("fontSize"),
//   				Position: jsii.String("position"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			TimecodeInsertion: jsii.String("timecodeInsertion"),
//   			TreeblockSize: jsii.String("treeblockSize"),
//   		},
//   		Mpeg2Settings: &Mpeg2SettingsProperty{
//   			AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			AfdSignaling: jsii.String("afdSignaling"),
//   			ColorMetadata: jsii.String("colorMetadata"),
//   			ColorSpace: jsii.String("colorSpace"),
//   			DisplayAspectRatio: jsii.String("displayAspectRatio"),
//   			FilterSettings: &Mpeg2FilterSettingsProperty{
//   				TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   					PostFilterSharpening: jsii.String("postFilterSharpening"),
//   					Strength: jsii.String("strength"),
//   				},
//   			},
//   			FixedAfd: jsii.String("fixedAfd"),
//   			FramerateDenominator: jsii.Number(123),
//   			FramerateNumerator: jsii.Number(123),
//   			GopClosedCadence: jsii.Number(123),
//   			GopNumBFrames: jsii.Number(123),
//   			GopSize: jsii.Number(123),
//   			GopSizeUnits: jsii.String("gopSizeUnits"),
//   			ScanType: jsii.String("scanType"),
//   			SubgopLength: jsii.String("subgopLength"),
//   			TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   				FontSize: jsii.String("fontSize"),
//   				Position: jsii.String("position"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   			TimecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   	},
//   	Height: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	RespondToAfd: jsii.String("respondToAfd"),
//   	ScalingBehavior: jsii.String("scalingBehavior"),
//   	Sharpness: jsii.Number(123),
//   	Width: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videodescription.html
//
type CfnChannelPropsMixin_VideoDescriptionProperty struct {
	// The video codec settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videodescription.html#cfn-medialive-channel-videodescription-codecsettings
	//
	CodecSettings interface{} `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// The output video height, in pixels.
	//
	// This must be an even number. For most codecs, you can keep this field and width blank in order to use the height and width (resolution) from the source. Note that we don't recommend keeping the field blank. For the Frame Capture codec, height and width are required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videodescription.html#cfn-medialive-channel-videodescription-height
	//
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The name of this VideoDescription.
	//
	// Outputs use this name to uniquely identify this description. Description names should be unique within this channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videodescription.html#cfn-medialive-channel-videodescription-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates how to respond to the AFD values in the input stream.
	//
	// RESPOND causes input video to be clipped, depending on the AFD value, input display aspect ratio, and output display aspect ratio, and (except for the FRAMECAPTURE codec) includes the values in the output. PASSTHROUGH (does not apply to FRAMECAPTURE codec) ignores the AFD values and includes the values in the output, so input video is not clipped. NONE ignores the AFD values and does not include the values through to the output, so input video is not clipped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videodescription.html#cfn-medialive-channel-videodescription-respondtoafd
	//
	RespondToAfd *string `field:"optional" json:"respondToAfd" yaml:"respondToAfd"`
	// STRETCHTOOUTPUT configures the output position to stretch the video to the specified output resolution (height and width).
	//
	// This option overrides any position value. DEFAULT might insert black boxes (pillar boxes or letter boxes) around the video to provide the specified output resolution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videodescription.html#cfn-medialive-channel-videodescription-scalingbehavior
	//
	ScalingBehavior *string `field:"optional" json:"scalingBehavior" yaml:"scalingBehavior"`
	// Changes the strength of the anti-alias filter used for scaling.
	//
	// 0 is the softest setting, and 100 is the sharpest. We recommend a setting of 50 for most content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videodescription.html#cfn-medialive-channel-videodescription-sharpness
	//
	Sharpness *float64 `field:"optional" json:"sharpness" yaml:"sharpness"`
	// The output video width, in pixels.
	//
	// It must be an even number. For most codecs, you can keep this field and height blank in order to use the height and width (resolution) from the source. Note that we don't recommend keeping the field blank. For the Frame Capture codec, height and width are required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videodescription.html#cfn-medialive-channel-videodescription-width
	//
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

