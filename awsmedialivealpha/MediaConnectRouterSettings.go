package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Transit-encryption settings for a MediaConnect Router output group, applied per channel pipeline.
//
// Omit the output group's `routerSettings` entirely for AUTOMATIC (service-managed) encryption on
// every pipeline. Use `shared()` to apply one secret across all pipelines, or `perPipeline()` to
// control each pipeline independently.
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
type MediaConnectRouterSettings interface {
}

// The jsii proxy struct for MediaConnectRouterSettings
type jsiiProxy_MediaConnectRouterSettings struct {
	_ byte // padding
}

// Experimental.
func NewMediaConnectRouterSettings_Override(m MediaConnectRouterSettings) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.MediaConnectRouterSettings",
		nil, // no parameters
		m,
	)
}

// Configure each pipeline independently.
//
// An omitted pipeline uses AUTOMATIC encryption.
// Experimental.
func MediaConnectRouterSettings_PerPipeline(settings *MediaConnectRouterPerPipelineSettings) MediaConnectRouterSettings {
	_init_.Initialize()

	if err := validateMediaConnectRouterSettings_PerPipelineParameters(settings); err != nil {
		panic(err)
	}
	var returns MediaConnectRouterSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MediaConnectRouterSettings",
		"perPipeline",
		[]interface{}{settings},
		&returns,
	)

	return returns
}

// Apply the same settings to every pipeline.
// Experimental.
func MediaConnectRouterSettings_Shared(settings *MediaConnectRouterPipelineConfig) MediaConnectRouterSettings {
	_init_.Initialize()

	if err := validateMediaConnectRouterSettings_SharedParameters(settings); err != nil {
		panic(err)
	}
	var returns MediaConnectRouterSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MediaConnectRouterSettings",
		"shared",
		[]interface{}{settings},
		&returns,
	)

	return returns
}

