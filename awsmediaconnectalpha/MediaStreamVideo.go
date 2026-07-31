package awsmediaconnectalpha


// A media stream represents one component of your content, such as video, audio, or ancillary data.
//
// After you add a media stream to your flow, you can associate it with sources and outputs that use
// the ST 2110 JPEG XS or CDI protocol.
//
// Example:
//   var stack Stack
//   var role IRole
//   var securityGroup ISecurityGroup
//   var subnet ISubnet
//
//
//   efaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("efa-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Subnet: subnet,
//   	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_EFA(),
//   })
//
//   videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
//   	MediaStreamId: jsii.Number(1),
//   	MediaStreamName: jsii.String("video"),
//   	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_HD_1080P(),
//   	Fmtp: &FmtpVideo{
//   		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_29_97(),
//   		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
//   	},
//   })
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyCdiFlow"), &FlowProps{
//   	FlowSize: awsmediaconnectalpha.FlowSize_LARGE_4X(),
//   	 // Required for CDI and JPEG XS
//   	VpcInterfaces: []VpcInterfaceConfig{
//   		efaInterface,
//   	},
//   	MediaStreams: []MediaStream{
//   		videoStream,
//   	},
//   	Source: awsmediaconnectalpha.SourceConfiguration_Cdi(&SourceCdi{
//   		FlowSourceName: jsii.String("cdi-source"),
//   		VpcInterface: efaInterface,
//   		Port: jsii.Number(5000),
//   		MaxSyncBuffer: jsii.Number(100),
//   		MediaStreamSourceConfigurations: []MediaStreamSourceConfigurationCdi{
//   			&MediaStreamSourceConfigurationCdi{
//   				Encoding: awsmediaconnectalpha.Encoding_RAW(),
//   				MediaStream: videoStream,
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type MediaStreamVideo struct {
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
	// Attributes that are related to the media stream.
	// Experimental.
	Fmtp *FmtpVideo `field:"required" json:"fmtp" yaml:"fmtp"`
	// The resolution of the video.
	// Experimental.
	VideoFormat MediaVideoFormat `field:"required" json:"videoFormat" yaml:"videoFormat"`
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
}

