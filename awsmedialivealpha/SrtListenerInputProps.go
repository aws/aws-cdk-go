package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// Properties for an SRT listener input.
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
type SrtListenerInputProps struct {
	// The input security groups.
	//
	// Required for SRT listener inputs.
	// Experimental.
	InputSecurityGroups *[]interfacesawsmedialive.IInputSecurityGroupRef `field:"required" json:"inputSecurityGroups" yaml:"inputSecurityGroups"`
	// Decryption settings for the SRT connection.
	// Default: - no decryption.
	//
	// Experimental.
	Decryption *SrtDecryptionProps `field:"optional" json:"decryption" yaml:"decryption"`
	// Whether this is a STANDARD (two-pipeline) or SINGLE_PIPELINE input.
	//
	// A STANDARD input creates
	// two listener endpoints for pipeline redundancy.
	// Default: ChannelClass.SINGLE_PIPELINE
	//
	// Experimental.
	InputClass ChannelClass `field:"optional" json:"inputClass" yaml:"inputClass"`
	// The minimum latency.
	// Default: - service default.
	//
	// Experimental.
	MinimumLatency awscdk.Duration `field:"optional" json:"minimumLatency" yaml:"minimumLatency"`
	// The stream ID that the upstream system uses when connecting to this listener.
	// Default: - no stream ID.
	//
	// Experimental.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

