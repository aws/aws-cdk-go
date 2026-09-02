package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A DASH role to assign to a captions output (used when the output carries DVB DASH accessibility signaling).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   captionDashRole := medialive_alpha.CaptionDashRole_ALTERNATE()
//
// Experimental.
type CaptionDashRole interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionDashRole
type jsiiProxy_CaptionDashRole struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionDashRole) Value() *string {
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
func CaptionDashRole_Of(value *string) CaptionDashRole {
	_init_.Initialize()

	if err := validateCaptionDashRole_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionDashRole

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionDashRole_ALTERNATE() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"ALTERNATE",
		&returns,
	)
	return returns
}

func CaptionDashRole_CAPTION() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"CAPTION",
		&returns,
	)
	return returns
}

func CaptionDashRole_COMMENTARY() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"COMMENTARY",
		&returns,
	)
	return returns
}

func CaptionDashRole_DESCRIPTION() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"DESCRIPTION",
		&returns,
	)
	return returns
}

func CaptionDashRole_DUB() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"DUB",
		&returns,
	)
	return returns
}

func CaptionDashRole_EASYREADER() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"EASYREADER",
		&returns,
	)
	return returns
}

func CaptionDashRole_EMERGENCY() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"EMERGENCY",
		&returns,
	)
	return returns
}

func CaptionDashRole_FORCED_SUBTITLE() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"FORCED_SUBTITLE",
		&returns,
	)
	return returns
}

func CaptionDashRole_KARAOKE() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"KARAOKE",
		&returns,
	)
	return returns
}

func CaptionDashRole_MAIN() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"MAIN",
		&returns,
	)
	return returns
}

func CaptionDashRole_METADATA() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"METADATA",
		&returns,
	)
	return returns
}

func CaptionDashRole_SUBTITLE() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"SUBTITLE",
		&returns,
	)
	return returns
}

func CaptionDashRole_SUPPLEMENTARY() CaptionDashRole {
	_init_.Initialize()
	var returns CaptionDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionDashRole",
		"SUPPLEMENTARY",
		&returns,
	)
	return returns
}

