package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// SRT caller destination properties.
//
// Example:
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//   var passphrase ISecret
//
//
//   // SRT caller to a remote listener
//   medialive.OutputGroupConfiguration_Srt(&SrtOutputGroupProps{
//   	Name: jsii.String("srt_out"),
//   	Outputs: []SrtOutputDefinition{
//   		&SrtOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("srt_caller"),
//   			Destinations: []SrtDestination{
//   				medialive.SrtDestination_Caller(&SrtCallerDestinationProps{
//   					Address: jsii.String("203.0.113.20"),
//   					Port: jsii.Number(5000),
//   					EncryptionPassphraseSecret: passphrase,
//   				}),
//   			},
//   		},
//   	},
//   })
//
//   // SRT listener — MediaLive waits for the downstream system to connect
//   medialive.OutputGroupConfiguration_Srt(&SrtOutputGroupProps{
//   	Name: jsii.String("srt_listen"),
//   	Outputs: []SrtOutputDefinition{
//   		&SrtOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("srt_listener"),
//   			Destinations: []SrtDestination{
//   				medialive.SrtDestination_Listener(&SrtListenerDestinationProps{
//   					ListenerPort: jsii.Number(5000),
//   					EncryptionPassphraseSecret: passphrase,
//   				}),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type SrtCallerDestinationProps struct {
	// The address (IP or host) of the SRT listener to connect to.
	// Experimental.
	Address *string `field:"required" json:"address" yaml:"address"`
	// The Secrets Manager secret containing the encryption passphrase.
	//
	// [disable-awslint:prefer-ref-interface].
	// Experimental.
	EncryptionPassphraseSecret awssecretsmanager.ISecret `field:"required" json:"encryptionPassphraseSecret" yaml:"encryptionPassphraseSecret"`
	// The port of the SRT listener to connect to.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The stream ID for the SRT connection.
	// Default: - no stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

