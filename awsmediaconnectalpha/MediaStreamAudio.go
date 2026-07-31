package awsmediaconnectalpha


// A media stream represents one component of your content, such as video, audio, or ancillary data.
//
// After you add a media stream to your flow, you can associate it with sources and outputs that use
// the ST 2110 JPEG XS or CDI protocol.
//
// Example:
//   var stack Stack
//
//
//   // Create media streams
//   videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
//   	MediaStreamId: jsii.Number(1),
//   	MediaStreamName: jsii.String("video-stream"),
//   	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_HD_1080P(),
//   	Fmtp: &FmtpVideo{
//   		Colorimetry: awsmediaconnectalpha.Colorimetry_BT709(),
//   		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_29_97(),
//   		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
//   		VideoRange: awsmediaconnectalpha.VideoRange_NARROW(),
//   		ScanMode: awsmediaconnectalpha.ScanMode_PROGRESSIVE(),
//   		Tcs: awsmediaconnectalpha.Tcs_SDR(),
//   	},
//   })
//
//   audioStream := awsmediaconnectalpha.MediaStream_Audio(&MediaStreamAudio{
//   	MediaStreamId: jsii.Number(2),
//   	MediaStreamName: jsii.String("audio-stream"),
//   	ChannelOrder: awsmediaconnectalpha.AudioStreamOrderOptions_STANDARD_STEREO(),
//   })
//
//   // Add to flow
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_Router(),
//   	MediaStreams: []MediaStream{
//   		videoStream,
//   		audioStream,
//   	},
//   })
//
// Experimental.
type MediaStreamAudio struct {
	// A unique identifier for the media stream.
	// Experimental.
	MediaStreamId *float64 `field:"required" json:"mediaStreamId" yaml:"mediaStreamId"`
	// A name that helps you distinguish one media stream from another.
	// Experimental.
	MediaStreamName *string `field:"required" json:"mediaStreamName" yaml:"mediaStreamName"`
	// A description that can help you quickly identify what your media stream is used for.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The format of the audio channel.
	// Default: - the MediaConnect service default.
	//
	// Experimental.
	ChannelOrder AudioStreamOrderOptions `field:"optional" json:"channelOrder" yaml:"channelOrder"`
	// The sample rate for the stream.
	//
	// This value is measured in Hz.
	// Default: - default clock rate for the media stream type.
	//
	// Experimental.
	ClockRate *float64 `field:"optional" json:"clockRate" yaml:"clockRate"`
	// The format type number (sometimes referred to as RTP payload type) of the media stream.
	//
	// MediaConnect assigns this value to the media stream. For ST 2110 JPEG XS outputs, you need to provide this value to the receiver.
	// Default: - assigned by MediaConnect.
	//
	// Experimental.
	Fmt *float64 `field:"optional" json:"fmt" yaml:"fmt"`
	// The audio language, in a format that is recognized by the receiver.
	// Default: - no language specified.
	//
	// Experimental.
	Lang *string `field:"optional" json:"lang" yaml:"lang"`
}

