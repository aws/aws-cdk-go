package awsmediaconnectalpha


// Configuration for NDI (SpeedHQ) source.
//
// NDI sources are ingested from NDI senders inside your VPC. The flow must be
// configured with `flowSize: FlowSize.LARGE` and `ndiConfig.ndiState = ENABLED`
// with at least one NDI discovery server. The VPC connectivity for NDI is
// configured on the discovery server entries in `NdiConfig`, not on the source
// itself.
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
type SourceNdi struct {
	// A description of the source.
	//
	// This description appears only on the MediaConnect
	// console and will not be seen by the end user.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the source.
	// Default: - a name is generated automatically.
	//
	// Experimental.
	FlowSourceName *string `field:"optional" json:"flowSourceName" yaml:"flowSourceName"`
	// The exact name of an existing NDI sender that's registered with your discovery server.
	//
	// If included, the format of this name must be `MACHINENAME (ProgramName)`.
	//
	// If not specified, you can select the upstream NDI sender from the console
	// after starting the flow.
	// Default: - select the NDI sender after starting the flow.
	//
	// Experimental.
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
}

