package awsmedialivealpha


// Per-pipeline settings for `MediaConnectRouterSettings.perPipeline()`.
//
// MediaLive maps a channel's pipelines to MediaConnect Router output destinations positionally;
// the console labels them "Destination A" (pipeline 0) and "Destination B" (pipeline 1). An
// omitted pipeline uses AUTOMATIC transit encryption.
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//   var passphrase ISecret
//   var passphrase1 ISecret
//
//
//   // AUTOMATIC encryption on every pipeline (MPEG-TS container, like UDP)
//   medialive.OutputGroupConfiguration_MediaConnectRouter(&MediaConnectRouterOutputGroupProps{
//   	Name: jsii.String("router_out"),
//   	AvailabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   	},
//   	Outputs: []MediaConnectRouterOutputDefinition{
//   		&MediaConnectRouterOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("router_ts"),
//   		},
//   	},
//   })
//
//   // One shared Secrets Manager passphrase across all pipelines (SECRETS_MANAGER encryption)
//   medialive.OutputGroupConfiguration_MediaConnectRouter(&MediaConnectRouterOutputGroupProps{
//   	Name: jsii.String("router_out"),
//   	AvailabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   	},
//   	RouterSettings: medialive.MediaConnectRouterSettings_Shared(&MediaConnectRouterPipelineConfig{
//   		EncryptionSecret: passphrase,
//   	}),
//   	Outputs: []MediaConnectRouterOutputDefinition{
//   		&MediaConnectRouterOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("router_ts"),
//   		},
//   	},
//   })
//
//   // Distinct encryption per pipeline — an omitted pipeline stays AUTOMATIC (STANDARD channels)
//   medialive.OutputGroupConfiguration_MediaConnectRouter(&MediaConnectRouterOutputGroupProps{
//   	Name: jsii.String("router_out"),
//   	AvailabilityZones: []*string{
//   		jsii.String("us-east-1a"),
//   		jsii.String("us-east-1b"),
//   	},
//   	RouterSettings: medialive.MediaConnectRouterSettings_PerPipeline(&MediaConnectRouterPerPipelineSettings{
//   		Pipeline1: &MediaConnectRouterPipelineConfig{
//   			EncryptionSecret: passphrase1,
//   		},
//   	}),
//   	Outputs: []MediaConnectRouterOutputDefinition{
//   		&MediaConnectRouterOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("router_ts"),
//   		},
//   	},
//   })
//
// Experimental.
type MediaConnectRouterPerPipelineSettings struct {
	// Settings for pipeline 0 ("Destination A" in the MediaLive console).
	// Default: - AUTOMATIC transit encryption.
	//
	// Experimental.
	Pipeline0 *MediaConnectRouterPipelineConfig `field:"optional" json:"pipeline0" yaml:"pipeline0"`
	// Settings for pipeline 1 ("Destination B" in the console).
	//
	// STANDARD channels only —
	// SINGLE_PIPELINE channels have no second pipeline.
	// Default: - AUTOMATIC transit encryption.
	//
	// Experimental.
	Pipeline1 *MediaConnectRouterPipelineConfig `field:"optional" json:"pipeline1" yaml:"pipeline1"`
}

