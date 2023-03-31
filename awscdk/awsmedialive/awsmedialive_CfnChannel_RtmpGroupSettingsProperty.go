package awsmedialive


// The configuration of an RTMP output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtmpGroupSettingsProperty := &rtmpGroupSettingsProperty{
//   	adMarkers: []*string{
//   		jsii.String("adMarkers"),
//   	},
//   	authenticationScheme: jsii.String("authenticationScheme"),
//   	cacheFullBehavior: jsii.String("cacheFullBehavior"),
//   	cacheLength: jsii.Number(123),
//   	captionData: jsii.String("captionData"),
//   	inputLossAction: jsii.String("inputLossAction"),
//   	restartDelay: jsii.Number(123),
//   }
//
type CfnChannel_RtmpGroupSettingsProperty struct {
	// Choose the ad marker type for this output group.
	//
	// MediaLive will create a message based on the content of each SCTE-35 message, format it for that marker type, and insert it in the datastream.
	AdMarkers *[]*string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// An authentication scheme to use when connecting with a CDN.
	AuthenticationScheme *string `field:"optional" json:"authenticationScheme" yaml:"authenticationScheme"`
	// Controls behavior when the content cache fills up.
	//
	// If a remote origin server stalls the RTMP connection and doesn't accept content fast enough, the media cache fills up. When the cache reaches the duration specified by cacheLength, the cache stops accepting new content. If set to disconnectImmediately, the RTMP output forces a disconnect. Clear the media cache, and reconnect after restartDelay seconds. If set to waitForServer, the RTMP output waits up to 5 minutes to allow the origin server to begin accepting data again.
	CacheFullBehavior *string `field:"optional" json:"cacheFullBehavior" yaml:"cacheFullBehavior"`
	// The cache length, in seconds, that is used to calculate buffer size.
	CacheLength *float64 `field:"optional" json:"cacheLength" yaml:"cacheLength"`
	// Controls the types of data that pass to onCaptionInfo outputs.
	//
	// If set to all, 608 and 708 carried DTVCC data is passed. If set to field1AndField2608, DTVCC data is stripped out, but 608 data from both fields is passed. If set to field1608, only the data carried in 608 from field 1 video is passed.
	CaptionData *string `field:"optional" json:"captionData" yaml:"captionData"`
	// Controls the behavior of this RTMP group if the input becomes unavailable.
	//
	// emitOutput: Emit a slate until the input returns. pauseOutput: Stop transmitting data until the input returns. This does not close the underlying RTMP connection.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// If a streaming output fails, the number of seconds to wait until a restart is initiated.
	//
	// A value of 0 means never restart.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
}

