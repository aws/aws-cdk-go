package awsmedialive


// Encoding information for one output video.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoDescriptionProperty := &VideoDescriptionProperty{
//   	CodecSettings: &VideoCodecSettingsProperty{
//   		FrameCaptureSettings: &FrameCaptureSettingsProperty{
//   			CaptureInterval: jsii.Number(123),
//   			CaptureIntervalUnits: jsii.String("captureIntervalUnits"),
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
//   			MinIInterval: jsii.Number(123),
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
//   				Hdr10Settings: &Hdr10SettingsProperty{
//   					MaxCll: jsii.Number(123),
//   					MaxFall: jsii.Number(123),
//   				},
//   				Rec601Settings: &Rec601SettingsProperty{
//   				},
//   				Rec709Settings: &Rec709SettingsProperty{
//   				},
//   			},
//   			FilterSettings: &H265FilterSettingsProperty{
//   				TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   					PostFilterSharpening: jsii.String("postFilterSharpening"),
//   					Strength: jsii.String("strength"),
//   				},
//   			},
//   			FixedAfd: jsii.String("fixedAfd"),
//   			FlickerAq: jsii.String("flickerAq"),
//   			FramerateDenominator: jsii.Number(123),
//   			FramerateNumerator: jsii.Number(123),
//   			GopClosedCadence: jsii.Number(123),
//   			GopSize: jsii.Number(123),
//   			GopSizeUnits: jsii.String("gopSizeUnits"),
//   			Level: jsii.String("level"),
//   			LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   			MaxBitrate: jsii.Number(123),
//   			MinIInterval: jsii.Number(123),
//   			ParDenominator: jsii.Number(123),
//   			ParNumerator: jsii.Number(123),
//   			Profile: jsii.String("profile"),
//   			QvbrQualityLevel: jsii.Number(123),
//   			RateControlMode: jsii.String("rateControlMode"),
//   			ScanType: jsii.String("scanType"),
//   			SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   			Slices: jsii.Number(123),
//   			Tier: jsii.String("tier"),
//   			TimecodeInsertion: jsii.String("timecodeInsertion"),
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
type CfnChannel_VideoDescriptionProperty struct {
	// The video codec settings.
	CodecSettings interface{} `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// The output video height, in pixels.
	//
	// This must be an even number. For most codecs, you can keep this field and width blank in order to use the height and width (resolution) from the source. Note that we don't recommend keeping the field blank. For the Frame Capture codec, height and width are required.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The name of this VideoDescription.
	//
	// Outputs use this name to uniquely identify this description. Description names should be unique within this channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates how to respond to the AFD values in the input stream.
	//
	// RESPOND causes input video to be clipped, depending on the AFD value, input display aspect ratio, and output display aspect ratio, and (except for the FRAMECAPTURE codec) includes the values in the output. PASSTHROUGH (does not apply to FRAMECAPTURE codec) ignores the AFD values and includes the values in the output, so input video is not clipped. NONE ignores the AFD values and does not include the values through to the output, so input video is not clipped.
	RespondToAfd *string `field:"optional" json:"respondToAfd" yaml:"respondToAfd"`
	// STRETCHTOOUTPUT configures the output position to stretch the video to the specified output resolution (height and width).
	//
	// This option overrides any position value. DEFAULT might insert black boxes (pillar boxes or letter boxes) around the video to provide the specified output resolution.
	ScalingBehavior *string `field:"optional" json:"scalingBehavior" yaml:"scalingBehavior"`
	// Changes the strength of the anti-alias filter used for scaling.
	//
	// 0 is the softest setting, and 100 is the sharpest. We recommend a setting of 50 for most content.
	Sharpness *float64 `field:"optional" json:"sharpness" yaml:"sharpness"`
	// The output video width, in pixels.
	//
	// It must be an even number. For most codecs, you can keep this field and height blank in order to use the height and width (resolution) from the source. Note that we don't recommend keeping the field blank. For the Frame Capture codec, height and width are required.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

