package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS program date time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsProgramDateTime := medialive_alpha.HlsProgramDateTime_Of(jsii.String("value"))
//
// Experimental.
type HlsProgramDateTime interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsProgramDateTime
type jsiiProxy_HlsProgramDateTime struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsProgramDateTime) Value() *string {
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
func HlsProgramDateTime_Of(value *string) HlsProgramDateTime {
	_init_.Initialize()

	if err := validateHlsProgramDateTime_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsProgramDateTime

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsProgramDateTime",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsProgramDateTime_EXCLUDE() HlsProgramDateTime {
	_init_.Initialize()
	var returns HlsProgramDateTime
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsProgramDateTime",
		"EXCLUDE",
		&returns,
	)
	return returns
}

func HlsProgramDateTime_INCLUDE() HlsProgramDateTime {
	_init_.Initialize()
	var returns HlsProgramDateTime
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsProgramDateTime",
		"INCLUDE",
		&returns,
	)
	return returns
}

