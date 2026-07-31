package awsmediaconnectalpha


// Configuration for Jpeg XS.
//
// Example:
//   var stack Stack
//   var role IRole
//   var sg1 ISecurityGroup
//   var sg2 ISecurityGroup
//   var subnet ISubnet
//
//
//   efaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("efa-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		sg1,
//   	},
//   	Subnet: subnet,
//   	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_EFA(),
//   })
//
//   enaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("ena-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		sg2,
//   	},
//   	Subnet: subnet,
//   	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_ENA(),
//   })
//
//   videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
//   	MediaStreamId: jsii.Number(1),
//   	MediaStreamName: jsii.String("video"),
//   	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_UHD_2160P(),
//   	Fmtp: &FmtpVideo{
//   		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_59_94(),
//   		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
//   		Colorimetry: awsmediaconnectalpha.Colorimetry_BT2020(),
//   		VideoRange: awsmediaconnectalpha.VideoRange_FULL(),
//   		ScanMode: awsmediaconnectalpha.ScanMode_PROGRESSIVE(),
//   		Tcs: awsmediaconnectalpha.Tcs_PQ(),
//   	},
//   })
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyJpegXsFlow"), &FlowProps{
//   	FlowSize: awsmediaconnectalpha.FlowSize_LARGE_4X(),
//   	 // Required for JPEG XS
//   	VpcInterfaces: []VpcInterfaceConfig{
//   		efaInterface,
//   		enaInterface,
//   	},
//   	MediaStreams: []MediaStream{
//   		videoStream,
//   	},
//   	Source: awsmediaconnectalpha.SourceConfiguration_JpegXs(&SourceJpegXs{
//   		FlowSourceName: jsii.String("jpegxs-source"),
//   		MaxSyncBuffer: jsii.Number(100),
//   		MediaStreamSourceConfigurations: []MediaStreamSourceConfigurationJpegXs{
//   			&MediaStreamSourceConfigurationJpegXs{
//   				Encoding: awsmediaconnectalpha.Encoding_JXSV(),
//   				Port: jsii.Number(5000),
//   				InputInterface: []VpcInterfaceConfig{
//   					efaInterface,
//   					enaInterface,
//   				},
//   				 // 2 interfaces for redundancy
//   				MediaStream: videoStream,
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type SourceJpegXs struct {
	// The media stream that is associated with the source, and the parameters for that association.
	// Experimental.
	MediaStreamSourceConfigurations *[]*MediaStreamSourceConfigurationJpegXs `field:"required" json:"mediaStreamSourceConfigurations" yaml:"mediaStreamSourceConfigurations"`
	// The name of the source.
	// Default: - a name is generated automatically.
	//
	// Experimental.
	FlowSourceName *string `field:"optional" json:"flowSourceName" yaml:"flowSourceName"`
	// The size of the buffer (in ms) to use to sync incoming source data.
	// Default: 100.
	//
	// Experimental.
	MaxSyncBuffer *float64 `field:"optional" json:"maxSyncBuffer" yaml:"maxSyncBuffer"`
	// The VPC interface attachment to use for this bridge source.
	// Default: - no VPC interface.
	//
	// Experimental.
	VpcInterface *[]*VpcInterfaceConfig `field:"optional" json:"vpcInterface" yaml:"vpcInterface"`
}

