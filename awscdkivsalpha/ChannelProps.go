package awscdkivsalpha


// Properties for creating a new Channel.
//
// Example:
//   myChannel := ivs.NewChannel(this, jsii.String("Channel"), &ChannelProps{
//   	Authorized: jsii.Boolean(true),
//   })
//
// Experimental.
type ChannelProps struct {
	// Whether the channel is authorized.
	//
	// If you wish to make an authorized channel, you will need to ensure that
	// a PlaybackKeyPair has been uploaded to your account as this is used to
	// validate the signed JWT that is required for authorization.
	// Default: false.
	//
	// Experimental.
	Authorized *bool `field:"optional" json:"authorized" yaml:"authorized"`
	// A name for the channel.
	// Default: Automatically generated name.
	//
	// Experimental.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Channel latency mode.
	// Default: LatencyMode.LOW
	//
	// Experimental.
	LatencyMode LatencyMode `field:"optional" json:"latencyMode" yaml:"latencyMode"`
	// The channel type, which determines the allowable resolution and bitrate.
	//
	// If you exceed the allowable resolution or bitrate, the stream will disconnect immediately.
	// Default: ChannelType.STANDARD
	//
	// Experimental.
	Type ChannelType `field:"optional" json:"type" yaml:"type"`
}

