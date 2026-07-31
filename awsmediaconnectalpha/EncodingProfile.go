package awsmediaconnectalpha


// Encoding profiles supported when transcoding an NDI source to a transport stream.
//
// Example:
//   var stack Stack
//   var ndiVpcInterface VpcInterfaceConfig
//   var efaInterface VpcInterfaceConfig
//   var videoStream MediaStream
//
//
//   // NDI requires LARGE, an encoding profile, and at least one discovery server
//   // NDI requires LARGE, an encoding profile, and at least one discovery server
//   awsmediaconnectalpha.NewFlow(stack, jsii.String("NdiFlow"), &FlowProps{
//   	FlowSize: awsmediaconnectalpha.FlowSize_LARGE(),
//   	NdiConfig: &NdiConfig{
//   		NdiState: awsmediaconnectalpha.State_ENABLED,
//   		NdiDiscoveryServers: []NdiDiscoveryServerConfig{
//   			&NdiDiscoveryServerConfig{
//   				DiscoveryServerAddress: jsii.String("10.0.0.10"),
//   				VpcInterface: ndiVpcInterface,
//   			},
//   		},
//   	},
//   	EncodingConfig: &EncodingConfig{
//   		EncodingProfile: awsmediaconnectalpha.EncodingProfile_CONTRIBUTION_H264_DEFAULT,
//   	},
//   	Source: awsmediaconnectalpha.SourceConfiguration_Ndi(&SourceNdi{
//   		FlowSourceName: jsii.String("ndi-source"),
//   	}),
//   })
//
//   // CDI and JPEG XS require LARGE_4X
//   // CDI and JPEG XS require LARGE_4X
//   awsmediaconnectalpha.NewFlow(stack, jsii.String("CdiFlow"), &FlowProps{
//   	FlowSize: awsmediaconnectalpha.FlowSize_LARGE_4X(),
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
type EncodingProfile string

const (
	// H.264 profile tuned for contribution workflows — higher quality, higher bitrate.
	// Experimental.
	EncodingProfile_CONTRIBUTION_H264_DEFAULT EncodingProfile = "CONTRIBUTION_H264_DEFAULT"
	// H.264 profile tuned for distribution workflows — lower bitrate, wider compatibility.
	// Experimental.
	EncodingProfile_DISTRIBUTION_H264_DEFAULT EncodingProfile = "DISTRIBUTION_H264_DEFAULT"
)

