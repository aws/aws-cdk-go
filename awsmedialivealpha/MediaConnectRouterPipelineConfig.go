package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Per-pipeline settings for a MediaConnect Router output destination.
//
// Today this carries only transit encryption, but it is a struct so that future per-pipeline
// MediaConnect Router settings can be added without an API break.
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
type MediaConnectRouterPipelineConfig struct {
	// A Secrets Manager secret holding the transit-encryption passphrase.
	//
	// When set, the pipeline
	// uses `SECRETS_MANAGER` encryption; the channel role is granted read access to the secret.
	// Default: - AUTOMATIC (service-managed) transit encryption.
	//
	// Experimental.
	EncryptionSecret awssecretsmanager.ISecret `field:"optional" json:"encryptionSecret" yaml:"encryptionSecret"`
}

