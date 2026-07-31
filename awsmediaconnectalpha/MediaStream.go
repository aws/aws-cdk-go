package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Simplified configuration for Media Streams.
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
type MediaStream interface {
}

// The jsii proxy struct for MediaStream
type jsiiProxy_MediaStream struct {
	_ byte // padding
}

// Configuration for MediaStream ancillary data.
// Experimental.
func MediaStream_AncillaryData(config *MediaStreamAncillaryData) MediaStream {
	_init_.Initialize()

	if err := validateMediaStream_AncillaryDataParameters(config); err != nil {
		panic(err)
	}
	var returns MediaStream

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStream",
		"ancillaryData",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Configuration for MediaStream audio.
// Experimental.
func MediaStream_Audio(config *MediaStreamAudio) MediaStream {
	_init_.Initialize()

	if err := validateMediaStream_AudioParameters(config); err != nil {
		panic(err)
	}
	var returns MediaStream

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStream",
		"audio",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Configuration for MediaStream Video.
// Experimental.
func MediaStream_Video(config *MediaStreamVideo) MediaStream {
	_init_.Initialize()

	if err := validateMediaStream_VideoParameters(config); err != nil {
		panic(err)
	}
	var returns MediaStream

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.MediaStream",
		"video",
		[]interface{}{config},
		&returns,
	)

	return returns
}

