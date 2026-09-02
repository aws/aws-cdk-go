package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 bitstream mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3BitstreamMode := medialive_alpha.Eac3BitstreamMode_COMMENTARY()
//
// Experimental.
type Eac3BitstreamMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3BitstreamMode
type jsiiProxy_Eac3BitstreamMode struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3BitstreamMode) Value() *string {
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
func Eac3BitstreamMode_Of(value *string) Eac3BitstreamMode {
	_init_.Initialize()

	if err := validateEac3BitstreamMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3BitstreamMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3BitstreamMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3BitstreamMode_COMMENTARY() Eac3BitstreamMode {
	_init_.Initialize()
	var returns Eac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3BitstreamMode",
		"COMMENTARY",
		&returns,
	)
	return returns
}

func Eac3BitstreamMode_COMPLETE_MAIN() Eac3BitstreamMode {
	_init_.Initialize()
	var returns Eac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3BitstreamMode",
		"COMPLETE_MAIN",
		&returns,
	)
	return returns
}

func Eac3BitstreamMode_EMERGENCY() Eac3BitstreamMode {
	_init_.Initialize()
	var returns Eac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3BitstreamMode",
		"EMERGENCY",
		&returns,
	)
	return returns
}

func Eac3BitstreamMode_HEARING_IMPAIRED() Eac3BitstreamMode {
	_init_.Initialize()
	var returns Eac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3BitstreamMode",
		"HEARING_IMPAIRED",
		&returns,
	)
	return returns
}

func Eac3BitstreamMode_VISUALLY_IMPAIRED() Eac3BitstreamMode {
	_init_.Initialize()
	var returns Eac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3BitstreamMode",
		"VISUALLY_IMPAIRED",
		&returns,
	)
	return returns
}

