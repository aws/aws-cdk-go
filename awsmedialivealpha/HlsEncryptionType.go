package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS encryption type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsEncryptionType := medialive_alpha.HlsEncryptionType_Of(jsii.String("value"))
//
// Experimental.
type HlsEncryptionType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsEncryptionType
type jsiiProxy_HlsEncryptionType struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsEncryptionType) Value() *string {
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
func HlsEncryptionType_Of(value *string) HlsEncryptionType {
	_init_.Initialize()

	if err := validateHlsEncryptionType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsEncryptionType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsEncryptionType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsEncryptionType_AES128() HlsEncryptionType {
	_init_.Initialize()
	var returns HlsEncryptionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsEncryptionType",
		"AES128",
		&returns,
	)
	return returns
}

func HlsEncryptionType_SAMPLE_AES() HlsEncryptionType {
	_init_.Initialize()
	var returns HlsEncryptionType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsEncryptionType",
		"SAMPLE_AES",
		&returns,
	)
	return returns
}

