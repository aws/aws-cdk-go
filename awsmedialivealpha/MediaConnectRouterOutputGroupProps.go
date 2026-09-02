package awsmedialivealpha


// Properties for a MediaConnect Router output group.
//
// Delivers each channel pipeline to a MediaConnect Router. The downstream routing (which router
// input each pipeline feeds) is configured on the MediaConnect side, referencing this group by
// `name` and pipeline id.
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
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mediaconnectroutergroupsettings.html
//
// Experimental.
type MediaConnectRouterOutputGroupProps struct {
	// The Availability Zones in which to write output to the MediaConnect Router.
	//
	// Provide exactly
	// one AZ for a `SINGLE_PIPELINE` channel, or two (one per pipeline) for a `STANDARD` channel.
	// Experimental.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// The name of this output group.
	//
	// Used as the destination reference ID. Underscores are normalised to hyphens internally.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The outputs for this output group.
	//
	// MediaConnect Router outputs use an MPEG-TS container, so each output may multiplex one video
	// encode and one or more audio encodes (as with UDP).
	// Experimental.
	Outputs *[]*MediaConnectRouterOutputDefinition `field:"required" json:"outputs" yaml:"outputs"`
	// Transit-encryption settings, applied per channel pipeline.
	// Default: - AUTOMATIC (service-managed) transit encryption on every pipeline.
	//
	// Experimental.
	RouterSettings MediaConnectRouterSettings `field:"optional" json:"routerSettings" yaml:"routerSettings"`
}

