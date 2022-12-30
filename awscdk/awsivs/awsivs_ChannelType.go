package awsivs


// The channel type, which determines the allowable resolution and bitrate.
//
// If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
// Experimental.
type ChannelType string

const (
	// Multiple qualities are generated from the original input, to automatically give viewers the best experience for their devices and network conditions.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html
	//
	// Experimental.
	ChannelType_STANDARD ChannelType = "STANDARD"
	// delivers the original input to viewers.
	//
	// The viewer’s video-quality choice is limited to the original input.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html
	//
	// Experimental.
	ChannelType_BASIC ChannelType = "BASIC"
)

