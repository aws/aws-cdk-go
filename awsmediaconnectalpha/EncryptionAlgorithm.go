package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Encryption Algorithms used in AWS Elemental MediaConnect.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var role IRole
//   var secret ISecret
//
//
//   entitlement := awsmediaconnectalpha.NewFlowEntitlement(stack, jsii.String("MyEntitlement"), &FlowEntitlementProps{
//   	Flow: flow,
//   	Description: jsii.String("Grant partner access to live feed"),
//   	Subscribers: []*string{
//   		jsii.String("111122223333"),
//   	},
//   	Encryption: &StaticKeyEncryption{
//   		Role: *Role,
//   		Secret: *Secret,
//   		Algorithm: awsmediaconnectalpha.EncryptionAlgorithm_AES256(),
//   	},
//   })
//
// Experimental.
type EncryptionAlgorithm interface {
	// The encryption algorithm string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EncryptionAlgorithm
type jsiiProxy_EncryptionAlgorithm struct {
	_ byte // padding
}

func (j *jsiiProxy_EncryptionAlgorithm) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom encryption algorithm value.
// Experimental.
func EncryptionAlgorithm_Of(value *string) EncryptionAlgorithm {
	_init_.Initialize()

	if err := validateEncryptionAlgorithm_OfParameters(value); err != nil {
		panic(err)
	}
	var returns EncryptionAlgorithm

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.EncryptionAlgorithm",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func EncryptionAlgorithm_AES128() EncryptionAlgorithm {
	_init_.Initialize()
	var returns EncryptionAlgorithm
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.EncryptionAlgorithm",
		"AES128",
		&returns,
	)
	return returns
}

func EncryptionAlgorithm_AES192() EncryptionAlgorithm {
	_init_.Initialize()
	var returns EncryptionAlgorithm
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.EncryptionAlgorithm",
		"AES192",
		&returns,
	)
	return returns
}

func EncryptionAlgorithm_AES256() EncryptionAlgorithm {
	_init_.Initialize()
	var returns EncryptionAlgorithm
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.EncryptionAlgorithm",
		"AES256",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EncryptionAlgorithm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

