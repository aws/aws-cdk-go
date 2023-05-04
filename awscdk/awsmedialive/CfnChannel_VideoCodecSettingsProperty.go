package awsmedialive


// The settings for the video codec in the output.
//
// The parent of this entity is VideoDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoCodecSettingsProperty := &VideoCodecSettingsProperty{
//   	FrameCaptureSettings: &FrameCaptureSettingsProperty{
//   		CaptureInterval: jsii.Number(123),
//   		CaptureIntervalUnits: jsii.String("captureIntervalUnits"),
//   		TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   			FontSize: jsii.String("fontSize"),
//   			Position: jsii.String("position"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	H264Settings: &H264SettingsProperty{
//   		AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		AfdSignaling: jsii.String("afdSignaling"),
//   		Bitrate: jsii.Number(123),
//   		BufFillPct: jsii.Number(123),
//   		BufSize: jsii.Number(123),
//   		ColorMetadata: jsii.String("colorMetadata"),
//   		ColorSpaceSettings: &H264ColorSpaceSettingsProperty{
//   			ColorSpacePassthroughSettings: &ColorSpacePassthroughSettingsProperty{
//   			},
//   			Rec601Settings: &Rec601SettingsProperty{
//   			},
//   			Rec709Settings: &Rec709SettingsProperty{
//   			},
//   		},
//   		EntropyEncoding: jsii.String("entropyEncoding"),
//   		FilterSettings: &H264FilterSettingsProperty{
//   			TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   				PostFilterSharpening: jsii.String("postFilterSharpening"),
//   				Strength: jsii.String("strength"),
//   			},
//   		},
//   		FixedAfd: jsii.String("fixedAfd"),
//   		FlickerAq: jsii.String("flickerAq"),
//   		ForceFieldPictures: jsii.String("forceFieldPictures"),
//   		FramerateControl: jsii.String("framerateControl"),
//   		FramerateDenominator: jsii.Number(123),
//   		FramerateNumerator: jsii.Number(123),
//   		GopBReference: jsii.String("gopBReference"),
//   		GopClosedCadence: jsii.Number(123),
//   		GopNumBFrames: jsii.Number(123),
//   		GopSize: jsii.Number(123),
//   		GopSizeUnits: jsii.String("gopSizeUnits"),
//   		Level: jsii.String("level"),
//   		LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   		MaxBitrate: jsii.Number(123),
//   		MinIInterval: jsii.Number(123),
//   		NumRefFrames: jsii.Number(123),
//   		ParControl: jsii.String("parControl"),
//   		ParDenominator: jsii.Number(123),
//   		ParNumerator: jsii.Number(123),
//   		Profile: jsii.String("profile"),
//   		QualityLevel: jsii.String("qualityLevel"),
//   		QvbrQualityLevel: jsii.Number(123),
//   		RateControlMode: jsii.String("rateControlMode"),
//   		ScanType: jsii.String("scanType"),
//   		SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   		Slices: jsii.Number(123),
//   		Softness: jsii.Number(123),
//   		SpatialAq: jsii.String("spatialAq"),
//   		SubgopLength: jsii.String("subgopLength"),
//   		Syntax: jsii.String("syntax"),
//   		TemporalAq: jsii.String("temporalAq"),
//   		TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   			FontSize: jsii.String("fontSize"),
//   			Position: jsii.String("position"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		TimecodeInsertion: jsii.String("timecodeInsertion"),
//   	},
//   	H265Settings: &H265SettingsProperty{
//   		AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		AfdSignaling: jsii.String("afdSignaling"),
//   		AlternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   		Bitrate: jsii.Number(123),
//   		BufSize: jsii.Number(123),
//   		ColorMetadata: jsii.String("colorMetadata"),
//   		ColorSpaceSettings: &H265ColorSpaceSettingsProperty{
//   			ColorSpacePassthroughSettings: &ColorSpacePassthroughSettingsProperty{
//   			},
//   			DolbyVision81Settings: &DolbyVision81SettingsProperty{
//   			},
//   			Hdr10Settings: &Hdr10SettingsProperty{
//   				MaxCll: jsii.Number(123),
//   				MaxFall: jsii.Number(123),
//   			},
//   			Rec601Settings: &Rec601SettingsProperty{
//   			},
//   			Rec709Settings: &Rec709SettingsProperty{
//   			},
//   		},
//   		FilterSettings: &H265FilterSettingsProperty{
//   			TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   				PostFilterSharpening: jsii.String("postFilterSharpening"),
//   				Strength: jsii.String("strength"),
//   			},
//   		},
//   		FixedAfd: jsii.String("fixedAfd"),
//   		FlickerAq: jsii.String("flickerAq"),
//   		FramerateDenominator: jsii.Number(123),
//   		FramerateNumerator: jsii.Number(123),
//   		GopClosedCadence: jsii.Number(123),
//   		GopSize: jsii.Number(123),
//   		GopSizeUnits: jsii.String("gopSizeUnits"),
//   		Level: jsii.String("level"),
//   		LookAheadRateControl: jsii.String("lookAheadRateControl"),
//   		MaxBitrate: jsii.Number(123),
//   		MinIInterval: jsii.Number(123),
//   		ParDenominator: jsii.Number(123),
//   		ParNumerator: jsii.Number(123),
//   		Profile: jsii.String("profile"),
//   		QvbrQualityLevel: jsii.Number(123),
//   		RateControlMode: jsii.String("rateControlMode"),
//   		ScanType: jsii.String("scanType"),
//   		SceneChangeDetect: jsii.String("sceneChangeDetect"),
//   		Slices: jsii.Number(123),
//   		Tier: jsii.String("tier"),
//   		TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   			FontSize: jsii.String("fontSize"),
//   			Position: jsii.String("position"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		TimecodeInsertion: jsii.String("timecodeInsertion"),
//   	},
//   	Mpeg2Settings: &Mpeg2SettingsProperty{
//   		AdaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		AfdSignaling: jsii.String("afdSignaling"),
//   		ColorMetadata: jsii.String("colorMetadata"),
//   		ColorSpace: jsii.String("colorSpace"),
//   		DisplayAspectRatio: jsii.String("displayAspectRatio"),
//   		FilterSettings: &Mpeg2FilterSettingsProperty{
//   			TemporalFilterSettings: &TemporalFilterSettingsProperty{
//   				PostFilterSharpening: jsii.String("postFilterSharpening"),
//   				Strength: jsii.String("strength"),
//   			},
//   		},
//   		FixedAfd: jsii.String("fixedAfd"),
//   		FramerateDenominator: jsii.Number(123),
//   		FramerateNumerator: jsii.Number(123),
//   		GopClosedCadence: jsii.Number(123),
//   		GopNumBFrames: jsii.Number(123),
//   		GopSize: jsii.Number(123),
//   		GopSizeUnits: jsii.String("gopSizeUnits"),
//   		ScanType: jsii.String("scanType"),
//   		SubgopLength: jsii.String("subgopLength"),
//   		TimecodeBurninSettings: &TimecodeBurninSettingsProperty{
//   			FontSize: jsii.String("fontSize"),
//   			Position: jsii.String("position"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		TimecodeInsertion: jsii.String("timecodeInsertion"),
//   	},
//   }
//
type CfnChannel_VideoCodecSettingsProperty struct {
	// The settings for the video codec in a frame capture output.
	FrameCaptureSettings interface{} `field:"optional" json:"frameCaptureSettings" yaml:"frameCaptureSettings"`
	// The settings for the H.264 codec in the output.
	H264Settings interface{} `field:"optional" json:"h264Settings" yaml:"h264Settings"`
	// Settings for video encoded with the H265 codec.
	H265Settings interface{} `field:"optional" json:"h265Settings" yaml:"h265Settings"`
	// Settings for video encoded with the MPEG-2 codec.
	Mpeg2Settings interface{} `field:"optional" json:"mpeg2Settings" yaml:"mpeg2Settings"`
}

