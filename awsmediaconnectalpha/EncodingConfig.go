package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Encoding configuration applied to an NDI source when transcoding it to a transport stream for downstream distribution.
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
type EncodingConfig struct {
	// The encoding profile to use when transcoding the NDI source content to a transport stream.
	//
	// You can change this value while the flow is running.
	// Default: - the MediaConnect service default.
	//
	// Experimental.
	EncodingProfile EncodingProfile `field:"optional" json:"encodingProfile" yaml:"encodingProfile"`
	// The maximum video bitrate to use when transcoding the NDI source to a transport stream.
	//
	// The supported range is 10 Mbps – 50 Mbps.
	// Default: - undefined; when omitted, MediaConnect applies 20 Mbps at deploy time.
	//
	// Experimental.
	VideoMaxBitrate awscdk.Bitrate `field:"optional" json:"videoMaxBitrate" yaml:"videoMaxBitrate"`
}

