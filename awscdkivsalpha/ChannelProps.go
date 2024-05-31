package awscdkivsalpha


// Properties for creating a new Channel.
//
// Example:
//   myChannel := ivs.NewChannel(this, jsii.String("myChannel"), &ChannelProps{
//   	Type: ivs.ChannelType_ADVANCED_HD,
//   	Preset: ivs.Preset_CONSTRAINED_BANDWIDTH_DELIVERY,
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
	// An optional transcode preset for the channel.
	//
	// Can be used for ADVANCED_HD and ADVANCED_SD channel types.
	// When LOW or STANDARD is used, the preset will be overridden and set to none regardless of the value provided.
	// Default: - Preset.HIGHER_BANDWIDTH_DELIVERY if channelType is ADVANCED_SD or ADVANCED_HD, none otherwise
	//
	// Experimental.
	Preset Preset `field:"optional" json:"preset" yaml:"preset"`
	// The channel type, which determines the allowable resolution and bitrate.
	//
	// If you exceed the allowable resolution or bitrate, the stream will disconnect immediately.
	// Default: ChannelType.STANDARD
	//
	// Experimental.
	Type ChannelType `field:"optional" json:"type" yaml:"type"`
}

