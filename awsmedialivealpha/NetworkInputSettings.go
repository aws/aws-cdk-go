package awsmedialivealpha


// Network input settings for URL pull inputs.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   			NetworkInputSettings: &NetworkInputSettings{
//   				ServerValidation: medialive.ServerValidation_CHECK_CRYPTOGRAPHY_AND_VALIDATE_NAME(),
//   				HlsInputSettings: &HlsInputSettings{
//   					Bandwidth: awscdk.Bitrate_Mbps(jsii.Number(5)),
//   					Scte35Source: medialive.HlsScte35Source_MANIFEST(),
//   				},
//   			},
//   			LogicalInterfaceNames: []*string{
//   				jsii.String("eth0"),
//   				jsii.String("eth1"),
//   			},
//   		},
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type NetworkInputSettings struct {
	// HLS input settings (bandwidth selection, buffer segments, retries).
	// Default: - no HLS input settings.
	//
	// Experimental.
	HlsInputSettings *HlsInputSettings `field:"optional" json:"hlsInputSettings" yaml:"hlsInputSettings"`
	// For a multicast input, filter to content from a specific source IP address (source-specific multicast).
	// Default: - no source IP filter.
	//
	// Experimental.
	MulticastSourceIp *string `field:"optional" json:"multicastSourceIp" yaml:"multicastSourceIp"`
	// HTTPS server certificate validation mode.
	// Default: - service default.
	//
	// Experimental.
	ServerValidation ServerValidation `field:"optional" json:"serverValidation" yaml:"serverValidation"`
}

