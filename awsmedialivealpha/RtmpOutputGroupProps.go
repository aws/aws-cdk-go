package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for an RTMP output group.
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Rtmp(&RtmpOutputGroupProps{
//   	Name: jsii.String("social"),
//   	Outputs: []RtmpOutputDefinition{
//   		&RtmpOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("live"),
//   			Destinations: []RtmpDestination{
//   				medialive.RtmpDestination_Url(jsii.String("rtmp://rtmp.example.com/live"), jsii.String("your-stream-key")),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type RtmpOutputGroupProps struct {
	// The name of this output group.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The outputs for this RTMP output group.
	//
	// Each output includes its own RTMP destination.
	// Experimental.
	Outputs *[]*RtmpOutputDefinition `field:"required" json:"outputs" yaml:"outputs"`
	// Choose the ad marker type for this output group.
	// Default: - no ad markers.
	//
	// Experimental.
	AdMarkers *[]RtmpAdMarkers `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// The authentication scheme for the RTMP connection.
	// Default: RtmpAuthenticationScheme.COMMON
	//
	// Experimental.
	AuthenticationScheme RtmpAuthenticationScheme `field:"optional" json:"authenticationScheme" yaml:"authenticationScheme"`
	// Controls behavior when the content cache fills up.
	// Default: - service default.
	//
	// Experimental.
	CacheFullBehavior RtmpCacheFullBehavior `field:"optional" json:"cacheFullBehavior" yaml:"cacheFullBehavior"`
	// The cache length, in seconds, that is used to calculate buffer size.
	// Default: - service default.
	//
	// Experimental.
	CacheLength awscdk.Duration `field:"optional" json:"cacheLength" yaml:"cacheLength"`
	// Controls the types of data that pass to onCaptionInfo outputs.
	// Default: - service default.
	//
	// Experimental.
	CaptionData RtmpCaptionData `field:"optional" json:"captionData" yaml:"captionData"`
	// Controls whether filler NAL units are included in the output.
	// Default: - service default.
	//
	// Experimental.
	IncludeFillerNalUnits RtmpIncludeFillerNalUnits `field:"optional" json:"includeFillerNalUnits" yaml:"includeFillerNalUnits"`
	// Controls the behavior of this RTMP group if the input becomes unavailable.
	// Default: - service default.
	//
	// Experimental.
	InputLossAction RtmpInputLossAction `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// The delay before restarting after a streaming output failure.
	//
	// A value of
	// `Duration.seconds(0)` means never restart.
	// Default: Duration.seconds(1)
	//
	// Experimental.
	RestartDelay awscdk.Duration `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

