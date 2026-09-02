package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth certificate mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothCertificateMode := medialive_alpha.MsSmoothCertificateMode_Of(jsii.String("value"))
//
// Experimental.
type MsSmoothCertificateMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothCertificateMode
type jsiiProxy_MsSmoothCertificateMode struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothCertificateMode) Value() *string {
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
func MsSmoothCertificateMode_Of(value *string) MsSmoothCertificateMode {
	_init_.Initialize()

	if err := validateMsSmoothCertificateMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothCertificateMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothCertificateMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothCertificateMode_SELF_SIGNED() MsSmoothCertificateMode {
	_init_.Initialize()
	var returns MsSmoothCertificateMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothCertificateMode",
		"SELF_SIGNED",
		&returns,
	)
	return returns
}

func MsSmoothCertificateMode_VERIFY_AUTHENTICITY() MsSmoothCertificateMode {
	_init_.Initialize()
	var returns MsSmoothCertificateMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothCertificateMode",
		"VERIFY_AUTHENTICITY",
		&returns,
	)
	return returns
}

