package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// RTMP TLS certificate verification mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   rtmpCertificateMode := medialive_alpha.RtmpCertificateMode_Of(jsii.String("value"))
//
// Experimental.
type RtmpCertificateMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RtmpCertificateMode
type jsiiProxy_RtmpCertificateMode struct {
	_ byte // padding
}

func (j *jsiiProxy_RtmpCertificateMode) Value() *string {
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
func RtmpCertificateMode_Of(value *string) RtmpCertificateMode {
	_init_.Initialize()

	if err := validateRtmpCertificateMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RtmpCertificateMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RtmpCertificateMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RtmpCertificateMode_SELF_SIGNED() RtmpCertificateMode {
	_init_.Initialize()
	var returns RtmpCertificateMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpCertificateMode",
		"SELF_SIGNED",
		&returns,
	)
	return returns
}

func RtmpCertificateMode_VERIFY_AUTHENTICITY() RtmpCertificateMode {
	_init_.Initialize()
	var returns RtmpCertificateMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpCertificateMode",
		"VERIFY_AUTHENTICITY",
		&returns,
	)
	return returns
}

