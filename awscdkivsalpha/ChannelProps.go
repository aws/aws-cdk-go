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
	// Indicates which content-packaging format is used (MPEG-TS or fMP4).
	//
	// If `multitrackInputConfiguration` is specified, only fMP4 can be used.
	// Otherwise, `containerFormat` may be set to `ContainerFormat.TS` or `ContainerFormat.FRAGMENTED_MP4`.
	// Default: - `ContainerFormat.FRAGMENTED_MP4` is automatically set when the `multitrackInputConfiguration` is specified. If not specified, it remains undefined and uses the IVS default setting (TS).
	//
	// Experimental.
	ContainerFormat ContainerFormat `field:"optional" json:"containerFormat" yaml:"containerFormat"`
	// Whether the channel allows insecure RTMP ingest.
	// Default: false.
	//
	// Experimental.
	InsecureIngest *bool `field:"optional" json:"insecureIngest" yaml:"insecureIngest"`
	// Channel latency mode.
	// Default: LatencyMode.LOW
	//
	// Experimental.
	LatencyMode LatencyMode `field:"optional" json:"latencyMode" yaml:"latencyMode"`
	// Object specifying multitrack input configuration. You must specify `multitrackInputConfiguration` if you want to use MultiTrack Video.
	//
	// `multitrackInputConfiguration` is only supported for `ChannelType.STANDARD`.
	// See: https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/multitrack-video.html
	//
	// Default: undefined - IVS default setting is not use MultiTrack Video.
	//
	// Experimental.
	MultitrackInputConfiguration *MultitrackInputConfiguration `field:"optional" json:"multitrackInputConfiguration" yaml:"multitrackInputConfiguration"`
	// An optional transcode preset for the channel.
	//
	// Can be used for ADVANCED_HD and ADVANCED_SD channel types.
	// When LOW or STANDARD is used, the preset will be overridden and set to none regardless of the value provided.
	// Default: - Preset.HIGHER_BANDWIDTH_DELIVERY if channelType is ADVANCED_SD or ADVANCED_HD, none otherwise
	//
	// Experimental.
	Preset Preset `field:"optional" json:"preset" yaml:"preset"`
	// A recording configuration for the channel.
	// Default: - recording is disabled.
	//
	// Experimental.
	RecordingConfiguration IRecordingConfiguration `field:"optional" json:"recordingConfiguration" yaml:"recordingConfiguration"`
	// The channel type, which determines the allowable resolution and bitrate.
	//
	// If you exceed the allowable resolution or bitrate, the stream will disconnect immediately.
	// Default: ChannelType.STANDARD
	//
	// Experimental.
	Type ChannelType `field:"optional" json:"type" yaml:"type"`
}

