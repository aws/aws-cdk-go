package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// SRT output encryption type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   srtEncryptionType := medialive_alpha.SrtEncryptionType_AES128()
//
// Experimental.
type SrtEncryptionType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for SrtEncryptionType
type jsiiProxy_SrtEncryptionType struct {
	_ byte // padding
}

func (j *jsiiProxy_SrtEncryptionType) Value() *string {
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
func SrtEncryptionType_Of(value *string) SrtEncryptionType {
	_init_.Initialize()

	if err := validateSrtEncryptionType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SrtEncryptionType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SrtEncryptionType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SrtEncryptionType_AES128() SrtEncryptionType {
	_init_.Initialize()
	var returns SrtEncryptionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtEncryptionType",
		"AES128",
		&returns,
	)
	return returns
}

func SrtEncryptionType_AES192() SrtEncryptionType {
	_init_.Initialize()
	var returns SrtEncryptionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtEncryptionType",
		"AES192",
		&returns,
	)
	return returns
}

func SrtEncryptionType_AES256() SrtEncryptionType {
	_init_.Initialize()
	var returns SrtEncryptionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtEncryptionType",
		"AES256",
		&returns,
	)
	return returns
}

