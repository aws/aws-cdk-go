package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for QVBR rate control.
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
type QvbrRateControlProps struct {
	// The maximum bitrate.
	// Experimental.
	MaxBitrate awscdk.Bitrate `field:"required" json:"maxBitrate" yaml:"maxBitrate"`
	// The QVBR quality level (1-10).
	//
	// Leave unset to let MediaLive infer the target quality from the
	// output resolution and max bitrate.
	// Default: - MediaLive infers the quality level from the resolution and max bitrate.
	//
	// Experimental.
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
}

