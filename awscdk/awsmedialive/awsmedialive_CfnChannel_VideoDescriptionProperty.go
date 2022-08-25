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
//   videoDescriptionProperty := &videoDescriptionProperty{
//   	codecSettings: &videoCodecSettingsProperty{
//   		frameCaptureSettings: &frameCaptureSettingsProperty{
//   			captureInterval: jsii.Number(123),
//   			captureIntervalUnits: jsii.String("captureIntervalUnits"),
//   		},
//   		h264Settings: &h264SettingsProperty{
//   			adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			afdSignaling: jsii.String("afdSignaling"),
//   			bitrate: jsii.Number(123),
//   			bufFillPct: jsii.Number(123),
//   			bufSize: jsii.Number(123),
//   			colorMetadata: jsii.String("colorMetadata"),
//   			colorSpaceSettings: &h264ColorSpaceSettingsProperty{
//   				colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   				},
//   				rec601Settings: &rec601SettingsProperty{
//   				},
//   				rec709Settings: &rec709SettingsProperty{
//   				},
//   			},
//   			entropyEncoding: jsii.String("entropyEncoding"),
//   			filterSettings: &h264FilterSettingsProperty{
//   				temporalFilterSettings: &temporalFilterSettingsProperty{
//   					postFilterSharpening: jsii.String("postFilterSharpening"),
//   					strength: jsii.String("strength"),
//   				},
//   			},
//   			fixedAfd: jsii.String("fixedAfd"),
//   			flickerAq: jsii.String("flickerAq"),
//   			forceFieldPictures: jsii.String("forceFieldPictures"),
//   			framerateControl: jsii.String("framerateControl"),
//   			framerateDenominator: jsii.Number(123),
//   			framerateNumerator: jsii.Number(123),
//   			gopBReference: jsii.String("gopBReference"),
//   			gopClosedCadence: jsii.Number(123),
//   			gopNumBFrames: jsii.Number(123),
//   			gopSize: jsii.Number(123),
//   			gopSizeUnits: jsii.String("gopSizeUnits"),
//   			level: jsii.String("level"),
//   			lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   			maxBitrate: jsii.Number(123),
//   			minIInterval: jsii.Number(123),
//   			numRefFrames: jsii.Number(123),
//   			parControl: jsii.String("parControl"),
//   			parDenominator: jsii.Number(123),
//   			parNumerator: jsii.Number(123),
//   			profile: jsii.String("profile"),
//   			qualityLevel: jsii.String("qualityLevel"),
//   			qvbrQualityLevel: jsii.Number(123),
//   			rateControlMode: jsii.String("rateControlMode"),
//   			scanType: jsii.String("scanType"),
//   			sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   			slices: jsii.Number(123),
//   			softness: jsii.Number(123),
//   			spatialAq: jsii.String("spatialAq"),
//   			subgopLength: jsii.String("subgopLength"),
//   			syntax: jsii.String("syntax"),
//   			temporalAq: jsii.String("temporalAq"),
//   			timecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   		h265Settings: &h265SettingsProperty{
//   			adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			afdSignaling: jsii.String("afdSignaling"),
//   			alternativeTransferFunction: jsii.String("alternativeTransferFunction"),
//   			bitrate: jsii.Number(123),
//   			bufSize: jsii.Number(123),
//   			colorMetadata: jsii.String("colorMetadata"),
//   			colorSpaceSettings: &h265ColorSpaceSettingsProperty{
//   				colorSpacePassthroughSettings: &colorSpacePassthroughSettingsProperty{
//   				},
//   				hdr10Settings: &hdr10SettingsProperty{
//   					maxCll: jsii.Number(123),
//   					maxFall: jsii.Number(123),
//   				},
//   				rec601Settings: &rec601SettingsProperty{
//   				},
//   				rec709Settings: &rec709SettingsProperty{
//   				},
//   			},
//   			filterSettings: &h265FilterSettingsProperty{
//   				temporalFilterSettings: &temporalFilterSettingsProperty{
//   					postFilterSharpening: jsii.String("postFilterSharpening"),
//   					strength: jsii.String("strength"),
//   				},
//   			},
//   			fixedAfd: jsii.String("fixedAfd"),
//   			flickerAq: jsii.String("flickerAq"),
//   			framerateDenominator: jsii.Number(123),
//   			framerateNumerator: jsii.Number(123),
//   			gopClosedCadence: jsii.Number(123),
//   			gopSize: jsii.Number(123),
//   			gopSizeUnits: jsii.String("gopSizeUnits"),
//   			level: jsii.String("level"),
//   			lookAheadRateControl: jsii.String("lookAheadRateControl"),
//   			maxBitrate: jsii.Number(123),
//   			minIInterval: jsii.Number(123),
//   			parDenominator: jsii.Number(123),
//   			parNumerator: jsii.Number(123),
//   			profile: jsii.String("profile"),
//   			qvbrQualityLevel: jsii.Number(123),
//   			rateControlMode: jsii.String("rateControlMode"),
//   			scanType: jsii.String("scanType"),
//   			sceneChangeDetect: jsii.String("sceneChangeDetect"),
//   			slices: jsii.Number(123),
//   			tier: jsii.String("tier"),
//   			timecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   		mpeg2Settings: &mpeg2SettingsProperty{
//   			adaptiveQuantization: jsii.String("adaptiveQuantization"),
//   			afdSignaling: jsii.String("afdSignaling"),
//   			colorMetadata: jsii.String("colorMetadata"),
//   			colorSpace: jsii.String("colorSpace"),
//   			displayAspectRatio: jsii.String("displayAspectRatio"),
//   			filterSettings: &mpeg2FilterSettingsProperty{
//   				temporalFilterSettings: &temporalFilterSettingsProperty{
//   					postFilterSharpening: jsii.String("postFilterSharpening"),
//   					strength: jsii.String("strength"),
//   				},
//   			},
//   			fixedAfd: jsii.String("fixedAfd"),
//   			framerateDenominator: jsii.Number(123),
//   			framerateNumerator: jsii.Number(123),
//   			gopClosedCadence: jsii.Number(123),
//   			gopNumBFrames: jsii.Number(123),
//   			gopSize: jsii.Number(123),
//   			gopSizeUnits: jsii.String("gopSizeUnits"),
//   			scanType: jsii.String("scanType"),
//   			subgopLength: jsii.String("subgopLength"),
//   			timecodeInsertion: jsii.String("timecodeInsertion"),
//   		},
//   	},
//   	height: jsii.Number(123),
//   	name: jsii.String("name"),
//   	respondToAfd: jsii.String("respondToAfd"),
//   	scalingBehavior: jsii.String("scalingBehavior"),
//   	sharpness: jsii.Number(123),
//   	width: jsii.Number(123),
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

