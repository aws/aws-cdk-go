package awscdkivsalpha


// The channel type, which determines the allowable resolution and bitrate.
//
// If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
//
// Example:
//   myChannel := ivs.NewChannel(this, jsii.String("myChannel"), &ChannelProps{
//   	Type: ivs.ChannelType_ADVANCED_HD,
//   	Preset: ivs.Preset_CONSTRAINED_BANDWIDTH_DELIVERY,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html
//
// Experimental.
type ChannelType string

const (
	// Multiple qualities are generated from the original input, to automatically give viewers the best experience for their devices and network conditions.
	//
	// Transcoding allows higher playback quality across a range of download speeds. Resolution can be up to 1080p and bitrate can be up to 8.5 Mbps.
	// Audio is transcoded only for renditions 360p and below; above that, audio is passed through.
	// Experimental.
	ChannelType_STANDARD ChannelType = "STANDARD"
	// Delivers the original input to viewers.
	//
	// The viewer’s video-quality choice is limited to the original input.
	// Experimental.
	ChannelType_BASIC ChannelType = "BASIC"
	// Multiple qualities are generated from the original input, to automatically give viewers the best experience for their devices and network conditions.
	//
	// Input resolution can be up to 1080p and bitrate can be up to 8.5 Mbps; output is capped at SD quality (480p).
	// Audio for all renditions is transcoded, and an audio-only rendition is available.
	// Experimental.
	ChannelType_ADVANCED_SD ChannelType = "ADVANCED_SD"
	// Multiple qualities are generated from the original input, to automatically give viewers the best experience for their devices and network conditions.
	//
	// Input resolution can be up to 1080p and bitrate can be up to 8.5 Mbps; output is capped at HD quality (720p).
	// Audio for all renditions is transcoded, and an audio-only rendition is available.
	// Experimental.
	ChannelType_ADVANCED_HD ChannelType = "ADVANCED_HD"
)

