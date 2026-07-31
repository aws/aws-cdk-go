package awsmediaconnectalpha


// Options for FMTP Video.
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
type FmtpVideo struct {
	// The format used for the representation of color.
	// Default: - no colorimetry specified.
	//
	// Experimental.
	Colorimetry Colorimetry `field:"optional" json:"colorimetry" yaml:"colorimetry"`
	// The frame rate for the video stream.
	// Default: - no frame rate specified.
	//
	// Experimental.
	ExactFramerate Framerate `field:"optional" json:"exactFramerate" yaml:"exactFramerate"`
	// The pixel aspect ratio (PAR) of the video — the shape of a single pixel, not the display aspect ratio.
	//
	// Use `PixelAspectRatio.SQUARE` for most modern content.
	// Default: - no pixel aspect ratio specified.
	//
	// Experimental.
	Par PixelAspectRatio `field:"optional" json:"par" yaml:"par"`
	// The type of compression that was used to smooth the video's appearance.
	// Default: - no scan mode specified.
	//
	// Experimental.
	ScanMode ScanMode `field:"optional" json:"scanMode" yaml:"scanMode"`
	// The transfer characteristic system (TCS) that is used in the video.
	// Default: - no TCS specified.
	//
	// Experimental.
	Tcs Tcs `field:"optional" json:"tcs" yaml:"tcs"`
	// The encoding range of the video.
	// Default: - no video range specified.
	//
	// Experimental.
	VideoRange VideoRange `field:"optional" json:"videoRange" yaml:"videoRange"`
}

