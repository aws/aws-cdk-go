package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for SRT decryption.
//
// Example:
//   var stack Stack
//   var passphrase ISecret
//
//
//   sg := medialive.NewInputSecurityGroup(stack, jsii.String("SrtSg"), &InputSecurityGroupProps{
//   	AllowlistRules: []*string{
//   		jsii.String("203.0.113.0/24"),
//   	},
//   })
//
//   medialive.NewInput(stack, jsii.String("SrtListenerInput"), &InputProps{
//   	InputName: jsii.String("srt-listener"),
//   	Input: medialive.InputConfiguration_SrtListener(&SrtListenerInputProps{
//   		InputSecurityGroups: []IInputSecurityGroupRef{
//   			sg,
//   		},
//   		MinimumLatency: awscdk.Duration_Millis(jsii.Number(500)),
//   		StreamId: jsii.String("my-stream-id"),
//   		Decryption: &SrtDecryptionProps{
//   			Algorithm: medialive.SrtDecryptionAlgorithm_AES256(),
//   			PassphraseSecret: passphrase,
//   		},
//   	}),
//   })
//
// Experimental.
type SrtDecryptionProps struct {
	// The encryption algorithm.
	// Default: - no algorithm.
	//
	// Experimental.
	Algorithm SrtDecryptionAlgorithm `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The Secrets Manager secret containing the passphrase used to decrypt the content.
	// Default: - no passphrase.
	//
	// Experimental.
	PassphraseSecret awssecretsmanager.ISecret `field:"optional" json:"passphraseSecret" yaml:"passphraseSecret"`
}

