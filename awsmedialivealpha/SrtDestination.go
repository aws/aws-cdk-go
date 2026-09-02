package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A destination for an SRT output group.
//
// Use the static factory methods to create.
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
type SrtDestination interface {
}

// The jsii proxy struct for SrtDestination
type jsiiProxy_SrtDestination struct {
	_ byte // padding
}

// Experimental.
func NewSrtDestination_Override(s SrtDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.SrtDestination",
		nil, // no parameters
		s,
	)
}

// Create a caller-mode SRT destination.
//
// MediaLive connects to the remote listener.
// Experimental.
func SrtDestination_Caller(props *SrtCallerDestinationProps) SrtDestination {
	_init_.Initialize()

	if err := validateSrtDestination_CallerParameters(props); err != nil {
		panic(err)
	}
	var returns SrtDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SrtDestination",
		"caller",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a caller-mode SRT destination from a full SRT URL.
//
// Use this when you already have a URL rather than a separate host and port — for example a
// MediaConnect Router Input's ingest endpoint (`routerInput.endpoints[0].url`).
// Experimental.
func SrtDestination_CallerUrl(url *string, options *SrtCallerUrlOptions) SrtDestination {
	_init_.Initialize()

	if err := validateSrtDestination_CallerUrlParameters(url, options); err != nil {
		panic(err)
	}
	var returns SrtDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SrtDestination",
		"callerUrl",
		[]interface{}{url, options},
		&returns,
	)

	return returns
}

// Create a listener-mode SRT destination.
//
// MediaLive listens for incoming connections.
// Experimental.
func SrtDestination_Listener(props *SrtListenerDestinationProps) SrtDestination {
	_init_.Initialize()

	if err := validateSrtDestination_ListenerParameters(props); err != nil {
		panic(err)
	}
	var returns SrtDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SrtDestination",
		"listener",
		[]interface{}{props},
		&returns,
	)

	return returns
}

