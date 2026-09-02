package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// SRT listener destination properties.
//
// In listener mode, MediaLive opens a socket on `listenerPort` and waits for the downstream
// system to connect. The downstream system needs the channel's outbound IP and this port —
// AWS does not require (or use) a destination URL in listener mode.
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
type SrtListenerDestinationProps struct {
	// The Secrets Manager secret containing the encryption passphrase.
	//
	// [disable-awslint:prefer-ref-interface].
	// Experimental.
	EncryptionPassphraseSecret awssecretsmanager.ISecret `field:"required" json:"encryptionPassphraseSecret" yaml:"encryptionPassphraseSecret"`
	// The port that MediaLive will listen on.
	//
	// AWS reserves the range 5000–5200 for SRT listener output.
	// Experimental.
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
	// The stream ID for the SRT connection.
	// Default: - no stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

