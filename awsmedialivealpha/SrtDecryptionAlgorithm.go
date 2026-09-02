package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The encryption algorithm for SRT decryption.
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
type SrtDecryptionAlgorithm interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for SrtDecryptionAlgorithm
type jsiiProxy_SrtDecryptionAlgorithm struct {
	_ byte // padding
}

func (j *jsiiProxy_SrtDecryptionAlgorithm) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func SrtDecryptionAlgorithm_Of(value *string) SrtDecryptionAlgorithm {
	_init_.Initialize()

	if err := validateSrtDecryptionAlgorithm_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SrtDecryptionAlgorithm

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SrtDecryptionAlgorithm",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SrtDecryptionAlgorithm_AES128() SrtDecryptionAlgorithm {
	_init_.Initialize()
	var returns SrtDecryptionAlgorithm
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtDecryptionAlgorithm",
		"AES128",
		&returns,
	)
	return returns
}

func SrtDecryptionAlgorithm_AES192() SrtDecryptionAlgorithm {
	_init_.Initialize()
	var returns SrtDecryptionAlgorithm
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtDecryptionAlgorithm",
		"AES192",
		&returns,
	)
	return returns
}

func SrtDecryptionAlgorithm_AES256() SrtDecryptionAlgorithm {
	_init_.Initialize()
	var returns SrtDecryptionAlgorithm
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtDecryptionAlgorithm",
		"AES256",
		&returns,
	)
	return returns
}

