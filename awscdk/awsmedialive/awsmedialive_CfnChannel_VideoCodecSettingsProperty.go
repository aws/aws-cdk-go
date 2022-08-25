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
//   videoCodecSettingsProperty := &videoCodecSettingsProperty{
//   	frameCaptureSettings: &frameCaptureSettingsProperty{
//   		captureInterval: jsii.Number(123),
//   		captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   	},
//   	h264Settings: &h264SettingsProperty{
//   		adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		afdSignaling: jsii.String("afdSignaling"),
//   		bitrate: jsii.Number(123),
//   		bufFillPct: jsii.Number(123),
//   		bufSize: jsii.Number(123),
//   		colorMetadata: jsii.String("colorMetadata"),
//   		colorSpaceSettings: &h264ColorSpaceSettingsProperty{
//   			colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   			},
//   			rec601Settings: &rec601SettingsProperty{
//   			},
//   			rec709Settings: &rec709SettingsProperty{
//   			},
//   		},
//   		entropyEncoding: jsii.String("entropyEncoding"),
//   		filterSettings: &h264FilterSettingsProperty{
//   			temporalFilterSettings: &temporalFilterSettingsProperty{
//   				postFilterSharpening: jsii.String("postFilterSharpening"),
//   				strength: jsii.String("strength"),
//   			},
//   		},
//   		fixedAfd: jsii.String("fixedAfd"),
//   		flickerAq: jsii.String("flickerAq"),
//   		forceFieldPictures: jsii.String("forceFieldPictures"),
//   		framerateControl: jsii.String("framerateControl"),
//   		framerateDenominator: jsii.Number(123),
//   		framerateNumerator: jsii.Number(123),
//   		gopBReference: jsii.String("gopBReference"),
//   		gopClosedCadence: jsii.Number(123),
//   		gopNumBFrames: jsii.Number(123),
//   		gopSize: jsii.Number(123),
//   		gopSizeUnits: jsii.String("gopSizeUnits"),
//   		level: jsii.String("level"),
//   		lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   		maxBitrate: jsii.Number(123),
//   		minIInterval: jsii.Number(123),
//   		numRefFrames: jsii.Number(123),
//   		parControl: jsii.String("parControl"),
//   		parDenominator: jsii.Number(123),
//   		parNumerator: jsii.Number(123),
//   		profile: jsii.String("profile"),
//   		qualityLevel: jsii.String("qualityLevel"),
//   		qvbrQualityLevel: jsii.Number(123),
//   		rateControlMode: jsii.String("rateControlMode"),
//   		scanType: jsii.String("scanType"),
//   		sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   		slices: jsii.Number(123),
//   		softness: jsii.Number(123),
//   		spatialAq: jsii.String("spatialAq"),
//   		subgopLength: jsii.String("subgopLength"),
//   		syntax: jsii.String("syntax"),
//   		temporalAq: jsii.String("temporalAq"),
//   		timecodeInsertion: jsii.String("timecodeInsertion"),
//   	},
//   	h265Settings: &h265SettingsProperty{
//   		adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		afdSignaling: jsii.String("afdSignaling"),
//   		alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   		bitrate: jsii.Number(123),
//   		bufSize: jsii.Number(123),
//   		colorMetadata: jsii.String("colorMetadata"),
//   		colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   			colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   			},
//   			hdr10Settings: &hdr10SettingsProperty{
//   				maxCll: jsii.Number(123),
//   				maxFall: jsii.Number(123),
//   			},
//   			rec601Settings: &rec601SettingsProperty{
//   			},
//   			rec709Settings: &rec709SettingsProperty{
//   			},
//   		},
//   		filterSettings: &h265FilterSettingsProperty{
//   			temporalFilterSettings: &temporalFilterSettingsProperty{
//   				postFilterSharpening: jsii.String("postFilterSharpening"),
//   				strength: jsii.String("strength"),
//   			},
//   		},
//   		fixedAfd: jsii.String("fixedAfd"),
//   		flickerAq: jsii.String("flickerAq"),
//   		framerateDenominator: jsii.Number(123),
//   		framerateNumerator: jsii.Number(123),
//   		gopClosedCadence: jsii.Number(123),
//   		gopSize: jsii.Number(123),
//   		gopSizeUnits: jsii.String("gopSizeUnits"),
//   		level: jsii.String("level"),
//   		lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   		maxBitrate: jsii.Number(123),
//   		minIInterval: jsii.Number(123),
//   		parDenominator: jsii.Number(123),
//   		parNumerator: jsii.Number(123),
//   		profile: jsii.String("profile"),
//   		qvbrQualityLevel: jsii.Number(123),
//   		rateControlMode: jsii.String("rateControlMode"),
//   		scanType: jsii.String("scanType"),
//   		sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   		slices: jsii.Number(123),
//   		tier: jsii.String("tier"),
//   		timecodeInsertion: jsii.String("timecodeInsertion"),
//   	},
//   	mpeg2Settings: &mpeg2SettingsProperty{
//   		adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   		afdSignaling: jsii.String("afdSignaling"),
//   		colorMetadata: jsii.String("colorMetadata"),
//   		colorSpace: jsii.String("colorSpace"),
//   		displayAspectRatio: jsii.String("displayAspectRatio"),
//   		filterSettings: &mpeg2FilterSettingsProperty{
//   			temporalFilterSettings: &temporalFilterSettingsProperty{
//   				postFilterSharpening: jsii.String("postFilterSharpening"),
//   				strength: jsii.String("strength"),
//   			},
//   		},
//   		fixedAfd: jsii.String("fixedAfd"),
//   		framerateDenominator: jsii.Number(123),
//   		framerateNumerator: jsii.Number(123),
//   		gopClosedCadence: jsii.Number(123),
//   		gopNumBFrames: jsii.Number(123),
//   		gopSize: jsii.Number(123),
//   		gopSizeUnits: jsii.String("gopSizeUnits"),
//   		scanType: jsii.String("scanType"),
//   		subgopLength: jsii.String("subgopLength"),
//   		timecodeInsertion: jsii.String("timecodeInsertion"),
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

