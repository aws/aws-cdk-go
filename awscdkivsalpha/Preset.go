package awscdkivsalpha


// An optional transcode preset for the channel.
//
// This is selectable only for ADVANCED_HD and ADVANCED_SD channel types.
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
type Preset string

const (
	// Use a lower bitrate than STANDARD for each quality level.
	//
	// Use it if you have low download bandwidth and/or simple video content (e.g., talking heads).
	// Experimental.
	Preset_CONSTRAINED_BANDWIDTH_DELIVERY Preset = "CONSTRAINED_BANDWIDTH_DELIVERY"
	// Use a higher bitrate for each quality level.
	//
	// Use it if you have high download bandwidth and/or complex video content (e.g., flashes and quick scene changes).
	// Experimental.
	Preset_HIGHER_BANDWIDTH_DELIVERY Preset = "HIGHER_BANDWIDTH_DELIVERY"
)

