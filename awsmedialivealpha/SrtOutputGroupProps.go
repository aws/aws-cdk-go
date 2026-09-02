package awsmedialivealpha


// Properties for an SRT output group.
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
type SrtOutputGroupProps struct {
	// The name of this output group.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The outputs for this SRT output group.
	//
	// Each output includes its own SRT destination.
	// Experimental.
	Outputs *[]*SrtOutputDefinition `field:"required" json:"outputs" yaml:"outputs"`
	// Controls the behavior of this SRT group if the input becomes unavailable.
	// Default: SrtInputLossAction.EMIT_PROGRAM
	//
	// Experimental.
	InputLossAction SrtInputLossAction `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
}

